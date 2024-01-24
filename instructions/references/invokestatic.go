package references

import (
	"leiyichen/jvmgo/instructions/base"
	"leiyichen/jvmgo/rtda"
	"leiyichen/jvmgo/rtda/heap"
)

type INVOKE_STATIC struct {
	base.Index16Instruction
}

func (is INVOKE_STATIC) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(is.Index).(*heap.MethodRef)
	resolvedMethod := methodRef.ResolvedMethod()

	if !resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	base.InvokeMethod(frame, resolvedMethod)

}
