package heap

import (
	"fmt"
	"leiyichen/jvmgo/classfile"
	"leiyichen/jvmgo/classpath"
)

type ClassLoader struct {
	cp          *classpath.Classpath
	verboseFlag bool
	classMap    map[string]*Class
}

func NewClassLoader(cp *classpath.Classpath, verboseFlag bool) *ClassLoader {
	return &ClassLoader{
		cp:          cp,
		verboseFlag: verboseFlag,
		classMap:    make(map[string]*Class),
	}
}

func (c *ClassLoader) LoadClass(name string) *Class {
	if class, ok := c.classMap[name]; ok {
		return class
	}
	return c.loadNonArrayClass(name)
}

func (c *ClassLoader) loadNonArrayClass(name string) *Class {
	data, entry := c.readClass(name) //拿到二进制class
	class := c.defineClass(data)     //开始解析文件结构
	link(class)
	if c.verboseFlag {
		fmt.Printf("[Loaded %s from %s\n]", name, entry)
	}
	return class
}

func link(class *Class) {
	verify(class)  //链接
	prepare(class) //准备
}

func prepare(class *Class) {
	calcInstanceFieldSlotIds(class)
	calcStaticFieldSlotIds(class)
	allocAndInitStaticVars(class)
}

func allocAndInitStaticVars(class *Class) {
	class.staticVars = newSlots(class.staticSlotCount)
	for _, field := range class.fields {
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalVar(class, field)
		}
	}
}

// initStaticFinalVar 初始化静态+常量
func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.ConstValueIndex()
	slotId := field.SlotId()

	if cpIndex > 0 {
		switch field.Descriptor() {
		case "Z", "B", "C", "S", "I":
			val := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotId, val)
		case "J":
			val := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotId, val)
		case "F":
			val := cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotId, val)
		case "D":
			val := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotId, val)
		case "Ljava/lang/String;":
			panic("todo")
		}
	}
}

// calcStaticFieldSlotIds 处理静态变量字段
func calcStaticFieldSlotIds(class *Class) {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.staticSlotCount = slotId
}

// calcInstanceFieldSlotIds 载入slot大小包括父类的字段
func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount
	}
	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.instanceSlotCount = slotId
}

func verify(class *Class) {

}

func (c *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	data, entry, err := c.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + name)
	}
	return data, entry
}

func (c *ClassLoader) defineClass(data []byte) *Class {
	class := parseClass(data)
	class.loader = c
	resolveSuperClass(class) //单继承
	resolveInterfaces(class) //多接口
	c.classMap[class.name] = class
	return class
}

func resolveInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}
}

func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}

func parseClass(data []byte) *Class {
	cf, err := classfile.Parse(data)
	if err != nil {
		panic("java,lang.ClassFormatError")
	}
	return newClass(cf)
}
