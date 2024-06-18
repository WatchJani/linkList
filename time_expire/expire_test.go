package timeExpire

import "testing"

func Benchmark(b *testing.B) {
	b.StopTimer()

	expire := NewTimeExpire()

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		expire.AppendData(RandomData())
	}
}
