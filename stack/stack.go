package stack

type Stack[T any] struct {
	store []T
}

func NewStack[T any](capacity int) *Stack[T] {
	return &Stack[T]{
		store: make([]T, 0, capacity),
	}
}

func (s *Stack[T]) Push(value T) {
	s.store = append(s.store, value)
}

func (s *Stack[T]) Pop() T {
	len := len(s.store)
	if len == 0 {
		var zero T
		return zero
	}

	value := s.store[len-1]   // get last value
	s.store = s.store[:len-1] //remove last value
	return value
}

func (s *Stack[T]) Flush() {
	s.store = s.store[:0]
}

func (s *Stack[T]) Len() int {
	return len(s.store)
}
