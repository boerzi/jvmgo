package rtda

// Frame
// 有自己的局部变量标 和 操作数栈
//
// /**
type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
}

func newFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}
