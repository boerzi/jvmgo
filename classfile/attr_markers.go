package classfile

/*
用于指出类、接口、字段或者方法不建议使用
@deprecated_attribute

	Deprecated_attribute {
	    u2 attribute_name_index;
	    u4 attribute_length;
	}
*/
type DeprecatedAttribute struct {
	MarkerAttribute
}

/*
标记源文件不存在、由编辑器生产的类成员

	Synthetic_attribute {
	    u2 attribute_name_index;
	    u4 attribute_length;
	}
*/
type SyntheticAttribute struct {
	MarkerAttribute
}

type MarkerAttribute struct {
}

func (self *MarkerAttribute) readInfo(reader *ClassReader) {

}
