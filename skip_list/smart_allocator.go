package skipList

import "root/stack"

type SmartAllocator[T any] struct {
	Memory []T
	*stack.Stack[int]
}

func NewSmartAllocator[T any](capacity int) SmartAllocator[T] {
	return SmartAllocator[T]{
		Memory: make([]T, 0, capacity),
		Stack:  stack.NewStack[int](capacity),
	}
}

func (s *SmartAllocator[T]) GetNode(node T) {
	if s.Len() != 0 {
		position := s.Pop()
		s.Memory[position] = node
		return
	}

	s.Memory = append(s.Memory, node)
}
