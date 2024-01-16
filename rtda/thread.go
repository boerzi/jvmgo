package rtda

// Thread 有自己的栈和pc寄存器
type Thread struct {
	pc    int
	stack *Stack
}

func (s *Thread) Pc() int {
	return s.pc
}

func (s *Thread) SetPc(pc int) {
	s.pc = pc
}

func NewThread() *Thread {
	return &Thread{}
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
