package skipList

import (
	"root/stack"
)

type SmartAllocator struct {
	Memory   []*Node
	FreeList []int
	Tracker  int
	stack.Stack[int]
}

func NewSmartAllocator(capacity int) SmartAllocator {
	return SmartAllocator{
		Memory:   make([]*Node, capacity),
		FreeList: make([]int, 0, capacity),
		Tracker:  -1,
		Stack:    stack.NewStack[int](capacity),
	}
}

func (s *SmartAllocator) GetNode(key, value int) *Node {
	if s.Tracker == len(s.Memory)-1 {
		s.Memory = append(s.Memory, make([]*Node, len(s.Memory))...)
	}

	s.Tracker++
	s.Memory[s.Tracker] = NewNode(key, value)
	return s.Memory[s.Tracker]
}
