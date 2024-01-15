package classfile

/*
	    用于表达常量表达式的值
		ConstantValue_attribute {
		    u2 attribute_name_index;
		    u4 attribute_length;
		}
*/
type ConstantValueAttribute struct {
	constantValueIndex uint16
}

func (self *ConstantValueAttribute) readInfo(reader *ClassReader) {
	self.constantValueIndex = reader.readUint16()
}

func (self *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return self.constantValueIndex
}
