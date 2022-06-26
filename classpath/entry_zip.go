package classpath

import (
	"archive/zip"
	"errors"
	"io"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	absPath string
}

func NewZipEntry(absPath string) *ZipEntry {
	abs, err := filepath.Abs(absPath)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath: abs}
}

func (z ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(z.absPath)
	if err != nil {
		return nil, nil, err
	}

	defer func(r *zip.ReadCloser) {
		err := r.Close()
		if err != nil {
			panic(err)
		}
	}(r)
	for _, f := range r.File {
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}

			defer func(rc io.ReadCloser) {
				err := rc.Close()
				if err != nil {
					panic(err)
				}
			}(rc)
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}

			return data, z, nil
		}
	}

	return nil, nil, errors.New("class not found: " + className)
}

func (z ZipEntry) String() string {
	return z.absPath
}
