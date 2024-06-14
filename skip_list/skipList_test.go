package skipList

import "testing"

func BenchmarkSpeedInsert(b *testing.B) {
	b.StopTimer()

	skipList := NewSkipList(32, 0.2)

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		skipList.Add(123, 3)
	}
}
