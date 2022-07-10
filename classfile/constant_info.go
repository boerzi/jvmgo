package classfile

import (
	"fmt"
	"math"
	"unicode/utf16"
)

// Constant pool tags
const (
	CONSTANT_Class              = 7
	CONSTANT_Fieldref           = 9
	CONSTANT_Methodref          = 10
	CONSTANT_InterfaceMethodref = 11
	CONSTANT_String             = 8
	CONSTANT_Integer            = 3
	CONSTANT_Float              = 4
	CONSTANT_Long               = 5
	CONSTANT_Double             = 6
	CONSTANT_NameAndType        = 12
	CONSTANT_Utf8               = 1
	CONSTANT_MethodHandle       = 15
	CONSTANT_MethodType         = 16
	CONSTANT_InvokeDynamic      = 18
)

type ConstantInfo struct {
	tag  uint8
	info interface{}
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
//	cp          ConstantPool
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
	case CONSTANT_Integer:
		bytes := r.readUint32()
		info = &ConstantIntegerInfo{val: int32(bytes)}
	case CONSTANT_Float:
		bytes := r.readUint32()
		info = &ConstantFloatInfo{val: math.Float32frombits(bytes)}
	case CONSTANT_Long:
		bytes := r.readUint64()
		info = &ConstantLongInfo{val: int64(bytes)}
	case CONSTANT_Double:
		bytes := r.readUint64()
		info = &ConstantDoubleInfo{val: math.Float64frombits(bytes)}
	case CONSTANT_Utf8:
		length := uint32(r.readUint16())
		bytes := r.readBytes(length)
		info = &ConstantUtf8Info{str: decodeMUTF8(bytes)}
	case CONSTANT_String:
		stringIndex := r.readUint16()
		info = &ConstantStringInfo{stringIndex: stringIndex}
	case CONSTANT_Class:
		nameIndex := r.readUint16()
		info = &ConstantClassInfo{nameIndex: nameIndex}
	case CONSTANT_Fieldref:
		nameIndex := r.readUint16()
		descriptorIndex := r.readUint16()
		info = &ConstantNameAndTypeInfo{nameIndex: nameIndex, descriptorIndex: descriptorIndex}
	case CONSTANT_Methodref:
		classIndex := r.readUint16()
		nameAndTypeIndex := r.readUint16()
		info = &ConstantMethodRef{classIndex: classIndex, nameAndTypeIndex: nameAndTypeIndex}
	case CONSTANT_InterfaceMethodref:
		classIndex := r.readUint16()
		nameAndTypeIndex := r.readUint16()
		info = &ConstantMethodRef{classIndex: classIndex, nameAndTypeIndex: nameAndTypeIndex}
	case CONSTANT_NameAndType:
		classIndex := r.readUint16()
		nameAndTypeIndex := r.readUint16()
		info = &ConstantMethodRef{classIndex: classIndex, nameAndTypeIndex: nameAndTypeIndex}
	case CONSTANT_MethodType:
		classIndex := r.readUint16()
		nameAndTypeIndex := r.readUint16()
		info = &ConstantMethodRef{classIndex: classIndex, nameAndTypeIndex: nameAndTypeIndex}
	case CONSTANT_MethodHandle:
		classIndex := r.readUint16()
		nameAndTypeIndex := r.readUint16()
		info = &ConstantMethodRef{classIndex: classIndex, nameAndTypeIndex: nameAndTypeIndex}
	case CONSTANT_InvokeDynamic:
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
