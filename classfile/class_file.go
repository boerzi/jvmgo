package classfile

import "fmt"

type ClassFile struct {
	minorVersion      uint16
	majorVersion      uint16
	constantPoolCount uint16
	constantPool      ConstantPool
	accessFlags       uint16
	thisClass         uint16
	superClass        uint16
	interfaces        []uint16
	fieldsCount       uint16
	fields            []*FieldsInfo
	methodsCount      uint16
	methods           []*MethodsInfo
	attributesCount   uint16
	attributes        []AttributeInfo
}

func (cf *ClassFile) read(cr *ClassReader) {
	cf.readAndCheckMagic(cr)                                         //魔数固定的
	cf.readAndCheckVersion(cr)                                       //版本号
	cf.constantPoolCount = cr.readUint16()                           //常量池
	cf.constantPool = cr.readConstantPool(int(cf.constantPoolCount)) //类和超类索引
	cf.accessFlags = cr.readUint16()                                 //类访问标识
	cf.thisClass = cr.readUint16()                                   //类索引
	cf.superClass = cr.readUint16()                                  //超类索引
	cf.interfaces = cr.readUint16s()                                 //接口索引
	cf.fieldsCount = cr.readUint16()                                 //字段数量
	cf.fields = cr.readFields(cf.fieldsCount, cf.constantPool)       //字段
	cf.methodsCount = cr.readUint16()                                //方法数量
	cf.methods = cr.readMethods(cf.methodsCount, cf.constantPool)    //方法表
	cf.attributesCount = cr.readUint16()
	cf.attributes = cr.readAttributes(cf.attributesCount, cf.constantPool)
}

func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			} else {
				print(err.Error())
			}
		}
	}()
	cr := &ClassReader{data: classData}
	cf = &ClassFile{}
	cf.read(cr)
	return cf, err
}

func (cf *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
	fmt.Println(fmt.Sprintf("%x", magic))
}

func (cf *ClassFile) readAndCheckVersion(reader *ClassReader) {
	cf.minorVersion = reader.readUint16() //次版本号
	cf.majorVersion = reader.readUint16() //主版本号
	//todo 书里面版本号写死了 这里先不写
}

//type constantPool struct {
//	tag  uint8
//	info interface{} //可能有多种类型
//}

type FieldsInfo struct {
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

type MemberInfo struct {
	cp              ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}
