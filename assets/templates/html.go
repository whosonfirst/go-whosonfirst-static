// Code generated by go-bindata.
// sources:
// www/templates/html/alt.html
// www/templates/html/error.html
// www/templates/html/example.html
// www/templates/html/id.html
// www/templates/html/inc_alt_basics.html
// www/templates/html/inc_common_css.html
// www/templates/html/inc_common_footer.html
// www/templates/html/inc_common_header.html
// www/templates/html/inc_common_meta.html
// www/templates/html/inc_id_basics.html
// www/templates/html/inc_id_status.html
// www/templates/html/index.html
// www/templates/html/notfound.html
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

var _wwwTemplatesHtmlAltHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x58\x4d\x8f\xdb\x36\x13\x3e\xdb\xbf\x82\x2f\x81\x17\x3d\x59\x6a\x91\x53\xb7\x92\xdb\x20\x1f\xc5\x02\x49\x63\x24\x2d\x8a\x9e\x16\xb3\xd2\x58\xe2\x96\x1f\x02\x39\xb6\xd7\x75\xfd\xdf\x0b\x8a\x92\x2c\xd9\xf2\xee\xba\x6e\x9b\x4b\xc4\xe1\x3c\x0f\x1f\xce\x0c\x87\xf4\xee\x76\x2c\xc7\xa5\xd0\xc8\x38\x48\xe2\x6c\xbf\x9f\x26\xff\x7b\xfb\xe9\xcd\xcf\xbf\x2d\xde\xb1\x92\x94\x9c\x4f\x13\xff\x1f\x93\xa0\x8b\x94\xa3\xe6\xf3\x29\x63\x49\x89\x90\xfb\x0f\xff\x2f\x21\x41\x12\xe7\xbb\x1d\x8b\xbe\x2c\x3e\x47\x3f\x81\x42\xb6\xdf\xb3\x3f\x59\x6b\xb9\xcd\xc3\xf8\xd7\xd2\x7c\xe5\xd8\x27\xcd\xde\x0b\xeb\x28\x89\x03\xae\x61\xd9\xed\x18\xa1\xaa\x24\x10\x32\x2e\x74\x76\x97\x19\xa5\x8c\xbe\x53\x48\x50\xcb\x6a\x16\xf3\x63\x56\x59\x53\xa1\xa5\x6d\xca\xbd\xb6\x1b\x09\x8e\x94\xc9\xc5\x52\x60\xce\x59\x66\x34\xa1\xa6\x94\xfb\xf5\x3f\x80\xa3\x8f\xcd\x14\xdb\xef\x39\x8b\xe7\xd3\x09\x63\x67\xd8\x4c\x71\x43\xdb\x0a\x7b\x1c\x60\x49\x64\x12\x6b\xdc\x59\x8c\x13\x84\x77\x1a\x54\x1f\x38\xdc\xec\xd3\xf8\x3a\x10\x47\xc2\xfb\xa1\x14\x8e\x41\x17\xcd\x85\x84\x0c\xbd\xca\x7a\x46\x77\xf6\x37\x66\xa5\xc9\x6e\x9b\x3d\x9e\x5d\x2b\x47\x97\x59\x51\x91\x30\xba\xb7\xe2\x2d\xb9\xa3\xf4\xb0\xdb\xb7\x7e\xdd\x41\x0e\x9f\x66\x16\x0a\x8a\xe3\x5d\xbc\xd3\x79\x65\x84\x26\x17\x2d\x74\xc1\xf6\xfb\x01\xdd\xf7\x4e\xfc\x81\xe9\xa3\x53\x3d\xde\x33\xf4\xb4\x11\x44\x68\x6f\x32\xb0\x39\x67\x6b\x90\x2b\x4c\xb9\x5b\x29\x05\x76\xfb\x84\xaa\x16\xe6\x13\xd4\x53\xf6\x03\x48\x69\x96\x54\x62\xe5\x63\xe9\x5e\x40\x90\x59\x04\x32\xf6\x2a\x8e\x95\x95\x3d\x7c\x2c\xf2\xf8\xa5\xc1\x6d\x19\xfe\xb3\x3a\x69\x17\xfc\x97\x8a\xa5\xa5\xff\x47\x2b\x26\x28\x65\xce\x66\x29\x8f\x1f\x60\x0d\xc1\x10\x4b\x93\x81\x5c\x1a\x0b\x05\x46\x4a\xe8\xe8\xc1\xf1\x79\x12\x87\xc9\xf9\xe4\x79\xbc\x93\xa2\xaa\xb6\x0a\xaa\x28\xb3\xc6\xb9\x12\x84\x75\x43\x92\x17\x0a\xd9\x94\xc6\x19\xbd\xf4\xf1\x8a\x36\x9b\xcd\xb1\x90\x4b\xe0\x16\x75\x8e\x76\x54\xc5\x8b\xf0\x4d\x2a\x04\x5e\xbf\x93\x0c\xb2\x12\xaf\xd9\xcb\xca\x8a\xbf\xbf\x11\x8d\x74\x05\x18\x94\x58\x6e\x87\x78\x36\x61\x9d\xfc\x4b\xb8\x0a\x34\x0f\xce\x9c\x54\xd7\x25\x14\x12\x61\x29\x91\xa2\x15\x09\x39\x9e\x97\x8b\x68\x1c\x6d\xe5\x99\xfc\x5e\xc4\x53\x82\xce\x25\x1e\xd7\x7c\x88\xd3\x25\x7c\xf7\xd6\x6c\x1c\xda\x48\x41\x75\x85\xaa\x96\x25\xbc\x0e\x4e\x0e\xf3\xc5\x79\x6b\xf9\x40\x5e\x51\x48\x7d\x12\xa1\x05\x9d\xa9\x82\x67\x5e\x3a\x99\x73\xf5\x43\xa7\x5d\x59\x0a\xfd\x3b\xb3\x28\x53\x1e\x52\x59\x22\x12\x67\xa5\xc5\x65\xca\xe3\xcc\xb9\x51\x09\x22\x8f\x3c\x4f\xef\x24\x27\x71\xfb\x58\x4b\xee\x4d\xbe\x65\x39\x10\xcc\x7a\xd0\xd9\xca\x8a\x19\x36\x5d\xf7\xb8\x07\xbf\x05\x02\xdf\xcb\xe7\x9d\xaa\x5c\xac\x59\x26\xc1\xb9\x94\xfb\xbe\x0d\x42\xa3\xe5\xcf\x3d\xe2\xbc\x02\xb4\x9c\x45\xbd\xed\x29\x10\xba\x65\xea\xc9\xe1\x4c\xe4\x03\xc3\xac\xbe\x60\xf9\xa9\x6e\xef\x77\x74\xe3\x8c\xed\x2d\x38\xfd\xf2\xf9\x76\xdc\x43\xb8\x19\x48\x42\xab\x81\x30\xe5\xdf\x8c\x78\x80\xa4\x99\x33\x2b\x9b\x61\x47\xf5\xda\x16\x2e\x7a\x2d\xe9\x47\x34\x2a\xfa\x52\xcf\x8d\xb3\x7b\xec\x72\xa5\x33\x7f\x83\x8e\xa3\xdf\x37\xb3\xe7\xf1\xf8\x48\x16\x52\x3e\x32\x5b\x01\x95\x87\x18\x2c\x80\xca\x71\x96\x0a\x2c\xea\x61\xc0\x16\xb5\xa9\x1f\x36\x09\x24\x68\x95\xe3\xc1\xe7\x43\x63\xe9\xf9\x18\x5d\x1c\x3b\xb5\xa6\x83\x97\x12\x7a\x84\xed\xa3\xd0\x2d\xe1\xd0\xf3\x94\xd3\xbb\x8e\xd0\xc2\xe3\x18\x2d\x3c\x9e\xd2\x7a\xcf\x11\x5a\x78\x1c\xd0\xce\xa7\xfe\xa7\xc0\xb9\x8a\x9e\x4e\x7a\x15\xef\x43\xa7\xa0\xe2\xad\xa3\x35\x1b\x3e\x9f\x4e\x3c\x5c\xa8\x22\x74\x87\xe1\xc9\xf9\xb2\x6e\x5f\x2f\x6d\xe9\xd5\xa7\x38\xe5\x1b\x91\x53\x79\x23\x74\x89\x56\xd0\x77\xac\x44\x51\x94\x74\x18\x7b\xed\x8d\xed\xdb\xaf\xff\x1f\xc6\x01\xe2\x87\xbc\x55\x32\x73\xeb\xa2\xfd\x31\x53\xab\x8c\x73\xb1\x3e\x12\x7d\xaa\xb5\x7c\x15\x3e\x9a\xb3\xda\x7f\x3c\x76\xf6\xc4\x29\x90\x72\x9e\x64\x26\x0f\x3f\xea\x82\xfe\x24\xae\x0d\x49\x1c\xa6\x5b\xef\xc0\x1b\x97\xaf\x6a\xcb\x91\x9a\xce\xe7\x8c\xa8\x00\x99\x8c\xf4\x0d\x90\x74\x77\x0f\x4e\x64\x2e\x34\x8d\xcb\x78\xa7\x41\xd5\x20\xb1\x92\x1f\xb6\xde\xa5\x74\x70\x48\xba\x37\x51\x97\x65\x87\x19\xd5\xb0\xc9\xb3\x98\x59\x65\x91\x68\xeb\xdb\xbe\x97\x18\x36\xf6\x02\x98\x85\x4d\x8b\x39\xc8\x3b\x0c\x7b\x9f\x23\x99\x6e\x3e\x87\xb7\x5e\xec\x1b\x6b\xd7\xaf\xcf\x75\xe4\xa5\x31\x34\xe8\xc8\x81\xab\x77\x6b\xf8\xcb\x62\x3e\x4d\xe2\xf0\x47\x80\xdd\x8e\xa1\xf6\xdd\x62\xfa\x57\x00\x00\x00\xff\xff\x62\x47\xca\x2e\x32\x10\x00\x00")

func wwwTemplatesHtmlAltHtmlBytes() ([]byte, error) {
	return bindataRead(
		_wwwTemplatesHtmlAltHtml,
		"www/templates/html/alt.html",
	)
}

func wwwTemplatesHtmlAltHtml() (*asset, error) {
	bytes, err := wwwTemplatesHtmlAltHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "www/templates/html/alt.html", size: 4146, mode: os.FileMode(420), modTime: time.Unix(1579803216, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wwwTemplatesHtmlErrorHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x91\xbd\x6e\x03\x21\x10\x84\xeb\xdc\x53\x4c\x68\xd2\x1d\x2f\x80\x69\xf2\xa3\x74\x4e\x11\x29\x4a\x65\xe1\x63\xed\x43\xe2\xe7\x04\x2b\x47\x16\xba\x77\x8f\x10\xb9\x28\x95\x43\xc3\x8f\x66\xbf\x1d\x76\x6a\x85\xa5\x93\x8b\x04\x41\x39\xa7\x2c\xb0\xae\x83\xba\x7f\xda\x3f\xbe\x7f\xbe\x3d\x63\xe6\xe0\xf5\xa0\xda\x06\x6f\xe2\x79\x27\x28\x0a\x3d\x00\x6a\x26\x63\xdb\xa1\x2d\xc5\x8e\x3d\xe9\x8f\x39\x3d\x14\xec\x23\x5e\x5c\x2e\xac\x64\x7f\xfd\xd1\xd4\xca\x14\x16\x6f\x98\x20\x5c\x9c\x0e\x53\x0a\x21\xc5\x43\x20\x36\x62\x5d\x6f\x8b\xa6\x52\xba\x46\xc9\xad\xad\x3a\x26\x7b\xd5\xc3\x66\xc0\xba\x0b\x26\x6f\x4a\xd9\x89\x29\x45\x36\x2e\x52\x16\xff\xb4\x6e\x28\xca\x0d\xbc\x51\x82\x71\x51\x0f\x77\x80\x5a\xf4\x6b\x0e\xe3\x38\x82\x67\xca\x84\x2f\x53\x60\xb0\xe4\x74\xf4\x14\xc0\xf9\xea\xe2\x19\x9c\x60\x13\x78\x36\x3c\x2a\xb9\xfc\xce\x42\x76\x4a\xbf\xdd\x76\x70\x4a\x89\xbb\x83\xad\xd6\xba\xcb\x9f\x52\x25\xfb\x2f\x95\xec\x39\xd4\x0a\x8a\xb6\x25\xf4\x1d\x00\x00\xff\xff\xc2\x11\x53\xc7\xb7\x01\x00\x00")

func wwwTemplatesHtmlErrorHtmlBytes() ([]byte, error) {
	return bindataRead(
		_wwwTemplatesHtmlErrorHtml,
		"www/templates/html/error.html",
	)
}

func wwwTemplatesHtmlErrorHtml() (*asset, error) {
	bytes, err := wwwTemplatesHtmlErrorHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "www/templates/html/error.html", size: 439, mode: os.FileMode(420), modTime: time.Unix(1576865520, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wwwTemplatesHtmlExampleHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x90\xbd\x6e\xc7\x20\x0c\xc4\x77\x9e\xc2\x65\xe9\xc8\x0b\xb8\x2c\xfd\x58\xd3\xa1\x52\xd5\x29\xa2\xc1\x69\x90\xc0\x54\xc1\x43\x2b\xc4\xbb\x57\x94\xe4\x3f\x86\x05\x8c\x7e\xe7\x3b\x5d\xad\xe0\x69\x0d\x4c\xa0\x03\x7b\xfa\xd1\xd0\x9a\xc2\xbb\xa7\xe9\xf1\xed\xe3\xf5\x19\x36\x49\xd1\x2a\xec\x17\x44\xc7\x5f\x0f\x9a\x58\x5b\x05\x80\x1b\x39\xdf\x1f\xfd\xa0\x04\x89\x64\xdf\xb7\x7c\x5f\x60\x62\x78\x09\x7b\x11\x34\xe3\xf7\x60\x6a\x15\x4a\xdf\xd1\xc9\xbf\xd3\x32\x2f\x39\xa5\xcc\x73\x22\x71\xba\xb5\x6b\x68\x29\x65\x30\x68\x4e\x5b\xfc\xcc\xfe\xd7\xaa\x6b\x5d\x87\x69\xef\xd2\x33\x68\x72\x81\x6f\x2a\x34\x63\x1c\xd3\xf5\xaa\x35\x67\x19\xab\x0e\x16\xcd\x48\x80\x66\x74\x54\x2b\x10\xfb\xde\xde\x5f\x00\x00\x00\xff\xff\x3f\x52\xd4\x63\x53\x01\x00\x00")

func wwwTemplatesHtmlExampleHtmlBytes() ([]byte, error) {
	return bindataRead(
		_wwwTemplatesHtmlExampleHtml,
		"www/templates/html/example.html",
	)
}

func wwwTemplatesHtmlExampleHtml() (*asset, error) {
	bytes, err := wwwTemplatesHtmlExampleHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "www/templates/html/example.html", size: 339, mode: os.FileMode(420), modTime: time.Unix(1576865520, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wwwTemplatesHtmlIdHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x57\x51\x8f\xe2\x36\x10\x7e\x86\x5f\xe1\x5a\xaa\xfa\x04\x39\xe9\x9e\xba\x0d\x69\xa5\xdb\xab\xb4\xd2\x5d\x0f\xdd\xb5\xaa\xfa\xb4\x9a\x8d\x07\x32\x5b\xc7\x8e\xec\x01\x96\x52\xfe\x7b\x65\x4c\x20\x81\x64\x77\x29\xed\xed\xcb\xc6\x93\xf9\xbe\xf9\xc6\x33\x1e\x9c\xcd\x46\x28\x9c\x91\x41\x21\x49\x49\xb1\xdd\x0e\xd3\x6f\x6e\x3f\xbd\xfb\xf5\x8f\xe9\x7b\x51\x70\xa9\xb3\x61\x1a\xfe\x09\x0d\x66\x3e\x91\x68\x64\x36\x14\x22\x2d\x10\x54\x78\x08\x7f\x29\x13\x6b\xcc\x36\x1b\x31\xfe\x32\xfd\x3c\xfe\x05\x4a\x14\xdb\xad\xf8\x5b\xd4\x96\x3b\x15\xd7\xbf\x17\xf6\x3b\x2f\x3e\x19\xf1\x33\x39\xcf\x69\x12\x71\x7b\x96\xcd\x46\x30\x96\x95\x06\x0e\x4a\x4c\x7e\x9f\xdb\xb2\xb4\xe6\xbe\x44\x86\x9d\xac\x7d\xb0\xb0\x16\x95\xb3\x15\x3a\x5e\x4f\x64\xd0\x76\xa3\xc1\x73\x69\x15\xcd\x08\x95\x14\xb9\x35\x8c\x86\x27\x32\xc4\xff\x00\x9e\x3f\xee\x5f\x89\xed\x56\x8a\x24\x1b\x0e\x84\xe8\x61\xb3\xf3\x1b\x5e\x57\xd8\xe0\x00\xc7\x94\x6b\xdc\xe1\x7a\x31\x9e\x18\xef\x0d\x94\x4d\x60\x3b\xd9\xe7\xf1\xbb\x8d\x38\x11\xde\xdc\x4a\xf2\x02\x0e\xbb\x39\xd5\x90\x63\x50\xb9\x7b\x63\x0e\xf6\x77\x76\x61\xd8\xad\xf7\x39\xf6\xc6\x52\xe8\x73\x47\x15\x93\x35\x8d\x88\x77\xec\x4f\xca\x23\xee\x6e\x43\xdc\x56\x0d\x9f\x67\xa6\x12\xe6\xa7\x59\xbc\x37\xaa\xb2\x64\xd8\x8f\xa7\x66\x2e\xb6\xdb\x16\xdd\x8f\x9e\xfe\xc2\xc9\x93\x2f\x1b\xbc\x3d\xf4\xbc\x22\x66\x74\x37\x39\x38\x25\xc5\x12\xf4\x02\x27\xd2\x2f\xca\x12\xdc\xfa\x19\x55\x35\x2c\x14\xa8\xa1\xec\x27\xd0\xda\xce\xb8\xc0\x2a\xec\xa5\x7f\x05\x41\xee\x10\xd8\xba\xab\x38\x16\x4e\x37\xf0\x09\xa9\xe4\xb5\x9b\x5b\x33\x7c\xb5\x3e\xa9\x03\xfe\x4f\xcd\x52\xd3\xff\xa7\x1d\x13\x95\x0a\xef\xf2\x89\x4c\x1e\x61\x09\xd1\x90\x68\x9b\x83\x9e\x59\x07\x73\x1c\x97\x64\xc6\x8f\x5e\x66\x69\x12\x5f\x66\x83\x97\xf1\x5e\x53\x55\xad\x4b\xa8\xc6\xb9\xb3\xde\x17\x40\xce\xb7\x49\x5e\x29\x64\x55\x58\x6f\xcd\x2c\xec\xd7\x78\xb5\x5a\x9d\x0a\xb9\x04\xee\xd0\x28\x74\x9d\x2a\x5e\x85\xdf\x97\x82\xf0\xfa\x4c\x72\xc8\x0b\xbc\x26\x97\x85\xa3\x7f\x9f\x88\x41\xbe\x02\x0c\x25\xcd\xd6\x6d\xbc\x18\x88\x83\xfc\x4b\xb8\xe6\x68\x1f\xbd\x3d\xeb\xae\x4b\x28\x34\xc2\x4c\x23\x8f\x17\x4c\xba\xbb\x2e\x17\xd1\x78\x5e\xeb\x9e\xfa\x5e\xc4\x53\x80\x51\x1a\x4f\x7b\x3e\xee\xd3\x25\x7c\x0f\xce\xae\x3c\xba\x71\x09\xd5\x15\xaa\x6a\x96\x78\x3b\x38\x3b\xcc\x17\xd7\xad\xe6\x23\x75\xbd\x26\x52\x63\x32\xc4\x3d\x3d\xf0\xc2\x3d\x27\xf7\x7e\x77\xcd\xa9\x03\x6b\x32\x7f\x0a\x87\x7a\x22\x63\x21\x0b\x44\x96\xa2\x70\x38\x9b\xc8\x24\xf7\xbe\x4f\x41\xe0\x69\x9c\xe3\x34\xa9\xaf\x6a\xe9\x83\x55\x6b\xa1\x80\x61\xd4\x80\x8e\x16\x8e\x46\xb8\x9f\xb9\xa7\x13\xf8\x16\x18\xc2\x24\xcf\x0e\xaa\x14\x2d\x45\xae\xc1\xfb\x89\x0c\x53\x1b\xc8\xa0\x93\x2f\x5d\xe1\x82\x02\x74\x52\x8c\x1b\xe9\x95\x40\xa6\x66\x6a\xc8\x91\x82\x54\xcb\x30\xda\xfd\xbc\xca\x73\xdd\xc1\xef\xe4\xf7\xa6\x2b\xb7\xe8\xf4\xdb\xe7\xbb\x6e\x0f\xf2\x23\xd0\x3c\x91\x6f\x3a\xde\x55\xc0\xc5\x31\xc6\x14\xb8\xe8\xe6\xa8\xc0\xa1\x69\x0b\x9a\xee\x4c\x4d\x59\x1a\x98\x78\xa1\xf0\xe8\xf3\x61\x6f\x69\xf8\x58\x33\x3f\x75\xaa\x4d\x47\xaf\x92\x4c\x07\xdb\x47\x32\x35\x61\xdb\xf3\x9c\x33\xb8\x76\xd0\xc2\x53\x17\x2d\x3c\x9d\xd3\x06\xcf\x0e\x5a\x78\x6a\xd1\x66\xc3\xa7\x70\xd3\xee\x6b\x99\xe1\xa0\xd1\x52\x61\xef\x4a\xa8\x64\xed\xe8\xec\x4a\x66\xc3\x41\x80\x53\x39\x8f\xa7\xaf\xdd\x9a\x5f\x96\xf5\xe5\xa0\xae\xed\xee\x98\x4c\xe4\x8a\x14\x17\x37\x64\x0a\x74\xc4\x3f\x88\x02\x69\x5e\xf0\x71\x1d\xc4\xef\x6d\xdf\xbf\xf9\x36\xae\x23\x24\x2c\x65\xad\x64\xe4\x97\xf3\xfa\x5b\x61\xa7\x32\x51\xb4\x3c\x11\x7d\xae\xb5\x78\x1b\x1f\xf6\x87\xa1\x79\x37\x3b\xd8\x53\x5f\x82\xd6\x59\x9a\x5b\x15\xbf\x99\xa2\xfe\x34\xd9\x19\xd2\x24\xbe\x8e\x7c\x49\xf1\x36\x0b\x0f\x27\x2a\xf6\xab\x5e\x31\x11\x32\xe8\x38\x90\xa4\xee\x1f\xc0\x53\xee\xe3\x61\xbc\x8c\x76\x18\x45\xb5\xea\xa9\x65\x2b\xe3\xb3\x68\x9e\x81\x17\x31\xda\xe0\xc0\x3f\x38\x86\xac\x95\xf6\xb3\x1e\xfa\xa3\x75\xe4\x0e\xf7\x97\x43\xcb\x78\xcc\x79\x07\x1b\xbc\x88\x19\x55\x0e\x99\xd7\x61\x48\x07\x15\x51\xc3\x2b\x60\x0e\x56\x35\xe6\x28\xef\xb8\x6c\xe5\x74\xd6\x36\xfb\xc7\xf6\x2f\x54\x12\xc6\xe0\x61\xba\xf6\xcd\xcf\x99\xb5\xdc\x9a\x9f\x91\xab\x31\xe3\xc3\x68\xcf\x86\x69\x12\x3f\xd8\x37\x1b\x81\x26\xcc\x9e\xe1\x3f\x01\x00\x00\xff\xff\x5e\xae\x5c\x21\xdd\x0f\x00\x00")

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

	info := bindataFileInfo{name: "www/templates/html/id.html", size: 4061, mode: os.FileMode(420), modTime: time.Unix(1594859184, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wwwTemplatesHtmlInc_alt_basicsHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x93\x41\x8f\xd3\x30\x14\x84\xcf\xe4\x57\x8c\xcc\x81\x0b\x4d\xe1\x9e\xcd\x05\x58\xa9\x12\x08\xc4\x22\xed\x71\xe5\xc4\x2f\xe9\x93\xcc\x73\xb0\xdd\x52\x14\xf9\xbf\xa3\x34\x49\xdb\xa0\x48\x20\xf6\xe8\xc9\x64\xfc\x8d\xed\xd7\xf7\x86\x1a\x16\x82\x62\xa9\x9f\xb4\x8d\x4f\x95\x0e\x5c\x07\x95\x52\x56\x18\x3e\x82\xcd\x9d\xfa\xb9\x77\xc1\x49\xc3\x3e\xc4\x8d\xa1\xa8\xd9\x06\x55\x66\xc5\xd6\xf0\xb1\xcc\x5e\x00\xeb\xce\x29\xa8\xcc\x00\xa0\xe8\xca\x6f\x7b\x0e\xe0\x00\x2d\xd0\x36\x92\x17\x1d\x09\x2d\xb9\xef\x14\xfd\x2f\x34\xce\xe3\x71\xef\x5e\x05\x7c\x16\xdc\x0f\x01\xd8\xbd\x47\xa1\xb1\xf7\xd4\xdc\xa9\xbe\x47\xfe\x41\x4c\xe7\x58\x62\xc8\x77\x06\x29\x0d\xd2\xc3\x97\xaf\xe3\x42\xa1\xb6\x3a\x84\x25\x01\x1b\x55\x2e\x5c\xc5\x56\x97\x79\xb1\xed\x46\xa6\x19\x6c\x17\x03\x2a\x77\x10\xc3\xd2\xa2\x72\xa7\x01\xf2\xb2\xf1\xcb\xd5\xe4\x0b\xf6\x9a\xb8\x99\xc3\x2a\x77\xba\x02\x7c\x62\xf9\xa8\x23\xc7\x83\xa1\x94\x5e\xe3\x56\x76\xd2\x9e\x75\x2c\x3e\xe8\xd3\xaa\x5f\x9f\x6e\xfd\x43\xa3\x73\xa1\xdb\x32\x9d\x67\xa9\xb9\xd3\x16\x35\x49\xf4\x8e\xcd\xb3\x2b\xcd\x41\xd7\x3e\x33\xdc\x02\xfa\x4f\xb4\xf1\xb0\xb3\x62\x3a\xf1\xbe\x07\x37\xa0\x1f\xd3\x95\x84\x77\x07\xef\x49\x62\xfe\x10\x3d\x4b\x7b\x6f\x75\x0b\xf5\x56\x21\xa5\xb1\x4d\xe8\xb4\xac\xa1\x86\xa8\xe3\x21\x6c\xea\xf1\x6f\x55\x3e\x12\x2a\xb2\x4c\x47\x42\x1c\x1e\x99\xa7\xda\x79\x83\xe8\x50\x11\x26\x57\x5e\x6c\x87\xb8\x0b\x07\xd9\x40\x7f\x87\x79\xf3\xcf\x30\xe2\xe2\x02\x48\x1c\xac\x93\x96\xfc\x7f\xa2\x4d\xdb\x0e\x4b\x19\x9e\xee\x79\xcc\x2e\x17\x3d\x0f\x5f\xdf\x93\x98\x94\xb2\xdf\x01\x00\x00\xff\xff\xd7\xcd\xc2\xef\xc7\x03\x00\x00")

func wwwTemplatesHtmlInc_alt_basicsHtmlBytes() ([]byte, error) {
	return bindataRead(
		_wwwTemplatesHtmlInc_alt_basicsHtml,
		"www/templates/html/inc_alt_basics.html",
	)
}

func wwwTemplatesHtmlInc_alt_basicsHtml() (*asset, error) {
	bytes, err := wwwTemplatesHtmlInc_alt_basicsHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "www/templates/html/inc_alt_basics.html", size: 967, mode: os.FileMode(420), modTime: time.Unix(1579628589, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wwwTemplatesHtmlInc_common_cssHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\xcd\xe1\x09\xc2\x30\x10\x06\xd0\xdf\x66\x8a\x70\x03\xb4\x0b\x58\x57\x29\x98\x5e\x48\x30\xbd\x83\xfb\x02\x87\x84\xec\x2e\xd2\x09\x74\x80\xc7\x1b\xe3\xe0\x5c\x85\x23\x55\x49\x7b\xd2\xf3\x54\xd9\x13\x40\x73\x86\x7b\xab\xf2\x8a\xc6\x6d\x23\xf4\x77\x63\x14\xe6\x4e\xb1\x18\xe7\x8d\xd6\x04\xac\x5e\x14\x2a\xb9\x1a\xfa\xe2\xee\xcb\x17\x3e\x7e\x74\xd7\xf9\x17\x7d\x9a\x3a\xd8\x2e\x7b\x8b\x61\x0c\x96\x63\xce\xf0\x09\x00\x00\xff\xff\xf6\x2a\xad\xcc\xd5\x00\x00\x00")

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

	info := bindataFileInfo{name: "www/templates/html/inc_common_css.html", size: 213, mode: os.FileMode(420), modTime: time.Unix(1576865520, 0)}
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

	info := bindataFileInfo{name: "www/templates/html/inc_common_footer.html", size: 150, mode: os.FileMode(420), modTime: time.Unix(1576865520, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wwwTemplatesHtmlInc_common_headerHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x91\xc1\x6a\xdc\x30\x10\x40\xcf\xd6\x57\x0c\xd3\x43\x4f\x8e\xa1\xe4\xb4\x2b\xeb\x50\x9a\x96\x40\xa1\x97\x42\x8f\x41\x96\xc6\xb5\x58\x79\xc6\x48\xb2\xb3\x8b\xd9\x4f\xe8\xad\xf7\xfe\x62\x3f\xa1\x6c\x9c\x86\x90\xbd\x88\x41\x20\xbd\xa7\xa7\x75\xf5\xd4\x07\x26\xc0\xc0\xee\xc1\xc9\x38\x0a\x3f\x0c\x64\x3d\x25\x3c\x9f\x95\xde\x46\x08\xbe\xc5\xe7\x5d\xa3\x00\x00\x9e\x16\xcd\x76\x31\x4a\x55\x7a\xb8\x05\x17\x6d\xce\x2d\xb2\x5d\xea\x50\x68\x44\xa3\x2d\x0c\x89\xfa\x16\xd7\x15\x6e\xee\xd8\x4f\x12\xb8\xe4\x9b\x7b\xf6\x74\x84\xf3\x19\xcd\x8f\x41\xde\x67\xf8\xc6\xf0\x39\xa4\x5c\xe0\x63\x92\xc7\x4c\x49\x37\xd6\xe8\x66\xb8\x35\xaa\x52\x95\xf6\x61\xb9\xba\xf9\x95\x4c\x9d\xc9\x26\x37\xa0\x51\xd5\x93\x4f\x2f\x69\x84\x91\xca\x20\xbe\xc5\x2f\x77\xdf\x11\xac\x2b\x41\xf8\x4a\xc2\x5f\x0c\x20\x97\x53\xa4\x16\x03\xc7\xc0\x54\x77\x51\xdc\x01\x4d\xa5\xaa\x4a\x07\x9e\xe6\x02\xe5\x34\x51\x8b\x85\x8e\x05\x61\xb1\x71\xa6\x16\x37\x7a\xf0\x08\x6c\x47\xda\xa6\x29\x5a\x47\x83\x44\x4f\xa9\xc5\xaf\x22\x87\x79\x02\x0b\x6f\x9e\x77\xff\xe9\x85\xe7\x43\x9e\xa2\x3d\xed\x36\xee\x1e\x1b\x73\x41\x76\x73\x29\xc2\xcf\xcc\x3c\x77\x63\x28\x2f\x27\x3a\xeb\x0e\x3f\x93\xcc\xec\x6b\x27\x51\xd2\xee\x5d\xdf\xf7\xfb\x4e\x92\xa7\xb4\x03\x16\xa6\xfd\x68\x8f\xf5\x63\xf0\x65\xd8\x7d\x48\x34\xee\xdf\x32\xcc\xdf\x3f\xbf\x7f\xe9\x66\x83\xfc\xcf\xd5\x5c\x7a\x19\x55\xe9\xc6\x87\x65\x0b\xde\x25\x70\x91\x6c\x6a\xd1\xc6\x88\xd0\x18\xb5\xfd\x74\xb3\x05\x37\x6a\x5d\x81\xf8\x92\x4f\xfd\x0b\x00\x00\xff\xff\x7f\x09\xff\xc1\x3b\x02\x00\x00")

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

	info := bindataFileInfo{name: "www/templates/html/inc_common_header.html", size: 571, mode: os.FileMode(420), modTime: time.Unix(1576865520, 0)}
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

	info := bindataFileInfo{name: "www/templates/html/inc_common_meta.html", size: 500, mode: os.FileMode(420), modTime: time.Unix(1576865520, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wwwTemplatesHtmlInc_id_basicsHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x96\x4d\x6f\xe3\x36\x10\x86\xcf\xc9\xaf\x18\x68\x17\xc8\x65\x6d\x35\xdd\x3d\x05\x8a\x0e\xdd\xc5\x02\x06\xfa\x11\x6c\x5a\xec\x31\xa0\xc4\xb1\x34\x28\x45\xb2\xe4\x28\x8e\x61\xe8\xbf\x17\x94\x64\x7d\x24\xf2\x36\x4d\xdb\x00\xbd\x18\x36\x3d\x7c\xf9\xce\xc3\xe1\x90\x87\x83\xc4\x2d\x69\x84\x88\x74\x7e\x47\xf2\x2e\x13\x9e\x72\x1f\x35\xcd\x79\x22\xe9\x1e\x48\x5e\x47\xbb\xd2\x78\xa3\xb7\xe4\x3c\xaf\x24\xb2\x20\xe5\xa3\xf4\xfc\x1c\x00\xa0\xfd\x38\x1c\x80\xb6\x80\x7f\xc0\xfa\xf6\xe6\xcb\x7a\xe3\x3f\xa1\x75\x98\x0b\x46\xb9\xbe\x65\x47\xba\xf8\xac\x44\x01\xd1\x65\x04\x4d\xd3\x4e\x48\xbc\x15\x1a\x72\x25\xbc\x9f\xab\x7b\x16\x5c\xfb\x95\x1c\x04\xa2\xf4\xd7\x92\x3c\x38\xcc\x8d\x93\x50\x0a\x0f\x19\xa2\x86\x4a\xb8\xdf\x51\x82\xf0\x30\x86\xae\x93\x38\xc8\xa6\x47\x4b\xa8\xe5\x71\xbd\x45\x97\xb7\xb5\x45\xe7\x51\x2e\xba\x3c\x03\x18\x26\x2e\x3a\xf0\xc3\x6c\xc8\xf6\xc0\x25\xc2\xd6\x28\x65\x76\xa4\x8b\x3e\xd6\x5f\x75\xb9\xd6\x6a\x29\x53\x6b\x48\x33\xba\x00\xf2\xec\x70\x00\x27\x74\x81\xf0\x96\xde\xc1\x5b\x0f\x57\xd7\x9d\xc7\xd1\xe1\x0f\xfb\xd6\x54\xa2\x28\x4d\x04\x94\x0e\xb7\xd7\x51\x4c\x32\x3e\x1c\x42\x7c\xd3\x44\x4b\x4b\x90\x84\xe9\x4f\x2d\x2a\xda\xee\x23\x90\x82\xc5\x6a\x1e\x77\x1d\x0d\x42\xe9\xf1\x5b\x12\x8b\x34\x89\x15\x75\xfe\x26\x30\x93\xb8\x56\x1d\xe4\xf5\x0c\x6e\x1f\x32\x25\x77\x0a\x39\xe9\xe2\x44\x65\x3c\x01\x3e\x70\xf6\xff\x35\x64\xff\xbf\x41\x9c\xc4\x92\xee\xd3\x40\x7a\xf9\x8c\xf6\x47\xb8\x53\x48\x6c\x58\xaf\x4d\xf5\x67\x51\x21\x34\xcd\x85\x87\xaf\xa5\xb9\xf0\xf0\x8b\x86\xcf\x61\x02\x6c\x3e\x01\xf9\xd3\xa7\x92\x64\x34\x68\x6c\x64\x6b\xbc\x3d\x69\x20\xb4\x04\xe2\xb0\x57\x4a\x30\xdd\x23\xfc\xf6\x65\xf3\x4d\x25\x87\xca\x0a\x2e\x47\xb9\x1b\xc1\xe5\x28\xb8\x4e\x62\xdb\x37\x96\xc4\xa6\xe7\x67\x1b\x0e\x6a\xe2\xb4\x9e\x55\x22\x47\xde\x5b\x9c\x28\x1e\x87\x26\x3e\x49\xcf\x25\xc6\xe8\x8f\xa6\xd6\xec\xf6\xa3\x83\x53\xfd\xe2\x46\x38\xd4\xbc\x91\x10\xad\x3e\x0c\xd5\x3a\x29\xd4\x0b\x0f\xb6\x0d\x39\xc2\xcc\x8d\xc4\x74\xf5\x21\x89\xdb\x2f\xb0\x2b\x29\x2f\x81\xb4\xa4\xd0\xac\x3c\x10\xb7\xbd\x64\x28\xb5\x37\x25\xa1\x13\x2e\x2f\xf7\x51\x5a\xd5\x8a\xc9\x2a\xec\x15\x7d\xa8\x13\xc8\x30\x17\xb5\xc7\x71\x46\xc9\x6c\xfd\x55\x1c\x17\xc4\x65\x9d\xad\x73\x53\xc5\x13\x34\xf1\x0c\x93\x33\x16\x1d\x13\xfa\x98\x1d\x62\x5c\x09\xcf\xe8\xe2\xc9\xf0\xce\x6c\xdf\x74\xab\xdd\x85\xdd\x6e\xdd\x59\x8b\xda\x83\x37\x15\x32\x55\xd8\xba\x58\x0f\x95\xa8\x3c\x9e\xc0\xf3\xbe\x6b\xa0\x7f\xcd\xe6\xfd\x9c\x4d\x85\x42\xb7\x5c\x2a\x2a\x4a\x06\x85\x05\x31\x55\x82\x51\xed\xa1\x14\xf7\xf8\x02\x56\x16\x4d\xf8\xeb\x55\x90\x09\x57\xd4\x08\x22\x33\x35\x03\x87\xdc\xbd\x71\x0c\x66\x1b\x7e\xe8\xe2\xb9\xf0\xbe\xff\x1b\xb5\x75\xb9\xc4\x6f\x87\x2d\x2c\x7d\xc1\x90\x19\x2e\xd1\xa1\x04\x76\xfb\xd0\x31\xd9\x80\x25\x0d\xd2\xec\x74\xe7\x70\x50\x5f\x24\x4b\x55\x85\x92\x04\x1f\xd1\xbe\x7e\x15\x2a\xda\x62\xc8\x37\x37\x95\x55\xed\xb9\x91\xc1\xc4\x33\x30\x5e\x7e\xbb\x06\x37\x4f\x19\x06\x7a\x01\xd2\x0e\x41\xb8\x57\xca\xaf\xd6\x39\x3a\x16\xa4\x5b\xb4\xdd\x2e\xf6\xd7\x5e\x5f\xcd\xdd\x3e\x19\x8d\xf3\xda\x09\xb9\x75\xbd\xb1\x8b\xeb\x1e\x21\x8f\xaf\xad\x19\x94\x7f\xe7\x02\x7b\x2c\x99\x3e\xbd\x48\x9e\x46\xcd\x6b\xbf\xbb\xc2\xfa\x6b\xce\xa6\x63\xbf\x4d\x6c\xba\x61\x0f\x99\xa9\x75\x78\x20\x40\x66\x1e\xda\x62\x1f\x6a\x73\x31\x81\x02\x43\x7f\x72\x7b\x58\x1a\x5c\x1d\xc5\x32\xf3\x30\xf6\xfc\x9f\x48\xff\x28\x98\xb8\x96\xd8\x34\xef\x60\x3a\x6c\x74\xd1\x8e\xc3\xec\x0f\xf1\xb0\x18\x2f\x1e\xa6\xf1\xdd\x65\x3e\xb9\xbe\x42\x32\xd6\x91\xce\xc9\x0a\x05\x39\x6a\x76\x86\xe4\x3f\x4e\xe9\x28\x34\xe6\x73\x34\x37\x33\xfd\xd8\x5a\x7f\xb5\x26\x36\x5d\x7c\x99\x7d\xac\x5d\xd8\xae\x17\xbe\xd7\xf3\x6e\x76\x94\x7e\x45\xc8\x50\x11\xde\xe3\xb4\xc3\x84\xce\x93\x21\xf4\x51\x4f\xde\xe9\x8f\x8e\xf1\xb2\x99\xef\x9e\x6d\x46\x1b\x9e\x19\xd2\x06\x94\xd1\x05\xba\x17\x5a\xeb\x97\x1d\x6a\xb7\x7d\x76\x0d\x1b\x7d\x7c\x8c\x1d\x0e\xa8\x65\xd3\x9c\xff\x19\x00\x00\xff\xff\xf6\x5c\x1e\xac\x50\x0d\x00\x00")

func wwwTemplatesHtmlInc_id_basicsHtmlBytes() ([]byte, error) {
	return bindataRead(
		_wwwTemplatesHtmlInc_id_basicsHtml,
		"www/templates/html/inc_id_basics.html",
	)
}

func wwwTemplatesHtmlInc_id_basicsHtml() (*asset, error) {
	bytes, err := wwwTemplatesHtmlInc_id_basicsHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "www/templates/html/inc_id_basics.html", size: 3408, mode: os.FileMode(420), modTime: time.Unix(1576865520, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wwwTemplatesHtmlInc_id_statusHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x95\x41\x8f\x9b\x30\x10\x85\xcf\xf0\x2b\x2c\xdf\x13\x0e\x7b\x6b\x1d\x5f\xba\xad\xb4\xea\x65\xd5\xf4\xbe\x72\xf1\x24\xb8\x32\x36\xf2\x18\xb2\x11\xe2\xbf\x57\xd8\x90\x92\xdd\xa4\x69\x60\x4f\x49\x9c\x79\x9f\x79\x6f\xd0\x4c\xdb\x4a\xd8\x29\x03\x84\x2a\x93\xbf\x28\xf9\x82\x5e\xf8\x1a\x69\xd7\xa5\x29\x93\xaa\x21\x4a\x6e\xe8\xa1\xb0\x68\xcd\x4e\x39\xf4\x2b\x74\x39\x25\xb9\x16\x88\x1b\x8a\x90\x7b\xca\x53\x96\x49\xd5\xf0\x2b\xf5\x07\xbb\xbb\xab\x3e\xde\xbf\x7a\x7d\x2b\x4a\x09\x21\x84\x49\xf0\x42\x69\x24\xb6\x02\xc3\xd3\x84\x61\x5d\x96\xc2\x1d\x79\x9a\x84\xbf\x8b\x07\x82\xfe\xa8\x61\x43\xa5\xc2\x4a\x8b\xe3\x27\x65\xb4\x32\xf0\x99\x86\x7b\x5e\x07\x3a\xe5\xf1\x93\x65\xc5\x43\x4f\xc9\xfe\x62\xc6\x1b\x56\x25\x98\x7a\xc0\x0e\xf0\xfe\x69\x87\x67\xf2\xe2\x97\x86\x95\x03\xac\xac\x41\xd5\x00\xe5\x69\x92\xb0\x70\x7a\x56\x42\x62\x21\x96\xa1\x20\x50\xbc\xeb\xbf\x26\xcc\x17\x5c\x21\xc9\x6b\xe7\xc0\x78\x96\xf9\x62\x38\x96\x23\xe0\x7d\x28\x94\xb7\x2d\x59\x6f\x9f\x7f\xac\x9f\xf0\x4b\x14\xae\xb7\xde\x29\xb3\xff\xa6\xc5\x9e\x74\x1d\xcb\xbc\x9c\x87\x79\xc2\x9f\xae\x86\x85\x88\xef\xc6\x1e\xcc\x84\x11\xfc\x66\xd1\xf0\x25\xef\x20\x10\xe4\x0c\xeb\x41\xb7\xd8\x79\xa4\x2c\x31\x3e\x12\xee\xf4\x2d\xa1\x72\x90\x0b\x3f\xc7\xfb\xe3\x49\xbb\xd4\xff\x84\x74\x31\x83\x6b\xa5\x77\x9a\xc5\xba\x02\x87\x20\xe7\x98\xdd\x9e\xb4\x4b\xcd\x4e\x48\x0b\x1a\x7e\x46\x99\x99\x83\x32\xfb\xf9\x41\x28\xb3\xff\xa8\x24\x7a\xd4\x07\x44\x11\x31\x57\xb3\x48\x48\xd2\x8f\xcf\x84\x65\x61\x10\x8e\x63\x3a\xce\xff\x61\xb6\xb2\xec\x7c\xe4\x0e\x15\xf1\x2c\xfe\xfc\xf7\xc6\xd8\x59\x57\x0a\x8f\x6f\x17\xc6\x74\x5f\xdc\xb5\x2a\xf8\x00\xbc\xbd\x1d\x26\xcb\xa1\xd6\x61\x05\x68\xc5\x99\x20\x85\x83\xdd\x86\xf6\x69\x7d\x35\xb2\xb2\xca\x78\x5c\x3f\x9b\xbe\x61\xa7\x04\x25\xe9\x3a\xca\xab\xfe\x6d\x10\x9c\x65\x5a\xdd\x90\x6f\x9b\x0b\x72\x6c\xfe\x57\xfe\x28\xbc\x78\xaf\xdf\x83\xfd\x8d\xd6\x4c\x18\x31\xfc\x60\xe6\x46\x5f\xc6\x96\xb4\x2d\x18\xd9\x75\xe9\x9f\x00\x00\x00\xff\xff\xa6\x5a\xf3\x2e\x49\x08\x00\x00")

func wwwTemplatesHtmlInc_id_statusHtmlBytes() ([]byte, error) {
	return bindataRead(
		_wwwTemplatesHtmlInc_id_statusHtml,
		"www/templates/html/inc_id_status.html",
	)
}

func wwwTemplatesHtmlInc_id_statusHtml() (*asset, error) {
	bytes, err := wwwTemplatesHtmlInc_id_statusHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "www/templates/html/inc_id_status.html", size: 2121, mode: os.FileMode(420), modTime: time.Unix(1579628589, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wwwTemplatesHtmlIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\xbd\x6e\xc3\x30\x0c\x84\xe7\xfa\x29\x58\x2d\xdd\xac\x17\x60\xb4\xf4\x67\x4d\x87\x02\x45\xa7\x40\x95\x98\x5a\x80\x44\x15\x16\x11\xb4\x10\xfc\xee\x05\xe3\x18\xc8\x14\xf4\x16\x13\x87\xf3\x7d\x14\x7b\x87\x48\xc7\xc4\x04\x26\x71\xa4\x1f\x03\xcb\x32\xe0\xfd\xd3\xfe\xf1\xed\xe3\xf5\x19\x26\x29\xd9\x0d\xa8\x1f\xc8\x9e\xbf\x76\x86\xd8\xb8\x01\x00\x27\xf2\x51\x07\x15\x4a\x92\x4c\xee\x7d\xaa\x0f\x0d\xf6\x0c\x2f\x69\x6e\x82\x76\x75\x2f\x99\xde\x85\xca\x77\xf6\x72\x26\x85\x43\xa8\xa5\x54\x3e\x14\x12\x6f\x96\xe5\x76\x28\xb4\xa6\x19\xc5\xda\x8d\x8b\x9f\x35\xfe\xba\x61\xdb\x20\xa6\x13\x84\xec\x5b\xdb\x99\x50\x59\x7c\x62\x9a\x8d\x83\x2b\xdd\x46\x68\x2d\xcd\x06\x46\xe5\xdc\x01\x60\xf1\x89\x9d\x4e\xaa\x71\x1c\xcf\xa6\x5d\xdd\xff\x14\x1e\x6b\x95\xad\xf0\xb2\xa3\x8d\xe9\x74\xf5\x33\xda\xf5\x09\x68\xd7\x2b\xf7\x0e\xc4\x51\xef\xff\x17\x00\x00\xff\xff\xe0\x53\x6d\xb7\x95\x01\x00\x00")

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

	info := bindataFileInfo{name: "www/templates/html/index.html", size: 405, mode: os.FileMode(420), modTime: time.Unix(1576865520, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wwwTemplatesHtmlNotfoundHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x91\x3f\x4f\x34\x21\x10\xc6\xeb\x77\x3f\xc5\xbc\x34\x76\x62\x71\xe5\x1c\x8d\xa7\x89\xd5\x59\x98\x18\xab\x0b\xc2\xac\x4b\x02\xc3\x66\x19\x2f\x31\x64\xbf\xbb\x21\xb8\x46\x9b\x93\x86\x3f\x79\x9e\xe7\xc7\xcc\xd4\x0a\x9e\xc6\xc0\x04\x8a\xb3\x8c\xf9\x9d\xbd\x82\x75\x1d\xf0\xff\xe1\x78\xfb\xf4\xf2\x78\x07\x93\xa4\x68\x06\x6c\x1b\x44\xcb\x6f\x7b\x45\xac\xcc\x00\x80\x13\x59\xdf\x0e\x6d\xa1\x04\x89\x64\x9e\xa7\x7c\x55\xe0\xc8\x70\x1f\x96\x22\xa8\xfb\xeb\x97\xa6\x56\xa1\x34\x47\x2b\x04\x2a\xb0\x3b\xb9\x9c\x52\xe6\x53\x22\xb1\x6a\x5d\x2f\x8b\x5c\x29\x5d\x83\x7a\xc3\xe2\x6b\xf6\x1f\x66\xd8\x3e\xe0\xc3\x19\x5c\xb4\xa5\xec\x95\xcb\x2c\x36\x30\x2d\xea\x0f\x74\x8b\xa2\xa5\x05\x6f\x29\xc9\x06\x36\xc3\x3f\x00\x9c\xcd\xee\x66\x07\xa1\x80\x4c\x04\xc5\x7a\x4f\x45\xe0\x77\x7d\xf0\x70\xb8\x46\x3d\x7f\xf7\x40\x77\x77\xbf\x5d\x26\x8f\x39\x4b\x27\x6f\x5e\x1f\xce\x3f\xac\xa8\x7b\x75\xa8\x7b\xff\x6b\x05\x62\xdf\x26\xf3\x19\x00\x00\xff\xff\x66\x74\x74\xed\xb2\x01\x00\x00")

func wwwTemplatesHtmlNotfoundHtmlBytes() ([]byte, error) {
	return bindataRead(
		_wwwTemplatesHtmlNotfoundHtml,
		"www/templates/html/notfound.html",
	)
}

func wwwTemplatesHtmlNotfoundHtml() (*asset, error) {
	bytes, err := wwwTemplatesHtmlNotfoundHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "www/templates/html/notfound.html", size: 434, mode: os.FileMode(420), modTime: time.Unix(1576865520, 0)}
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
	"www/templates/html/alt.html": wwwTemplatesHtmlAltHtml,
	"www/templates/html/error.html": wwwTemplatesHtmlErrorHtml,
	"www/templates/html/example.html": wwwTemplatesHtmlExampleHtml,
	"www/templates/html/id.html": wwwTemplatesHtmlIdHtml,
	"www/templates/html/inc_alt_basics.html": wwwTemplatesHtmlInc_alt_basicsHtml,
	"www/templates/html/inc_common_css.html": wwwTemplatesHtmlInc_common_cssHtml,
	"www/templates/html/inc_common_footer.html": wwwTemplatesHtmlInc_common_footerHtml,
	"www/templates/html/inc_common_header.html": wwwTemplatesHtmlInc_common_headerHtml,
	"www/templates/html/inc_common_meta.html": wwwTemplatesHtmlInc_common_metaHtml,
	"www/templates/html/inc_id_basics.html": wwwTemplatesHtmlInc_id_basicsHtml,
	"www/templates/html/inc_id_status.html": wwwTemplatesHtmlInc_id_statusHtml,
	"www/templates/html/index.html": wwwTemplatesHtmlIndexHtml,
	"www/templates/html/notfound.html": wwwTemplatesHtmlNotfoundHtml,
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
				"alt.html": &bintree{wwwTemplatesHtmlAltHtml, map[string]*bintree{}},
				"error.html": &bintree{wwwTemplatesHtmlErrorHtml, map[string]*bintree{}},
				"example.html": &bintree{wwwTemplatesHtmlExampleHtml, map[string]*bintree{}},
				"id.html": &bintree{wwwTemplatesHtmlIdHtml, map[string]*bintree{}},
				"inc_alt_basics.html": &bintree{wwwTemplatesHtmlInc_alt_basicsHtml, map[string]*bintree{}},
				"inc_common_css.html": &bintree{wwwTemplatesHtmlInc_common_cssHtml, map[string]*bintree{}},
				"inc_common_footer.html": &bintree{wwwTemplatesHtmlInc_common_footerHtml, map[string]*bintree{}},
				"inc_common_header.html": &bintree{wwwTemplatesHtmlInc_common_headerHtml, map[string]*bintree{}},
				"inc_common_meta.html": &bintree{wwwTemplatesHtmlInc_common_metaHtml, map[string]*bintree{}},
				"inc_id_basics.html": &bintree{wwwTemplatesHtmlInc_id_basicsHtml, map[string]*bintree{}},
				"inc_id_status.html": &bintree{wwwTemplatesHtmlInc_id_statusHtml, map[string]*bintree{}},
				"index.html": &bintree{wwwTemplatesHtmlIndexHtml, map[string]*bintree{}},
				"notfound.html": &bintree{wwwTemplatesHtmlNotfoundHtml, map[string]*bintree{}},
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

