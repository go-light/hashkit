package hashkit

import "testing"

const str = "abcdefghijklmnopqrstuvwxyz0123456789"

func Benchmark_Fnv32(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Fnv32(str)
	}
}

func Benchmark_Fnv32a(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Fnv32a(str)
	}
}

func Benchmark_Fnv64(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Fnv64(str)
	}
}

func Benchmark_Fnv64a(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Fnv64a(str)
	}
}
