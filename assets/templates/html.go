// Code generated by go-bindata.
// sources:
// www/templates/html/id.html
// www/templates/html/inc_common_css.html
// www/templates/html/inc_common_footer.html
// www/templates/html/inc_common_footer.html~
// www/templates/html/inc_common_header.html
// www/templates/html/inc_common_header.html~
// www/templates/html/inc_common_meta.html
// www/templates/html/index.html
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

var _wwwTemplatesHtmlIdHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x5a\xef\x6e\xdc\xb8\x11\xff\xbc\x7e\x8a\x89\x72\x68\xbe\x9c\x56\x49\x7d\xed\x07\x47\xab\x16\x4d\xee\x00\xa3\x77\x3d\x23\x49\x7b\xe8\xa7\x80\x2b\xce\x4a\xcc\x51\xa4\x4a\x52\x5e\x6f\x5d\xbf\x56\x1f\xa0\x4f\x56\x50\x12\x25\x4a\x2b\xd9\xfb\x27\x35\x50\x03\xc1\x6a\x29\xce\x6f\x66\x7e\x1c\x72\x66\xb8\xb9\xbf\x07\x8a\x1b\x26\x10\x02\x46\x03\x78\x78\xb8\x88\x5f\xbc\xff\xf9\xdd\xa7\xbf\xdf\x7c\x0f\xb9\x29\x78\x72\x11\xdb\x0f\xe0\x44\x64\xab\x00\x45\x90\x5c\x00\xc4\x39\x12\x6a\x1f\xec\x5f\x6c\x98\xe1\x98\xdc\xdf\xc3\xf2\xe3\xcd\x87\xe5\x5f\x48\x81\xf0\xf0\x00\xff\x02\x37\x72\x4d\x9b\xef\xbf\xe4\xf2\x95\x86\x9f\x05\xfc\xc0\x94\x36\x71\xd4\xc8\xb5\x28\xf7\xf7\x60\xb0\x28\x39\x31\xd6\x12\x91\x7e\x4e\x65\x51\x48\xf1\xb9\x40\x43\x6a\xb3\x5a\x65\xf6\x3b\x94\x4a\x96\xa8\xcc\x6e\x15\x58\xdb\xae\x38\xd1\xa6\x90\x94\x6d\x18\xd2\x00\x52\x29\x0c\x0a\xb3\x0a\xac\xfe\x1f\x89\x36\x3f\xb5\xaf\xe0\xe1\x21\x80\x28\xb9\x58\x00\xcc\xa0\xc9\xec\xca\xec\x4a\xf4\x30\x88\x32\x2c\xe5\x58\xcb\xcd\xca\x68\x66\xf0\xb3\x20\x85\x2f\x38\x74\xf6\x71\xf9\x9a\x88\x91\xe1\x3e\x95\x4c\x03\xe9\xd8\xbc\xe1\x24\x45\x6b\x65\xfd\x46\x74\xe3\xef\x64\x25\x8c\xda\xb5\x3e\xce\xea\xa2\xa8\x53\xc5\x4a\xc3\xa4\xf0\x34\x5e\x1b\x3d\x5a\x1e\xb8\x7e\x6f\xf5\x0e\xd6\xf0\x71\x64\x56\x90\xcc\xf7\x22\x2a\x45\x16\x0d\xe4\xff\xa0\xd9\x3f\x71\x75\xa7\x0b\x0f\x68\x06\xcf\x6c\x99\x31\xa8\xae\x52\xa2\x68\x00\xb7\x84\x57\xb8\x0a\x74\x55\x14\x44\xed\x1e\x31\xc3\x89\xd9\x15\xf1\x4c\xf9\x23\xe1\x5c\x6e\x4c\x8e\xa5\x25\x4f\x1f\x00\x90\x2a\x24\x46\xaa\xb3\x30\x2a\xc5\x7d\x3a\x18\x8d\x0e\x65\xd3\x21\x3c\x5b\x60\x38\x85\xff\xa3\xe8\x70\xf0\xe7\x85\x48\x63\x1a\x68\x95\xae\x82\xe8\x0b\xb9\x25\xcd\x40\xc4\x65\x4a\xf8\x46\x2a\x92\xe1\xb2\x60\x62\xf9\x45\x07\x49\x1c\x35\x2f\x93\xc5\xd3\xf2\x9a\xb3\xb2\xdc\x15\xa4\x5c\xa6\x4a\x6a\x9d\x13\xa6\xf4\x10\xe4\x40\x43\xb6\xb9\xd4\x52\x6c\x2c\x41\xcb\xed\x76\x3b\x36\xe4\x18\x71\x85\x82\xa2\x9a\xb4\xe2\x20\xf9\x96\x7b\x86\xe7\x7b\x92\x92\x34\xc7\x73\x7c\xa9\x14\x3b\xdd\x11\x81\xe6\x0c\x61\x52\xb0\xcd\x6e\x28\x0f\x0b\xe8\xcc\x3f\x06\x2b\x43\xf9\x45\xcb\xbd\xe8\x3a\x06\x82\x23\xd9\x70\x34\xcb\xca\x30\x3e\xbd\x2e\x47\xc1\x68\xb3\xe3\x33\xeb\x7b\x14\x4e\x4e\x04\xe5\x38\x8e\xf9\x86\xa7\x63\xf0\xd6\x4a\x6e\x35\xaa\x65\x41\x4a\x7d\x0e\x4d\x0e\x87\xd1\xd3\x7d\xf3\x30\x98\x60\x66\xc6\x9c\x27\x6a\x8f\x54\x6b\xaf\xf4\xb0\x75\x4f\xe4\x0a\x9f\x78\x2d\xe9\x0e\x28\x31\x24\xf4\xd4\x86\x95\x62\x21\x0a\x5a\x4a\xe6\x8e\xeb\xf7\xc4\x90\xef\xdb\x11\x7b\x48\x26\x1d\xa5\x94\xdd\x42\xca\x89\xd6\xab\xc0\x1e\x88\x84\x09\x54\x41\x72\x90\x61\xd6\x0a\x54\xb5\x6d\xdd\x89\x4b\x98\x70\x70\x9e\x45\x01\x30\x3a\x18\x08\xeb\xf4\x15\xec\x9b\x6e\xe7\x8d\xce\xf3\xbd\x39\x25\x31\x79\x3f\xeb\x86\x98\x7c\x6e\x9e\x42\x31\x84\xbc\xa9\x87\x7c\x60\x4e\x0c\x33\x15\xc5\x7e\xce\x8f\xed\x88\x37\x47\x8a\x6c\x3c\xc9\x0d\xf5\xb3\x0a\x26\x26\xd0\x7e\x62\xc2\x01\x0e\x67\xee\x63\xda\xa9\x13\xb0\xe4\x6e\x0a\x96\xdc\xed\xc3\xda\x99\x13\xb0\xe4\x6e\x00\x3b\x5c\x7b\xcb\x4e\x41\xca\xc0\xad\x9a\x92\xdb\x00\xea\x3d\xbd\x0a\xd6\x52\x51\x5b\xc6\x48\xce\x28\x28\xa4\x6f\x83\xe4\x62\x11\xbf\x08\x43\x88\x59\x91\xb5\xa1\x3f\x2e\x28\x96\xfa\x36\x0b\x20\x47\x96\xe5\x66\x15\x5c\xfe\xee\xb5\xcd\xa1\x10\x86\xdd\xc6\x89\x28\xbb\x9d\x8c\xbf\x27\x55\xdb\x92\xff\xb2\xfe\x6c\x76\x60\x49\xa6\x82\x2d\xac\x6b\xe0\xe4\x62\xb1\x00\x18\x95\x2a\x71\x64\x65\x92\x56\x74\x32\x26\x43\x4a\x0c\x36\xdb\xb4\x9e\x5a\x17\x37\xb3\xaa\x4a\x57\xed\x04\xc9\x54\x01\xd4\x81\x88\x21\x44\x3f\xbb\x2b\x8b\x6c\x85\x60\x2d\x6e\x45\x96\xf5\x97\x66\x44\x17\x84\xf3\xc6\x9f\x76\x3b\xb2\x0d\xe0\x3f\x5a\xce\xf5\xbb\x4a\xd9\xa0\x5e\x7e\x34\x8a\x89\xec\x07\x4e\x32\x08\xde\xd4\xbb\x72\xf1\x14\x53\xda\x10\x53\xe9\x30\x6d\x10\x82\xe4\x17\x84\x35\x72\x86\xb7\x08\x26\x67\x1a\x14\xa6\x52\x51\x30\x12\xd6\x08\xed\xac\x65\x6b\xa1\x6f\x0f\x72\x8d\x4f\x1b\xf5\xfa\x28\xa3\x84\x34\x03\xc3\x84\x04\x1b\xdc\xa8\xce\x30\xd1\x53\x6f\x87\x84\x8d\x57\x8f\xf8\x9a\xe7\x45\x1d\x67\x91\x0d\xb4\x36\x46\x17\xde\x56\xf1\x0d\xa5\x68\x08\xe3\x3a\xe8\x22\x72\xbc\x32\xef\xb1\x54\x98\x12\x83\x74\x6a\x71\x0e\xa3\x81\x76\x18\x41\xf2\xc9\x73\x38\x27\x1a\xd6\x88\x02\x0a\xa2\x7e\x45\x0a\x44\x43\x3f\xb5\x27\x60\xe4\xac\x1b\x98\xb3\xf8\x63\x55\xa2\xd2\x48\x1f\xb3\xd8\x7d\x4e\x5a\xa3\x3b\x00\x58\xef\xc0\xe4\x08\x1b\xc9\xb9\xdc\x32\x91\xb5\x73\xf5\x55\xe7\x7a\xc5\x27\xf7\x94\x4d\x52\xa8\x74\xbf\x85\x15\x11\x19\xc2\x37\xec\x5b\xf8\x46\xc3\xd5\xaa\x31\xb6\x37\xf5\x4f\x3b\xb7\xb0\x31\x67\x49\x4c\x20\x57\xb8\xe9\x4e\xa6\x6f\x74\x7d\x96\x4e\x28\x62\x14\x46\xc7\x06\xdb\xec\xe6\x73\x52\x03\x94\xb8\xa7\x38\x22\x49\x1c\x71\xd6\x59\x39\x64\x38\x8e\x2a\xde\xf1\xbf\x9c\xe0\xfd\xb8\x05\x61\x22\x3b\x6e\x45\xba\x85\xd0\xcf\xb3\x0a\xfa\xff\x7b\x0d\xda\xac\xb4\x68\x5e\x5b\x4e\xc6\xbb\x7d\x4d\x34\x4b\xbd\xcd\x6e\xfd\x9c\x6d\x49\x67\x77\x35\xa3\xfd\xd9\x5f\xa7\x4b\x97\x22\x88\xa0\xc0\x8c\x5d\x3e\x9b\xec\x6f\x11\xfe\xfa\xe1\xba\x46\x72\x54\xe6\xc6\x94\xfa\x2a\x8a\x2c\x35\x4b\xbf\xc8\x94\xaa\xef\x5e\xbb\x62\x68\x42\xb5\x42\x6e\x2b\x27\x2f\x53\x35\x93\x2d\x89\xcb\x96\xc5\xc7\x82\xb1\xab\x9d\x82\xf0\xbb\xc1\x09\xc6\x99\x7f\x32\xbd\xd2\xd0\x14\x5e\x8e\x8a\x54\x52\x4c\xc2\xef\xe2\xa8\x7e\x80\x6d\xce\xd2\x1c\x98\xa0\xcc\x1e\x55\x1a\x98\xa9\x4f\x8f\xce\xcf\x97\x39\x43\x45\x54\x9a\xef\x82\xa4\xa8\xb8\x61\x25\xc7\x16\x51\x5b\x53\x61\x8d\x29\xa9\x34\xee\x33\x93\x31\x93\x57\xeb\x65\x2a\x0b\xbf\x08\x8f\x06\x51\xdd\x35\xa4\x91\x51\x88\x51\x41\xb4\x41\x15\x79\xc3\x5b\xb9\x79\xd9\x68\xfb\x6c\xd7\xaa\xb6\xae\x2c\x51\x68\xd0\xb2\x40\xc3\x0a\xd4\xfb\x84\x4d\xa4\x40\x8f\xad\xcb\xe3\xd9\xba\x1c\xb2\x55\x20\x11\x35\x53\x85\xad\xaa\x80\x63\xc6\x0c\x2b\x88\x41\xbe\x83\x9c\xdc\xe2\x09\xec\x95\x28\xed\xab\x67\x21\x91\xa8\xac\x42\x20\x6b\x59\x99\x26\x69\x6b\xa9\x0c\xc8\x8d\xfd\x22\xb2\x63\xe9\xfc\xed\xf1\x74\xbe\x99\xa2\x73\x8b\x35\x77\xe2\x95\x81\xb5\x34\x39\x2a\xa4\x60\xd4\xce\x9e\x91\x46\x42\xc9\x04\x50\xb9\x15\x7e\x95\xf1\x6a\x26\x4c\x59\x51\x20\x65\xb6\x3d\x6a\x94\x3f\x7f\x98\x72\xb6\x41\xeb\x6f\x2a\x8b\x92\xd7\x1b\x8b\x7a\x67\xe3\x21\xac\xbe\x39\x94\xd5\xeb\x7d\x56\x2d\x9f\x96\xb6\x2d\x02\x51\xcf\xe4\x71\x25\x52\x54\xb6\x5d\xad\xc9\x6e\xd6\xb5\x4d\x7d\x6d\xb8\x37\x2b\x27\x05\x4e\x07\xd7\xd0\xdb\x6b\x63\xe9\x6b\x24\x9b\xda\x65\x9c\xc3\x06\x84\x7d\x9d\x6c\x36\x86\x4c\xf6\x73\xc9\xfe\xac\xe9\xed\x32\x59\x4c\xb8\x14\xb5\x96\x95\xb0\x15\x04\xac\xe5\xdd\x20\xa7\xbc\x9c\xf4\x22\x43\x7b\xd2\xa9\x1d\x4c\x0d\x86\x0e\x6c\x2d\xef\xfa\x54\x32\xe8\x72\xbf\x85\x99\x8e\xd6\x7b\xe1\xb7\xaf\xc3\x61\x6f\x7e\xbd\xb4\x2e\x31\x96\x8a\x89\x94\x95\x84\x43\x8a\xc2\x28\xc9\xe8\xd9\xae\x38\xa0\xde\x0f\xaf\xf7\xef\xad\x1a\x9b\xb4\x9f\x2c\xbb\x42\xa3\xfe\x77\xb1\xe8\xda\x05\x8e\x1b\x13\xea\x9c\xa8\x5f\x6d\xe1\xd0\x88\x4c\xf6\x12\x7d\xa0\x77\x7e\x68\x4c\x8d\x57\x6e\x3c\x2e\x15\x96\x0a\x8d\xd9\xd9\xbe\xd5\x16\x32\x0b\x2f\x08\x9e\x10\x54\x64\xeb\xa4\x5a\xfb\xda\x67\xd7\xa8\xf7\xde\x28\x9b\x7d\x3a\x77\x2e\xe6\xfa\xa2\xa6\x73\x19\xfb\xd1\x82\xe7\x97\xae\xc5\x37\x78\x67\x42\xc2\x59\x26\xae\x2c\x4d\xf0\x82\x15\xa5\x54\x86\x08\xf3\xb6\xb9\x2f\x6a\x71\x92\xe6\xd3\x5b\xea\x91\x02\x22\xd2\x5c\xaa\x20\xf9\xcf\xbf\x9b\x33\xaf\xbb\x23\x68\x55\x1a\xb2\xe6\xd8\x13\xf9\x22\x0c\x9b\x42\xd5\x28\xd7\x21\xc6\x26\x4f\x7e\xf3\xf2\xcd\xef\x5f\xbf\x8d\x23\x93\xfb\xa3\x1f\x6b\x55\xe3\xd1\xbf\x11\x5e\xe1\x78\xf0\x9d\xf5\x9f\xa2\x48\xbd\x37\x71\xe4\x94\x84\x61\x6f\x41\x33\x66\x65\xec\x91\xdd\xf4\xac\xad\x8c\x9b\x41\xe7\x5b\x42\xaf\x80\x9c\x6a\xb3\x6d\x8c\x1a\x7a\x3a\xd4\xb5\xfe\xa4\x2a\xfc\x0a\x30\x7f\x16\x36\x7d\x8e\x71\x1a\x42\x7c\x22\xdc\x73\xc3\x06\x12\x6d\x73\xd7\x49\x64\xd4\xb2\x5f\x85\x8b\x06\xe9\x5c\x2a\x1c\xca\xa1\x4c\xb8\x90\xe8\x1b\x79\x17\x49\x07\xeb\x9c\xbe\x72\xe8\x54\x9f\x04\x34\xa6\xc1\x82\xcc\xcd\x3c\xd6\xd5\xfe\xbe\xe0\x68\x57\xa7\xef\x2a\x4e\x70\xd5\x03\x9a\x72\xf5\x04\x90\x53\x59\xa8\xab\xe1\x13\x69\x18\xdd\x10\x9c\xe1\x82\x45\x3a\x9f\x88\x06\x65\x9e\x89\x05\x74\x49\x2a\x8e\xfa\x23\xba\xcb\x3a\x33\xc9\xa5\x2f\xbb\x67\xf3\x8b\x15\xf1\xaa\xf3\xee\x71\xb2\x78\x3f\x28\x89\xb8\x94\xe8\xae\x07\xbc\x4b\xeb\x49\xed\x5a\x56\x2a\x45\x9b\xba\x9a\x07\x3f\x77\xb5\xaf\x0e\x4d\x5e\xde\x0d\xc6\xe0\x66\xe5\xa8\xeb\x80\xe4\x98\xd9\x7b\x8d\xc3\xa4\xde\xe9\x9a\x3e\xb4\x0a\x3a\xc0\x0f\x58\x4a\x78\x78\x88\xd6\x5c\xae\x5d\x45\x3f\x98\xb0\x67\xdf\xd7\x44\x1d\xf9\xd1\xd5\x68\xfe\x55\x4f\xfb\x0c\xde\xdf\xc4\xef\x13\x07\x14\x6b\x4f\xfd\x76\x31\xc0\x6e\x0b\x47\x37\x56\x10\x26\x3a\x85\x73\x3f\xb6\x6d\xa4\x34\xfe\x8f\x6d\x0d\x54\x87\x14\x47\x6b\x49\x77\xc9\x45\x1c\x35\xff\x5d\xaa\xef\x09\xfe\x1b\x00\x00\xff\xff\x8a\x46\xd6\x69\x5b\x25\x00\x00")

func wwwTemplatesHtmlIdHtmlBytes() ([]byte, error) {
	return bindataRead(
		_wwwTemplatesHtmlIdHtml,
		"www/templates/html/id.html",
	)
}

func wwwTemplatesHtmlIdHtml() (*asset, error) {
	bytes, err := wwwTemplatesHtmlIdHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "www/templates/html/id.html", size: 9563, mode: os.FileMode(420), modTime: time.Unix(1576395447, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wwwTemplatesHtmlInc_common_cssHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\xcd\xe1\x09\xc3\x20\x10\x06\xd0\xdf\xcd\x14\xe2\x00\x66\x81\xa6\xab\x84\x72\x9e\x78\xd4\x9c\xe0\x77\x20\x45\xdc\xbd\x94\x64\x81\x0c\xf0\x78\x63\x44\x4e\xa2\xec\xbc\x28\xed\x54\x8f\xa3\xea\x4e\x80\x9f\x73\x79\x16\xd1\x8f\x6b\x5c\x36\x0f\xfb\x16\x46\x66\x36\xef\x72\xe3\xb4\xf9\x95\x80\xb5\xe7\x8a\xaa\x49\x1a\x2c\xf4\xde\xc3\x1f\xbe\x6e\xba\xf3\x3c\xe9\xc3\xdd\xc4\xb0\xb7\x09\x05\x89\x57\x3d\x06\x6b\x9c\x73\xf9\x05\x00\x00\xff\xff\xa4\x70\x59\xda\xd7\x00\x00\x00")

func wwwTemplatesHtmlInc_common_cssHtmlBytes() ([]byte, error) {
	return bindataRead(
		_wwwTemplatesHtmlInc_common_cssHtml,
		"www/templates/html/inc_common_css.html",
	)
}

func wwwTemplatesHtmlInc_common_cssHtml() (*asset, error) {
	bytes, err := wwwTemplatesHtmlInc_common_cssHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "www/templates/html/inc_common_css.html", size: 215, mode: os.FileMode(420), modTime: time.Unix(1576393981, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wwwTemplatesHtmlInc_common_footerHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\xcc\xb1\x0a\xc3\x20\x10\x80\xe1\xfd\x9e\xe2\xea\x6e\xdd\x8b\x39\xfa\x26\x41\xcc\x89\x82\x7a\xc1\x1c\x74\x10\xdf\xbd\x43\xba\x74\xfb\xf9\x86\x7f\xce\x83\x53\xe9\x8c\xa6\xf4\xb8\x47\x69\x4d\xfa\x9e\x44\x94\x87\x59\x0b\xc0\xdf\x8d\xe5\xd8\xcc\x8f\x09\x10\x11\xfd\xc3\x5a\xf4\x01\xf3\xe0\xb4\x99\xac\x7a\x5e\x2f\xe7\xf4\x53\x54\x79\x3c\xa3\x34\x17\x6a\x95\xa4\x99\xcf\x1a\x22\x5f\x86\xde\xff\xe0\x5d\x20\xb4\x96\xc0\xbb\x7b\x4c\x00\x73\x72\x3f\xd6\x82\x6f\x00\x00\x00\xff\xff\xd8\x68\xd3\xe6\x96\x00\x00\x00")

func wwwTemplatesHtmlInc_common_footerHtmlBytes() ([]byte, error) {
	return bindataRead(
		_wwwTemplatesHtmlInc_common_footerHtml,
		"www/templates/html/inc_common_footer.html",
	)
}

func wwwTemplatesHtmlInc_common_footerHtml() (*asset, error) {
	bytes, err := wwwTemplatesHtmlInc_common_footerHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "www/templates/html/inc_common_footer.html", size: 150, mode: os.FileMode(420), modTime: time.Unix(1576394134, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wwwTemplatesHtmlInc_common_footerHtml2 = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\xcc\xb1\x0a\x03\x21\x0c\x80\xe1\x3d\x4f\x91\xba\xdb\xec\x25\x17\xfa\x26\x45\xce\x88\x82\xd5\xc3\x0b\x74\x10\xdf\xbd\xc3\x75\xe9\xf6\xf3\x0d\xff\x9c\x51\x53\x69\x8a\xae\xb4\xfd\x95\x7a\x37\x1d\x6e\x2d\x00\xbe\x1a\x4b\xdc\xdc\x8f\x05\x10\x11\xf9\xe6\x3d\x72\xc0\x3c\x34\x6d\x2e\x9b\x1d\xe7\x83\xc8\x3e\xc5\x4c\xc7\x7d\xef\x6f\x0a\xb5\xf6\x64\x59\x8f\x1a\x76\x3d\x9d\x3c\xff\x81\x29\x08\x7a\x2f\xc0\x74\x8d\x05\x60\x4e\x6d\x71\x2d\xf8\x06\x00\x00\xff\xff\xb1\x0b\xad\xd3\x8f\x00\x00\x00")

func wwwTemplatesHtmlInc_common_footerHtml2Bytes() ([]byte, error) {
	return bindataRead(
		_wwwTemplatesHtmlInc_common_footerHtml2,
		"www/templates/html/inc_common_footer.html~",
	)
}

func wwwTemplatesHtmlInc_common_footerHtml2() (*asset, error) {
	bytes, err := wwwTemplatesHtmlInc_common_footerHtml2Bytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "www/templates/html/inc_common_footer.html~", size: 143, mode: os.FileMode(420), modTime: time.Unix(1576393709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wwwTemplatesHtmlInc_common_headerHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x91\xbd\xae\xd4\x30\x10\x46\xeb\xf8\x29\x46\x43\x41\x15\x59\xa2\xdc\x8d\x5d\xf1\x23\x24\x24\x1a\x24\xca\x2b\xc7\x9e\xdc\x58\x6b\xcf\x04\xc7\x59\xee\x2a\xca\x23\xd0\xd1\xf3\x8a\x3c\x02\xda\x04\xb6\xa0\xfb\x8a\x23\x1d\xfb\xcc\xba\x06\x1a\x22\x13\x60\x64\xff\xe4\x25\x67\xe1\xa7\x91\x5c\xa0\x82\xdb\xa6\xba\x63\x42\x0c\x06\x03\x55\x17\xd3\xdc\x66\xe2\x05\xad\x02\x00\xe8\xd8\x5d\xad\x6a\x54\x73\x1f\x3b\x34\x93\x2b\x7e\x6c\xd9\x5d\x63\x5b\xe2\xf3\x58\xd1\xaa\x66\x47\x07\x29\x19\x32\xd5\x51\x82\xc1\x0f\xef\xbe\x20\x38\x5f\xa3\xb0\x41\x1d\x83\x46\x98\xeb\x2d\x91\xc1\xc8\x29\x32\xb5\x7d\x12\x7f\x41\xdb\xa8\xa6\xe9\xfa\xa5\x56\x61\xa8\xb7\x89\x0c\xce\x4b\x9f\x63\x7d\xe0\xbd\xf3\x97\xe7\x22\x0b\x87\xd6\x4b\x92\x72\x7a\x35\x0c\xc3\xb9\x97\x12\xa8\x9c\x80\x85\xe9\x9c\xdd\x4b\xfb\x3d\x86\x3a\x9e\xde\x14\xca\xe7\x10\xe7\x29\xb9\xdb\xe9\x10\x9d\xd1\xfe\xfe\xf5\xf3\x47\xa7\x0f\x89\xbd\xfb\x22\x4f\x4b\xfd\xab\xab\xf4\x52\x11\xae\x2e\x2d\x64\x10\xf7\x2f\x7e\x43\x60\x97\x69\x1f\x53\x72\x9e\x46\x49\x81\x8a\xc1\x4f\x22\x97\x65\x02\x07\x5f\x47\x79\x3d\xc3\x67\x86\xf7\xb1\xcc\x15\x3e\xbe\x7d\x3c\xf7\x7f\xb9\xfe\x57\x47\xdf\xf3\x58\xd5\x74\xfa\x91\xb4\x2f\xe0\x13\xb9\x62\xd0\xa5\x84\xa0\xad\x3a\x9a\xeb\xe3\x26\x56\xad\x2b\x10\x07\xd8\x36\xf5\x27\x00\x00\xff\xff\x56\x56\xba\x56\xc6\x01\x00\x00")

func wwwTemplatesHtmlInc_common_headerHtmlBytes() ([]byte, error) {
	return bindataRead(
		_wwwTemplatesHtmlInc_common_headerHtml,
		"www/templates/html/inc_common_header.html",
	)
}

func wwwTemplatesHtmlInc_common_headerHtml() (*asset, error) {
	bytes, err := wwwTemplatesHtmlInc_common_headerHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "www/templates/html/inc_common_header.html", size: 454, mode: os.FileMode(420), modTime: time.Unix(1576394126, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wwwTemplatesHtmlInc_common_headerHtml2 = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x91\xbf\x6a\xe4\x30\x10\x87\x6b\xeb\x29\x86\xb9\xe2\x2a\x23\xb8\xd2\xb6\x54\xdd\x1f\x0e\x0e\xae\x09\xa4\x0c\xb2\x34\x5e\x8b\x95\x47\x8e\x2c\x6f\x76\x31\x7e\x84\x74\xe9\xf3\x8a\x79\x84\xb0\x76\xb2\x45\xba\x5f\xf1\xc1\x27\x7d\xb3\x2c\x8e\x3a\xcf\x04\xe8\xd9\x3e\xf4\x64\x1c\x25\x5c\x57\xd1\xec\x13\xbc\x53\xe8\x28\x1b\x1f\xa6\x72\x20\x9e\x51\x0b\x00\x80\x86\xcd\x49\x8b\x42\x14\xd7\xb1\x41\x13\x99\x64\xfb\x92\xcd\xc9\x97\xc9\x1f\xfa\x8c\x5a\x14\x1b\xda\xc5\x34\xc0\x40\xb9\x8f\x4e\xe1\x9f\x5f\x77\x08\xc6\x66\x1f\x59\xa1\xf4\x4e\x22\x4c\xf9\x12\x48\xa1\xe7\xe0\x99\xca\x36\x44\x7b\x44\x5d\x88\xa2\x68\xda\x39\xe7\xc8\x90\x2f\x23\x29\x9c\xe6\x76\xf0\xf9\x86\xb7\xc6\x1e\x0f\x29\xce\xec\x4a\x1b\x43\x4c\xd5\xb7\xae\xeb\xea\x36\x26\x47\xa9\x02\x8e\x4c\xf5\x60\xce\xe5\x93\x77\xb9\xaf\x7e\x24\x1a\x6a\xe7\xa7\x31\x98\x4b\xb5\x8b\x6a\xd4\x6f\xaf\x2f\xcf\x8d\xdc\x25\xfa\xea\xf3\x3c\xce\xf9\x43\x97\xe9\x9c\x11\x4e\x26\xcc\xa4\x10\xb7\x2f\x3e\x22\xb0\x19\x68\x1b\x63\x30\x96\xfa\x18\x1c\x25\x85\xff\x62\x3c\xce\x23\x18\xb8\xef\xe3\xf7\x09\xfe\x33\xfc\xf6\x69\xca\xf0\xf7\xe7\xed\xb9\x5f\xe5\xf2\xb3\x8e\xbc\xe6\xd1\xa2\x68\xe4\x2d\x69\x9b\xc0\x06\x32\x49\xa1\x09\x01\x41\x6a\xb1\x37\x97\xfb\x4d\xb4\x58\x16\x20\x76\xb0\xae\xe2\x3d\x00\x00\xff\xff\xd8\x12\x90\x5b\xbf\x01\x00\x00")

func wwwTemplatesHtmlInc_common_headerHtml2Bytes() ([]byte, error) {
	return bindataRead(
		_wwwTemplatesHtmlInc_common_headerHtml2,
		"www/templates/html/inc_common_header.html~",
	)
}

func wwwTemplatesHtmlInc_common_headerHtml2() (*asset, error) {
	bytes, err := wwwTemplatesHtmlInc_common_headerHtml2Bytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "www/templates/html/inc_common_header.html~", size: 447, mode: os.FileMode(420), modTime: time.Unix(1576393215, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wwwTemplatesHtmlInc_common_metaHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\xc1\x6a\xf3\x30\x10\x84\xef\x79\x0a\xa1\x73\x44\xf8\x6f\x7f\x69\x55\x28\xa1\xa1\xbd\x27\xd0\x5b\x58\x4b\x9b\x7a\xa9\xb4\x52\xa5\x75\x12\x63\xfc\xee\x25\x69\x29\x36\xb9\xf4\xa6\x19\x66\xbe\x41\x3b\x0c\x1e\x0f\xc4\xa8\x34\xb1\xdb\xbb\x14\x63\xe2\x7d\x44\x01\x3d\x8e\x8b\x87\xcb\x43\xb5\x22\xd9\xe0\x67\x47\x47\xab\xd7\x89\x05\x59\xcc\xb6\xcf\xa8\x95\xfb\x56\x56\x0b\x9e\x65\xd5\x4a\x0c\xf7\xca\xb5\x50\x2a\x8a\xdd\x6d\x37\xe6\xbf\x7e\xfc\x61\x30\x44\xb4\xba\xe0\x01\x4b\xc1\x32\x69\xa6\x42\xef\xc4\xbf\xb9\xe9\xd6\x9b\xd9\x3d\x99\x75\x8a\x19\x84\x9a\x30\x9d\x7b\x7d\xb6\x77\x5a\xad\xe6\x70\xc8\x39\xa0\x89\xa9\xa1\x80\xe6\x84\x8d\x81\x9c\x8d\x83\x0c\xf3\x6e\x8f\xf5\x6f\xd5\x2a\x20\x5d\x35\x0d\x14\x53\xa5\x9f\x31\x9a\x00\xee\xe3\x86\xf2\x02\xec\x5b\x0c\x7e\x53\x08\xd9\x87\x7e\x7a\x9f\xd2\xe1\x4d\xfe\x48\x78\xca\xa9\xc8\x24\x77\x22\x2f\xad\xf5\x78\x24\x87\xe6\x2a\x96\x8a\x98\x84\x20\x98\xea\x20\xa0\xfd\xb7\x54\x11\xce\x14\xbb\x38\x31\x88\xe7\x46\x57\xb1\x5c\xd5\xe5\xef\x96\xd3\x75\x7a\x18\x14\xb2\x57\xe3\xb8\xf8\x0a\x00\x00\xff\xff\x65\xf3\xbe\xc6\xf4\x01\x00\x00")

func wwwTemplatesHtmlInc_common_metaHtmlBytes() ([]byte, error) {
	return bindataRead(
		_wwwTemplatesHtmlInc_common_metaHtml,
		"www/templates/html/inc_common_meta.html",
	)
}

func wwwTemplatesHtmlInc_common_metaHtml() (*asset, error) {
	bytes, err := wwwTemplatesHtmlInc_common_metaHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "www/templates/html/inc_common_meta.html", size: 500, mode: os.FileMode(420), modTime: time.Unix(1576394094, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wwwTemplatesHtmlIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x90\xbd\x4e\x05\x21\x10\x85\x6b\x79\x8a\x91\xc6\x92\x17\x18\x29\xfc\x6b\xaf\x85\x89\xb1\xba\xc1\x65\xae\x4b\x02\x83\x81\x49\xd4\x10\xde\xdd\x20\xbb\x85\xcd\xd2\x00\x93\x6f\xf2\x9d\x9c\xd6\xc0\xd3\x25\x30\x81\x0e\xec\xe9\x5b\x43\xef\x0a\xaf\x1f\x4e\xf7\x2f\x6f\xcf\x8f\xb0\x4a\x8a\x56\xe1\xb8\x20\x3a\xfe\xb8\xd5\xc4\xda\x2a\x00\x5c\xc9\xf9\xf1\x18\x07\x25\x48\x24\xfb\xba\xe6\x9b\x0a\x27\x86\xa7\x50\xaa\xa0\x99\xd3\x8d\x69\x4d\x28\x7d\x46\x27\x7f\xa6\xe5\xbc\xe4\x94\x32\x9f\x13\x89\xd3\xbd\x1f\x43\x4b\xad\x93\x41\xb3\x6b\xf1\x3d\xfb\x1f\xab\x8e\xf7\x06\x4c\x65\xac\xee\x41\x93\x0b\x6c\xd5\x15\xc0\xff\xac\x70\x57\xf2\x57\xa5\xb2\x63\x66\x72\xf3\x77\xec\xb8\xe4\x2c\xd3\xb1\xb1\x68\x66\x34\x34\xb3\xbc\xd6\x80\xd8\x8f\x5a\x7f\x03\x00\x00\xff\xff\xdf\x2a\xbf\x41\x6c\x01\x00\x00")

func wwwTemplatesHtmlIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_wwwTemplatesHtmlIndexHtml,
		"www/templates/html/index.html",
	)
}

func wwwTemplatesHtmlIndexHtml() (*asset, error) {
	bytes, err := wwwTemplatesHtmlIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "www/templates/html/index.html", size: 364, mode: os.FileMode(420), modTime: time.Unix(1576394020, 0)}
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
	"www/templates/html/id.html": wwwTemplatesHtmlIdHtml,
	"www/templates/html/inc_common_css.html": wwwTemplatesHtmlInc_common_cssHtml,
	"www/templates/html/inc_common_footer.html": wwwTemplatesHtmlInc_common_footerHtml,
	"www/templates/html/inc_common_footer.html~": wwwTemplatesHtmlInc_common_footerHtml2,
	"www/templates/html/inc_common_header.html": wwwTemplatesHtmlInc_common_headerHtml,
	"www/templates/html/inc_common_header.html~": wwwTemplatesHtmlInc_common_headerHtml2,
	"www/templates/html/inc_common_meta.html": wwwTemplatesHtmlInc_common_metaHtml,
	"www/templates/html/index.html": wwwTemplatesHtmlIndexHtml,
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
	"www": &bintree{nil, map[string]*bintree{
		"templates": &bintree{nil, map[string]*bintree{
			"html": &bintree{nil, map[string]*bintree{
				"id.html": &bintree{wwwTemplatesHtmlIdHtml, map[string]*bintree{}},
				"inc_common_css.html": &bintree{wwwTemplatesHtmlInc_common_cssHtml, map[string]*bintree{}},
				"inc_common_footer.html": &bintree{wwwTemplatesHtmlInc_common_footerHtml, map[string]*bintree{}},
				"inc_common_footer.html~": &bintree{wwwTemplatesHtmlInc_common_footerHtml2, map[string]*bintree{}},
				"inc_common_header.html": &bintree{wwwTemplatesHtmlInc_common_headerHtml, map[string]*bintree{}},
				"inc_common_header.html~": &bintree{wwwTemplatesHtmlInc_common_headerHtml2, map[string]*bintree{}},
				"inc_common_meta.html": &bintree{wwwTemplatesHtmlInc_common_metaHtml, map[string]*bintree{}},
				"index.html": &bintree{wwwTemplatesHtmlIndexHtml, map[string]*bintree{}},
			}},
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

