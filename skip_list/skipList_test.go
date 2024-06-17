package skipList

import (
	"math/rand"
	"testing"
)

func BenchmarkSpeedInsert(b *testing.B) {
	b.StopTimer()

	skipList := NewSkipList(32, 0.5)

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		skipList.Add(rand.Intn(b.N), 3)
	}
}

func BenchmarkSearchList(b *testing.B) {
	b.StopTimer()

	skipList := NewSkipList(32, 0.4)
	for i := range 1000000 {
		skipList.Add(i, 3)
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		skipList.Search(15321)
	}
}
