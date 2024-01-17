package stack

import (
	"leiyichen/jvmgo/instructions/base"
	"leiyichen/jvmgo/rtda"
)

type SWAP struct{ base.NoOperandsInstruction }

/*
Execute

	bottom -> top
	[...][c][b][a]
			  \/
			  /\
			 V  V
	[...][c][a][b]
*/
func (self *SWAP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}
