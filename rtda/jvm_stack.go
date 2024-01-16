package rtda

// Stack
// linked list 实现java虚拟机栈
// Stack   |-> Frame    |-> Frame
// _top  ---   lower  ---   lower
// *
type Stack struct {
	maxSize uint
	size    uint
	_top    *Frame
}

func (s *Stack) push(frame *Frame) {
	if s.size >= s.maxSize {
		panic("java.lang.StackOverflowError")
	}
	if s._top != nil {
		frame.lower = s._top
	}
	s._top = frame
	s.size++
}

func (s *Stack) pop() *Frame {
	top := s._top
	if top == nil {
		panic("jvm stack is empty!")
	}
	s._top = top.lower
	top.lower = nil
	s.size--
	return top
}

func (s *Stack) top() *Frame {
	if s._top == nil {
		panic("jvm stack is empty!")
	}
	return s._top
}

func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}
