package skipList

import (
	"math/rand"
	"testing"
)

// 1355 ns/op
func BenchmarkSpeedInsert(b *testing.B) {
	b.StopTimer()

	skipList := NewSkipList[int, int](32, 0.5)

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		skipList.Add(rand.Intn(b.N), 3)
	}
}

// 767.6 ns/op
func BenchmarkSearchList(b *testing.B) {
	b.StopTimer()

	skipList := NewSkipList[int, int](32, 0.4)
	for i := range 1000000 {
		skipList.Add(i, 3)
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		skipList.Search(15321)
	}
}

// 1004 ns/op
func BenchmarkDeleteElement(b *testing.B) {
	b.StopTimer()

	skipList := NewSkipList[int, int](32, 0.4)
	for i := range b.N {
		skipList.Add(i, 3)
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		skipList.Delete(i)
	}
}
