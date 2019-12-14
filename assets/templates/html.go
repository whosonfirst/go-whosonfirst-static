// Code generated by go-bindata.
// sources:
// templates/html/id.html
// DO NOT EDIT!

package templates

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesHtmlIdHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x5a\xef\x92\xdb\xb6\x11\xff\x7c\xf7\x14\x6b\x3a\x53\x27\x33\xa6\x18\xc7\x69\xa7\x3d\x53\xec\x9f\xb3\x3d\xbd\x69\xd2\x78\x62\xbb\x69\x3f\x65\x20\x62\x45\xc2\x06\x01\x06\x00\xa5\x53\xdd\x7b\xad\x3e\x40\x9f\xac\x03\x80\x7f\x40\x8a\xba\xd3\xc9\xae\xeb\x19\x9f\x48\x70\xf7\xb7\xbb\x3f\x2e\xb0\x0b\x48\xe9\x83\xe7\x3f\x5c\xbe\xf9\xc7\xab\x17\x50\x9a\x8a\x67\xe7\xa9\xfd\x00\x4e\x44\xb1\x8c\x50\x44\xd9\x39\x40\x5a\x22\xa1\xd9\xf9\x19\x40\x6a\x98\xe1\x98\x7d\xf8\x00\x8b\xd7\xaf\x7e\x5c\xfc\x95\x54\x08\x37\x37\xf0\x2f\xe8\x46\xae\xa8\xbf\xff\xa9\x94\x8f\x34\xfc\x20\xe0\x25\x53\xda\xa4\x89\xd7\x73\x10\x15\x1a\x02\xa5\x31\x75\x8c\xbf\x34\x6c\xb3\x8c\x2e\xa5\x30\x28\x4c\xfc\x66\x57\x63\x04\xb9\xbf\x5b\x46\x06\xaf\x4d\x62\x9d\x79\x06\x79\x49\x94\x46\xb3\x7c\xfb\xe6\x65\xfc\xdb\x28\x80\x11\xa4\xc2\x65\xa4\x70\x8d\x4a\xa1\x0a\x94\xa5\x62\x05\x13\xd1\x01\x8b\x7f\x8f\xdf\xfe\x31\xbe\x94\x55\x4d\x0c\x5b\xf1\xd0\xe8\xd5\x8b\xe5\xef\x22\x48\xf6\x4c\x90\xba\xe6\x18\x57\x72\xc5\x38\xc6\x5b\x5c\xc5\xa4\xae\xe3\x9c\xd4\x64\xac\xbe\x43\x7d\xb4\xb6\x36\xc4\x34\x3a\x5e\x11\x15\x6b\xb3\x1b\xc1\xac\x38\xc9\xdf\xcf\x01\xfd\x99\x08\x5a\x22\xa7\x2f\x15\x43\x41\xf9\x2e\xa4\x4b\x35\x38\xa7\xb2\x61\xb8\xad\xa5\x32\x81\xe8\x96\x51\x53\x2e\x29\x6e\x58\x8e\xb1\xbb\x79\x0c\x4c\x30\xc3\x08\x8f\x75\x4e\x38\x2e\x9f\x3c\x86\x8a\x5c\xb3\xaa\xa9\x82\x01\x26\xc6\x03\x8d\x46\xe5\xee\x2c\x09\x4b\x21\x9d\xf5\xc1\x7c\xad\x64\x8d\xca\xec\x96\x91\x7d\x89\x17\x9c\x68\x53\x49\xca\xd6\x0c\x69\xe0\x8b\x4d\x9c\xef\x88\x36\xdf\xb7\x8f\xe0\xe6\xa6\x8b\x62\x0e\x4a\x16\x17\x66\x9c\x26\x44\x19\x96\xf3\x49\xe8\x23\x05\xcd\x0c\xfe\x6c\xc9\x08\xb4\xc6\xf9\x79\x8b\xb2\x4b\xdc\x89\xbf\x61\xea\x33\x0d\xa4\xcf\xfe\x57\x9c\xe4\x68\xfd\x73\x4f\x44\x3f\x7e\x29\x1b\x61\xd4\x6e\x08\x6d\xc6\x10\x45\x9d\x2b\x56\x1b\x26\x45\x98\x8f\x46\x4f\xe6\x12\x5c\x3d\xb7\x46\x47\x13\xee\x16\x58\x56\x91\x22\xf4\x3f\x61\x34\x19\xe9\x2e\x6a\x51\xfc\x5e\xb3\x7f\xe2\xf2\x5a\x57\x07\x5f\xa1\xd9\x32\x63\x50\x5d\xe4\x44\xd1\x08\x36\x84\x37\xb8\x8c\x74\x53\x55\x44\xed\x0e\x59\xef\x74\x2c\xff\x81\x07\x7f\x20\x9c\xcb\xb5\x29\xb1\xb6\x6c\xe9\xbb\xb4\x73\x85\xc4\x48\x75\x3a\x40\xa3\xf8\x6d\xf1\xdf\xa5\xfe\x79\x12\xa0\xb3\xf6\xbf\xc8\x82\x0e\xfb\xb4\x54\x48\xbd\x43\xa0\x55\xbe\x8c\x92\x77\x64\x43\xfc\x40\xc2\x65\x4e\xf8\x5a\x2a\x52\xe0\xa2\x62\x62\xf1\x4e\x47\x59\x9a\xf8\x87\x99\x9f\xbe\x07\x75\x35\x67\x75\xbd\xab\x48\xbd\xc8\x95\xd4\xba\x24\x4c\xe9\x31\xc0\xad\xa6\xb7\xa5\xd4\x52\xac\x2d\x11\x8b\xed\x76\x3b\x35\x7d\xac\xaa\x42\x41\x51\x4d\xec\x1e\xa9\xdb\xd2\xcb\x70\xea\xf7\xf1\xe6\x73\x92\x97\x78\xaa\xef\x8d\x62\xa7\x39\x2e\xd0\x9c\xa8\x48\x2a\xb6\xde\x8d\x75\xe1\x0c\xce\xe0\xe8\x57\x55\xa0\x7c\xa7\xe5\x5e\xa6\x1c\xab\xce\x91\xac\x39\x9a\x45\x63\x18\xdf\x67\xfd\x5e\x10\xae\xe0\x7e\x24\x46\x49\x04\xe5\x38\xcd\xdb\x9e\x91\x73\x38\x83\x94\x33\xf1\x1e\x14\xf2\x65\xe4\x2d\x96\x88\x26\x82\x52\xe1\x7a\x19\x25\xb9\xd6\x7b\x99\x9c\x6b\x6d\x5b\xae\xfb\xaa\xe6\xb2\xaa\xa4\xf0\xda\x8e\xd0\xf3\xfb\x22\xd8\x6e\x84\xe5\x0b\x46\x5b\x17\x8e\x24\x63\x50\x3b\x38\x7b\xed\x72\xd8\x76\x73\x03\x88\xef\xcc\x60\xcb\x04\x95\xdb\x05\xa1\xf4\xc5\x06\x85\xf9\x8e\x69\x83\x02\xd5\x97\x11\x97\x84\x46\x8f\x61\xdd\x88\xdc\xae\x85\x60\xef\xbf\x44\x2b\xf3\xd5\x07\xa7\x79\x36\xef\x85\x6d\x63\xbe\xfc\xea\x99\x13\xb9\x71\x9f\x83\x4f\x00\x69\xe2\x3b\x59\x80\x74\x25\xe9\x0e\x28\x31\x24\x0e\x80\xe2\x46\xb1\x18\x05\xad\x25\xeb\x96\xfa\xe7\xc4\x90\x17\xed\x88\x5d\x63\xad\x32\x80\xfb\x93\x52\xb6\x81\x9c\x13\xad\x97\x51\x00\x12\x01\xa3\xa3\x81\xd8\xd5\xa9\x68\xdf\x9a\x95\x9b\xac\xe0\x7b\x32\x35\x31\xe5\x20\xf5\x8a\x98\xf2\x90\x9c\xb2\x3d\x74\x08\xf9\xca\x0d\x85\xc0\x9c\x18\x66\x1a\x8a\x83\xcc\x77\xed\x48\x20\x23\x45\x31\x15\xea\x86\x06\xa9\x8a\x89\x19\xb4\xef\x99\xe8\x00\xc7\x92\xfb\x98\x56\x74\x06\x96\x5c\xcf\xc1\x92\xeb\x7d\x58\x2b\x39\x03\x4b\xae\x47\xb0\xd9\xb9\x7b\x55\xed\xcb\xb2\xec\x54\xa4\xb6\xc9\x97\x3e\x88\x63\x48\x59\x55\xb4\xf9\xbd\x57\x0b\xf5\xa6\x88\xa0\x44\x56\x94\x66\x19\x3d\xfd\xf5\xd7\xb6\x1e\x42\x1c\x67\x1d\x5e\x42\xd9\xa6\xbb\xe9\xc6\xca\x6f\x2c\xb2\xae\x89\x98\xc9\x8b\xd8\xb5\xa0\x2e\xef\x27\x3d\x44\x9a\x58\x95\x0c\xbc\xe6\x6c\xf6\xc4\x94\x18\xf4\x33\xcc\x89\xba\xae\xe3\xa0\xa5\xba\x6b\x43\xa2\x6c\xae\x33\xe9\x41\xc4\x18\x62\x90\xee\xfb\x15\x57\xc7\x5b\xf9\xc5\xb9\x0d\xae\x22\x9c\x77\x41\xb0\x35\xe0\x2f\x2d\x69\xfa\xb2\x51\x36\xdf\x16\xaf\x8d\x62\xa2\x78\xc9\x49\x01\xd1\x93\x08\x6e\x6e\x5c\x4f\x72\xc8\xd5\x76\x17\x94\x7b\xe5\x28\xfb\x09\x61\x85\x9c\xe1\x06\xc1\x94\x4c\x83\xc2\x5c\x2a\x0a\x46\xc2\x0a\xa1\x95\x5a\xb4\x0e\xb5\x5e\x20\xd7\x78\xb7\x2b\x5f\x1f\xeb\x8a\x90\x66\xe4\x8e\x90\x60\x13\x0d\xd5\x69\x8e\x79\xa3\xf6\x4e\xd8\xbc\xea\x08\x75\x34\x9e\xf5\xc9\x64\x33\xa7\xbd\x39\xeb\x53\x35\x74\x8e\xa2\x21\x8c\xeb\x68\x96\xfa\xe7\x58\x2b\xcc\x89\x41\x7a\x1a\xfb\xb4\xd7\x8f\xb2\x37\x41\x70\x25\xd1\xb0\x42\x14\x50\x11\xf5\x1e\x29\x10\x0d\x83\xe8\x34\x58\x17\xde\xf9\xd9\x9c\x7f\xaf\x9b\x1a\x95\x46\x7a\xc8\x3f\xfb\x7f\xd6\xae\xee\x15\x61\xb5\x03\x53\x22\xac\x25\xe7\x72\xcb\x44\xd1\xca\xea\x0b\xb7\xd6\x43\xda\xf0\xd9\x99\x60\xd7\x6d\x54\x2d\x6f\xf6\xdf\x87\x0f\xa0\x88\x28\x10\xbe\x60\x8f\xe1\x0b\x0d\x17\x4b\xef\xe5\xe0\xe3\x9f\x76\xad\x5b\x0e\x97\xb3\x2c\x25\x5d\xd5\xf4\xab\xc4\x17\xda\xad\x58\x33\xe6\x18\x85\xc9\x8c\x67\xeb\xdd\xe1\x95\xdf\x03\x65\xdd\x55\x9a\x90\x2c\x4d\x38\x0b\x7d\xed\x68\xf5\x69\xd2\xf0\xf6\xd9\xc2\x7f\x8c\xb9\xb7\xe4\x1f\xa2\x9e\x89\x62\x8e\x7b\x0f\x13\x72\xdf\x53\xae\x6f\xe1\xfb\x28\xb6\xef\x64\x5a\x77\xe9\xf9\xff\xe2\x78\xc4\x6f\xcf\xee\xc2\xa7\x64\xf0\xb0\x5d\xec\xfb\xe9\xda\xf0\xbd\xf9\xb9\x22\x9a\xe5\x2e\x70\x1b\xcd\xc1\x0d\xdb\xc1\x99\xc8\xe8\xb0\xfa\xba\x02\xd4\x2d\xd2\x44\x50\x60\xc6\xbe\x1e\x5b\x18\x37\x08\x6f\x7f\xbc\x72\x48\x1d\x61\xa5\x31\xb5\xbe\x48\x12\x4b\xc0\x22\xec\x8a\xa4\x2a\x92\xbd\xc6\x61\xc6\xb4\x42\x6e\xbb\x8c\xa0\x56\x78\x61\x4b\xd5\xc2\x73\x35\xcd\xab\xbe\xb5\x88\xe2\x6f\x7d\x22\xd9\xa8\x83\x34\x7a\xa4\xc1\x77\x24\x5d\xdc\xb9\xa4\x98\xc5\xdf\xa6\x89\xbb\x80\x6d\xc9\xf2\x12\x98\xa0\xcc\xae\x25\x1a\x98\x71\x93\xbe\x0f\xea\x61\xc9\x50\x11\x95\x97\xbb\x28\xab\x1a\x6e\x58\xcd\xb1\x45\xd4\xd6\x2f\x58\x61\x4e\x1a\x8d\xfb\x34\x14\xcc\x94\xcd\xca\x76\xc4\x61\xa3\x9a\x8c\xd2\xb4\xdf\xbb\x25\x46\x21\x26\x15\xd1\x06\x55\x12\x0c\x6f\xe5\xfa\xa1\xb7\xf6\xb3\x7d\x31\xce\xbb\xba\x46\xa1\x41\xcb\x0a\x0d\xab\x50\x07\xec\x9c\xcd\x14\xa1\x80\xa1\xa7\xf7\x61\xe8\xe9\x98\xa1\x0a\x89\x70\xec\x54\xb6\x11\x01\x8e\x05\x33\xac\x22\x06\xf9\x0e\x4a\xb2\xc1\x13\x18\xab\x51\xda\x47\x9f\x85\x38\xa2\x8a\x06\x81\xac\x64\x63\x7c\xd5\xd4\x52\x19\x90\x6b\x7b\x23\x8a\xe3\x29\xfc\xe6\x3e\x14\x3e\x99\xa3\x70\x8b\x8e\x2f\xf1\xc8\xc0\x4a\x9a\x12\x15\x52\x30\x6a\x67\x17\x36\x23\xa1\x66\x02\xa8\xdc\x8a\xb0\xb4\x3f\x3a\x90\x8e\xac\xaa\x90\x32\x62\x3a\x76\x3f\x7f\x3a\x72\xb6\x46\x1b\x6f\x2e\xab\x9a\xbb\x09\x44\x83\x65\xed\x76\x26\x9f\xdc\xcd\xe4\xd5\x3e\x93\x96\x43\x4b\xd5\x16\x81\xa8\xcf\x14\x65\x23\x72\x54\x86\x30\xe1\x08\xf6\xef\xb2\xad\x51\x6d\x5a\xfb\xb7\x25\x05\x4e\x93\xa8\x8b\xf0\xca\x58\x9a\xbc\xb4\x6f\x21\xa6\x65\x66\x44\xcf\xa7\x29\x38\x53\xc8\x6c\xbf\x10\xec\x4b\x4d\xa7\x42\xdf\x4c\x75\xd5\x64\x25\x1b\x61\x8b\x38\xac\xe4\xf5\x68\xf9\x7f\x38\xeb\x73\x81\x76\x9d\x52\x3b\x98\x1b\x8c\x3b\xb0\x95\xbc\x1e\x56\xfd\xd1\xe6\xed\x31\x1c\xd8\xa8\x05\x0f\xc2\x5d\xd9\x78\x38\x90\x77\x2f\xaf\xab\x61\xb5\x62\x22\x67\x35\xe1\x90\xa3\x30\x4a\x32\xfa\xd1\xa1\x74\x40\x43\x1c\xc1\x96\x76\xf0\x6a\xea\x52\x57\xd7\xba\xfe\xdb\x96\xfe\x60\xa7\x68\x2b\x5e\xdf\x86\x73\x5c\x9b\x58\x97\x44\xbd\xb7\xe5\xbd\x6d\xc5\xe6\x7a\xf4\x21\x97\xfb\x40\x34\xe6\x66\xe8\x3d\xef\xd0\x8a\x6b\x85\xc6\xec\xec\x2e\xcf\xfa\xd0\x36\xc6\xc7\x28\x2a\xb2\xed\xb4\xba\x56\xd1\x5f\xb7\x9f\x41\x34\xca\x16\x92\x3e\x1c\xd7\xf8\xcc\x81\xfb\x9d\xc1\x34\x8e\x16\xbc\x7c\x0a\xee\x48\xc9\x1f\xeb\xc4\x84\xb3\x42\x5c\x58\x9a\xe0\x01\xab\x6a\xa9\x0c\x11\xe6\x99\x3f\x07\x69\x71\x32\xff\x19\xbc\xeb\x89\x01\x22\xf2\x52\xaa\x28\xfb\xcf\xbf\xfd\x52\x56\x3e\xcd\x86\x16\x17\x20\x35\x64\xd5\x7e\x81\xe8\x6e\x1f\xc4\xf1\xf9\x99\xfb\x42\x52\x65\xee\xc2\x5e\x96\xd9\xaf\x1e\x3e\xf9\xcd\xd7\xcf\xd2\xc4\x94\xe1\xe8\x6b\x67\x6a\x3a\xfa\x37\xc2\x1b\x9c\x0e\x5e\xda\xf8\x29\x8a\x3c\x78\x92\x26\x9d\x11\xbb\xfd\xef\x3c\xf0\x63\x56\xc7\xae\xc4\x7e\xff\xd7\xea\x74\x12\xf4\xf0\x96\x2b\x68\xf6\xe6\x76\xab\x36\x49\x0d\x3d\x1d\xea\x4a\xbf\x51\x0d\x7e\x02\x98\xbf\x08\x5b\x15\xa7\x38\x9e\x90\x90\x88\xee\xda\xb3\x81\x44\xdb\x92\x74\x12\x19\x4e\xf7\x93\x70\xe1\x91\x3e\x96\x8a\x0e\xe5\x58\x26\xba\x94\x18\x36\xca\x5d\x26\x1d\x6d\x73\x7e\x3b\xdf\x9b\x3e\x09\x68\x4a\x83\x05\x39\x24\x79\xdf\x50\x87\x5d\xfa\xbd\x43\x9d\x3f\x19\x38\x21\xd4\x00\x68\x2e\xd4\x13\x40\x4e\x65\xc1\x35\xb6\x27\xd2\x30\xd9\xa5\x7f\x44\x08\x16\xe9\xe3\x89\xf0\x28\x87\x99\x38\x83\xbe\x48\xa5\xc9\xb0\x44\xf7\x55\xe7\x40\x71\x19\xba\xe9\x83\xf5\xc5\xaa\x04\x4d\x77\x7f\x39\xdb\x93\x1f\x55\x44\xba\x92\xd8\xfd\xb2\x20\x38\xbd\x9f\xb5\xae\x65\xa3\x72\xb4\xa5\xcb\x5f\x84\xb5\xab\x7d\x74\x6c\xf1\xea\x8f\x6d\x26\x67\x1d\xf7\xda\xba\x67\xf7\x91\xde\x3b\x4a\x9a\xb5\x3b\xdf\xb6\xc7\xd6\x40\x0f\xf8\x23\xd6\x12\x6e\x6e\x92\x15\x97\xab\xae\x69\x1f\x09\xec\xf9\xf7\x29\x51\x27\x71\xf4\x67\x34\x43\xaf\xd3\xf7\x38\xa3\x73\xf9\xe9\xa9\xff\x5a\x4a\x83\x6a\x38\xf8\x9f\x32\xd1\x7e\x47\xee\x9c\x9e\xfc\xbc\x20\x9b\xfc\xde\xc0\x75\xb4\xb7\x7d\x11\x10\x8c\xb8\xef\x9c\x56\x92\xee\xb2\xf3\x34\xf1\xbf\xaf\xfa\x6f\x00\x00\x00\xff\xff\x2c\x7f\xd7\xc0\x70\x25\x00\x00")

func templatesHtmlIdHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesHtmlIdHtml,
		"templates/html/id.html",
	)
}

func templatesHtmlIdHtml() (*asset, error) {
	bytes, err := templatesHtmlIdHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/html/id.html", size: 9584, mode: os.FileMode(420), modTime: time.Unix(1576359800, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"templates/html/id.html": templatesHtmlIdHtml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"templates": &bintree{nil, map[string]*bintree{
		"html": &bintree{nil, map[string]*bintree{
			"id.html": &bintree{templatesHtmlIdHtml, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

