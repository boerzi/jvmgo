package jvmgo

import "testing"

func TestStart(t *testing.T) {
	startJVM("/Library/Java/JavaVirtualMachines/jdk1.8.0_181.jdk/Contents/Home/jre",
		"/Users/leiyuchen/go/src/leiyichen/jvmgo/testfile/GaussTest.class",
		"InvokeDemo",
		true)
}
