package rtda

import "leiyichen/jvmgo/rtda/heap"

// Slot
// ref引用
//
// /**
type Slot struct {
	num int32
	ref *heap.Object
}
