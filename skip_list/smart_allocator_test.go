package skipList

import "testing"

func BenchmarkSmartAllocator(b *testing.B) {

	for i := 0; i < b.N; i++ {
		NewNode(1, 1)
	}
}
