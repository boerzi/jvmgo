package classfile

import (
	"encoding/binary"
	"fmt"
)

type ClassReader struct {
	data []byte
}

func (c *ClassReader) readUint8() uint8 {
	val := c.data[0]
	c.data = c.data[1:]
	return val
}

func (c *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(c.data)
	c.data = c.data[2:]
	return val
}
func (c *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(c.data)
	c.data = c.data[4:]
	return val
}
func (c *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(c.data)
	c.data = c.data[8:]
	return val
}

func (c *ClassReader) readUint16s() []uint16 {
	n := c.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = c.readUint16()
	}
	return s
}

func (c *ClassReader) readBytes(n uint32) []byte {
	bytes := c.data[:n]
	c.data = c.data[n:]
	return bytes
}

//fieldsInfo
func (c *ClassReader) readFields(cp []ConstantInfo) []*fieldsInfo {
	list := make([]*fieldsInfo, len(cp))
	for i := 0; i < len(cp); i++ {
		list[i] = &fieldsInfo{
			accessFlags:     c.readUint16(),
			nameIndex:       c.readUint16(),
			descriptorIndex: c.readUint16(),
			attributesCount: c.readUint16(),
			attributes:      c.readAttributes(int(c.readUint16())),
		}

	}
	return list
}

func (c *ClassReader) readMethods(cp []ConstantInfo) []*MethodsInfo {
	list := make([]*MethodsInfo, len(cp))
	for i := 0; i < len(cp); i++ {
		list[i] = &MethodsInfo{
			accessFlags:     c.readUint16(),
			nameIndex:       c.readUint16(),
			descriptorIndex: c.readUint16(),
			attributesCount: c.readUint16(),
			attributes:      c.readAttributes(int(c.readUint16())),
		}
	}
	return list
}

func (c *ClassReader) readAttributes(count int) []AttributeInfo {
	attributeInfos := make([]AttributeInfo, count)
	//for i := 0; i < count; i++ {
	//	attributeInfos[i] =
	//
	//}
	return attributeInfos
}

//todo 这里有个bug 不知道为什么会有
func (c *ClassReader) readConstantPool(count int) []ConstantInfo {
	cpTable := make([]ConstantInfo, count)
	//直接得到length
	for i := 0; i < count-1; i++ {
		tag := c.readUint8()
		cpTable[i] = c.readConstantInfo(tag)
		fmt.Println(i, cpTable[i].info)
		if tag == CONSTANT_Float || tag == CONSTANT_Long {
			i++
			cpTable[i] = ConstantInfo{}
			fmt.Println(i, cpTable[i])
		}
	}
	return cpTable
}
