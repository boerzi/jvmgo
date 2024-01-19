package heap

import (
	"fmt"
	"leiyichen/jvmgo/classfile"
	"os"
	"testing"
)

func TestClass(t *testing.T) {
	file, _ := os.ReadFile("../../testfile/GaussTest.class")
	a, _ := classfile.Parse(file)
	class := newClass(a)
	fmt.Println(class.getPackageName())
}
