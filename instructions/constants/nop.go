package constants

import (
	"leiyichen/jvmgo/instructions/base"
	"leiyichen/jvmgo/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

//func (n *NOP) FetchOperands(reader *base.BytecodeReader) {
//	//TODO implement me
//	panic("implement me")
//}

func (n *NOP) Execute(frame *rtda.Frame) {

}
