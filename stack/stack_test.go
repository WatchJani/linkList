package stack

import (
	"testing"
)

func BenchmarkPopPushSpeed(b *testing.B) {
	b.StopTimer()
	stack := NewStack[int](26)

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		stack.Push(5)
		stack.Pop()
	}
}
