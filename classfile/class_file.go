package classfile

import "fmt"

type ClassFile struct {
	minorVersion      uint16
	majorVersion      uint16
	constantPoolCount uint16
	constantPool      []ConstantPool
	accessFlags       uint16
	thisClass         uint16
	superClass        uint16
	interfaces        []uint16
	fields            []*fieldsInfo
	methods           []*MethodsInfo
	attributes        []AttributeInfo
}

func (cf ClassFile) read(cr *ClassReader) {
	cf.readAndCheckMagic(cr)               //魔数固定的
	cf.minorVersion = cr.readUint16()      //次版本号
	cf.majorVersion = cr.readUint16()      //主版本号
	cf.constantPoolCount = cr.readUint16() //常量池
	//cf.constantPool = cr.readConstantPool(int(cf.constantPoolCount)) //类和超类索引
	cf.accessFlags = cr.readUint16()
	cf.thisClass = cr.readUint16()
	cf.superClass = cr.readUint16()
	cf.interfaces = cr.readUint16s()
	//cf.fields = readMembers(cr, cf.constantPool)
	//cf.methods = readMembers(cr, cf.constantPool)
	//cf.attributes = readAttributes(cr, cf.constantPool)
	fmt.Println(cf)
}

//func readAttributes(cr *ClassReader, pool []ConstantPool) []AttributeInfo {
//
//}
//
//func readMembers(reader interface{}, pool interface{}) []*fieldsInfo {
//
//}

func Parse(classData []byte) (cf *ClassFile, err error) {
	cr := &ClassReader{data: classData}
	cf = &ClassFile{}
	cf.read(cr)
	return nil, err
}

func (cf *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
	fmt.Println(fmt.Sprintf("%x", magic))
}

type ConstantPool struct {
	tag  uint8
	info interface{} //可能有多种类型
}

type fieldsInfo struct {
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributesCount uint16
	attributes      []AttributeInfo
}

type MethodsInfo struct {
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributesCount uint16
	attributes      []AttributeInfo
}

type AttributeInfo struct {
	attributeNameIndex uint16
	attributeLength    uint32
	info               interface{}
}
