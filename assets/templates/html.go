// Code generated by go-bindata.
// sources:
// www/templates/html/alt.html
// www/templates/html/alt.html~
// www/templates/html/error.html
// www/templates/html/example.html
// www/templates/html/id.html
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

var _wwwTemplatesHtmlAltHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x57\x4d\x8f\xe2\x46\x10\x3d\xc3\xaf\xe8\xb4\x14\xe5\x84\x1d\x69\x4f\x99\x80\x13\x69\x67\x23\x8d\xb4\x9b\x45\xd9\x48\x51\x4e\xa3\x1a\x77\x81\x6b\xd2\x1f\x56\x77\x01\x43\x08\xff\x3d\x6a\x6c\x83\x0d\x78\x06\x42\x92\xb9\x8c\xbb\x5d\xef\xf5\xab\x8f\x2e\x17\x9b\x8d\x50\x38\x23\x8b\x42\x82\x66\x29\xb6\xdb\xe1\xf8\xab\xfb\xcf\xef\x7f\xfd\x7d\xfa\x41\x14\x6c\x74\x36\x1c\xc7\x7f\x42\x83\x9d\x4f\x24\x5a\x99\x0d\x85\x18\x17\x08\x2a\x3e\xc4\xbf\x31\x13\x6b\xcc\x36\x1b\x91\x7c\x99\xfe\x92\xfc\x0c\x06\xc5\x76\x2b\xfe\x12\xcd\xce\x83\xaa\xd6\xbf\x15\xee\x9b\x20\x3e\x5b\xf1\x13\xf9\xc0\xe3\xb4\xc2\xd5\x2c\x9b\x8d\x60\x34\xa5\x06\x46\x21\xc9\xe6\x8f\xb9\x33\xc6\xd9\x47\x83\x0c\x3b\x59\xf5\x61\x71\x2d\x4a\xef\x4a\xf4\xbc\x9e\xc8\xa8\xed\x4e\x43\x60\xe3\x14\xcd\x08\x95\x14\xb9\xb3\x8c\x96\x27\x32\x9e\xff\x11\x02\x7f\xaa\x5f\x89\xed\x56\x8a\x34\x1b\x0e\x84\xe8\x61\x73\xf3\x3b\x5e\x97\xd8\xe2\x00\xcf\x94\x6b\xdc\xe1\x7a\x31\x81\x18\x1f\x2d\x98\x36\xb0\xeb\xec\xeb\xf8\x5d\x20\x8e\x84\xb7\x43\x49\x41\xc0\x3e\x9a\x53\x0d\x39\x46\x95\xbb\x37\x76\xbf\xff\xde\x2d\x2c\xfb\x75\xed\x63\xef\x59\x0a\x43\xee\xa9\x64\x72\xb6\x75\xe2\x03\x87\xa3\xf4\x88\x87\xfb\x78\x6e\x27\x87\xaf\x33\x93\x81\xf9\xb1\x17\x1f\xac\x2a\x1d\x59\x0e\xc9\xd4\xce\xc5\x76\xdb\xa1\xfb\x21\xd0\x9f\x38\x79\x09\xa6\xc5\xdb\x43\xcf\x2b\x62\x46\x7f\x97\x83\x57\x52\x2c\x41\x2f\x70\x22\xc3\xc2\x18\xf0\xeb\x57\x54\x35\xb0\x98\xa0\x96\xb2\x1f\x41\x6b\x37\xe3\x02\xcb\x18\xcb\x70\x01\x41\xee\x11\xd8\xf9\x9b\x38\x16\x5e\xb7\xf0\x29\xa9\xf4\xd2\xe0\x36\x0c\xff\x5b\x9d\x34\x07\xfe\x47\xc5\xd2\xd0\xff\xab\x15\x53\x29\x15\xc1\xe7\x13\x99\x3e\xc3\x12\xaa\x8d\x54\xbb\x1c\xf4\xcc\x79\x98\x63\x62\xc8\x26\xcf\x41\x66\xe3\xb4\x7a\x99\x0d\xde\xc6\x07\x4d\x65\xb9\x36\x50\x26\xb9\x77\x21\x14\x40\x3e\x74\x49\x2e\x14\xb2\x2a\x5c\x70\x76\x16\xe3\x95\xac\x56\xab\x63\x21\xd7\xc0\x3d\x5a\x85\xfe\xac\x8a\x8b\xf0\x75\x2a\x08\x6f\xf7\x24\x87\xbc\xc0\x5b\x7c\x59\x78\xfa\xe7\x8e\x58\xe4\x1b\xc0\x60\x68\xb6\xee\xe2\xc5\x40\xec\xe5\x5f\xc3\x35\x47\xf7\x1c\xdc\x49\x75\x5d\x43\xa1\x11\x66\x1a\x39\x59\x30\xe9\xf3\x79\xb9\x8a\x26\xf0\x5a\xf7\xe4\xf7\x2a\x9e\x02\xac\xd2\x78\x5c\xf3\x55\x9c\xae\xe1\x7b\xf2\x6e\x15\xd0\x27\x06\xca\x1b\x54\x35\x2c\xd5\x74\x70\x72\x99\xaf\xce\x5b\xc3\x07\xfa\x86\x42\x6a\x93\x90\x25\xee\xa9\x82\x37\x26\x9d\x3c\x84\xdd\xa0\xd3\x9c\xac\xc9\xfe\x21\x3c\xea\x89\xac\x52\x59\x20\xb2\x14\x85\xc7\xd9\x44\xa6\x79\x08\x67\x25\x90\x4a\x22\x4f\xeb\x26\x8f\xd3\x66\x58\x1b\x3f\x39\xb5\x16\x0a\x18\x46\x2d\xe8\x68\xe1\x69\x84\x75\xd7\x3d\xee\xc1\xf7\xc0\x10\x7b\x79\xb6\x57\xa5\x68\x29\x72\x0d\x21\x4c\x64\xec\xdb\x40\x16\xbd\x7c\x6b\x88\x8b\x0a\xd0\x4b\x91\xb4\xdc\x33\x40\xb6\x61\x6a\xc9\x91\x82\x54\x67\x63\xb4\xfb\xc0\xca\x53\xdd\xd1\xee\xe8\x8b\x73\x62\x53\x02\x17\x07\xab\x29\x70\xd1\x67\xe7\xd1\x76\x29\xa7\xbb\xad\x36\xb1\x06\x26\x5e\x28\x3c\xd8\x7c\xac\x77\x5a\x36\xce\xce\x8f\x8d\x9a\xad\x83\x95\x21\x7b\x86\xed\x13\xd9\x86\xb0\x6b\x79\xca\x19\x4d\xcf\xd0\xc2\xcb\x39\x5a\x78\x39\xa5\x8d\x96\x67\x68\xe1\xa5\x43\x9b\x0d\xe3\xb0\xdc\x97\xf3\xe1\xa0\x55\x13\x31\x74\x06\x4a\xd9\x18\x7a\xb7\x92\xd9\x70\x10\xe1\x64\xe6\xd5\xfd\xe9\xd6\xd6\x97\xe5\xc9\xf7\x5d\x8a\x5d\xa9\x4f\xe4\x8a\x14\x17\x77\x64\x0b\xf4\xc4\xdf\x8b\x02\x69\x5e\xf0\x61\x1d\xe5\xd7\x7b\xdf\x7d\xfb\x75\xb5\xae\x20\x71\x29\x1b\x31\xa3\xb0\x9c\x37\x13\xff\x4e\x68\xaa\x68\x79\xa4\xfb\x54\x6e\xf1\xae\x7a\xa8\x0b\xba\x3d\x61\x55\x06\x69\xf1\x2e\x8b\x0f\x47\xb4\xf5\xaa\x97\xbd\x82\x0c\xce\xdc\x12\x52\x8f\x4f\x10\x28\x0f\xd5\x0d\xb9\x8e\x76\x58\x89\xea\xe4\x48\xcb\x83\x0b\xfb\xec\x74\xea\x7d\x3f\x00\xec\x13\x16\x30\xe7\x1d\x6c\xf0\x26\x66\x54\x7a\x64\x5e\xc7\x1e\x17\x25\x56\x7e\x5d\x00\xf3\xb0\x6a\x30\x07\x79\x87\x65\xeb\xf1\x4c\xc6\xea\xc7\x6e\x8b\x4f\x63\x17\xd9\x37\xa7\xbe\xf6\x33\x73\x8e\x3b\xed\xa7\xe2\x6a\xb5\xc8\xd8\x19\xb3\xe1\x38\xad\x7e\xf1\x6e\x36\x02\x6d\xac\xc7\xe1\xdf\x01\x00\x00\xff\xff\x32\x4f\x44\xaf\x1f\x0f\x00\x00")

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

	info := bindataFileInfo{name: "www/templates/html/alt.html", size: 3871, mode: os.FileMode(420), modTime: time.Unix(1579283568, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wwwTemplatesHtmlAltHtml2 = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x57\x5b\x6f\xe3\x36\x13\x7d\xb6\x7f\x05\x3f\x02\x1f\xfa\x64\xab\xc0\x3e\x35\xb5\xd5\x02\x9b\x2d\x10\x60\xb7\x6b\x74\x0b\x14\x7d\x0a\x26\xe2\xd8\x9a\x94\x17\x81\x1c\xdb\x71\x5d\xff\xf7\x82\xd6\xc5\x92\x6c\x25\xf1\xa6\x6d\x5e\x22\x52\x73\xce\x9c\xb9\x70\x44\xef\xf7\x42\xe1\x92\x2c\x0a\x49\x4a\x8a\xc3\x61\x3c\xfb\xdf\xed\xe7\xf7\xbf\xfe\xbe\xf8\x20\x72\x36\x3a\x1d\xcf\xe2\x3f\xa1\xc1\xae\xe6\x12\xad\x4c\xc7\x42\xcc\x72\x04\x15\x1f\xe2\xdf\x8c\x89\x35\xa6\xfb\xbd\x98\x7e\x59\xfc\x32\xfd\x19\x0c\x8a\xc3\x41\xfc\x25\xea\x9d\x3b\x55\xae\x7f\xcb\xdd\x37\x41\x7c\xb6\xe2\x27\xf2\x81\x67\x49\x89\xab\x58\xf6\x7b\xc1\x68\x0a\x0d\x1c\x95\xd8\xec\x3e\x73\xc6\x38\x7b\x6f\x90\xe1\x28\xab\x72\x16\xd7\xa2\xf0\xae\x40\xcf\xbb\xb9\x8c\xda\x6e\x34\x04\x36\x4e\xd1\x92\x50\x49\x91\x39\xcb\x68\x79\x2e\xa3\xff\x8f\x10\xf8\x53\xf5\x4a\x1c\x0e\x52\x24\xe9\x78\x24\xc4\x00\x9b\x5b\xdd\xf0\xae\xc0\x16\x07\x78\xa6\x4c\xe3\x11\x37\x88\x09\xc4\x78\x6f\xc1\xb4\x81\xdd\x60\x9f\xc7\x1f\x13\xd1\x13\xde\x4e\x25\x05\x01\x4d\x36\x17\x1a\x32\x8c\x2a\x8f\x6f\x6c\xb3\xff\xde\xad\x2d\xfb\x5d\x15\xe3\xa0\x2f\x85\x21\xf3\x54\x30\x39\xdb\xf2\x78\xc7\xa1\x57\x1e\x71\x77\x1b\xfd\x76\x6a\xf8\x3c\x33\x19\x58\xf5\xa3\xf8\x60\x55\xe1\xc8\x72\x98\x2e\xec\x4a\x1c\x0e\x1d\xba\x1f\x02\xfd\x89\xf3\xa7\x60\x5a\xbc\x03\xf4\xbc\x25\x66\xf4\x37\x19\x78\x25\xc5\x06\xf4\x1a\xe7\x32\xac\x8d\x01\xbf\x7b\x46\x55\x0d\x8b\x05\x6a\x29\xfb\x11\xb4\x76\x4b\xce\xb1\x88\xb9\x0c\xaf\x20\xc8\x3c\x02\x3b\xff\x26\x8e\xb5\xd7\x2d\x7c\x42\x2a\x79\x6d\x72\x6b\x86\xff\xac\x4f\x6a\x87\xff\x52\xb3\xd4\xf4\xff\x68\xc7\x94\x4a\x45\xf0\xd9\x5c\x26\x8f\xb0\x81\x72\x23\xd1\x2e\x03\xbd\x74\x1e\x56\x38\x35\x64\xa7\x8f\x41\xa6\xb3\xa4\x7c\x99\x8e\x5e\xc6\x07\x4d\x45\xb1\x33\x50\x4c\x33\xef\x42\xc8\x81\x7c\xe8\x92\xbc\x52\xc8\x36\x77\xc1\xd9\x65\xcc\xd7\x74\xbb\xdd\xf6\x85\x5c\x03\xf7\x68\x15\xfa\x8b\x2a\x5e\x85\xaf\x4a\x41\xf8\xf6\x48\x32\xc8\x72\x7c\x4b\x2c\x6b\x4f\x5f\x1f\x88\x45\x7e\x03\x18\x0c\x2d\x77\x5d\xbc\x18\x89\x46\xfe\x35\x5c\x2b\x74\x8f\xc1\x9d\x75\xd7\x35\x14\x1a\x61\xa9\x91\xa7\x6b\x26\x7d\xb9\x2e\x57\xd1\x04\xde\xe9\x81\xfa\x5e\xc5\x93\x83\x55\x1a\xfb\x3d\x5f\xe6\xe9\x1a\xbe\x07\xef\xb6\x01\xfd\xd4\x40\x11\xde\x92\xa6\x9a\x87\xd4\xd7\xc7\xd6\xe2\x20\x4b\x3c\x20\xe7\x85\x9b\x49\x16\xc2\xf1\x62\x52\x3b\xd6\x64\xff\x10\x1e\xf5\x5c\x96\xa9\xcf\x11\x59\x8a\xdc\xe3\x72\x2e\x93\x2c\x84\x21\x05\x91\xa7\x75\xf2\x66\x49\x7d\xb9\x9a\x3d\x38\xb5\x13\x0a\x18\x26\x2d\xe8\x64\xed\x69\x82\xd5\x94\xec\xcf\xcc\x5b\x60\x88\xb3\x37\x6d\x54\x29\xda\x88\x4c\x43\x08\x73\x19\xe7\x2c\x90\x45\x2f\x5f\xba\x74\x45\x05\xe8\xa5\x98\xb6\xc2\x33\x40\xb6\x66\x6a\xc9\x91\x82\x54\x67\x63\x72\xfc\x20\xca\x73\xdd\xd1\xae\xf7\x85\x38\xb3\x29\x80\xf3\x93\xd5\x02\x38\x1f\xb2\xf3\x68\xbb\x94\x8b\xe3\x56\x9b\x58\x03\x13\xaf\x15\x9e\x6c\x3e\x56\x3b\x2d\x1b\x67\x57\x7d\xa3\x7a\xeb\x64\x65\xc8\x5e\x60\xfb\x44\xb6\x26\xec\x5a\x9e\x73\x46\xd3\x0b\xb4\xf0\x74\x89\x16\x9e\xce\x69\xa3\xe5\x05\x5a\x78\xea\xd0\xa6\xe3\x78\xb9\x1d\xaa\x79\x7c\xd7\xf4\x44\x4c\x9d\x81\x42\xd6\x86\xde\x6d\x65\x3a\x1e\x45\x38\x99\x55\x79\x7c\xba\xbd\xf5\x65\x73\xf6\x3d\x96\xe2\xd8\xea\x73\xb9\x25\xc5\xf9\x0d\xd9\x1c\x3d\xf1\xf7\x22\x47\x5a\xe5\x7c\x5a\x47\xf9\xd5\xde\x77\xdf\xfe\xbf\x5c\x97\x90\xb8\x94\xb5\x98\x49\xd8\xac\xea\x1b\xfa\x51\x68\xa2\x68\xd3\xd3\x7d\x2e\x37\x7f\x57\x3e\x54\x0d\xdd\xbe\x11\x95\x06\x49\xfe\x2e\x8d\x0f\x3d\xda\x6a\x35\xc8\x5e\x42\x46\x17\x4e\x09\xa9\xfb\x07\x08\x94\x85\xf2\x84\x5c\x47\x3b\x2e\x45\x75\x6a\xa4\x65\x27\x84\x33\x6f\x81\x81\xd7\xa5\xb7\x51\xc3\x3f\x3a\xb9\xac\x95\x0e\xb3\x36\x35\xef\x9c\xa2\xe6\x1a\xd0\xb4\x41\xc0\x8c\x8f\xb0\xd1\x8b\x98\x49\xe1\x91\x79\x17\x27\x67\x54\x51\x6a\x78\x05\xcc\xc3\xb6\xc6\x9c\xe4\x9d\x96\x9d\x98\xce\xfa\xa0\x7a\xec\x7e\xa0\x93\x38\x9b\x9a\x91\x37\x34\xd4\x96\xce\x71\x67\xa8\x95\x5c\xad\xc1\x1b\xe7\x6d\x3a\x9e\x25\xe5\xef\xde\xfd\x5e\xa0\x8d\x5d\x3e\xfe\x3b\x00\x00\xff\xff\x92\x9e\x9b\x24\x24\x0f\x00\x00")

func wwwTemplatesHtmlAltHtml2Bytes() ([]byte, error) {
	return bindataRead(
		_wwwTemplatesHtmlAltHtml2,
		"www/templates/html/alt.html~",
	)
}

func wwwTemplatesHtmlAltHtml2() (*asset, error) {
	bytes, err := wwwTemplatesHtmlAltHtml2Bytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "www/templates/html/alt.html~", size: 3876, mode: os.FileMode(420), modTime: time.Unix(1576865520, 0)}
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

var _wwwTemplatesHtmlIdHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x57\x4b\x6f\x22\x47\x10\x3e\xc3\xaf\xe8\xb4\x14\xe5\x04\x13\x69\x4f\x71\x60\x12\x69\xbd\x91\x2c\xed\x66\xd1\x6e\xa2\x28\x27\xab\x3c\x5d\x30\xe5\xf4\x63\xd4\x5d\x80\x09\xe1\xbf\x47\x3d\x0f\x98\x01\xc6\x36\x4b\xb2\xbe\x78\xba\xa7\xbe\xaf\xbe\x7a\x74\x4d\xb3\xdd\x0a\x85\x73\xb2\x28\x24\x29\x29\x76\xbb\xe1\xe4\x9b\xdb\x8f\x6f\x7f\xfb\x73\xf6\x4e\xe4\x6c\x74\x3a\x9c\xc4\x7f\x42\x83\x5d\x4c\x25\x5a\x99\x0e\x85\x98\xe4\x08\x2a\x3e\xc4\xbf\x09\x13\x6b\x4c\xb7\x5b\x31\xfe\x3c\xfb\x34\xfe\x15\x0c\x8a\xdd\x4e\xfc\x23\x9a\x9d\x3b\x55\xad\xff\xc8\xdd\x77\x41\x7c\xb4\xe2\x17\xf2\x81\x27\x49\x85\xab\x59\xb6\x5b\xc1\x68\x0a\x0d\x1c\x95\xd8\xec\x3e\x73\xc6\x38\x7b\x6f\x90\xa1\x94\x55\x3b\x8b\x6b\x51\x78\x57\xa0\xe7\xcd\x54\x46\x6d\x37\x1a\x02\x1b\xa7\x68\x4e\xa8\xa4\xc8\x9c\x65\xb4\x3c\x95\xd1\xff\x7b\x08\xfc\xa1\x7e\x25\x76\x3b\x29\x92\x74\x38\x10\xa2\x87\xcd\x2d\x6e\x78\x53\x60\x8b\x03\x3c\x53\xa6\xb1\xc4\xf5\x62\x02\x31\xde\x5b\x30\x6d\x60\x37\xd8\xe7\xf1\x65\x22\x8e\x84\xb7\x53\x49\x41\xc0\x3e\x9b\x33\x0d\x19\x46\x95\xe5\x1b\xbb\xdf\x7f\xeb\x96\x96\xfd\xa6\x8e\xb1\xd7\x97\xc2\x90\x79\x2a\x98\x9c\x6d\x79\xbc\xe3\x70\x54\x1e\x71\x77\x1b\xfd\x76\x6a\xf8\x3c\x33\x19\x58\x1c\x47\xf1\xce\xaa\xc2\x91\xe5\x30\x9e\xd9\x85\xd8\xed\x3a\x74\x3f\x05\xfa\x1b\xa7\x4f\xc1\xb4\x78\x7b\xe8\x79\x4d\xcc\xe8\x6f\x32\xf0\x4a\x8a\x15\xe8\x25\x4e\x65\x58\x1a\x03\x7e\xf3\x8c\xaa\x06\x16\x0b\xd4\x52\xf6\x33\x68\xed\xe6\x9c\x63\x11\x73\x19\x5e\x41\x90\x79\x04\x76\xfe\x2a\x8e\xa5\xd7\x2d\x7c\x42\x2a\x79\x6d\x72\x1b\x86\xaf\xd6\x27\x8d\xc3\xff\xa9\x59\x1a\xfa\xff\xb4\x63\x2a\xa5\x22\xf8\x6c\x2a\x93\x47\x58\x41\xb5\x91\x68\x97\x81\x9e\x3b\x0f\x0b\x1c\x1b\xb2\xe3\xc7\x20\xd3\x49\x52\xbd\x4c\x07\x2f\xe3\x83\xa6\xa2\xd8\x18\x28\xc6\x99\x77\x21\xe4\x40\x3e\x74\x49\x5e\x29\x64\x9d\xbb\xe0\xec\x3c\xe6\x6b\xbc\x5e\xaf\x8f\x85\x5c\x02\xf7\x68\x15\xfa\xb3\x2a\x5e\x85\xaf\x4b\x41\x78\x7d\x24\x19\x64\x39\x5e\x13\xcb\xd2\xd3\x97\x07\x62\x91\xaf\x00\x83\xa1\xf9\xa6\x8b\x17\x03\xb1\x97\x7f\x09\xd7\x02\xdd\x63\x70\x27\xdd\x75\x09\x85\x46\x98\x6b\xe4\xf1\x92\x49\x9f\xaf\xcb\x45\x34\x81\x37\xba\xa7\xbe\x17\xf1\xe4\x60\x95\xc6\xe3\x9e\xaf\xf2\x74\x09\xdf\x83\x77\xeb\x80\x7e\x6c\xa0\x08\xd7\xa4\xa9\xe1\x21\xf5\xe5\xb1\xb5\x38\xc8\x12\xf7\xc8\x79\xe1\x66\x92\x85\x50\x5e\x4c\x1a\xc7\x9a\xec\x5f\xc2\xa3\x9e\xca\x2a\xf5\x39\x22\x4b\x91\x7b\x9c\x4f\x65\x92\x85\xd0\xa7\x20\xf2\xb4\x4e\xde\x24\x69\x2e\x57\x93\x07\xa7\x36\x42\x01\xc3\xa8\x05\x1d\x2d\x3d\x8d\xb0\x9e\x92\xc7\x33\xf3\x16\x18\xe2\xec\x4d\xf7\xaa\x14\xad\x44\xa6\x21\x84\xa9\x8c\x73\x16\xc8\xa2\x97\x2f\x5d\xba\xa2\x02\xf4\x52\x8c\x5b\xe1\x19\x20\xdb\x30\xb5\xe4\x48\x41\xaa\xb3\x31\x2a\x3f\x88\xf2\x54\x77\xb4\x3b\xfa\x42\x9c\xd8\x14\xc0\xf9\xc1\x6a\x06\x9c\xf7\xd9\x79\xb4\x5d\xca\x59\xb9\xd5\x26\xd6\xc0\xc4\x4b\x85\x07\x9b\xf7\xf5\x4e\xcb\xc6\xd9\xc5\xb1\x51\xb3\x75\xb0\x32\x64\xcf\xb0\x7d\x20\xdb\x10\x76\x2d\x4f\x39\xa3\xe9\x19\x5a\x78\x3a\x47\x0b\x4f\xa7\xb4\xd1\xf2\x0c\x2d\x3c\x75\x68\xd3\x61\xbc\xdc\xf6\xd5\x3c\xbe\xdb\xf7\x44\x4c\x9d\x81\x42\x36\x86\xde\xad\x65\x3a\x1c\x44\x38\x99\x45\x75\x7c\xba\xbd\xf5\x79\xd5\x7c\x8f\x7f\xff\x74\x57\x06\x51\xf6\xf9\x54\xae\x49\x71\x7e\x43\x36\x47\x4f\xfc\xa3\xc8\x91\x16\x39\x1f\xd6\x51\x7b\xbd\xf7\xc3\xf7\xdf\x56\xeb\x0a\x12\x97\xb2\x51\x32\x0a\xab\x45\x73\x3d\x2f\x55\x26\x8a\x56\x47\xa2\x4f\xb5\xe6\x6f\xaa\x87\xba\x9b\xdb\xd7\xa1\xca\x20\xc9\xdf\xa4\xf1\xe1\x88\xb6\x5e\xf5\xb2\x57\x90\xc1\x99\x23\x42\xea\xfe\x01\x02\x65\xa1\x3a\x1e\x97\xd1\x0e\x2b\x51\x9d\x02\x69\xd9\x09\xe1\xc4\x5b\x60\xe0\x65\xe5\x6d\xb0\xe7\x1f\x1c\x5c\x36\x4a\xfb\x59\xf7\x05\xef\x1c\xa1\xfd\x1d\x60\xdf\x03\x01\x33\x2e\x61\x83\x17\x31\xa3\xc2\x23\xf3\x26\x8e\xcd\xa8\xa2\xd2\xf0\x0a\x98\x87\x75\x83\x39\xc8\x3b\x2c\x3b\x31\x9d\xf4\x41\xfd\xd8\xfd\x3a\x27\x71\x30\xed\xe7\x5d\xdf\x44\x9b\x3b\xc7\x9d\x89\x56\x71\xb5\xa6\x6e\x1c\xb6\xe9\x70\x92\x54\x3f\x7a\xb7\x5b\x81\x36\xce\x92\xe1\xbf\x01\x00\x00\xff\xff\xee\xe5\x05\x44\x21\x0f\x00\x00")

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

	info := bindataFileInfo{name: "www/templates/html/id.html", size: 3873, mode: os.FileMode(420), modTime: time.Unix(1579284527, 0)}
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

var _wwwTemplatesHtmlInc_id_statusHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x95\x41\x8f\x9b\x30\x10\x85\xcf\xf0\x2b\x2c\xdf\x13\x0e\x7b\x6b\x1d\x5f\xba\xad\xb4\xea\x65\xd5\xf4\xbe\x72\xe3\x09\xb8\x32\x36\xf2\x18\xb2\x11\xe2\xbf\x57\xd8\x90\x92\xdd\xa4\x69\x60\x4f\x49\x9c\x79\x1f\x7e\x6f\xd0\x4c\xdb\x4a\xd8\x2b\x03\x84\x2a\xb3\x7b\x51\xf2\x05\xbd\xf0\x35\xd2\xae\x4b\x53\x26\x55\x43\x94\xdc\xd0\x43\x61\xd1\x9a\xbd\x72\xe8\x57\x07\xbb\xa7\x64\xa7\x05\xe2\x86\x22\xec\x3c\xe5\x29\xcb\xa4\x6a\xf8\x95\xfa\xc8\x5b\xbd\xbe\x15\xa5\x84\x10\xc2\x24\x78\xa1\x34\x12\x5b\x81\xe1\x69\xc2\xb0\x2e\x4b\xe1\x8e\x3c\x4d\xc2\xdf\xc5\x03\x41\x7f\xd4\xb0\xa1\x52\x61\xa5\xc5\xf1\x93\x32\x5a\x19\xf8\x4c\xc3\x73\x5e\x07\x3a\xe5\xf1\x93\x65\xc5\x43\x4f\xc9\xfe\x62\xc6\x27\xac\x4a\x30\xf5\x80\x1d\xe0\xfd\x6d\x87\x3b\x79\xf1\x4b\xc3\xca\x01\x56\xd6\xa0\x6a\x80\xf2\x34\x49\x58\x38\x3d\x2b\x21\xb1\x10\xcb\x50\x10\x28\xde\xf5\x5f\x13\xe6\x0b\xae\x90\xec\x6a\xe7\xc0\x78\x96\xf9\x62\x38\x96\x23\xe0\x7d\x28\x94\xb7\x2d\x59\x6f\x9f\x7f\xac\x9f\xf0\x4b\x14\xae\xb7\xde\x29\x93\x7f\xd3\x22\x27\x5d\xc7\x32\x2f\xe7\x61\x9e\xf0\xa7\xab\x61\x21\xe2\xbb\xb1\x07\x33\x61\x04\xbf\x59\x34\x7c\xc9\x3b\x08\x04\x39\xc3\x7a\xd0\x2d\x76\x1e\x29\x4b\x8c\x8f\x84\x3b\x7d\x4b\xa8\x1c\xec\x84\x9f\xe3\xfd\xf1\xa4\x5d\xea\x7f\x42\xba\x98\xc1\xb5\xd2\x3b\xcd\x62\x5d\x81\x43\x90\x73\xcc\x6e\x4f\xda\xa5\x66\x27\xa4\x05\x0d\x3f\xa3\xcc\xcc\x41\x99\x7c\x7e\x10\xca\xe4\x1f\x95\x44\x8f\xfa\x80\x28\x22\xe6\x6a\x16\x09\x49\xfa\xf1\x99\xb0\x2c\x0c\xc2\x71\x4c\xc7\xf9\x3f\xcc\x56\x96\x9d\x8f\xdc\xa1\x22\x9e\xc5\x9f\xff\xde\x18\x7b\xeb\x4a\xe1\xf1\xed\xc2\x98\xee\x8b\xbb\x56\x05\x1f\x80\xb7\xb7\xc3\x64\x39\xd4\x3a\xac\x00\xad\x38\x13\xa4\x70\xb0\xdf\xd0\x3e\xad\xaf\x46\x56\x56\x19\x8f\xeb\x67\xd3\x37\xec\x94\xa0\x24\x5d\x47\x79\xd5\xbf\x0d\x82\xb3\x4c\xab\x1b\xf2\x6d\x73\x41\x8e\xcd\xff\xca\x1f\x85\x17\xef\xf5\x39\xd8\xdf\x68\xcd\x84\x11\xc3\x0f\x66\x6e\xf4\x65\x6c\x49\xdb\x82\x91\x5d\x97\xfe\x09\x00\x00\xff\xff\x3b\x14\x81\xdd\x19\x08\x00\x00")

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

	info := bindataFileInfo{name: "www/templates/html/inc_id_status.html", size: 2073, mode: os.FileMode(420), modTime: time.Unix(1576865520, 0)}
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
	"www/templates/html/alt.html~": wwwTemplatesHtmlAltHtml2,
	"www/templates/html/error.html": wwwTemplatesHtmlErrorHtml,
	"www/templates/html/example.html": wwwTemplatesHtmlExampleHtml,
	"www/templates/html/id.html": wwwTemplatesHtmlIdHtml,
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
				"alt.html~": &bintree{wwwTemplatesHtmlAltHtml2, map[string]*bintree{}},
				"error.html": &bintree{wwwTemplatesHtmlErrorHtml, map[string]*bintree{}},
				"example.html": &bintree{wwwTemplatesHtmlExampleHtml, map[string]*bintree{}},
				"id.html": &bintree{wwwTemplatesHtmlIdHtml, map[string]*bintree{}},
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

