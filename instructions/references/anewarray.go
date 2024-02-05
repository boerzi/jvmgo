package references

import (
	"leiyichen/jvmgo/instructions/base"
	"leiyichen/jvmgo/rtda"
	"leiyichen/jvmgo/rtda/heap"
)

type ANEW_ARRAY struct {
	base.Index16Instruction
}

func (a ANEW_ARRAY) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(a.Index).(*heap.ClassRef)
	componentClass := classRef.ResolvedClass()

	stack := frame.OperandStack()
	count := stack.PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}

	arrClass := componentClass.ArrayClass()
	arr := arrClass.NewArray(uint(count))
	stack.PushRef(arr)
}
