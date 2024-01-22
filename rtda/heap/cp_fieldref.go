package heap

import "leiyichen/jvmgo/classfile"

type FieldRef struct {
	MemberRef
	field *Field
}

func newFieldRef(cp *ConstantPool, refInfo *classfile.ConstantFieldrefInfo) *FieldRef {
	ref := &FieldRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

func (f *FieldRef) ResolvedField() *Field {
	if f.field == nil {
		f.resolvedFieldRef()
	}
	return f.field
}

func (f *FieldRef) resolvedFieldRef() {
	d := f.cp.class
	c := f.ResolvedClass()
	field := lookupField(c, f.name, f.descriptor)

	if field == nil {
		panic("java.lang.NoSuchFieldError")
	}
	if !field.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	f.field = field
}

func lookupField(c *Class, name string, descriptor string) *Field {
	for _, field := range c.fields {
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}

	for _, i := range c.interfaces {
		if field := lookupField(i, name, descriptor); field != nil {
			return field
		}
	}

	if c.superClass != nil {
		return lookupField(c.superClass, name, descriptor)
	}

	return nil
}
