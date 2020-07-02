package identity

import (
	"bytes"
	"compress/gzip"
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/lib/pq"
	"github.com/rudderlabs/rudder-server/config"
	"github.com/rudderlabs/rudder-server/services/filemanager"
	"github.com/rudderlabs/rudder-server/utils/logger"
	"github.com/rudderlabs/rudder-server/utils/misc"
	warehouseutils "github.com/rudderlabs/rudder-server/warehouse/utils"
	uuid "github.com/satori/go.uuid"
)

type HandleT struct {
	Warehouse warehouseutils.WarehouseT
	DbHandle  *sql.DB
	Upload    warehouseutils.UploadT
}

const (
	mergeRulesTable = "rudder_identity_merge_rules"
	mappingsTable   = "rudder_identity_mappings"
)

func (idr *HandleT) whMergeRulesTable() string {
	return warehouseutils.ToProviderCase(idr.Warehouse.Destination.DestinationDefinition.Name, warehouseutils.IdentityMergeRulesTable)
}

func (idr *HandleT) applyRule(txn *sql.Tx, ruleID int64, gzWriter *misc.GZipWriter) (err error) {
	sqlStatement := fmt.Sprintf(`SELECT merge_property_1_type, merge_property_1_value, merge_property_2_type, merge_property_2_value FROM %s WHERE id=%v`, mergeRulesTable, ruleID)

	var prop1Val, prop2Val, prop1Type, prop2Type sql.NullString
	err = txn.QueryRow(sqlStatement).Scan(&prop1Type, &prop1Val, &prop2Type, &prop2Val)
	if err != nil {
		return
	}

	var rudderIDs []string
	var additionalClause string
	if prop2Val.Valid && prop2Type.Valid {
		additionalClause = fmt.Sprintf(`OR (merge_property_type='%s' AND merge_property_value='%s')`, prop2Type.String, prop2Val.String)
	}
	sqlStatement = fmt.Sprintf(`SELECT ARRAY_AGG(DISTINCT(rudder_id)) FROM %s WHERE (merge_property_type='%s' AND merge_property_value='%s') %s`, mappingsTable, prop1Type.String, prop1Val.String, additionalClause)
	logger.Debugf(`IDR: Fetching all rudder_id's corresponding to the merge_rule: %v`, sqlStatement)
	err = txn.QueryRow(sqlStatement).Scan(pq.Array(&rudderIDs))
	if err != nil {
		panic(err)
	}

	currentTimeString := time.Now().Format(misc.RFC3339Milli)
	var buff bytes.Buffer
	csvWriter := csv.NewWriter(&buff)
	var csvRows [][]string

	// if no rudder_id is found with properties in merge_rule, create a new one
	// else if only one rudder_id is found with properties in merge_rule, use that rudder_id
	// else create a new rudder_id and assign it to all properties found with properties in the merge_rule
	if len(rudderIDs) <= 1 {
		// generate new one and assign to these two
		var rudderID string
		if len(rudderIDs) == 0 {
			rudderID = uuid.NewV4().String()
		} else {
			rudderID = rudderIDs[0]
		}
		row1 := []string{prop1Type.String, prop1Val.String, rudderID, currentTimeString}
		csvRows = append(csvRows, row1)
		row1Values := misc.SingleQuotedJoin(row1)

		var row2Values string
		if prop2Val.Valid && prop2Type.Valid {
			row2 := []string{prop2Type.String, prop2Val.String, rudderID, currentTimeString}
			csvRows = append(csvRows, row2)
			row2Values = fmt.Sprintf(`, (%s)`, misc.SingleQuotedJoin(row2))
		}

		sqlStatement = fmt.Sprintf(`INSERT INTO %s (merge_property_type, merge_property_value, rudder_id, updated_at) VALUES (%s) %s ON CONFLICT ON CONSTRAINT %s DO NOTHING`, mappingsTable, row1Values, row2Values, "unique_merge_property")
		logger.Debugf(`IDR: Inserting properties from merge_rule into mappings table: %v`, sqlStatement)
		_, err = txn.Exec(sqlStatement)
		if err != nil {
			return
		}
	} else {
		// generate new one and update all
		newID := uuid.NewV4().String()
		row1 := []string{prop1Type.String, prop1Val.String, newID, currentTimeString}
		csvRows = append(csvRows, row1)
		row1Values := misc.SingleQuotedJoin(row1)

		var row2Values string
		if prop2Val.Valid && prop2Type.Valid {
			row2 := []string{prop2Type.String, prop2Val.String, newID, currentTimeString}
			csvRows = append(csvRows, row2)
			row2Values = fmt.Sprintf(`, (%s)`, misc.SingleQuotedJoin(row2))
		}

		quotedRudderIDs := misc.SingleQuotedJoin(rudderIDs)
		sqlStatement := fmt.Sprintf(`SELECT merge_property_type, merge_property_value FROM %s WHERE rudder_id IN (%v)`, mappingsTable, quotedRudderIDs)
		logger.Debugf(`IDR: Get all merge properties from mapping table with rudder_id's %v: %v`, quotedRudderIDs, sqlStatement)
		var rows *sql.Rows
		rows, err = txn.Query(sqlStatement)
		if err != nil {
			return
		}

		for rows.Next() {
			var mergePropType, mergePropVal string
			err = rows.Scan(&mergePropType, &mergePropVal)
			if err != nil {
				return
			}
			csvRow := []string{mergePropType, mergePropVal, newID, currentTimeString}
			csvRows = append(csvRows, csvRow)
		}

		sqlStatement = fmt.Sprintf(`UPDATE %s SET rudder_id='%s', updated_at='%s' WHERE rudder_id IN (%v)`, mappingsTable, newID, currentTimeString, quotedRudderIDs)
		logger.Debugf(`IDR: Update rudder_id for all properties in mapping table with rudder_id's %v: %v`, quotedRudderIDs, sqlStatement)
		_, err = txn.Exec(sqlStatement)
		if err != nil {
			return
		}

		sqlStatement = fmt.Sprintf(`INSERT INTO %s (merge_property_type, merge_property_value, rudder_id, updated_at) VALUES (%s) %s ON CONFLICT ON CONSTRAINT %s DO NOTHING`, mappingsTable, row1Values, row2Values, "unique_merge_property")
		logger.Debugf(`IDR: Insert new mappings into %s: %v`, mappingsTable, sqlStatement)
		_, err = txn.Exec(sqlStatement)
		if err != nil {
			return
		}
	}
	for _, csvRow := range csvRows {
		csvWriter.Write(csvRow)
	}
	csvWriter.Flush()
	gzWriter.WriteGZ(buff.String())
	return
}

func (idr *HandleT) addRules(txn *sql.Tx, gzWriter *misc.GZipWriter) (ids []int64, err error) {
	loadFileNames, err := idr.downloadLoadFiles(idr.whMergeRulesTable())
	defer misc.RemoveFilePaths(loadFileNames...)
	if err != nil {
		logger.Errorf(`IDR: Failed to download load files for %s with error: %v`, mergeRulesTable, err)
		return
	}

	mergeRulesStagingTable := fmt.Sprintf(`rudder_identity_merge_rules_staging_%s`, strings.Replace(uuid.NewV4().String(), "-", "", -1))
	sqlStatement := fmt.Sprintf(`CREATE TEMP TABLE %s
						ON COMMIT DROP
						AS SELECT * FROM %s
						WITH NO DATA;`, mergeRulesStagingTable, mergeRulesTable)

	logger.Infof(`IDR: Creating temp table %s in postgres for loading %s: %v`, mergeRulesStagingTable, mergeRulesTable, sqlStatement)
	_, err = txn.Exec(sqlStatement)
	if err != nil {
		logger.Errorf(`IDR: Error creating temp table %s in postgres: %v`, mergeRulesStagingTable, err)
		return
	}

	sortedColumnNames := []string{"merge_property_1_type", "merge_property_1_value", "merge_property_2_type", "merge_property_2_value"}
	stmt, err := txn.Prepare(pq.CopyIn(mergeRulesStagingTable, sortedColumnNames...))
	if err != nil {
		logger.Errorf(`IDR: Error starting bulk copy using CopyIn: %v`, err)
		return
	}

	for _, loadFileName := range loadFileNames {
		var gzipFile *os.File
		gzipFile, err = os.Open(loadFileName)
		if err != nil {
			logger.Errorf(`IDR: Error opeining downloaded load file at %s: %v`, loadFileName, err)
			return
		}

		var gzipReader *gzip.Reader
		gzipReader, err = gzip.NewReader(gzipFile)
		if err != nil {
			logger.Errorf(`IDR: Error reading downloaded load file at %s: %v`, loadFileName, err)
			return
		}

		csvReader := csv.NewReader(gzipReader)
		for {
			var record []string
			record, err = csvReader.Read()
			if err != nil {
				if err == io.EOF {
					break
				} else {
					logger.Errorf("IDR: Error while reading csv file for loading in staging table locally:%s: %v", mergeRulesStagingTable, err)
					gzipReader.Close()
					gzipFile.Close()
					return
				}
			}
			var recordInterface []interface{}
			for _, value := range record {
				if strings.TrimSpace(value) == "" {
					recordInterface = append(recordInterface, nil)
				} else {
					recordInterface = append(recordInterface, value)
				}
			}
			_, err = stmt.Exec(recordInterface...)
		}
		gzipReader.Close()
		gzipFile.Close()
	}

	_, err = stmt.Exec()
	if err != nil {
		logger.Errorf(`IDR: Error bulk copy using CopyIn: %v`, err)
		return
	}

	sqlStatement = fmt.Sprintf(`DELETE FROM %s AS staging
					USING %s original
					WHERE
					(original.merge_property_1_type = staging.merge_property_1_type)
					AND
					(original.merge_property_1_value = staging.merge_property_1_value)
					AND
					(original.merge_property_2_type = staging.merge_property_2_type)
					AND
					(original.merge_property_2_value = staging.merge_property_2_value)`,
		mergeRulesStagingTable, mergeRulesTable)
	logger.Info(`IDR: Deleting from staging table %s using %s: %v`, mergeRulesStagingTable, mergeRulesTable, sqlStatement)
	_, err = txn.Exec(sqlStatement)
	if err != nil {
		logger.Errorf(`IDR: Error deleting from staging table %s using %s: %v`, mergeRulesStagingTable, mergeRulesTable, err)
		return
	}

	err = idr.writeTableToFile(mergeRulesStagingTable, txn, gzWriter)
	if err != nil {
		logger.Errorf(`IDR: Error writing staging table %s to file: %v`, mergeRulesStagingTable, err)
		return
	}

	sqlStatement = fmt.Sprintf(`INSERT INTO %s
						(merge_property_1_type, merge_property_1_value, merge_property_2_type, merge_property_2_value)
						SELECT DISTINCT ON (
							merge_property_1_type, merge_property_1_value, merge_property_2_type, merge_property_2_value
						) merge_property_1_type, merge_property_1_value, merge_property_2_type, merge_property_2_value
		FROM %s RETURNING id`, mergeRulesTable, mergeRulesStagingTable)
	logger.Infof(`IDR: Inserting into %s from %s: %v`, mergeRulesTable, mergeRulesStagingTable, sqlStatement)
	rows, err := txn.Query(sqlStatement)
	if err != nil {
		logger.Errorf(`IDR: Error inserting into %s from %s: %v`, mergeRulesTable, mergeRulesStagingTable, err)
		return
	}
	for rows.Next() {
		var id int64
		err = rows.Scan(&id)
		if err != nil {
			logger.Errorf(`IDR: Error reading id from inserted column in %s from %s: %v`, mergeRulesTable, mergeRulesStagingTable, err)
			return
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func (idr *HandleT) writeTableToFile(tableName string, txn *sql.Tx, gzWriter *misc.GZipWriter) (err error) {
	batchSize := int64(500)
	sqlStatement := fmt.Sprintf(`SELECT COUNT(*) FROM %s`, tableName)
	var totalRows int64
	err = txn.QueryRow(sqlStatement).Scan(&totalRows)
	if err != nil {
		return
	}

	var offset int64
	for {
		sqlStatement = fmt.Sprintf(`SELECT merge_property_1_type, merge_property_1_value, merge_property_2_type, merge_property_2_value FROM %s LIMIT %d OFFSET %d`, tableName, batchSize, offset)

		var rows *sql.Rows
		rows, err = txn.Query(sqlStatement)
		if err != nil {
			return
		}

		for rows.Next() {
			var buff bytes.Buffer
			csvWriter := csv.NewWriter(&buff)
			var csvRow []string

			var prop1Val, prop2Val, prop1Type, prop2Type sql.NullString
			err = rows.Scan(&prop1Type, &prop1Val, &prop2Type, &prop2Val)
			if err != nil {
				return
			}
			csvRow = append(csvRow, prop1Type.String, prop1Val.String, prop2Type.String, prop2Val.String)
			csvWriter.Write(csvRow)
			csvWriter.Flush()
			gzWriter.WriteGZ(buff.String())
		}

		offset += batchSize
		if offset >= totalRows {
			break
		}
	}
	return
}

func (idr *HandleT) downloadLoadFiles(tableName string) ([]string, error) {
	objectLocations, _ := warehouseutils.GetLoadFileLocations(idr.DbHandle, idr.Warehouse.Source.ID, idr.Warehouse.Destination.ID, tableName, idr.Upload.StartLoadFileID, idr.Upload.EndLoadFileID)
	var fileNames []string
	for _, objectLocation := range objectLocations {
		objectName, err := warehouseutils.GetObjectName(objectLocation, idr.Warehouse.Destination.Config, warehouseutils.ObjectStorageType(idr.Warehouse.Destination.DestinationDefinition.Name, idr.Warehouse.Destination.Config))
		if err != nil {
			logger.Errorf("IDR: Error in converting object location to object key for table:%s: %s,%v", tableName, objectLocation, err)
			return nil, err
		}
		dirName := "/rudder-warehouse-load-uploads-tmp/"
		tmpDirPath, err := misc.CreateTMPDIR()
		if err != nil {
			logger.Errorf("IDR: Error in creating tmp directory for downloading load file for table:%s: %s, %v", tableName, objectLocation, err)
			return nil, err
		}
		ObjectPath := tmpDirPath + dirName + fmt.Sprintf(`%s_%s_%d/`, idr.Warehouse.Destination.DestinationDefinition.Name, idr.Warehouse.Destination.ID, time.Now().Unix()) + objectName
		err = os.MkdirAll(filepath.Dir(ObjectPath), os.ModePerm)
		if err != nil {
			logger.Errorf("IDR: Error in making tmp directory for downloading load file for table:%s: %s, %s %v", tableName, objectLocation, err)
			return nil, err
		}
		objectFile, err := os.Create(ObjectPath)
		if err != nil {
			logger.Errorf("IDR: Error in creating file in tmp directory for downloading load file for table:%s: %s, %v", tableName, objectLocation, err)
			return nil, err
		}
		downloader, err := filemanager.New(&filemanager.SettingsT{
			Provider: warehouseutils.ObjectStorageType(idr.Warehouse.Destination.DestinationDefinition.Name, idr.Warehouse.Destination.Config),
			Config:   idr.Warehouse.Destination.Config,
		})
		err = downloader.Download(objectFile, objectName)
		if err != nil {
			logger.Errorf("IDR: Error in downloading file in tmp directory for downloading load file for table:%s: %s, %v", tableName, objectLocation, err)
			return nil, err
		}
		fileName := objectFile.Name()
		if err = objectFile.Close(); err != nil {
			logger.Errorf("IDR: Error in closing downloaded file in tmp directory for downloading load file for table:%s: %s, %v", tableName, objectLocation, err)
			return nil, err
		}
		fileNames = append(fileNames, fileName)
	}
	return fileNames, nil
}

func (idr *HandleT) uploadFile(tableName string, txn *sql.Tx, filePath string) (err error) {
	outputFile, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	uploader, err := filemanager.New(&filemanager.SettingsT{
		Provider: warehouseutils.ObjectStorageType(idr.Warehouse.Destination.DestinationDefinition.Name, idr.Warehouse.Destination.Config),
		Config:   idr.Warehouse.Destination.Config,
	})
	output, err := uploader.Upload(outputFile, config.GetEnv("WAREHOUSE_BUCKET_LOAD_OBJECTS_FOLDER_NAME", "rudder-warehouse-load-objects"), tableName, idr.Warehouse.Source.ID, tableName)
	if err != nil {
		return
	}

	sqlStatement := fmt.Sprintf(`UPDATE %s SET location='%s' WHERE wh_upload_id=%d AND table_name='%s'`, warehouseutils.WarehouseTableUploadsTable, output.Location, idr.Upload.ID, warehouseutils.ToProviderCase(idr.Warehouse.Destination.DestinationDefinition.Name, tableName))
	logger.Debugf(`IDR: Updating load file location for table: %s: %s `, tableName, sqlStatement)
	_, err = txn.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
	return
}

func (idr *HandleT) createTempGzFile(dirName string) (gzWriter misc.GZipWriter, path string) {
	tmpDirPath, err := misc.CreateTMPDIR()
	if err != nil {
		panic(err)
	}
	path = tmpDirPath + dirName + fmt.Sprintf(`%s_%s/%v/`, idr.Warehouse.Destination.DestinationDefinition.Name, idr.Warehouse.Destination.ID, idr.Upload.ID) + uuid.NewV4().String() + ".csv.gz"
	err = os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		panic(err)
	}
	gzWriter, err = misc.CreateGZ(path)
	if err != nil {
		panic(err)
	}
	return
}

// Resolve does the below things in a single pg txn
// 1. Fetch all new merge rules added in the upload
// 2. Append to local identity merge rules table
// 3. Apply each merge rule and update local identity mapping table
// 4. Upload the diff of each table to load files for both tables
func (idr *HandleT) Resolve() (err error) {
	txn, err := idr.DbHandle.Begin()
	if err != nil {
		panic(err)
	}

	// START: Add new merge rules to local pg table and also to file
	mergeRulesFileGzWriter, mergeRulesFilePath := idr.createTempGzFile(`/rudder-identity-merge-rules-tmp/`)
	defer os.Remove(mergeRulesFilePath)
	ruleIDs, err := idr.addRules(txn, &mergeRulesFileGzWriter)
	if err != nil {
		logger.Errorf(`IDR: Error adding rules to %s: %v`, mergeRulesTable, err)
		return
	}
	mergeRulesFileGzWriter.CloseGZ()
	// END: Add new merge rules to local pg table and also to file

	// START: Add new/changed identity mappings to local pg table and also to file
	mappingsFileGzWriter, mappingsFilePath := idr.createTempGzFile(`/rudder-identity-mappings-tmp/`)
	defer os.Remove(mappingsFilePath)
	for _, ruleID := range ruleIDs {
		err = idr.applyRule(txn, ruleID, &mappingsFileGzWriter)
		if err != nil {
			logger.Errorf(`IDR: Error applying rule %d in %s: %v`, ruleID, mergeRulesTable, err)
			return
		}
	}
	mappingsFileGzWriter.CloseGZ()
	// END: Add new/changed identity mappings to local pg table and also to file

	// upload new merge rules to object storage
	err = idr.uploadFile(mergeRulesTable, txn, mergeRulesFilePath)
	if err != nil {
		logger.Errorf(`IDR: Error uploading load file for %s at %s to object storage: %v`, mergeRulesTable, mergeRulesFilePath, err)
		return
	}

	// upload new/changed identity mappings to object storage
	err = idr.uploadFile(mappingsTable, txn, mappingsFilePath)
	if err != nil {
		logger.Errorf(`IDR: Error uploading load file for %s at %s to object storage: %v`, mappingsFilePath, mergeRulesFilePath, err)
		return
	}

	err = txn.Commit()
	if err != nil {
		logger.Errorf(`IDR: Error commiting transaction: %v`, err)
		return
	}
	return
}
