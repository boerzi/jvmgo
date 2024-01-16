package classfile

import (
	"io/ioutil"
	"testing"
)

func Test(t *testing.T) {
	bytes := []byte{1, 2, 3, 4}
	var a = &ClassReader{data: bytes}
	//println(a.readUint8())
	unit16 := a.readUint16()
	println(unit16)
	//print(a.readUint32())

}

func TestParse(t *testing.T) {
	file, _ := ioutil.ReadFile("./ClassFileTest.class")
	_, _ = Parse(file)
	//printClassInfo(cf)
}

//func printClassInfo(cf *ClassFile) {
//	fmt.Printf("version: %v.%v\n", cf.majorVersion, cf.minorVersion)
//	fmt.Printf("constants count: %v\n", cf.constantPoolCount)
//	fmt.Printf("access flags: 0x%x\n", cf.accessFlags)
//	fmt.Printf("this class: %v\n", cf.thisClass)
//	fmt.Printf("super class: %v\n", cf.superClass)
//	fmt.Printf("interfaces: %v\n", cf.interfaces)
//	fmt.Printf("fields count: %v\n", cf.fieldsCount)
//	for _, f := range cf.fields {
//		fmt.Printf("  %v\n", f)
//	}
//	fmt.Printf("methods count: %v\n", cf.methodsCount)
//	for _, m := range cf.methods {
//		fmt.Printf("  %v\n", m)
//	}
//}
