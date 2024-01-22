package rtda

import "leiyichen/jvmgo/rtda/heap"

// Frame
// 有自己的局部变量标 和 操作数栈
//
// /**
type Frame struct {
	lower        *Frame        //上一级栈
	localVars    LocalVars     //局部变量表
	operandStack *OperandStack //操作数栈
	thread       *Thread       //关联的线程
	method       *heap.Method  //方法区
	nextPC       int           //计数器
}

func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread:       thread,
		method:       method,
		localVars:    newLocalVars(method.MaxLocals()),
		operandStack: newOperandStack(method.MaxStack()),
	}
}

func (f *Frame) LocalVars() LocalVars {
	return f.localVars
}
func (f *Frame) OperandStack() *OperandStack {
	return f.operandStack
}

func (f *Frame) NextPC() int {
	return f.nextPC
}

func (f *Frame) SetNextPC(nextPC int) {
	f.nextPC = nextPC
}

func (f *Frame) Thread() *Thread {
	return f.thread
}

func (f *Frame) Method() *heap.Method {
	return f.method
}
