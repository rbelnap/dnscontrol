// Code generated by "esc "; DO NOT EDIT.

package js

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// _escFS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func _escFS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// _escDir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func _escDir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// _escFSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func _escFSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// _escFSMustByte is the same as _escFSByte, but panics if name is not present.
func _escFSMustByte(useLocal bool, name string) []byte {
	b, err := _escFSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// _escFSString is the string version of _escFSByte.
func _escFSString(useLocal bool, name string) (string, error) {
	b, err := _escFSByte(useLocal, name)
	return string(b), err
}

// _escFSMustString is the string version of _escFSMustByte.
func _escFSMustString(useLocal bool, name string) string {
	return string(_escFSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/helpers.js": {
		local:   "pkg/js/helpers.js",
		size:    14714,
		modtime: 0,
		compressed: `
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
Lfs/X3o5AAA=
`,
	},

	"/": {
		isDir: true,
		local: "pkg/js",
	},
}
