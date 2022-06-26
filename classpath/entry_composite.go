package classpath

import "errors"
import "strings"

type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	var compositeEntry []Entry
	//组合模式嵌套
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := NewEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}

	return compositeEntry
}

func (c CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range c {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, nil
		}
	}

	return nil, nil, errors.New("class not found: " + className)
}

func (c CompositeEntry) String() string {
	str := make([]string, len(c))

	for i, entry := range c {
		str[i] = entry.String()
	}

	return strings.Join(str, pathListSeparator)
}
