package references

import (
	"leiyichen/jvmgo/instructions/base"
	"leiyichen/jvmgo/rtda"
	"leiyichen/jvmgo/rtda/heap"
)

type NEW struct {
	base.Index16Instruction
}

func (n *NEW) Execute(frame *rtda.Frame) {

	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(n.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()

	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}

	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}
