// Code generated by go-bindata.
// sources:
// templates/html/id.html
// DO NOT EDIT!

package html

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

var _templatesHtmlIdHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x59\x6d\x73\xdc\xb6\x11\xfe\x7c\xfa\x15\x5b\xda\x53\x7f\x31\x8f\x71\xe2\xe9\xb4\x0a\x8f\xd3\x54\x8e\xa7\x9a\xc4\x8d\xc6\x2f\x4d\xfb\x29\x03\x12\x7b\x24\x62\x10\x60\x00\xf0\x4e\xac\xab\xff\x9e\x01\x40\xde\x81\x77\xa4\x5e\x4e\x8e\x3d\x63\x4b\x04\x77\x9f\xdd\x7d\xb0\xd8\xc5\xd2\xe9\x9f\x5e\xfd\x74\xf1\xfe\xbf\x57\xdf\x43\x65\x6a\x9e\x9d\xa5\xf6\x07\x70\x22\xca\x55\x84\x22\xca\xce\x00\xd2\x0a\x09\xcd\xce\x16\x00\xa9\x61\x86\x63\xf6\xe9\x13\x2c\xdf\x5d\xbd\x5d\xfe\x8b\xd4\x08\x37\x37\xf0\x7f\x18\x56\x2e\xa9\x7f\xfe\xb9\x92\xcf\x34\xfc\x24\xe0\x35\x53\xda\xa4\x89\xd7\x73\x10\x35\x1a\x02\x95\x31\x4d\x8c\xbf\xb5\x6c\xb3\x8a\x2e\xa4\x30\x28\x4c\xfc\xbe\x6b\x30\x82\xc2\x3f\xad\x22\x83\xd7\x26\xb1\xce\x7c\x0b\x45\x45\x94\x46\xb3\xfa\xf0\xfe\x75\xfc\xd7\x28\x80\x11\xa4\xc6\x55\xa4\x70\x8d\x4a\xa1\x0a\x94\xa5\x62\x25\x13\xd1\x8c\xc5\xff\xc4\x1f\xbe\x8b\x2f\x64\xdd\x10\xc3\x72\x1e\x1a\xbd\xfc\x7e\xf5\xb7\x08\x92\x23\x13\xa4\x69\x38\xc6\xb5\xcc\x19\xc7\x78\x8b\x79\x4c\x9a\x26\x2e\x48\x43\xc6\xea\x1d\xea\x7b\x6b\x6b\x43\x4c\xab\xe3\x9c\xa8\x58\x9b\x6e\x04\x93\x73\x52\x7c\x9c\x02\xfa\x27\x11\xb4\x42\x4e\x5f\x2b\x86\x82\xf2\x2e\xa4\x4b\xb5\x38\xa5\xb2\x61\xb8\x6d\xa4\x32\x81\xe8\x96\x51\x53\xad\x28\x6e\x58\x81\xb1\x7b\x78\x0e\x4c\x30\xc3\x08\x8f\x75\x41\x38\xae\x5e\x3c\x87\x9a\x5c\xb3\xba\xad\x83\x05\x26\xc6\x0b\xad\x46\xe5\x9e\x2c\x09\x2b\x21\x9d\xf5\xbd\xf9\x46\xc9\x06\x95\xe9\x56\x91\xdd\xc4\x73\x4e\xb4\xa9\x25\x65\x6b\x86\x34\xf0\xc5\x26\xce\x8f\x44\x9b\x37\xfd\x2b\xb8\xb9\x19\xa2\x98\x82\x92\xe5\xb9\x19\xa7\x09\x51\x86\x15\xfc\x20\xf4\x91\x82\x66\x06\x7f\xb1\x64\x04\x5a\xe3\xfc\xbc\x45\xd9\x25\xee\x81\xbf\x61\xea\x33\x0d\x64\x97\xfd\x57\x9c\x14\x68\xfd\x73\x6f\xc4\x6e\xfd\x42\xb6\xc2\xa8\x6e\x1f\xda\x84\x21\x8a\xba\x50\xac\x31\x4c\x8a\x30\x1f\x8d\x3e\x38\x4b\x70\xf9\xca\x1a\x1d\x1d\xb8\x5b\x60\x59\x4d\xca\xd0\xff\x84\xd1\x64\xa4\xbb\xd4\x9b\x72\x76\xe7\xcc\x96\x19\x83\xea\xbc\x20\x8a\x46\xb0\x21\xbc\xc5\x55\xa4\xdb\xba\x26\xaa\x9b\x33\x3a\xe8\x58\xda\x03\xc3\x7f\x27\x9c\xcb\xb5\xa9\xb0\xb1\x24\xe9\xbb\xb4\x0b\x85\xc4\x48\x75\x3a\x40\xab\xf8\x6d\x61\xdf\xa5\xfe\x65\xf6\x7d\xb0\xf6\x47\x6c\xfe\x80\xfd\x80\x0c\x00\x0b\xc3\x99\xf8\x08\x0a\xf9\x2a\x72\x65\x49\x57\x88\x26\x82\x4a\xe1\x7a\x15\x25\x85\xd6\x49\x4d\x9a\xff\xa1\x58\xfe\xaa\x97\x85\xd6\x7d\x89\xf5\xee\x83\x56\xc5\x2a\x4a\x7e\x25\x1b\xe2\x17\x06\xd9\x9a\x59\xf9\x28\x4b\x13\xbf\xee\xb2\x6d\x56\x89\xcb\x82\xf0\xb5\x54\xa4\xc4\x63\x4d\x5f\x18\x66\x75\x35\x67\x4d\xd3\xd5\xa4\x59\x16\x4a\x6a\x5d\x11\xa6\xf4\xa1\x69\x58\xdc\x23\xca\x6d\x25\xb5\x14\x6b\xcb\xfa\x72\xbb\xdd\xf6\xb1\x3e\x58\xb5\x90\x75\x2d\x85\xd7\x5e\x58\xcf\xcf\x1e\x8a\xc0\xe8\x8e\xe7\xd9\xa8\x0f\xe4\xc7\xf1\x2e\xc0\xf6\xf0\xc4\x37\x71\x80\x34\x97\xb4\x03\x4a\x0c\x89\xfd\xee\xc4\xa4\x61\x1f\xb1\xf3\x39\xfe\xc6\x2d\x7d\x77\x75\xf9\x03\x76\x36\xb9\xac\x06\x80\xfb\x27\xa5\x6c\x03\x05\x27\x5a\xaf\xa2\xc0\x60\x04\x8c\x8e\x16\x62\x77\x40\x23\x6f\x22\x5c\xb7\x72\x07\xa9\x7b\x24\xd3\x10\x53\xed\xa5\xae\x88\xa9\xe6\xe4\x94\xbd\x33\x84\x90\x57\x6e\x29\x04\xe6\xc4\x30\xd3\x52\xdc\xcb\xfc\xd8\xaf\x04\x32\x52\x94\x87\x42\xc3\xd2\x5e\xaa\x66\x62\x02\xed\x0d\x13\x03\xe0\x58\xf2\x18\xd3\x8a\x4e\xc0\x92\xeb\x29\x58\x72\x7d\x0c\x6b\x25\x27\x60\xc9\xf5\x08\xd6\x6f\x57\xbf\x57\x96\x9c\x9a\x34\x36\x17\x28\xdb\xd8\xc4\xf7\xef\xaa\xaf\xb3\xb3\x45\xaa\x1b\x22\x26\xb6\x33\x76\x9d\xd2\x9d\xea\x83\x9a\x97\x26\x56\x25\x03\xaf\x39\xb9\xe9\x31\x25\x06\x7d\xf2\x39\x51\x57\x25\x67\x2d\x35\x43\xd9\x8c\xb2\xa9\x4a\xba\x03\x11\xf3\x10\x82\xd4\x6c\xdd\x01\xd3\x32\x2e\x7c\xb1\xdd\x63\xed\xaa\xaf\x2b\x19\x3d\xda\xf2\xcc\x86\x5e\x13\xce\x87\x10\xd9\x1a\xf0\xb7\x3e\x29\xf5\x45\xab\x6c\x12\x2d\xdf\x19\xc5\x44\xf9\x9a\x93\x12\xa2\x17\x11\xdc\xdc\xf8\x2a\x37\xe3\x45\x7f\x95\x2b\xbc\x72\x94\xfd\x8c\x90\x23\x67\xb8\x41\x30\x15\xd3\xa0\xb0\x90\x8a\x82\x91\x90\x23\xf4\x52\xcb\xde\xa1\xde\x0b\xe4\x1a\xef\x76\xe5\xab\xfb\xba\x22\xa4\x19\xb9\x23\x24\xd8\xec\x41\x75\x9a\x63\xde\xa8\x7d\x12\xf6\x78\x0d\x84\x3a\x1a\x17\x43\x5a\x25\x36\xaf\xfa\x87\xc5\x2e\x01\x43\xe7\x28\x1a\xc2\xb8\x8e\x26\xa9\x7f\x85\x8d\xc2\x82\x18\xa4\xa7\xb1\x4f\x77\xfa\x51\xf6\x3e\x08\xae\x22\x1a\x72\x44\x01\x35\x51\x1f\x91\x02\xd1\xb0\x17\x3d\x0c\xd6\x85\x77\xb6\x98\xf2\xef\x5d\xdb\xa0\xd2\x48\xe7\xfc\xb3\x7f\x27\xed\xea\x9d\x22\xe4\x1d\x98\x0a\x61\x2d\x39\x97\x5b\x26\xca\x5e\x56\x9f\x5b\x65\x80\xb4\xe5\x93\xe7\x44\x32\x61\x50\xf5\xbc\xd9\x3f\x9f\x3e\x81\x22\xa2\x44\x78\xca\x9e\xc3\x53\x0d\xe7\x2b\xef\xe5\xde\xc7\x7f\x74\xbd\x5b\x0e\x97\xb3\x2c\x25\x43\x83\xf1\xfd\xff\xa9\x76\x65\x68\xc2\x1c\xa3\x70\x7c\xc4\xdc\xb1\x72\x3a\x69\x42\xb2\x34\xe1\x2c\x74\x66\xe0\xcd\xe7\x41\xcb\xfb\x77\x4b\xff\x63\x4c\xae\x65\x77\x8e\x5b\x26\xca\x29\x72\x3d\x4c\x48\xee\x8e\x53\x7d\x0b\xa1\xf7\xa2\xf3\x4e\x2a\xf5\x90\x7f\x7f\x18\x89\x23\x02\x77\xf4\x2d\x7d\x52\x05\x2f\xfb\x32\xbe\x3b\x70\x2d\x3f\x3a\x61\x39\xd1\xac\x70\x91\x59\x77\x67\x2f\x90\xb3\x67\x89\xd1\x7d\xfd\x74\x9d\x74\x28\xc2\x44\x50\x60\xc6\xf2\x6f\xfb\xd5\x06\xe1\xc3\xdb\x4b\x87\x34\x30\x62\xe7\x6b\x7d\x9e\x24\xb6\x5b\x2d\xc3\xfb\x88\x54\x65\x72\xd4\xcf\x27\x4c\x2b\xe4\xb6\xf9\x07\xbd\xc0\x0b\x5b\xaa\x96\x9e\xab\xc3\xc4\xd9\x75\xfc\x28\x7e\xe9\x33\xc5\x46\x1d\xe4\xc9\x33\x0d\xfe\xa2\x30\xc4\x5d\x48\x8a\x59\xfc\x32\x4d\xdc\x2f\xb0\xad\x58\x51\x01\x13\x94\xd9\x6a\xa0\x81\x19\x77\x6c\xeb\x96\x1b\xd6\x70\xec\x95\xed\x39\x2e\x48\xab\xf1\x38\xda\x92\x99\xaa\xcd\xed\x2d\x2f\xbc\x83\x25\xa3\x74\xf3\x37\x72\x86\x3a\x31\x0a\x31\xa9\x89\x36\xa8\x92\x60\x79\x2b\xd7\x4f\xbc\xa5\x5f\x2c\xff\xce\x89\xa6\x41\xa1\x41\xcb\x1a\x0d\xab\x51\x07\x24\x2c\x26\xba\x45\x40\xc4\x37\x0f\x21\xe2\x9b\x31\x11\x35\x12\xe1\x48\xa8\x59\x59\x19\xe0\x58\x32\xc3\x6a\x62\x90\x77\x50\x91\x0d\x1e\x13\xf3\x45\x08\x19\xd8\x6f\x50\x5a\xdb\x44\x95\x2d\x02\xc9\x65\x6b\x7c\x17\xd3\x52\x19\x90\x6b\xfb\x20\xca\xfb\x33\xf5\xf5\x43\x98\x7a\x31\xc5\xd4\x16\x1d\x2d\xe2\x99\x81\x5c\x9a\x0a\x15\x52\x30\xaa\xb3\x75\xc8\x48\x68\x98\x00\x2a\xb7\x22\x6c\xb5\xcf\x34\xb0\xba\x46\xca\x88\x19\x48\xfc\xb2\xc9\xc5\xd9\x1a\x6d\x58\x85\xac\x1b\xee\xb2\x9e\x06\xb5\xe8\x76\xc2\x5e\xdc\x4d\xd8\xe5\x31\x61\x96\x2a\xcb\xc8\xd6\xee\xdc\x17\x8a\xb2\x15\x05\x2a\x43\x98\xb0\xb1\xf5\x5b\xd6\x77\x8e\x21\x73\xdd\xa6\x48\x81\x87\xb9\x32\x44\x78\x69\x2c\x4d\x5e\xda\x77\xee\xc3\xe2\x3f\xa2\xe7\x41\x6d\xe0\xf6\xd1\x28\x80\xcc\x8e\xab\xf7\xb1\xd4\x61\xc6\xef\xee\x30\x43\x0b\xc8\x65\x2b\x6c\x6b\x85\x5c\x5e\x8f\x6a\xf6\x93\x49\x9f\x4b\xb4\x55\x47\x75\x30\xb5\x18\x0f\x60\xb9\xbc\xde\x97\xea\xd1\x20\xf4\x1c\x66\x86\x9e\xe0\x45\x38\xe1\x8c\x97\x03\x79\xb7\x79\x43\xe3\x69\x14\x13\x05\x6b\x08\x87\x02\x85\x51\x92\xd1\x47\x87\x32\x00\xed\xe3\x08\xc6\xc3\xbd\x57\x87\x2e\x0d\xcd\x68\xb8\xf6\xda\x7e\xed\x7f\x3f\x1c\xbe\x8e\x6f\xa9\x91\x6d\x62\x69\xf5\x32\x7b\xe7\x1e\xd3\xa4\x7a\x99\xb9\xad\x32\x24\x1f\x3e\x90\x1b\xd5\x5f\x9f\x52\x53\x65\x7f\x7e\xf2\xe2\x2f\x5f\x7d\x9b\x26\xa6\x0a\x16\x07\xe5\xd1\xe2\xbf\x09\x6f\xf1\x60\xed\xc2\x5a\xa7\x28\x8a\xfd\x8b\x34\xe9\xf1\xc7\x76\xfa\x69\x63\xa4\x4f\xe7\x6f\xdb\xc1\x2d\x61\x6a\x50\xb1\x44\x19\x7a\x2a\xd0\xa5\x7e\xaf\x5a\x7c\x34\xc8\x0f\xc2\x16\xdf\x00\x65\x2e\x74\x24\xda\x56\xc1\x13\x22\x77\x9a\x9f\x21\x70\x8f\xf3\xb8\xb8\x07\x8c\xfb\x86\xbd\x9f\xb7\x4e\x09\x7d\x7a\x5a\x3b\xcd\xf5\x00\x6b\x86\x82\x39\xe1\xa3\x58\x67\xa3\xdd\x4f\x46\xa7\x44\x3b\x3d\xfb\x9d\x16\x6d\x80\xf5\xa8\x0d\x1f\xe1\x3c\x9c\x08\x77\x55\x3a\x9d\x89\x83\x49\xed\x71\x21\x58\xb0\xcf\xc2\x85\x07\x9a\x3e\x01\x0b\xe8\x3f\x5b\xf4\xa5\x76\x28\xdf\xa3\x4f\x63\x7e\xc0\xb2\x05\xfa\x6a\x77\xb9\xf0\x45\xfa\x6c\xfa\x93\xc6\xfe\x0e\xd2\x7f\x09\xbf\x5d\x28\x56\x64\xbb\xfb\x1e\x77\x0f\xf1\x46\xa1\x31\xdd\xa0\xd1\x07\xe0\x94\x47\xee\x1f\xb5\x9f\x5d\x9b\x91\xad\x2a\xd0\xdd\x37\xf6\xbd\xa6\x1f\x30\x47\x03\xed\x83\xc6\xb7\xec\x21\xd2\xa3\x51\x77\xd2\xe6\xf4\x0d\x30\xb6\xe0\x3b\xb0\xb7\xd8\x48\xb8\xb9\x49\x72\x2e\xf3\xe1\xfe\x37\x12\x38\xf2\xed\x73\xa2\x06\x31\xf4\xf3\xf9\x31\xff\xc1\xaf\xee\x7b\x7b\x2e\x69\x97\x9d\xa5\x89\xff\x6f\xf5\xdf\x03\x00\x00\xff\xff\x50\x38\xbf\x80\x67\x1f\x00\x00")

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

	info := bindataFileInfo{name: "templates/html/id.html", size: 8039, mode: os.FileMode(420), modTime: time.Unix(1514052426, 0)}
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

