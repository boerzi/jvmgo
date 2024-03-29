package classpath

import (
	"fmt"
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

func (c *Classpath) parseBootAndExtClasspath(option string) {
	jreDir := getJreDir(option)

	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	fmt.Println("jreLibPath:", jreLibPath)
	c.bootClasspath = newWildcardEntry(jreLibPath)
	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	c.extClasspath = newWildcardEntry(jreExtPath)
	fmt.Println("jreExtPath", jreExtPath)
}

func (c *Classpath) parseUserClasspath(option string) {
	if option == "" {
		option = "."
	}
	c.userClasspath = NewEntry(option)
}

func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption) //运行时必要项目
	cp.parseUserClasspath(cpOption)        //用户自己的路径
	return cp
}

func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("Can not find jre folder!")
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (c *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := c.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := c.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	return c.userClasspath.readClass(className)
}

func (c *Classpath) String() string {
	return c.userClasspath.String()
}
