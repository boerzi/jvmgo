package rtda

import "math"

// OperandStack
// 操作数栈
//
// /**
type OperandStack struct {
	size  uint
	slots []Slot
}

func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slots: make([]Slot, maxStack),
		}
	}
	return nil
}

func (s *OperandStack) PushInt(val int32) {
	s.slots[s.size].num = val
	s.size++
}
func (s *OperandStack) PopInt() int32 {
	s.size--
	return s.slots[s.size].num
}

func (s *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	s.slots[s.size].num = int32(bits)
	s.size++
}
func (s *OperandStack) PopFloat() float32 {
	s.size--
	bits := uint32(s.slots[s.size].num)
	return math.Float32frombits(bits)
}

func (s *OperandStack) PushLong(val int64) {
	s.slots[s.size].num = int32(val)
	s.slots[s.size+1].num = int32(val >> 32)
	s.size += 2
}
func (s *OperandStack) PopLong() int64 {
	s.size -= 2
	low := uint32(s.slots[s.size].num)
	high := uint32(s.slots[s.size+1].num)
	return int64(high)<<32 | int64(low)
}

func (s *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	s.PushLong(int64(bits))
}
func (s *OperandStack) PopDouble() float64 {
	bits := uint64(s.PopLong())
	return math.Float64frombits(bits)
}

func (s *OperandStack) PushRef(ref *Object) {
	s.slots[s.size].ref = ref
	s.size++
}
func (s *OperandStack) PopRef() *Object {
	s.size--
	ref := s.slots[s.size].ref
	s.slots[s.size].ref = nil
	return ref
}
