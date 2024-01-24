package rtda

import "leiyichen/jvmgo/rtda/heap"

// Thread 有自己的栈和pc寄存器
type Thread struct {
	pc    int
	stack *Stack
}

func (s *Thread) PC() int {
	return s.pc
}

func (s *Thread) SetPC(pc int) {
	s.pc = pc
}

func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}

func (s *Thread) PushFrame(frame *Frame) {
	s.stack.push(frame)
}

func (s *Thread) PopFrame() *Frame {
	return s.stack.pop()
}

func (s *Thread) CurrentFrame() *Frame {
	return s.stack.top()
}

func (s *Thread) NewFrame(method *heap.Method) *Frame {
	return newFrame(s, method)
}

func (s *Thread) TopFrame() *Frame {
	return s.stack.pop()
}

func (s *Thread) IsStackEmpty() bool {
	return s.stack.isEmpty()
}
