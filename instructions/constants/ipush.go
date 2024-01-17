package constants

import (
	"leiyichen/jvmgo/instructions/base"
	"leiyichen/jvmgo/rtda"
)

type BIPUSH struct {
	val int8
}

func (B *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	B.val = reader.ReadInt8()
}

func (B BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(B.val)
	frame.OperandStack().PushInt(i)
}

type SIPUSH struct {
	val int16
}

func (self *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt16()
}
func (self *SIPUSH) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}
