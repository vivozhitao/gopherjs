// Code generated by vfsgen; DO NOT EDIT.

// +build !gopherjsdev

package gopherjspkg

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

// FS is a virtual filesystem that contains core GopherJS packages.
var FS = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2019, 3, 5, 14, 48, 17, 664668599, time.UTC),
		},
		"/js": &vfsgen۰DirInfo{
			name:    "js",
			modTime: time.Date(2019, 1, 3, 14, 55, 7, 232294000, time.UTC),
		},
		"/js/js.go": &vfsgen۰CompressedFileInfo{
			name:             "js.go",
			modTime:          time.Date(2019, 1, 3, 14, 55, 7, 232438709, time.UTC),
			uncompressedSize: 8002,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x59\x5f\x6f\xdc\x36\x12\x7f\x5e\x7d\x8a\x39\xa1\x40\x56\xcd\x56\xbe\xb6\x86\x51\x38\xe7\x87\xa4\xb9\xfa\xdc\x4b\xdc\x00\x6e\xd0\x07\x23\x30\xb8\xd2\x68\x97\xb1\x44\xea\x48\x6a\x37\x7b\xb6\xbf\xfb\x61\xf8\x47\x2b\xad\xa4\xc4\xbe\x24\x2f\x75\xc5\xe1\x6f\x7e\x9c\x19\xce\x1f\xee\xd1\x11\xbc\x63\xd9\x2d\x5b\x21\x7c\xd4\x50\x2b\xb9\xe1\x39\x6a\x28\x1a\x91\x19\x2e\x85\x86\x42\x2a\xe0\xc2\xa0\x62\x99\xe1\x62\x05\x5b\x6e\xd6\x20\x98\xe1\x1b\x84\xdf\xd9\x86\x5d\x65\x8a\xd7\x06\x5e\xbe\xbb\xd0\x29\xfc\xca\xca\x52\x83\x91\x60\xd6\xa8\xb1\x83\xc2\x14\x82\x51\xc8\x0c\xe6\xa0\x6b\xcc\x38\x2b\xcb\x1d\x2c\x77\x70\x2e\xeb\x35\xaa\xdf\xaf\x80\x89\x1c\x8c\x62\x42\x97\x56\x28\xe7\x0a\x33\x53\xee\x3c\x18\x57\x90\x49\xa5\x50\xd7\x52\xe4\x44\xa3\xa3\x5a\xef\x84\x61\x9f\xd2\xe8\xe8\x28\x3a\x3a\x82\xf7\x1a\xe1\x2d\xbb\xc5\xbf\x14\xab\x6b\x54\xb4\x1f\x3f\xd5\x52\x23\x54\x68\xd6\x32\xb7\xf4\xf6\xbb\x53\xf8\x6b\x8d\x02\x6a\xa6\x35\xc1\x6e\x58\xd9\xa0\x6e\xb5\x2f\x48\x37\x14\xb2\x2c\xe5\x96\x96\xcd\xae\x46\xc8\xa4\xd8\xa0\xd2\xed\xb9\x6a\x54\x85\x54\x15\xe6\xa7\x9e\x02\xdc\xc3\xb9\x74\xb2\xfd\x7f\xf7\x5d\xda\x9d\xf5\x7b\xf8\xb5\x83\xb9\x64\xd9\x2d\x91\xb4\x56\x2f\x58\x86\x77\x0f\x70\xef\x71\x7f\x18\xfb\xf7\xd4\xef\x5d\x09\x8f\xbb\x94\xb2\x84\xc1\xbf\x7b\x78\x25\x65\x89\x4c\x0c\xbe\x8f\xcb\x77\x24\x3c\x2e\x9d\x61\x85\x4a\x5b\xf7\x16\xa5\x64\x46\xdb\xfd\x97\x4d\xb5\x44\x35\xd4\x67\x45\x4e\x8e\xbf\x88\xab\x8d\x22\x7f\x0c\xf6\x5f\x4d\x7c\x1f\x97\x1f\xe2\x5e\x7f\xe0\xc2\xfc\x32\xdc\x7f\x21\xcc\x2f\x2f\x95\x62\xbb\x83\xef\xe3\xf2\x13\xb8\x3f\x9e\x8c\xe1\xfe\x78\x32\x00\x9e\x92\x9f\xc0\xfd\xf9\xa7\x85\xfb\xa3\x87\xfb\xf3\x4f\x53\xb8\xd3\x74\x3b\xb8\xcd\xc8\xc1\xee\xe1\x3d\x1f\x33\xc4\x94\xfc\x14\xee\xe1\xc1\x1c\xee\xd0\x10\x53\xf2\x53\xb8\xce\x10\x4d\x7b\x44\x87\x3b\x34\xc4\x7d\x4f\xea\xf3\xb8\x36\x22\x7f\xfe\xe9\x80\xef\x6f\xee\xeb\x01\xf0\x94\xfc\x24\xee\x41\xa4\x7b\xdc\x93\xe3\x29\xdc\xc9\x9b\x11\x70\x59\x59\x82\x34\x6b\x54\xa0\x4b\x9e\xa1\x0e\xfb\x87\xb1\x0b\xfb\x78\x68\xb3\xcc\x67\x70\x69\xbf\x1e\xee\xd7\x88\x4e\x53\x2f\xdd\x4d\x7d\x1f\xe2\xee\x2b\xc4\x81\x1d\xfc\xf7\x43\x7d\x24\x3f\x4f\xd3\xb4\xc3\x3a\x81\xef\x3f\xea\xf4\x8f\xe5\x47\xcc\x4c\x8b\x6b\x78\x85\xe9\x9f\xbc\xc2\x83\xfd\xaf\x99\x19\x63\x33\x21\x3f\xe4\xfb\xc3\xf8\x2a\x70\xa1\x0d\x13\x19\xca\x02\x2e\x65\xbe\xcf\xeb\x1d\x6a\x9f\xc5\xad\x58\xad\x17\x94\xa5\x9a\xcc\xe8\x71\xdc\x0e\x8c\x95\xbf\x76\x39\x6d\xdc\x81\xf7\xbe\x14\xbd\xcc\x73\x4e\x76\xa4\x72\xbb\xb0\xb5\x9c\x79\x2d\x54\xc6\x0c\xe3\x82\xd2\x22\xeb\xf2\x2c\x38\x96\xf9\x02\xa4\xa0\xe2\xbb\xb6\xe5\xce\xa0\x30\x20\x0b\x57\x0c\x69\x19\xb6\xbc\x2c\x61\x89\xb6\x6e\x62\xde\x2f\xa9\x36\xd7\x6f\xc8\xf7\x54\xd2\x58\x1a\xd5\x6d\x83\x11\x11\x27\xaf\x87\x6b\x60\x81\x04\x2a\xcf\x6d\xd8\x58\x48\x2b\xdd\x69\x2d\xb8\xd1\x6d\x29\xff\x06\x6d\xc5\xb0\x91\x80\x97\x20\x78\x09\xb5\xb4\x96\x25\xc9\x3d\x63\xfc\x4f\xc3\xca\xfe\x71\x9f\x69\x88\x45\x53\x96\x71\x1a\xe4\x32\x26\x40\x48\x43\xf6\x69\xc8\x3a\x8c\x4e\x5a\xb1\x1a\x6e\x71\x97\x46\xf6\x42\x78\x49\xe7\x8a\x3b\x7f\x48\xf8\xde\x7f\x7e\xb0\x76\x3a\x47\x03\x0a\x4d\xa3\x84\xb6\x96\x77\x42\xcf\x6c\x97\x56\xa3\x32\x3b\xd7\x8b\xd1\xd2\x8a\x6f\x50\x38\x78\xba\x21\x30\x97\x01\x2b\x21\x98\xf9\x2d\xee\x7c\x09\x4c\x5a\x25\x77\x1e\x1c\x64\xea\x6d\xec\x25\x13\xaf\xff\x0a\x0d\x50\x5b\xb4\xf2\xfa\x6d\x6f\xe4\x0d\xf7\xff\x92\xb9\xea\x91\x59\x78\xcc\xde\x6d\xbe\xdb\x13\xf2\xd2\x5e\x2c\xf0\x7a\x8d\x25\x1a\x04\x85\x95\xdc\xe0\x57\x99\xc6\x21\xf5\xac\xd3\xd1\xbe\x5f\x0d\x9a\xdf\xa0\x58\x99\xf5\xb8\x53\xe2\xd2\x2e\xc6\x2d\x85\x85\x6f\x14\x8d\xbb\x1f\x5c\x98\x11\x06\x0e\x71\x9e\xd0\xf2\x88\x47\xda\x65\xa7\xff\x42\xe4\xf8\xa9\xa7\x9e\x3f\x33\x6b\xc0\x12\x2b\x7f\x43\x99\x70\xa9\x7a\x44\x95\xdd\x3c\xe7\xa4\xe9\x73\x41\xe0\xc5\x3a\x41\xe0\xb4\x6a\x34\x4f\x56\x19\x36\x3b\xad\x8f\xf0\xb6\x97\x3e\x70\x38\x5d\x7d\xc8\xdc\xfd\xef\x9a\xdc\x65\x81\x43\x57\x0b\x56\xe1\x08\x17\x02\x99\xd3\x5a\x1b\x7b\x4c\xad\x34\x0c\x6a\xc9\xa4\x61\x5a\x00\xb7\x33\x4d\xd3\xbd\x5b\x36\xf2\x16\x07\x0c\x29\x53\x61\x59\xa4\xf0\xe7\x9a\x6b\x97\x31\x0b\xc6\x4b\xe0\x05\x70\x9b\x4c\x28\x47\xb0\xb6\x04\x8e\xba\x8c\x80\xe7\x4f\x24\xda\xd9\xd5\x21\x79\x89\x5b\xc8\x6c\xaa\xa4\x6c\x24\x70\xdb\xd6\x16\x97\xd9\xb9\x76\xa5\x3a\xe4\xdb\x51\xd2\x7d\xc6\x30\xcf\xa4\x70\x29\x4c\xaa\x64\x84\xff\x25\x6e\x9f\x4a\x3e\x6c\xe9\x30\xa7\x19\x64\xe4\xce\xf5\xaf\x97\x1d\x48\x58\x96\x49\x65\xc7\xc3\x7e\x41\x3a\x1c\xdb\x46\xa8\x92\x92\x79\xe2\x60\x86\xac\xfc\xaa\xbf\x12\x6e\x96\xf8\x12\x23\x3f\x72\x7c\x05\x27\xa7\x68\x9e\x04\xa8\x21\xaf\x56\x22\x04\xe2\x58\xc5\x18\xe4\xa1\x47\x73\x82\x79\xcd\x94\xc6\x0b\x61\xc6\xbc\x7b\x21\xcc\x64\xe2\x72\x6b\x2d\xab\x93\xe3\xc7\xf0\x3a\x39\xfe\x76\xcc\x4e\x8e\x1d\xb7\x93\xe3\x71\x76\x76\xdd\xf1\x7b\xcf\x1f\x45\xb0\xf9\x96\x0c\x9d\xce\x79\x12\x50\x87\x1c\x5b\x09\x47\xd2\x0e\x06\x5f\xe4\x18\x86\x84\x27\x92\xb4\xe0\x63\x34\xed\xc2\x3c\x69\x71\x87\x34\x83\x44\xeb\x6a\x77\xc9\x1f\xe3\xee\x90\x0e\x52\xb8\x42\x04\xc3\x96\x25\xd5\x06\x08\xdd\x62\x26\x2b\x5b\x62\xa8\x31\xcc\xd1\x30\x5e\x8e\xdd\x91\x56\xa3\x73\x77\xdb\x09\x8f\x3a\xbd\x95\xf4\x8e\x17\x9a\x15\xa3\x54\xa9\x63\x13\xd6\x37\xb5\x51\x0b\xd8\xae\x79\xb6\xb6\x6d\xdd\x12\x3b\xc7\xd8\x70\x06\x8d\xc5\x48\xdf\xb9\x66\x31\x85\x4b\x69\x2c\x0f\x91\x63\x6e\xa9\xd7\xcd\xb2\xe4\x19\x35\x82\x63\x61\x60\x77\xfb\x30\xa8\x8d\x1a\x8b\x83\x20\xe2\x38\xff\x53\x29\xa9\x00\x45\xc6\x6a\xdd\x94\x36\x9b\x77\xfc\x8b\xb4\xaa\x29\x79\x4b\x8d\xae\x3b\x6e\x94\xc0\x9c\x28\x49\x60\x70\x2e\xa1\x66\x82\x67\xb6\x2d\xae\xd8\x8e\xce\xa3\x30\x93\x1b\x54\x98\x2f\xa8\x80\xda\x94\x25\xe0\x7b\xa7\xc7\xac\x99\x81\xb5\x2c\x73\x67\x9d\x43\x4d\xa1\x58\xb8\x9e\xd6\x6d\xf1\xd3\xc5\x5d\x34\xf3\xa7\x8c\xba\xc4\xbb\xb6\xae\x50\x6b\x72\xb4\x1f\x2c\x3a\x67\xca\xa7\x35\x39\x13\xa2\x52\x9e\x62\xe2\x80\x3b\x49\x32\x9a\x79\x13\xc6\x87\x20\xa7\x10\xc3\x73\xfa\xd3\x76\xba\xb1\xd7\x1f\x27\x6d\x1a\x8d\x42\x82\x67\xd9\x6d\x8f\xaa\xb6\x5f\xda\xe6\xf2\x2b\x19\x5b\xfc\x31\xc6\x2d\x35\xab\x6f\x48\xec\xbc\x94\x4b\x56\xda\x3e\x47\xf7\x27\x90\x95\x5b\xf1\xe1\x3b\x8f\xb7\x5c\xe4\x72\x1b\xdb\x08\x5c\x2a\xb9\xd5\xe1\x0d\x2e\x3e\x7f\xf3\xc7\xab\x97\x6f\xdc\x0a\x8d\xaa\xe9\x47\x9d\xa4\xd1\x86\xa9\x80\x1e\xdc\x46\x0a\xdf\xca\xbc\x29\xd1\x2b\xdc\xcf\x00\xfe\xfc\x71\x65\x97\x63\xd8\x30\xc5\xed\xf5\xd5\x68\x68\xfa\xf2\xb8\x29\xfc\x8b\x0b\x73\xea\x06\x09\x70\xc2\xf6\x31\x56\x19\xd7\xb4\x3d\xfb\xa8\x53\xa7\xc2\x1d\xdb\xad\x69\x3a\xf8\xfe\x7f\x2f\x59\x85\xf1\x82\x5a\x88\xe4\x99\x23\xea\x59\x75\x89\xbe\x17\x39\x16\x9c\x22\x7d\xcf\xb5\xe3\x11\x47\x3b\x6e\x82\x54\xec\x80\xf6\xbb\xba\x58\xaf\x71\xd9\xac\x56\xa8\x60\x45\x2d\x6f\x26\xab\x9a\x97\x87\x33\x2e\x35\xfc\xb9\x97\x7b\x11\x53\x7c\x18\xdb\x10\x7b\x77\x07\x88\x79\x02\x77\x9d\xcc\x28\x58\xe9\x1b\x9f\x5e\x0f\xef\x97\x86\x53\xaf\xbb\x7f\x0a\x6b\x85\x1a\x85\xd1\xc0\x1f\x93\x60\xfa\xaa\x5c\xef\x3d\xd2\x7a\xb5\x51\x27\x78\xe9\xe3\xeb\x2d\xbb\xc5\xdf\x08\x62\xab\x58\xad\xbb\x9d\x1e\x85\x8e\xb3\x2c\xcb\x32\xd4\xe1\x8d\x3f\xbc\x97\xcb\xe2\xc0\x36\xd4\x4f\xc6\x2e\xe0\x98\x5a\x35\x64\x1a\x1d\xd3\x14\xb6\x95\x2a\x0f\x79\x3c\xa8\x9b\x17\xc2\x3d\xec\xd8\x2e\xd4\x13\xb4\x5d\xb6\xdb\x08\xd7\x1f\xda\x8c\xf9\x85\xb3\xb8\x18\x76\xbd\x7a\xfc\x5d\xe5\x15\xc4\x8b\x43\xa3\x14\x22\x09\x97\xea\xdf\xb8\xd3\x3d\x7f\xdc\xd2\x07\x1f\xe2\x6e\xa4\x18\x3e\x47\xb8\x03\xd0\xd6\x6e\x3a\xbf\xfe\xb0\xbf\xd2\xbc\x00\x09\x67\x67\xf6\x29\xe1\xfe\xde\xfd\xbd\x8f\xb7\xbb\x68\xd6\x35\xff\xec\x21\x9a\x31\x38\x3d\x0b\xfc\xed\x6d\x70\xa8\x71\xe2\x4f\x43\xb4\xe2\x05\xc8\x24\x9a\x69\x12\xa5\xc3\xcd\x83\xc6\x05\xb0\x76\x58\x4c\xa2\x99\xfd\xd1\x86\x84\xfe\xfe\x02\x38\xfc\xa3\xb3\xf8\x02\xf8\xf3\xe7\x56\xbd\xbe\xe6\x1f\xe0\x0c\x58\x3b\xf1\xed\xb3\x0d\xd1\xf1\xec\x74\x27\x34\xc2\x4f\x2a\xfb\x31\x62\x18\xb1\xae\x54\xae\x99\xb6\x31\x54\x53\xda\x29\x6c\x21\x09\x37\x1f\xf3\xf6\xf5\x46\x16\x14\xd0\xef\xb5\x5d\x2a\x79\xc6\x0d\x5d\x39\x83\xca\x06\x8e\x76\x7f\x76\x7e\xb5\xf1\xbf\xe3\xf8\x0a\x63\x1f\xa2\x0e\x7f\xcd\xd9\x07\x96\x27\xfb\x99\xf0\xdf\x90\x81\x0e\x2f\x4b\x12\xcd\xe4\xa4\x23\x68\x38\x21\x01\x97\x9e\x6e\x6e\xc2\xcd\xbd\x71\x87\xbf\xb9\x89\x17\xb0\x49\xa2\x59\xe0\x7c\x7a\x06\x1b\x07\xd1\x19\x94\xe2\x24\x94\x1f\x2b\x14\x8f\xb8\xcb\x2f\x8d\x38\xad\xb2\x9e\xf7\xcb\xc1\x71\xd1\x8c\xa2\xad\x72\xb0\xf5\xed\xaa\x53\x38\xe0\x6f\x67\x10\xc7\x70\x07\x47\x47\x76\x78\x0b\x3e\x88\x66\xb3\x59\x26\x85\xe1\xa2\xc1\x68\x46\xfe\xf6\xa7\xf2\x28\x34\xe7\x76\x60\x16\xee\x7e\x86\x59\xae\x0d\xf8\x8e\x35\x67\xe3\x57\x10\x3f\x39\x13\xf1\xff\x62\x78\xd3\x25\x23\x59\x2d\x81\xb1\x92\x75\x47\x57\xb2\x08\x47\x31\xbb\x3a\x4e\x16\x60\x54\x83\xe1\x12\xb0\xba\x2e\x77\x04\xe0\x86\x70\x3a\xfa\x43\x2f\x5e\x65\xd4\x8e\xbb\xf6\xcd\xfb\x55\x53\x14\x53\x21\xdb\x15\x28\x94\xac\x80\xc1\x72\x67\xfc\xc3\xb5\x0f\xa5\x3e\xce\x7c\x09\xd7\x1f\x48\xa6\x77\x74\xf7\xd0\x3d\x0c\xa6\x25\xc5\x4a\x51\x50\x51\x3c\x3d\xf3\xa8\xf6\x60\xdf\xb9\xaf\x71\xe2\xe6\xa4\x68\xe6\xde\x8e\x0e\xa5\xfc\x8b\x52\x2b\x15\xae\x64\x47\xc4\xbe\xbc\x84\x88\x5a\x5a\x8e\x6d\xc2\xb0\x72\x94\x31\xac\xb2\xf0\xdf\xe7\x0e\x35\x64\xbf\xb7\xee\x1d\x56\xf3\xaa\x2e\xd1\x3e\x52\x52\x2f\x97\xc2\x85\x7d\xa1\x68\x0b\x8d\x7d\xc2\xd4\x6b\xa9\xcc\xda\xfe\x92\x27\xd5\xf0\xee\x6b\x98\x2f\xb1\x90\xaa\x3b\x61\x24\xbe\x37\x7c\x3b\xf1\x62\xed\xfa\xad\x1e\x87\xfd\xcf\x06\x4f\x64\xe1\x7f\xa3\x98\x26\x71\xd5\xff\xb9\x23\x72\x1e\xe6\x82\xd3\x00\x73\x17\xcd\x8e\x8e\x80\x6d\x24\xcf\x21\x47\x96\x43\x26\x73\x04\x2c\x79\xc5\x05\xa3\xb0\x8d\x66\xd6\xc7\xb6\x87\xbb\x7b\x88\x66\x37\x70\x06\x18\x3d\x44\xff\x0b\x00\x00\xff\xff\x72\x0d\xcb\x80\x42\x1f\x00\x00"),
		},
		"/nosync": &vfsgen۰DirInfo{
			name:    "nosync",
			modTime: time.Date(2019, 3, 5, 13, 38, 20, 257702305, time.UTC),
		},
		"/nosync/map.go": &vfsgen۰CompressedFileInfo{
			name:             "map.go",
			modTime:          time.Date(2019, 1, 3, 14, 55, 7, 233338323, time.UTC),
			uncompressedSize: 1958,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x55\x4d\x8f\xdb\x46\x0c\x3d\x5b\xbf\x82\x3d\x55\x2e\x14\xe7\x9e\x62\x0f\x05\x7a\x29\xd0\x34\x40\xdb\x5b\x90\x03\x2d\x71\xac\x81\xe7\x43\x1d\x52\xeb\x2a\x8b\xfd\xef\x05\x39\xb2\x57\xde\x24\x45\x0f\xbd\xd9\x23\x0e\xf9\xf8\xde\x23\x67\xc2\xfe\x8c\x27\x82\x94\x79\x49\x7d\xd3\xbc\x7d\x0b\xef\x71\x02\xcf\x80\xd0\xe7\xd4\xcf\xa5\x50\x12\x88\x38\xc1\xc5\xcb\x08\x18\x73\x11\xff\x99\x86\x37\x7d\x4e\x2c\x98\xe4\x8d\xf8\x48\x10\x32\x0e\xdc\x01\x4b\x2e\xc4\x1d\x60\x1a\x60\xa0\x40\x42\x7c\xd0\x9c\xbf\x88\xa6\x64\x74\x04\x2e\x17\x88\x73\x10\x3f\x05\x82\x53\x2e\x79\x16\x9f\x88\x41\x32\xf4\x18\x02\xa0\x02\xf8\x9e\x21\x92\x8c\x79\xe0\x0d\x8a\xb0\x68\x2e\x4d\xf7\xe7\x48\xf0\x99\x4a\xbe\x62\x7d\xc4\xe0\x07\x2b\x4a\x71\x92\x5b\xd8\x4f\xf6\x3d\xce\x2c\x90\xb2\xc0\x91\xa0\xcf\x93\xa7\x01\xd0\x09\x15\x70\xbe\xb0\xc0\xcc\x74\x68\x64\x99\xc8\x82\x59\xca\xdc\x0b\x3c\x35\xbb\xa8\x4d\x7f\xf4\x49\xa8\x38\xec\xe9\xe9\xf9\xd3\xe6\x77\xf3\x6c\x54\xfd\x9a\x71\x80\x42\x32\x97\xc4\x20\x23\x29\x90\x99\x2a\x0b\x03\xf8\x64\x67\xca\x9d\x36\x8d\x70\xa6\xa5\x83\x5c\x20\xf9\x00\xde\x41\xca\x9a\xa3\x5e\xf1\x0c\x53\x21\xa6\x24\x87\x6b\x83\xf9\x0c\x85\x78\x0e\x02\x3e\x0d\xbe\x47\x21\x86\xcb\x48\x32\x52\x59\x2f\x5d\x90\xc1\xe5\x39\x6d\x4b\x1d\x1a\x37\xa7\x1e\xda\x08\x3f\xbc\xc7\x69\x6f\x10\xdb\x33\x2d\xb0\x41\xbf\x87\x76\xad\xfa\x72\xd6\x69\xbd\x63\xce\x61\xaf\xcd\xdb\x67\x3b\x7a\x80\x78\x88\x1f\xcf\xb4\x7c\x6a\x76\xb5\x53\xb8\x7d\x5c\x59\xf8\x43\xdb\x05\x26\xd9\x72\x70\xeb\xf8\x35\x20\x8b\x6e\x8d\x8a\x2f\x40\x58\x6d\xef\xb4\x24\x3c\x3c\x18\x4f\x4f\xcd\x6e\x67\x7f\x21\xe2\x99\xda\x7f\xd1\x64\xdf\xec\x9e\x9b\xdd\x15\x2d\x3c\xd4\xf4\x1b\xa5\x3e\x94\x8a\x74\x2b\x18\xfd\xed\x59\x7c\x3a\x6d\x50\xeb\xb1\x11\xe6\xee\x24\xf9\xa0\xc4\x5f\x3c\x53\x07\x5e\x56\xa3\x9b\xe5\xb6\xe9\x4e\xfe\x91\x56\x82\x6e\x3a\xea\x68\xd0\x70\xd3\x92\x41\x8a\x76\xed\x36\x64\xa9\x90\x35\xac\x03\x87\x81\xed\x73\x75\xd1\xd7\xf4\x5c\x1b\xf9\x26\x89\x2d\xf6\x32\x63\xb8\x97\x77\x85\x71\x93\xd8\xbb\x17\x21\xe1\xdd\x8b\xcc\x3f\xea\x7f\x65\xfd\x5e\x6d\x05\x6d\x04\xff\xcf\xf2\xbc\x2a\x63\xdd\xaf\x9a\xfd\x6c\x0b\xe4\xba\x47\xfe\x8b\xb7\xea\x8d\x2f\xed\xfe\x55\x57\xd5\xc2\x86\xaa\x96\x68\xe3\x21\x76\x9a\x76\xbf\x02\xf8\x1d\xd3\x89\x6c\x2b\x31\x38\x60\xfa\x6b\xa6\x24\x1e\x43\x58\x0c\x02\x61\x3f\x9a\x53\xd4\x05\x15\xd9\x6a\x98\xbb\x79\xd4\xf5\xe7\xc0\xdd\x7c\x62\x2d\x76\x50\x2c\x39\x4b\x9e\x6a\x6b\x5e\xa8\xa0\xf8\x9c\xae\xdb\xab\x56\x1f\x32\xb1\x6d\xaf\x44\x3d\x31\x63\xf1\x61\x81\x3e\x97\x42\x3c\xe5\x34\xe8\xda\xc4\xa4\x27\x89\x3d\x8b\xd6\xe6\x84\x13\x8f\x59\x20\x57\x8b\xd9\x3a\xd5\x84\x7d\x4e\x1a\xc0\xef\x20\x65\xc3\x7d\xf1\x21\xe8\x56\x7c\xf4\xec\x85\x06\x88\x3a\x1d\x32\x62\x82\x9c\x7a\xea\xe0\x38\xcb\xbd\x4f\x8d\xf8\xb4\xe8\x65\x4d\xa8\x2b\xbd\xae\xba\x5c\x56\x99\x86\xbb\x7d\xdd\xad\x4d\x44\x5c\xa0\x90\x0b\xd4\x8b\xdd\x8f\x38\x4d\x3a\x74\x75\xdc\x50\xae\x09\x5d\xc9\xd1\x02\xa6\xec\x93\xc0\x30\x17\x8d\xd2\xfa\x2f\x52\xdc\xd3\xa3\x99\x8f\x04\x1f\xda\xdf\xf6\xf5\x81\xd2\xe0\x34\xc7\x23\x15\xed\x9f\x02\x45\x6d\x79\xbb\x8b\x49\x47\xd4\x6f\x14\xb1\xca\x36\x75\xf5\x5d\xb0\x97\xcf\xde\xb6\x4d\x26\x73\xc1\x6b\xbf\x19\x86\xd6\x81\x9e\x7e\x73\x1a\x6f\x13\xa7\xdd\x9e\x3b\x78\xd4\x69\xab\xea\xab\x23\xd5\x8a\xde\xc1\x77\xae\xd5\x6f\x16\xb8\xdb\x1d\x0b\xe1\xb9\xd9\xa9\x37\xf5\xad\xf9\x27\x00\x00\xff\xff\xe8\x19\x65\x16\xa6\x07\x00\x00"),
		},
		"/nosync/mutex.go": &vfsgen۰CompressedFileInfo{
			name:             "mutex.go",
			modTime:          time.Date(2019, 3, 5, 13, 38, 20, 257752198, time.UTC),
			uncompressedSize: 2073,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x54\xcb\x6e\xdb\x30\x10\x3c\x4b\x5f\xb1\xc9\xc9\x4e\x62\xa5\xbd\xb6\xf5\xa1\x68\x81\x22\x40\x7a\x09\x50\xe4\x4c\x53\x2b\x99\xb0\x44\x1a\x24\x55\xd5\x4d\xf2\xef\xc5\xf2\x21\xcb\x92\xec\xc4\x2d\xaa\x93\xb0\xe4\xce\xce\xec\x0c\xb8\x65\x7c\xc3\x4a\x04\xa9\xcc\x4e\xf2\x34\xbd\xbd\x85\xef\x8d\xc5\x5f\x20\x0c\x30\xc8\x9b\xba\xde\x41\xbb\x16\x7c\x4d\x05\xa9\xe4\x62\x55\x29\xbe\x11\xb2\xcc\x52\xbb\xdb\x62\xb8\x6c\xac\x6e\xb8\x85\xa7\x34\xa1\x53\xcc\x61\xa5\x54\x95\xbe\x38\xb8\x7b\xc5\x37\x40\x65\x03\x75\x06\x77\xd6\x23\xeb\x46\x2e\xac\xa8\x11\x50\x6b\xa5\x41\x14\x50\xbb\x83\x4a\x23\xcb\x77\xe0\x61\xb2\xb4\x68\x24\x87\x59\x0d\x57\x6e\xce\xdc\x81\xcd\xe6\x34\x88\x3a\xb2\x30\xed\x29\x4d\x92\x2d\x93\x82\xcf\x2e\xbd\x8e\x0f\x50\x77\x22\x0e\x10\x2f\xe7\x69\xf2\x92\x26\x5d\xe7\x12\xac\x6e\x30\x30\xfd\x21\xa9\x0a\x8d\x7c\x2b\x5b\xa9\xec\x51\xa6\x1e\xac\xe3\x7a\x71\x8a\xac\x9f\x08\xaa\x08\x7f\x98\x7b\xfe\x63\xb6\x05\xab\x4c\xa4\xfb\xf0\x78\x96\x53\xf1\xfa\xde\xab\x56\x0b\x8b\xf7\x1e\x9a\x3e\x67\x5a\x42\xeb\xa2\xe2\x17\xd5\x48\x8b\x1a\x84\xb4\x13\x4e\x42\xa1\x34\x10\x00\x0d\x38\xb1\x27\xdd\x8e\x4d\x70\xbd\x54\x10\xb2\x84\x1e\x4c\xd8\xa1\x6e\xe1\x2a\x90\x1d\x18\xae\xdb\x6c\xc8\xee\x62\x09\xef\xe0\xf9\x99\x8e\xfa\x72\xce\x4e\xc4\xa0\xff\x54\x2e\x74\x7b\x9e\xf8\x7d\x4a\x0e\xfa\xa6\xd4\x0e\x43\xf3\xba\xaa\x57\xa2\x33\x92\x75\x10\xa0\x91\xa1\xc1\x94\xff\x69\xe8\xc3\xd0\xd1\x7f\xb5\x6d\x90\x88\xeb\xeb\xa8\xae\xb3\x2d\x57\x48\x5a\x8c\x90\x65\x85\x41\x35\x67\x55\xf5\x11\x84\x05\x77\x48\x16\xb1\xa2\x40\x6e\x41\xd9\x35\x6a\x30\xa2\x6e\x2a\xcb\x24\xaa\xc6\x38\x65\xa8\xcd\xd9\x4e\xc7\x6d\x4e\xae\x61\x60\xf5\x44\xb4\x97\x14\xed\xbf\xb2\x7c\x80\xb4\x58\x84\x95\x3c\x32\x61\xbf\x69\xd5\x6c\xdf\xfa\x66\xec\x1b\xf6\xaf\x06\x1f\xbd\x0b\x9f\xf3\x1c\x58\x9e\x1b\xc8\xb1\xb2\xec\x26\x20\xd6\x6c\x07\x2b\x04\x89\x25\xb3\xe2\x27\xde\x80\x55\x60\xd7\x7d\xcc\xbb\xc2\x15\x22\x60\xe9\x9c\xe8\xae\x13\xaa\x53\x6e\xe2\x02\xdb\x12\xae\xba\xee\x39\x5d\x98\xb9\x89\x44\xc5\xed\xb1\x2d\xb3\x08\x76\xbd\xf4\x6c\xdc\x72\x7b\xf5\x4f\x87\x3b\xf5\x1b\x8d\x43\x7b\xdc\xc2\x7d\xbf\x53\x2f\xf3\xab\x92\x08\x39\x72\x8d\x35\x4a\x6b\x06\x62\x42\xc3\x11\xae\xd4\x3b\x8b\x1c\x89\xf8\xe2\xfd\xbc\x67\x4a\x10\x4a\x49\x9a\x44\x8d\xe1\xfa\x8d\x5a\x1d\x99\x40\xbf\x5d\x9a\x7a\x82\x2f\x96\x53\x8a\xc7\x13\x22\x7c\x54\xfc\x27\x00\x00\xff\xff\xec\x95\x29\x83\x19\x08\x00\x00"),
		},
		"/nosync/once.go": &vfsgen۰CompressedFileInfo{
			name:             "once.go",
			modTime:          time.Date(2019, 1, 3, 14, 55, 7, 233609287, time.UTC),
			uncompressedSize: 1072,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x53\xcb\x92\xda\x40\x0c\x3c\xdb\x5f\xd1\xb5\x27\x9c\xa2\xe0\xbe\xa9\x1c\x52\xc5\x65\x4f\x39\xe4\x0b\xc4\x58\x03\xca\x0e\x1a\x32\x0f\x58\x67\x8b\x7f\x4f\x69\x6c\x08\xb9\xd9\x92\xba\xd5\x6a\x69\xce\xe4\xde\xe9\xc0\xd0\x98\x27\x75\x7d\xbf\xdd\xe2\x87\x3a\x86\x64\x90\x22\xee\x7f\xb1\x2b\x28\x47\x2a\xb8\x4a\x08\x38\x73\xf2\x31\x9d\xc0\x1f\xe4\x4a\x98\x10\x95\x41\xae\x48\xd4\x4d\x5f\xa6\x33\xcf\xe0\x5c\x52\x75\x05\x9f\x7d\x37\x46\xd1\x03\xf6\x31\x06\xfb\x56\xc6\xfc\x7d\x6b\x8d\x76\x11\x8e\x42\xc8\x28\x47\x86\xaf\xda\x78\xe0\x21\x1e\xa4\x23\xa2\x86\xc9\xbe\x77\xd1\xd4\xec\xd9\x98\xac\x9e\x47\xf8\x98\x0c\x64\x24\x5e\x52\x2e\x28\x72\xe2\x25\x2a\x19\xa2\xb9\x90\x09\x89\xbe\x09\xda\xe0\x4d\x11\xcb\x91\x13\xae\x31\x8d\x79\x8d\x83\x5c\x58\x0d\xde\x5d\x28\x21\x5a\xad\x15\x5a\x44\x7c\xfb\xdf\xec\xe2\xca\x0f\xd6\x79\xe9\x79\xaa\xa1\xc8\x39\x70\xeb\x95\xd7\xb3\xbc\xa6\xbc\x29\xb0\xaa\xd9\x23\xd1\x4b\x7c\x67\xf8\xb5\xb1\xf1\x85\xd5\x28\x3d\x8e\x94\x41\x18\xc5\x7b\x4e\xac\x05\x17\x0a\x95\x21\x0a\x26\x77\x6c\x20\x47\xcd\x48\xe0\x3b\x94\xaf\xcf\x53\x3c\xaf\x25\xf1\xef\x2a\x69\x31\xa1\x61\x1f\xd6\x95\x08\xfe\x60\x57\x0b\x6f\xfa\xed\x76\xb1\xb8\xf9\x51\x58\xc7\x05\x22\x2a\x45\x28\xc8\x1f\x9a\x31\xb6\xdb\x53\xcd\x05\x7b\x46\xaa\xfa\xb4\x5a\x33\x0e\x3f\xc5\xfa\x36\x05\x92\xa1\x12\x68\x14\xb7\x86\x14\x9c\x68\x32\x8c\xb2\xe3\x9c\x29\x4d\xd6\xbe\x66\x06\xfd\x13\x14\xa4\x70\xa2\x60\x19\x47\xe7\x52\x13\xdf\xd7\x46\xe9\x50\x4f\xac\x25\x5b\x8e\xfe\x1b\x61\xcf\x8b\x85\x23\xf6\x13\x76\xf1\xb5\xed\xc9\x45\xf5\x72\xd8\x3c\x56\x53\xd5\xad\x06\x7c\x62\x89\xdb\x54\x2b\x2f\x81\x95\x4e\x3c\xe0\x36\x2c\x06\xbc\x99\xf5\x8e\x6a\xe6\x6c\x66\xcc\xf4\xf3\x46\xdb\x10\xf3\x55\x93\x8a\xdb\x3c\x23\x5a\x24\xaf\xdb\x89\x46\xcd\x32\x72\xca\x56\x5e\x22\x8e\x74\x61\x24\x2e\x35\x29\x8f\x5f\xe1\x6b\x1b\x6b\x3e\xe4\xd8\xae\x75\x4e\x1a\xd7\x55\xca\x31\xd6\xf9\x38\xec\x7c\x7d\x6b\x62\xda\xb1\x8a\xf8\x62\x2b\x1d\x60\xd3\x60\x9e\x67\xb0\x37\x63\x07\xb8\x69\x8f\xe5\xb3\xef\xba\x85\xac\xbb\x3d\x12\x46\x64\x99\xa6\x71\xf5\x32\xbf\xdc\xd7\xfb\x6b\xe2\xb1\x75\x15\x85\x7f\x19\x1a\xec\x8e\xf9\x86\x92\x2a\xf7\xdd\xc8\x9e\x13\xee\x06\xf6\xdd\x53\x81\xa7\x90\x79\x89\x28\x3f\x10\xb7\xd5\xd0\x77\x7e\x35\xf4\xb7\xfe\x6f\x00\x00\x00\xff\xff\xf9\x72\xbe\xa9\x30\x04\x00\x00"),
		},
		"/nosync/pool.go": &vfsgen۰CompressedFileInfo{
			name:             "pool.go",
			modTime:          time.Date(2019, 1, 3, 14, 55, 7, 233714234, time.UTC),
			uncompressedSize: 2130,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x55\x3f\x93\xdb\xc6\x0f\xad\x4f\x9f\x02\xbf\xea\x77\xca\xe8\x74\x49\xeb\x99\x2b\x32\x29\x1c\x37\x89\x8b\x74\x1e\x17\x10\x09\x8a\x88\x97\x0b\x06\xc0\x4a\xa2\x3d\xf7\xdd\x33\x58\xfe\x39\x39\xee\x44\xee\xf2\xe1\xe1\xbd\x07\x68\xc4\xe6\x0b\x9e\x09\xb2\xd8\x94\x9b\xdd\xee\xf9\x19\x7e\x85\x8f\x22\x09\xd8\x00\xc1\xc8\x41\x3a\x70\x1a\x46\x51\xd4\x09\xe4\xf4\x37\x35\x6e\xe0\x3d\x3a\x0c\x38\xc1\x89\x80\x73\xcb\x17\x6e\x0b\xa6\x34\x81\xe1\x85\x5a\xc0\xdc\x06\x94\x92\x2b\xd3\x85\xda\xe3\xee\xf9\xb9\x62\xe7\x09\xd8\x69\x00\x73\x51\x6a\x81\x33\x78\x4f\x73\xc1\x05\x4d\x69\x90\x0a\x51\x5c\x06\x74\x6e\x2a\x2c\x3a\x60\x9e\xc0\x79\x20\xb8\xb2\xf7\x52\x3c\xf0\xb2\x38\x77\xdc\xa0\xb3\xe4\x23\x7c\xe8\xde\xd0\x7a\x49\xad\xd5\x47\xc9\x69\x02\xa5\x8e\x94\x72\x43\x70\xed\x29\x8a\xb2\x41\x8f\xe3\x48\xd9\x0e\x71\x2b\xc0\x2a\xb1\x81\xcf\xbd\x07\x8f\x96\x30\x25\x69\xd0\xef\xd8\x6f\xca\x18\x76\x04\x9d\x28\x14\x23\x38\x4d\x30\x94\xe4\x3c\x26\x82\xb3\xa8\x14\xe7\x4c\x06\xc6\xf1\x16\x33\x49\xb1\x34\xad\x18\x81\xf0\x7f\x83\xb1\xe8\x28\x46\x81\xe5\x02\x0d\x36\x3d\xc1\x56\x0f\x4e\xc5\xa1\xe4\x62\xa1\x90\xd3\x60\xb5\x54\x42\x27\x05\xa5\x62\x74\x98\xc5\x4d\x4c\x17\xce\x67\x18\x95\xcc\x8a\x46\xab\xb5\xe3\x33\xea\x29\x4c\x6d\x24\x25\x6a\x5c\xf4\x08\x7f\x85\x5f\x6c\x07\xe0\xb0\xed\x0b\x59\xfc\x20\xb4\x09\x5c\x02\xec\x54\x38\xb5\x40\x5d\xc7\x0d\x53\xf6\xd0\x44\x09\xdb\xa7\xb9\x51\x25\x82\xc4\xe6\x76\x84\xdf\xe5\x4a\x17\xd2\x0a\xc4\x16\x06\x80\x15\x76\x3c\xa5\x59\x10\x4c\x29\xf0\xee\x3e\xd9\xac\x07\x1c\x47\x95\x51\x19\x9d\xaa\x70\xd2\x01\x6e\x92\xba\xc0\x80\x39\x68\x23\x9c\x55\xca\xf8\x7d\xf0\xaa\x0e\x81\x63\x9c\x28\x7b\x24\xad\xc7\x88\x10\x0e\x92\xcf\x11\x38\x18\xc5\x29\x3b\xd7\xbc\x54\x99\xda\xb0\xa6\x91\xdc\x14\x55\xca\x1e\x41\xa5\x91\x72\x4b\xb9\x86\xa7\x49\xd1\xaa\xcd\x34\x96\x41\x38\xce\x7c\x46\x95\x0b\xb7\x14\x23\x70\xc5\xd0\x28\xca\xa8\xf3\xd7\xcd\x25\x96\x0c\x72\x21\xed\x09\x6b\xd4\xb1\x51\x31\x8b\x16\xa6\x15\xf8\xae\x73\xba\xe1\x10\xf1\x90\x0e\xce\x22\xed\x8f\xdd\x2f\x83\xd0\x0d\xbe\x32\x39\xc0\xb5\xe7\xa6\x87\x01\x39\x3b\x72\x36\xc0\x00\x6b\xa7\x8c\xc3\x3c\x14\x4f\xc6\x5f\xa9\x9d\x47\xe9\x3f\x53\x5a\x7c\x2c\x0e\xa7\xd2\x75\xa4\x16\xee\xd3\x72\xcd\x1a\x4c\x64\x50\x72\x4b\x1a\x70\x49\xb0\x85\xc7\x3a\x13\x95\xfa\x5d\x7e\x51\x09\xb0\x71\xbe\x50\x9a\x60\x54\xce\xce\xf9\xbc\xaf\x4a\x5b\xaf\x9c\xbf\x58\x9d\xa5\x40\xf9\xa7\x30\x59\x43\xd9\xd7\x96\xff\x9c\xdb\x11\xef\x49\xa1\xc7\xdc\x1e\x00\xdf\x32\xb1\xf5\x14\xf6\x19\x8c\xa8\x3e\xab\x61\xbd\xa8\x3f\x25\x8e\xf9\x9f\x37\x0d\xb0\x2d\x73\x1e\xc7\x6b\xd0\x42\xbe\x1a\xb6\xaa\xdf\x01\x8c\x63\xb2\x6b\xc5\xc5\x12\x68\x85\xe6\x74\x6e\xc6\x5d\x29\x25\xe0\xca\xb7\x6e\xaf\x20\x8c\xca\x72\x84\x0f\x35\xca\x43\xe8\xb3\x4d\x40\x78\xde\xe3\x85\xc0\x4a\xd3\x6f\x6b\x8f\xc3\xc5\xa1\x1e\xf7\xc4\x0a\x72\xcd\xdf\xa5\xbd\xf6\xef\xd3\xb8\x2c\x21\x73\x2d\x8d\xc3\xb7\xdd\xc3\xac\xfe\xa7\xcf\x9c\x9d\xb4\xc3\x86\xbe\xbd\xee\x1e\xfe\xa0\x2b\x00\x74\x25\x37\x8f\x7b\xb8\x3f\x79\xad\x8b\xf8\x3d\x39\x18\xa5\x5a\x18\x33\xa0\x9e\xd8\xb7\x59\x80\x4e\x65\xd8\xd6\xdd\x61\x59\x9b\x75\xac\xd7\x93\x75\xdd\x1c\xaa\x67\x4a\x5e\x34\xd7\x0b\x2e\xf5\xc3\x08\x11\xe9\x71\x2d\x15\xfb\xb7\xe9\x25\xb6\x92\x0b\xf0\x39\x07\xe3\xb8\x37\x46\x2b\x01\xe1\x4a\xb1\x45\x3c\x4c\xa3\x61\xf4\xba\xd4\xe0\xb7\x0a\x63\x61\x5e\x49\xed\xac\xb9\x59\x19\xa8\x6e\x6c\xa5\x34\x0f\xcb\x89\xfc\x4a\x94\xe1\x82\xa9\x50\x98\x6e\x31\xa0\x2e\xf0\xb1\xf8\xfa\x7f\x11\xd5\x96\xf3\x99\xee\x3c\xc2\xef\x69\x0b\xd6\x87\xae\x72\xbd\xd6\x52\x35\x5e\x57\x36\x5a\x6e\x43\xe6\x99\xe8\x78\x0c\x69\xeb\x7a\xca\x4f\x99\xd3\xa1\x7e\xb4\x28\xb0\x16\x52\xb2\x92\x6a\xf0\x42\x88\xba\x47\xe3\xb3\xe3\x2e\x0c\x81\xc7\x11\x7e\x0a\xf1\xf6\xf1\xe9\xf7\xf6\x84\x9f\xdc\x41\xa2\xfc\x38\x1e\xab\xb1\x7b\x78\x79\x81\x9f\xe3\x7d\x1c\xcc\xd5\xff\xf7\x52\xe9\xc4\xbb\x87\x85\x5e\x3d\x78\xdc\xef\x1e\x1e\x5e\x77\xdb\xcb\xcc\x69\x17\xcf\x37\x78\xf7\x02\x0b\xde\xa7\x7b\xec\xa7\x5f\x3e\xef\x1e\x96\x07\x78\xbb\xf2\xee\x87\x3b\x0b\xe0\x6d\x89\x4f\xd5\xb5\x6d\x0d\x6e\xab\xe1\x61\xe4\x0f\xed\x7d\x2c\xfe\x78\xbb\x6f\x6f\xbf\xf4\x77\x8b\xa6\xd6\x16\x66\xec\x4a\xf4\x8d\x4a\xfd\xff\x6c\x57\x12\x07\xb8\xed\x77\xaf\xbb\x7f\x03\x00\x00\xff\xff\x07\xba\x3e\x57\x52\x08\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/js"].(os.FileInfo),
		fs["/nosync"].(os.FileInfo),
	}
	fs["/js"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/js/js.go"].(os.FileInfo),
	}
	fs["/nosync"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/nosync/map.go"].(os.FileInfo),
		fs["/nosync/mutex.go"].(os.FileInfo),
		fs["/nosync/once.go"].(os.FileInfo),
		fs["/nosync/pool.go"].(os.FileInfo),
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
