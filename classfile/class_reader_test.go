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
	Parse(file)
}
