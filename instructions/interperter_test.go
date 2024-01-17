package instructions

import (
	"leiyichen/jvmgo/classfile"
	"os"
	"testing"
)

func TestInterpret(t *testing.T) {
	file, _ := os.ReadFile("/Users/leiyuchen/go/src/leiyichen/jvmgo/testfile/GaussTest.class")
	a, _ := classfile.Parse(file)
	method := getMainMethod(a)
	interpret(method)
}

func getMainMethod(cf *classfile.ClassFile) *classfile.MethodsInfo {
	for _, m := range cf.Methods() {
		if cf.GetMethodsInfoName(m) == "main" && cf.GetMethodsInfoDescriptor(m) == "([Ljava/lang/String;)V" {
			return m
		}
	}
	return nil
}
