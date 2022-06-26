package classpath

import (
	"io/ioutil"
	"path/filepath"
)

type DirEntry struct {
	absDir string
}

func NewDirEntry(absDir string) *DirEntry {
	abs, err := filepath.Abs(absDir)
	if err != nil {
		panic(err)
	}
	return &DirEntry{abs}
}

func (d DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(d.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, d, err
}

func (d DirEntry) String() string {
	return d.absDir
}
