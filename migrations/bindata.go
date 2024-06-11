// Code generated by go-bindata. DO NOT EDIT.
// sources:
// migrations/001_initial.sql (672B)

package migrations

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	clErr := gz.Close()
	if clErr != nil {
		return nil, clErr
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var __001_initialSql = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x92\x4f\x8f\x9b\x30\x10\xc5\xef\x91\xf2\x1d\xde\x61\x0f\x89\xd4\x44\xdb\x33\xea\x81\xc0\x90\xa2\x12\x83\x8c\x51\x76\x7b\x41\x56\xec\x52\x4b\x60\x58\xfe\x68\xfb\xf1\x2b\x67\x29\x6d\xaa\x36\x7b\x7e\xbf\x99\x37\x7a\x6f\x76\x3b\x3c\x3e\x7e\x2c\x8d\x35\xa3\x91\xf5\x7e\xea\xf6\xc3\x4b\xbd\x5e\x05\x9c\x7c\x41\x88\x0a\x16\x88\x38\x65\x98\x3a\x25\x47\x5d\x36\xad\x32\xdf\x8c\x56\x9b\x2d\x38\x89\x82\xb3\x1c\x82\xc7\xc7\x23\x71\xf8\x39\x1e\x1e\x70\xa0\x63\xcc\xc0\xe8\xbc\xff\x85\xe2\x13\x6c\xfb\xba\xd9\x7a\xf3\x84\x13\x3d\x10\x0b\x3d\xc7\xd7\xd2\x56\x93\xac\x34\xba\xba\xab\x86\x97\xda\x5b\xaf\x16\x73\xe1\x1f\x12\xc2\x34\xe8\x7e\xc0\x66\xbd\x02\x8c\xc2\x34\x19\x85\x8c\xc7\x27\x9f\x3f\xe3\x0b\x3d\x23\xa4\xc8\x2f\x12\x81\x4a\xdb\xb2\x97\x56\xb5\x4d\xe9\x98\xcd\xf6\x83\x9b\x50\x66\xb8\xb4\xbd\x2a\x8d\xc2\xa8\x7f\x8c\x60\xa9\x00\x2b\x92\xe4\x46\x74\x0e\x56\x36\xfa\x0e\x32\xca\xea\x1f\xea\xa5\xd7\x72\xd4\x0a\x22\x3e\x51\x2e\xfc\x53\x26\xbe\x2e\xc0\x72\x18\x4b\xcf\xf3\x31\x4b\x22\xef\xf3\xeb\xd5\xf6\x26\x89\x39\xe3\xb9\x85\x6b\x24\x4b\x17\x38\x50\x94\x72\x42\x91\x85\x8e\x75\x65\x5d\x23\x8b\x52\x0e\xf2\x83\xcf\xe0\xe9\x19\xf4\x44\x41\x71\xb7\xd0\xab\xdf\x5f\xdf\xa0\xda\x57\xfb\xf6\x0f\xbb\x1d\x42\x9e\x66\x73\x27\x71\x04\x7a\x8a\x73\x91\xbf\x59\x79\xbf\xf5\x65\xff\x1f\xc8\xad\x93\xf7\xff\x65\xc3\xe5\xbb\x6e\x64\xd9\x98\xaa\x97\xa3\x69\xad\x5b\xfc\x33\x00\x00\xff\xff\x1d\x69\x32\xcc\xa0\x02\x00\x00"

func _001_initialSqlBytes() ([]byte, error) {
	return bindataRead(
		__001_initialSql,
		"001_initial.sql",
	)
}

func _001_initialSql() (*asset, error) {
	bytes, err := _001_initialSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "001_initial.sql", size: 672, mode: os.FileMode(0666), modTime: time.Unix(1717592782, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd0, 0x77, 0x52, 0x3d, 0x5a, 0xe7, 0xc0, 0xd2, 0xcb, 0x49, 0xbb, 0x41, 0x82, 0x91, 0x69, 0x5c, 0xd8, 0x8c, 0x4b, 0x1a, 0x69, 0xf2, 0x78, 0x36, 0x2e, 0xc3, 0x46, 0x2d, 0x5b, 0x9b, 0xf7, 0x1a}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"001_initial.sql": _001_initialSql,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"001_initial.sql": {_001_initialSql, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	err = os.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
