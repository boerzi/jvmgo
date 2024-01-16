package classfile

import (
	"fmt"
	"math"
	"unicode/utf16"
)

// Constant pool tags
const (
	ConstantClass              = 7
	ConstantFieldref           = 9
	ConstantMethodref          = 10
	ConstantInterfacemethodref = 11
	ConstantString             = 8
	ConstantInteger            = 3
	ConstantFloat              = 4
	ConstantLong               = 5
	ConstantDouble             = 6
	ConstantNameandtype        = 12
	ConstantUtf8               = 1
	ConstantMethodhandle       = 15
	ConstantMethodtype         = 16
	ConstantInvokedynamic      = 18
)

type ConstantInfo struct {
	tag  uint8
	info interface{}
}

func (self ConstantInfo) getUtf8(index uint16) string {
	utf8Info := self.info.(*ConstantUtf8Info)
	return utf8Info.str
}

type ConstantIntegerInfo struct {
	val int32
}

type ConstantLongInfo struct {
	val int64
}

type ConstantFloatInfo struct {
	val float32
}

type ConstantDoubleInfo struct {
	val float64
}

type ConstantUtf8Info struct {
	str string
}

//type ConstantStringInfo struct {
//	cp          constantPool
//	stringIndex uint16
//}

type ConstantStringInfo struct {
	stringIndex uint16
}

type ConstantClassInfo struct {
	nameIndex uint16
}

type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

type ConstantMethodHandleInfo struct {
	referenceKind  uint8
	referenceIndex uint16
}

type ConstantMethodTypeInfo struct {
	descriptorIndex uint16
}

type ConstantInvokeDynamicInfo struct {
	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex         uint16
}

type ConstantMethodRef struct {
	classIndex       uint16
	nameAndTypeIndex uint16
}

type ConstantFieldRefInfo struct{ ConstantMethodRef }
type ConstantMethodRefInfo struct{ ConstantMethodRef }
type ConstantInterfaceMethodRefInfo struct{ ConstantMethodRef }

func (r *ClassReader) readConstantInfo(tag uint8) ConstantInfo {
	var info interface{}

	switch tag {
	case ConstantInteger:
		bytes := r.readUint32()
		info = &ConstantIntegerInfo{val: int32(bytes)}
	case ConstantFloat:
		bytes := r.readUint32()
		info = &ConstantFloatInfo{val: math.Float32frombits(bytes)}
	case ConstantLong:
		bytes := r.readUint64()
		info = &ConstantLongInfo{val: int64(bytes)}
	case ConstantDouble:
		bytes := r.readUint64()
		info = &ConstantDoubleInfo{val: math.Float64frombits(bytes)}
	case ConstantUtf8:
		length := uint32(r.readUint16())
		bytes := r.readBytes(length)
		info = &ConstantUtf8Info{str: decodeMUTF8(bytes)}
	case ConstantString:
		stringIndex := r.readUint16()
		info = &ConstantStringInfo{stringIndex: stringIndex}
	case ConstantClass:
		nameIndex := r.readUint16()
		info = &ConstantClassInfo{nameIndex: nameIndex}
	case ConstantFieldref:
		nameIndex := r.readUint16()
		descriptorIndex := r.readUint16()
		info = &ConstantNameAndTypeInfo{nameIndex: nameIndex, descriptorIndex: descriptorIndex}
	case ConstantMethodref:
		classIndex := r.readUint16()
		nameAndTypeIndex := r.readUint16()
		info = &ConstantMethodRef{classIndex: classIndex, nameAndTypeIndex: nameAndTypeIndex}
	case ConstantInterfacemethodref:
		classIndex := r.readUint16()
		nameAndTypeIndex := r.readUint16()
		info = &ConstantMethodRef{classIndex: classIndex, nameAndTypeIndex: nameAndTypeIndex}
	case ConstantNameandtype:
		classIndex := r.readUint16()
		nameAndTypeIndex := r.readUint16()
		info = &ConstantMethodRef{classIndex: classIndex, nameAndTypeIndex: nameAndTypeIndex}
	case ConstantMethodtype:
		classIndex := r.readUint16()
		nameAndTypeIndex := r.readUint16()
		info = &ConstantMethodRef{classIndex: classIndex, nameAndTypeIndex: nameAndTypeIndex}
	case ConstantMethodhandle:
		classIndex := r.readUint16()
		nameAndTypeIndex := r.readUint16()
		info = &ConstantMethodRef{classIndex: classIndex, nameAndTypeIndex: nameAndTypeIndex}
	case ConstantInvokedynamic:
		classIndex := r.readUint16()
		nameAndTypeIndex := r.readUint16()
		info = &ConstantMethodRef{classIndex: classIndex, nameAndTypeIndex: nameAndTypeIndex}
	default:
		panic("java.lang.ClassFormatError: constant pool tag!")
	}

	return ConstantInfo{tag, info}
}

func decodeMUTF8(bytearr []byte) string {
	utflen := len(bytearr)
	chararr := make([]uint16, utflen)

	var c, char2, char3 uint16
	count := 0
	chararr_count := 0

	for count < utflen {
		c = uint16(bytearr[count])
		if c > 127 {
			break
		}
		count++
		chararr[chararr_count] = c
		chararr_count++
	}

	for count < utflen {
		c = uint16(bytearr[count])
		switch c >> 4 {
		case 0, 1, 2, 3, 4, 5, 6, 7:
			/* 0xxxxxxx*/
			count++
			chararr[chararr_count] = c
			chararr_count++
		case 12, 13:
			/* 110x xxxx   10xx xxxx*/
			count += 2
			if count > utflen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(bytearr[count-1])
			if char2&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", count))
			}
			chararr[chararr_count] = c&0x1F<<6 | char2&0x3F
			chararr_count++
		case 14:
			/* 1110 xxxx  10xx xxxx  10xx xxxx*/
			count += 3
			if count > utflen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(bytearr[count-2])
			char3 = uint16(bytearr[count-1])
			if char2&0xC0 != 0x80 || char3&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", (count - 1)))
			}
			chararr[chararr_count] = c&0x0F<<12 | char2&0x3F<<6 | char3&0x3F<<0
			chararr_count++
		default:
			/* 10xx xxxx,  1111 xxxx */
			panic(fmt.Errorf("malformed input around byte %v", count))
		}
	}
	// The number of chars produced may be less than utflen
	chararr = chararr[0:chararr_count]
	runes := utf16.Decode(chararr)
	return string(runes)
}
