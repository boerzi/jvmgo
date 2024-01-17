package stack

import (
	"leiyichen/jvmgo/instructions/base"
	"leiyichen/jvmgo/rtda"
)

type POP struct{ base.NoOperandsInstruction }

/*
Execute

	bottom -> top
	[...][c][b][a]
				|
				V
	[...][c][b]
*/
func (p *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

type POP2 struct{ base.NoOperandsInstruction }

/*
Execute

	bottom -> top
	[...][c][b][a]
			 |  |
			 V  V
	[...][c]
*/
func (p *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
