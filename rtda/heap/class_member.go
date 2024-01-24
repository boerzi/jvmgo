package heap

import "leiyichen/jvmgo/classfile"

// ClassMember 字段信息
type ClassMember struct {
	accessFlags uint16
	name        string
	descriptor  string
	class       *Class
}

func (cm *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	cm.accessFlags = memberInfo.AccessFlags()
	cm.name = memberInfo.Name()
	cm.descriptor = memberInfo.Descriptor()
}

func (cm *ClassMember) IsPublic() bool {
	return 0 != cm.accessFlags&ACC_PUBLIC
}
func (cm *ClassMember) IsPrivate() bool {
	return 0 != cm.accessFlags&ACC_PRIVATE
}
func (cm *ClassMember) IsProtected() bool {
	return 0 != cm.accessFlags&ACC_PROTECTED
}
func (cm *ClassMember) IsStatic() bool {
	return 0 != cm.accessFlags&ACC_STATIC
}
func (cm *ClassMember) IsFinal() bool {
	return 0 != cm.accessFlags&ACC_FINAL
}
func (cm *ClassMember) IsSynthetic() bool {
	return 0 != cm.accessFlags&ACC_SYNTHETIC
}

// getters
func (cm *ClassMember) Name() string {
	return cm.name
}
func (cm *ClassMember) Descriptor() string {
	return cm.descriptor
}
func (cm *ClassMember) Class() *Class {
	return cm.class
}

// jvms 5.4.4
func (self *ClassMember) isAccessibleTo(d *Class) bool {
	if self.IsPublic() {
		return true
	}
	c := self.class
	if self.IsProtected() {
		return d == c || d.isSubClassOf(c) ||
			c.getPackageName() == d.getPackageName()
	}
	if !self.IsPrivate() {
		return c.getPackageName() == d.getPackageName()
	}
	return d == c
}
