package jvmgo

import "testing"

func TestStart(t *testing.T) {
	strings := make([]string, 0)
	strings = append(strings, "hello world")
	strings = append(strings, "hello world123")
	startJVM("/Library/Java/JavaVirtualMachines/jdk1.8.0_181.jdk/Contents/Home/jre",
		"/Users/leiyuchen/go/src/leiyichen/jvmgo/testfile/PrintArgs.class",
		"PrintArgs",
		true, strings)
}
