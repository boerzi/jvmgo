package heap

import (
	"leiyichen/jvmgo/classfile"
	"strings"
)

type ClassLoader struct {
}

type Slots struct {
}

type Class struct {
	accessFlags       uint16 //类的访问标识包括私有公有枚举类型啥的
	name              string
	superClassName    string
	interfaceNames    []string
	constantPool      *ConstantPool
	fields            []*Field
	methods           []*Method
	loader            *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        Slots
}

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}

func (c *Class) IsPublic() bool {
	return 0 != c.accessFlags&ACC_PUBLIC
}
func (c *Class) IsFinal() bool {
	return 0 != c.accessFlags&ACC_FINAL
}
func (c *Class) IsSuper() bool {
	return 0 != c.accessFlags&ACC_SUPER
}
func (c *Class) IsInterface() bool {
	return 0 != c.accessFlags&ACC_INTERFACE
}
func (c *Class) IsAbstract() bool {
	return 0 != c.accessFlags&ACC_ABSTRACT
}
func (c *Class) IsSynthetic() bool {
	return 0 != c.accessFlags&ACC_SYNTHETIC
}
func (c *Class) IsAnnotation() bool {
	return 0 != c.accessFlags&ACC_ANNOTATION
}
func (c *Class) IsEnum() bool {
	return 0 != c.accessFlags&ACC_ENUM
}

func (c *Class) ConstantPool() *ConstantPool {
	return c.constantPool
}
func (c *Class) StaticVars() Slots {
	return c.staticVars
}

// jvms 5.4.4
func (c *Class) isAccessibleTo(other *Class) bool {
	return c.IsPublic() ||
		c.getPackageName() == other.getPackageName()
}

func (c *Class) getPackageName() string {
	if i := strings.LastIndex(c.name, "/"); i >= 0 {
		return c.name[:i]
	}
	return ""
}

func (c *Class) GetMainMethod() *Method {
	return c.getStaticMethod("main", "([Ljava/lang/String;)V")
}

func (c *Class) getStaticMethod(name, descriptor string) *Method {
	for _, method := range c.methods {
		if method.IsStatic() &&
			method.name == name &&
			method.descriptor == descriptor {

			return method
		}
	}
	return nil
}

//func (self *Class) NewObject() *Object {
//	return newObject(self)
//}
