package classpath

import (
	"os"
	"strings"
)

const (
	pathListSeparator = string(os.PathSeparator)
)

type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

// NewEntry 文件夹、jar文件、zip文件
func NewEntry(path string) Entry {

	//文件夹
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}

	if strings.Contains(path, "*") {
		return newWildcardEntry(path)
	}

	var supportSuffix = []string{".jar", ".JAR", ".zip", ".ZIP"}

	for _, suffix := range supportSuffix {
		if strings.HasSuffix(path, suffix) {
			return NewZipEntry(path)
		}
	}

	return NewDirEntry(path)

}
