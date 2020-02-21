package static

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var _static_1_html = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\x51\x74\xf1\x77\x0e\x89\x0c\x70\x55\xc8\x28\xc9\xcd\xb1\xe3\xe5\xb2\x01\xd1\x0a\x39\x89\x79\xe9\xb6\x4a\xa9\x79\x4a\x60\x91\xd4\xc4\x14\x3b\x5e\x2e\x05\x05\x05\x05\x9b\xdc\xd4\x92\x44\x85\xe4\x8c\xc4\xa2\xe2\xd4\x12\x5b\xa5\xd0\x10\x37\x5d\x0b\x25\x98\x5c\x49\x66\x49\x4e\xaa\x5d\x08\x88\xb4\xd1\x87\x70\x78\xb9\x6c\xf4\xa1\xda\x6d\x92\xf2\x53\x2a\xa1\x4a\x3d\x52\x73\x72\xf2\xf5\xd2\xf3\x41\xd2\x50\x61\x1b\x7d\xb0\x03\x00\x01\x00\x00\xff\xff\xbe\x32\x0b\xd5\x90\x00\x00\x00")

func static_1_html() ([]byte, error) {
	return bindata_read(
		_static_1_html,
		"static/1.html",
	)
}

var _static_static_go = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func static_static_go() ([]byte, error) {
	return bindata_read(
		_static_static_go,
		"static/static.go",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"static/1.html": static_1_html,
	"static/static.go": static_static_go,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"static/1.html": &_bintree_t{static_1_html, map[string]*_bintree_t{
	}},
	"static/static.go": &_bintree_t{static_static_go, map[string]*_bintree_t{
	}},
}}
