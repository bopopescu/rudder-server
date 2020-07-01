// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package migrator

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// MigrationAssets contains SQL migration scripts and templates
var MigrationAssets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2020, 6, 26, 8, 57, 18, 205508601, time.UTC),
		},
		"/jobsdb": &vfsgen۰DirInfo{
			name:    "jobsdb",
			modTime: time.Date(2020, 6, 26, 8, 57, 18, 205395399, time.UTC),
		},
		"/jobsdb/000001_create_tables.down.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "000001_create_tables.down.tmpl",
			modTime:          time.Date(2020, 6, 26, 8, 57, 18, 205093641, time.UTC),
			uncompressedSize: 265,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x8e\xb1\x8a\x83\x40\x10\x86\xfb\x7d\x8a\xbf\xb8\x56\x5f\xc0\xea\x0e\x3d\x10\x0e\x4e\xa2\x45\x52\x2d\x2b\x6e\x44\x91\x55\x76\x47\x50\x86\x79\xf7\x10\x13\x93\x90\x26\xe5\xcc\x7c\xff\xfc\x5f\x14\x21\xf5\xe3\x04\xe3\x56\xd8\xa5\x0b\xd4\xb9\x16\x8d\x21\x13\x2c\x05\xc5\xec\x8d\x6b\x2d\xe2\xf4\xbe\x11\x51\x00\x90\x1e\xfe\x0b\x54\xdf\x3f\x7f\x19\x98\xbf\xe2\xc2\xdb\x73\xb7\x88\xe8\x7e\xac\x83\x66\x8e\x45\x92\x4f\x9c\x0e\x64\x68\x7e\xd0\xcc\xd6\x35\x22\x4a\xed\x42\xb4\x4e\x16\x3b\x67\xf5\x75\x54\xb7\x77\xa7\x22\x7b\x3b\x24\xcf\x58\x3f\xce\xde\x99\x01\x64\xea\x61\x0f\x6c\xfd\xf9\x2f\xb2\x63\x5e\x56\x25\x98\x5f\x45\x36\x3c\x51\xea\x12\x00\x00\xff\xff\xa5\xa4\xe3\xe3\x09\x01\x00\x00"),
		},
		"/jobsdb/000001_create_tables.up.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "000001_create_tables.up.tmpl",
			modTime:          time.Date(2020, 6, 26, 8, 57, 18, 205214219, time.UTC),
			uncompressedSize: 1358,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x54\x51\x6f\xe2\x38\x10\x7e\xe7\x57\xcc\x03\x12\xad\x04\x9c\x74\xf7\x76\x7d\x0a\xe0\xdb\x4d\x0f\x02\x82\x74\xaf\x7d\x8a\x26\xf6\x00\xee\x1a\x3b\x67\x3b\x2d\x51\x94\xff\x7e\x72\x48\xba\x57\xb6\x2b\x35\x8f\x33\xdf\xf7\xcd\xcc\x97\x19\x4f\x26\x10\x6b\xe9\x25\x2a\x78\x21\xeb\xa4\xd1\x60\xf6\x70\x6f\x72\xb7\x98\x81\xc7\x5c\x91\x1b\x43\x8e\x8e\x04\x18\x0d\x9e\x4e\x85\x42\x4f\x20\xd0\x23\x14\xd6\xbc\x48\x41\x02\xf2\x0a\x08\xf9\xb1\xa7\x49\xed\x3c\x6a\x4e\x83\xc9\x04\xe6\x47\xe2\xdf\xe1\xd9\xe4\x4e\xe4\xbf\x39\xf2\x65\x31\x3d\x18\xd8\x1b\x0b\xa8\x14\xe0\x0b\x4a\x15\x8a\xbc\x57\x9e\x0e\x02\xf5\xde\x94\x56\xa3\x72\x97\x36\x06\xf3\x2d\x8b\x52\x06\x69\x34\x5b\x32\x88\xff\x82\x64\x9d\x02\x7b\x8c\x77\xe9\x0e\xea\x7a\xba\xb1\xb4\x97\xe7\xa6\xc9\x9e\x2f\x2c\xb8\x19\x00\x00\x48\x01\xb3\xf8\xcb\x8e\x6d\xe3\x68\x09\x9b\x6d\xbc\x8a\xb6\x4f\xf0\x37\x7b\x1a\xb7\x59\x53\x90\x45\x1f\x66\xfe\x16\x6d\xe7\x5f\xa3\xed\xcd\x1f\xbf\xdf\xb6\xc2\xc9\xc3\x72\x79\xc1\x08\xa3\x09\x66\xeb\xf5\x92\x45\xc9\x15\x2b\x2b\xb0\x52\x06\x05\xdc\xef\xd6\xc9\xec\x8a\xe7\x3c\x5a\x9f\x79\x79\x22\x48\xe3\x15\xdb\xa5\xd1\x6a\x73\x05\x21\x2d\xae\x00\xb7\x77\xdd\xe4\x79\xe0\x7b\x02\x5f\x15\x34\x86\xd2\x5d\x4c\x0e\x36\x76\x3f\x65\xb0\x58\xc3\x70\x08\x33\xf6\x25\x4e\x5a\xb1\xde\x9e\xa7\x0d\x0b\xb8\xac\xe5\x67\x81\xdf\xa6\xc3\x17\xed\x80\x25\x0f\xab\x9b\xb7\x40\xff\x8d\x5e\x51\x7a\xa9\x0f\xa3\xf1\xcf\x29\x3a\x13\x2f\x7f\x95\x74\x25\xe7\x44\x82\xc4\x47\xc9\x4e\x34\xb3\xe4\x6d\xf5\x11\x60\x8f\x52\x7d\x4c\xc5\xdc\x58\x4f\x62\x74\x7b\xd7\xe6\xd8\xe3\x9c\x6d\xd2\x78\x9d\xbc\x21\xff\xf9\xca\x12\x10\x65\xa1\x24\x0f\x63\x9a\xfc\x99\xb8\x87\x34\x44\x75\xa9\xd4\xdd\x80\x25\x0b\x18\x0e\x2f\x76\x2e\xd0\xa3\x23\xdf\x39\x07\x68\x09\xb4\xf1\xc0\x2d\xa1\x27\x01\x42\x5a\xe2\x5e\x55\xc1\xe1\x93\x3c\x74\x1b\xe1\xb8\x95\x85\x77\x63\xf0\x47\x6a\x29\x3d\xfc\x45\x62\xb7\xe8\x23\xd7\x05\x17\x3b\xd8\x97\x9a\xb7\xbc\xb0\xfc\x84\x62\x1a\x0a\xa7\x47\x82\x7f\x4b\xb2\xd5\xe5\x0f\x4a\xfd\x03\x7f\x2a\x9d\x07\x54\xaf\x58\xf5\x22\x6d\x21\xf1\xbe\xd5\xd2\x49\x7d\x68\x13\xe1\x36\x9c\x07\xc7\x8f\x74\xc2\xfe\x54\xdb\x22\xdf\xba\xb3\x2d\x0b\x11\x30\xed\x6d\xd1\x59\xba\xe0\xfd\xb5\x1e\x47\x0d\x39\x81\xa0\xbd\xd4\x24\xc6\xbd\xfe\xbb\xab\xfe\x4e\x15\x8c\x3a\xcb\xdc\x08\xbc\x01\xe9\xc3\xc2\xd3\xcf\xaa\x52\x0b\xc9\xc9\x75\xb3\x4a\x07\xd2\x01\x6a\xa0\x33\x9e\x0a\x45\xe1\x21\xc1\x37\xec\xb5\xb3\x7d\xf1\x1f\xf1\xbe\x0d\xf7\xe7\x60\x32\x09\x92\x75\x6d\x51\x1f\x08\xa6\x7d\x37\x4d\x13\xc2\xed\x2a\x2f\x53\xb6\xed\xde\x82\xba\x1e\xfe\xff\xfc\x73\x97\xd5\xf5\xb4\x69\x20\x5a\x2c\x60\xbe\x5e\x3e\xac\x12\xd0\xf4\x9a\x71\xa3\xca\x93\x86\x94\x3d\xa6\x77\x9f\x91\x69\x2f\xa8\xfc\xa4\x58\x5d\x93\x16\x4d\xf3\x5f\x00\x00\x00\xff\xff\x7d\x20\xde\x2b\x4e\x05\x00\x00"),
		},
		"/jobsdb/000002_alter_dataset_tables.down.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "000002_alter_dataset_tables.down.tmpl",
			modTime:          time.Date(2020, 6, 26, 8, 57, 18, 205331230, time.UTC),
			uncompressedSize: 802,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x92\x51\x6b\xc2\x30\x10\xc7\xdf\xfd\x14\xf7\x50\x70\x83\xe9\x17\xf0\xa9\xda\x6c\x13\x6a\x5a\x5c\xca\xb6\xa7\x90\x36\xa7\x44\x4a\x2b\x49\xca\x94\x90\xef\x3e\xda\x3a\xc1\x51\x75\xec\x1e\x02\xb9\xfb\xdf\xef\xc2\xff\xe2\x9c\x16\xd5\x16\x61\x1a\x09\x2b\x0c\x5a\xe3\xfd\x08\x00\x20\x8c\x19\x59\x03\x0b\xe7\x31\x01\xe7\x82\x69\xaa\x71\xa3\x0e\xde\xf3\x5d\x9d\x1b\xee\xdc\xd4\x7b\x88\xd6\x49\x0a\x8b\x24\xce\x56\x14\x1a\x83\x9a\x2b\x39\xfb\x73\x73\xaf\x39\x75\x17\x1a\x85\x45\xc9\x85\x05\x83\x16\x22\xf2\x1c\x66\x31\x03\x9a\xc5\xf1\x3f\x89\x78\xd8\x2b\x8d\xc3\xc0\x8e\x38\x99\x80\xc6\x7e\x2e\xec\xea\x9c\x1b\x2b\x2c\x72\x7b\xdc\x23\xb4\x47\xa7\x89\x12\x08\x02\x98\x93\x97\x25\xed\xee\x6d\x2c\xd6\x24\x64\x04\xd8\x67\x4a\x7e\xf5\x9d\x25\xdd\x83\xdf\x80\xd0\x6c\xf5\x70\x91\xfc\x89\xf1\x97\x50\x56\x55\xdb\xf1\xd3\x70\x19\x0f\x58\x34\xb7\x04\xa6\x29\x0a\x44\x89\xf2\x9a\xe0\x34\x80\x6b\xb4\xfa\x78\x4d\xb4\x11\xaa\xbc\x8e\x10\x79\xad\x2d\xca\xf1\xe3\xec\xa2\x4e\x3e\x16\x24\x65\xcb\x84\x5e\x64\xdf\x5f\x09\x05\xd9\xec\x4b\x55\xb4\x76\xd4\xf9\x0e\x0b\x0b\xac\xcd\x56\x4d\x59\xf6\x08\x42\x23\x08\x82\x93\xff\xb7\x36\xda\x99\xda\x0c\xee\xf5\x6c\xf9\xd0\x06\xee\xff\x95\x3b\x64\x25\x7b\xec\x92\xb2\xd9\xc8\x39\xac\xa4\xf7\xdf\x01\x00\x00\xff\xff\xea\x0e\x28\x43\x22\x03\x00\x00"),
		},
		"/jobsdb/000002_alter_dataset_tables.up.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "000002_alter_dataset_tables.up.tmpl",
			modTime:          time.Date(2020, 6, 26, 8, 57, 18, 205445325, time.UTC),
			uncompressedSize: 718,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x91\x51\x6b\x9c\x40\x14\x85\xdf\xfd\x15\xe7\x41\x48\x0a\xd9\x85\x42\xe9\x8b\x4f\xee\x3a\xc9\x0a\x66\x14\x77\xb6\xd9\x3e\xc9\xb8\xde\x46\x83\xa8\xcc\x8c\x69\xc2\x30\xff\xbd\x8c\xdb\x6e\x28\x14\x52\xda\x37\xef\xf5\xdc\xef\xde\x73\xc6\x5a\x25\x87\x47\xc2\x3a\x91\x46\x6a\x32\xda\xb9\x00\x00\xe2\x4c\xb0\x12\x22\xde\x64\x0c\xd6\x86\xeb\x42\xd1\xb7\xee\xc5\xb9\xea\x69\xac\x75\x65\xed\xda\xb9\x9f\x9a\x6d\x9e\x1d\xee\x39\x4e\x8a\xa4\xa1\xa6\x92\x06\x9a\x0c\x12\x76\x1b\x1f\x32\x01\x9e\x3f\x5c\x7f\x88\xfe\x0d\x49\x2f\x53\xa7\xe8\x7f\x89\x49\xf2\x8b\x97\xde\x82\xe7\x02\xec\x98\xee\xc5\x1e\xb3\x26\x55\x75\x0d\x04\x3b\x8a\xa5\xcf\x0f\x59\x76\xd9\x72\xb5\xfa\x78\x15\x05\xcb\x96\xd5\x0a\xdb\x71\x78\x26\x65\xf0\x34\xd6\x95\x36\xd2\x10\xcc\x88\x67\xa9\x4e\xad\x54\x37\x68\xd4\x38\xc1\xbc\x4e\xf4\xf6\xbf\xf2\xe5\xbb\x37\x2e\xda\xf9\x8f\xde\xdf\x36\x89\xaf\x05\xc3\x97\xb8\xdc\xee\xe2\xf2\xfa\xf3\xa7\xbf\xb0\xfe\x0e\xd6\x7b\xf6\xcc\x4d\x7a\x97\x72\x11\x05\x17\x97\xf7\xdd\x63\x6b\xa0\x4d\xd7\xf7\xa8\xc9\x07\xd4\xa0\x7e\xc5\x68\x5a\x52\x7e\x52\x37\x35\xba\x41\x1b\x39\x9c\xe8\x06\xbd\xd4\x06\xe3\x40\x98\xa7\xc6\xbf\x3c\xbe\xfb\xb9\x73\x16\x2d\x2d\x79\xac\x17\x72\x92\x23\x0c\xb1\x61\x77\x29\x5f\xea\xa5\x57\xe6\xc5\xf9\x8a\xdf\x23\x8b\x2e\x0a\x76\xdc\xb2\x42\xa4\x39\xc7\xc3\x8e\x71\xe4\x62\xc7\xca\x3d\x84\xff\x1e\xe6\xbe\x3f\x0b\x19\x4f\x10\x86\x51\x60\x2d\x0d\x8d\x73\xc1\x8f\x00\x00\x00\xff\xff\x3e\x89\xb1\x86\xce\x02\x00\x00"),
		},
		"/warehouse": &vfsgen۰DirInfo{
			name:    "warehouse",
			modTime: time.Date(2020, 6, 29, 6, 38, 36, 387554454, time.UTC),
		},
		"/warehouse/000001_create_tables.up.sql": &vfsgen۰CompressedFileInfo{
			name:             "000001_create_tables.up.sql",
			modTime:          time.Date(2020, 6, 26, 8, 57, 18, 205618043, time.UTC),
			uncompressedSize: 4017,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x57\x4d\x73\x9b\x30\x10\x3d\x97\x5f\xb1\x07\x1f\xe2\x99\xf8\xd6\xe9\x85\x13\xb6\xd5\x84\x16\x83\x07\x94\xd6\x39\x69\x54\x50\x12\xcd\x10\xf0\x80\xdc\xa6\xff\xbe\x03\x22\x20\x30\xc8\x4a\xe2\x74\xda\x9b\xc7\xbb\xfb\xd8\x8f\xb7\x7a\x92\xb5\x58\x58\x8b\x05\xfc\x7a\x20\xa5\xa0\xf7\x3c\xbb\x27\x77\x3c\x65\x65\xf5\xb7\xb5\x0a\x91\x83\x11\x60\x67\xe9\x21\x70\x3f\x83\x1f\x60\x40\x3b\x37\xc2\xd1\x91\x3f\x5c\x58\x00\x00\x3c\x81\xa5\x7b\x15\xa1\xd0\x75\x3c\xd8\x86\xee\xc6\x09\x6f\xe1\x2b\xba\xbd\xac\xad\x69\x1e\x53\xc1\xf3\x0c\x30\xda\xe1\x1a\xcd\xbf\xf1\x3c\x69\x2b\xf3\x43\x11\x33\xc2\x13\xf8\xe6\x84\xab\x6b\x27\xbc\xf8\xf4\x71\x3e\xf0\x49\x58\x29\x78\x56\x43\xe8\x1d\xcb\xf8\x81\x3d\x52\xf8\x12\x05\xfe\x72\x60\x62\x45\x91\x17\x75\x02\x8d\xab\xa0\xe2\x50\xaa\x58\xf2\xff\x3b\x5e\x94\x82\xb0\x9f\x2c\x13\x84\x0a\xc0\xee\x06\x45\xd8\xd9\x6c\x9b\x4a\xa8\xc6\x28\x72\x41\x53\x69\x2d\xab\x76\xb8\x7e\xf3\xad\xb8\x60\x54\xb0\xa4\x17\x32\x48\xef\xb0\x4f\xa6\x5d\xe6\xb6\x65\x39\x1e\x46\x61\x33\x92\xa3\xa1\x55\x08\xce\x7a\x0d\xab\xc0\xbb\xd9\xf8\x83\x91\xe9\x2b\x9a\x0c\xd3\x96\x3a\x19\x35\xd2\x03\xbb\x25\x94\xeb\xaf\xd1\xee\x04\xa1\x08\x4f\x08\xcf\x12\xf6\x04\x81\x3f\xc2\xb6\x96\x2e\x97\x03\x56\x9c\xea\x11\x48\x63\x93\x74\x33\x7d\x7c\xbb\x45\x2a\x05\x6c\xcb\x5a\x07\x30\x9b\xc1\x12\x5d\xb9\x7e\x5d\xe9\x3a\x0c\xb6\xd2\x4f\x41\xac\xc2\x19\x11\xbf\xf7\xcc\xae\x9d\xd0\x6e\x85\xb6\xd8\x0d\x7c\xf8\x7e\x8d\x7c\x08\xf0\x35\x0a\x23\xc0\xd5\xef\xec\x90\xa6\xb6\x85\xfc\x35\xcc\x66\xb6\xd5\x6d\x5d\x9a\xd3\xc4\x78\xe5\x3a\x67\xa3\x7d\x53\x0b\x27\xd2\xb5\xe5\xe2\x5f\xdb\x45\xd5\xb1\xea\x94\xc6\x55\xd0\x1f\x29\x23\x19\x7d\x64\x63\x59\xbd\x6e\xad\x46\xf8\xa0\x74\xb1\x47\x06\xf5\xf3\xd5\xa0\xab\x1c\xf4\xd1\xe7\x64\x7f\x87\x4b\x9a\xfe\xf7\x5b\x4c\xba\xf4\x7a\x8b\xa1\x72\x62\x72\x2b\x2e\x95\xe2\xe6\x2a\xfd\x0e\xfb\x2a\xde\x84\x7b\x8d\xa7\x19\xf1\x0c\x08\x54\xa5\x52\xee\x69\xac\x23\xc4\x7b\x90\xac\x14\xb4\x10\x44\xb7\x1a\x2c\x4b\xb4\x76\x89\xd0\xf6\x7d\x2c\x7c\xd2\x78\x2c\x37\x2f\x95\xae\xda\xf2\x66\x8d\x92\xc6\x27\x16\x8f\xe9\x17\x7f\xe4\xd9\x7d\xa9\x7e\xe9\x1d\x94\xeb\x99\x79\xff\x8e\x66\xe9\x7a\x32\xbd\xea\x6a\xb3\x4e\x2e\x79\x53\x34\x91\x34\xe8\xed\x71\xbb\x5f\xd2\x36\x37\xc7\x1a\x3d\x2d\xc6\xa1\x5f\xa0\x9a\xcf\x41\xe7\xd0\x4b\x89\x75\x0e\xb9\x94\xc7\x98\xe1\xa9\xf5\x61\x18\x60\x74\x78\x75\xf9\xf6\xfc\xb4\x52\x65\xb8\xdc\xc3\xcb\x67\xc7\x38\xc1\x2b\x24\xcd\x3d\xb2\x0b\x7a\x87\x5d\xec\x37\xe9\x85\xb2\xa6\xc7\x52\xd9\xd3\x94\xaf\x68\xab\x9e\xe1\x3d\x28\xa2\xce\x65\x4a\x0c\x07\xd3\x56\x43\x86\x12\x68\x9e\xf6\x2b\x49\xaf\x22\x9e\x83\xfa\x52\x19\x8c\x5e\x66\xd2\xf3\x0d\x6c\xff\x0f\x64\xdc\xfc\x89\x77\xf2\x7e\x58\x8f\xad\xe5\xe0\x51\x1f\x75\x27\xec\xe9\x47\x4d\x83\x31\x08\x6e\xdb\xd6\x7f\xe4\x3c\x0f\x6e\x78\x7b\x6b\xdd\xe7\xf6\x9f\x00\x00\x00\xff\xff\x89\x94\x7c\x27\xb1\x0f\x00\x00"),
		},
		"/warehouse/000002_create_id_resolution_tables.up.sql": &vfsgen۰CompressedFileInfo{
			name:             "000002_create_id_resolution_tables.up.sql",
			modTime:          time.Date(2020, 6, 30, 15, 47, 11, 202634791, time.UTC),
			uncompressedSize: 842,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x92\xc1\x4f\xc2\x30\x14\xc6\xef\xfb\x2b\xde\x71\x4b\xd8\x41\x63\xbc\x70\x2a\x50\xb4\xb1\x1b\xd8\x75\x0a\xa7\xa6\xd2\x06\x9b\xcc\xad\xd6\x56\xc3\x7f\x6f\x18\x04\x30\x6c\x24\x78\x6d\xbf\xf7\x7d\xef\xfb\xe5\x45\x69\x1a\xa5\x29\xfc\xbc\x0b\x2f\xdf\x2a\x2d\x82\xad\x1a\xa9\xbe\xb6\xcf\x11\xa2\x1c\x33\xe0\x68\x44\xf1\xb9\x00\x00\x00\x4d\x26\x30\x9e\xd1\x32\xcb\xa1\x6a\x56\xd2\x9b\xa6\x06\x8e\x17\x7c\x18\x1d\x6d\x8d\xd2\xb5\x37\x7e\x23\x3e\xb4\x5b\x6b\xe1\x42\xa5\x77\xee\x63\x86\x11\xc7\x7b\x7b\x32\x85\x7c\xc6\x01\x2f\x48\xc1\x0b\x70\x41\x29\xed\x3a\x47\x21\x6e\x93\x8d\x82\x11\x79\x28\x30\x23\x88\xc2\x9c\x91\x0c\xb1\x25\x3c\xe1\xe5\xa0\xfd\xdd\xe9\xad\x6b\xac\x76\x7e\x23\x6e\x84\xdf\x58\x0d\x2f\x88\x8d\x1f\x11\x8b\xef\xef\x92\x36\x2c\x2f\x29\xed\xd1\x7f\xcb\x2a\xe8\xb6\xca\x65\xe5\xed\x99\x73\x8f\xec\x68\xb8\x13\xac\x9c\x96\x5e\x2b\x21\x3d\x70\x92\xe1\x82\xa3\x6c\x7e\x88\x82\x09\x9e\xa2\x92\x6e\xb3\x5f\xe3\x24\x19\x46\x9d\x38\xa5\xb5\xa6\x5e\x5f\xcf\x72\x3f\xf7\x1f\x90\xd7\x61\xec\x85\x78\x58\xe9\x82\x57\xb0\xea\x2a\x40\xa7\xb7\xda\x57\xf9\xe4\x66\xf3\x82\x33\x44\x72\x0e\xa1\x36\x9f\x41\x8b\xbf\xab\x43\x99\x93\xe7\x12\x43\xdc\xd1\x7f\xd0\x59\x33\x19\x46\xbf\x01\x00\x00\xff\xff\x85\x01\x28\x73\x4a\x03\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/jobsdb"].(os.FileInfo),
		fs["/warehouse"].(os.FileInfo),
	}
	fs["/jobsdb"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/jobsdb/000001_create_tables.down.tmpl"].(os.FileInfo),
		fs["/jobsdb/000001_create_tables.up.tmpl"].(os.FileInfo),
		fs["/jobsdb/000002_alter_dataset_tables.down.tmpl"].(os.FileInfo),
		fs["/jobsdb/000002_alter_dataset_tables.up.tmpl"].(os.FileInfo),
	}
	fs["/warehouse"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/warehouse/000001_create_tables.up.sql"].(os.FileInfo),
		fs["/warehouse/000002_create_id_resolution_tables.up.sql"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
