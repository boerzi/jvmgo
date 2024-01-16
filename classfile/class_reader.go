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

// FieldsInfo
func (c *ClassReader) readFields(count uint16, cp ConstantPool) []*FieldsInfo {
	list := make([]*FieldsInfo, count)
	for i := range list {
		accessFlags := c.readUint16()
		nameIndex := c.readUint16()
		descriptorIndex := c.readUint16()
		attributesCount := c.readUint16()

		list[i] = &FieldsInfo{
			accessFlags:     accessFlags,
			nameIndex:       nameIndex,
			descriptorIndex: descriptorIndex,
			attributesCount: attributesCount,
			attributes:      c.readAttributes(attributesCount, cp),
		}

	}
	return list
}

func (c *ClassReader) readMethods(count uint16, cp []ConstantInfo) []*MethodsInfo {
	list := make([]*MethodsInfo, count)
	for i := range list {
		accessFlags := c.readUint16()
		nameIndex := c.readUint16()
		descriptorIndex := c.readUint16()
		attributesCount := c.readUint16()

		list[i] = &MethodsInfo{
			accessFlags:     accessFlags,
			nameIndex:       nameIndex,
			descriptorIndex: descriptorIndex,
			attributesCount: attributesCount,
			attributes:      c.readAttributes(attributesCount, cp),
		}

	}
	return list
}

//func (c *ClassReader) readAttribute(count int, cp []ConstantInfo) []AttributeInfo {
//	attributeInfos := make([]AttributeInfo, count)
//	for i := range attributeInfos {
//		attributeInfos[i] =
//	}
//	return attributeInfos
//}

type AttributeInfoCall interface {
	readInfo(reader *ClassReader)
}

func (c *ClassReader) readAttributes(count uint16, cp ConstantPool) []AttributeInfo {
	list := make([]AttributeInfo, count)
	for i := range list {
		attributeNameIndex := c.readUint16()
		attrName := cp[attributeNameIndex-1].info.(*ConstantUtf8Info).str
		attributeLength := c.readUint32()
		info := newAttributeInfo(attrName, attributeLength, cp)
		info.readInfo(c)
		fmt.Println(attrName)
		a := AttributeInfo{
			attributeNameIndex,
			attributeLength,
			info,
		}
		list[i] = a
	}
	return list
}

type ConstantPool []ConstantInfo

func (self ConstantPool) getUtf8(index uint16) string {
	utf8Info := self[index].info.(*ConstantUtf8Info)
	return utf8Info.str
}

func (c *ClassReader) readConstantPool(count int) ConstantPool {
	cpTable := make([]ConstantInfo, count)
	//直接得到length
	for i := 0; i < count-1; i++ {
		tag := c.readUint8()
		cpTable[i] = c.readConstantInfo(tag)
		fmt.Println(i, cpTable[i].info, tag)
		if tag == ConstantFloat || tag == ConstantLong {
			i++
			cpTable[i] = ConstantInfo{}
			fmt.Println(i, cpTable[i])
		}
	}
	return cpTable
}
