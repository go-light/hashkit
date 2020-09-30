package hashkit

import "testing"

const str = "abcdefghijklmnopqrstuvwxyz0123456789"

func Benchmark_hash_fnv1_32(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Hash_fnv1_32(str)
	}
}

func Benchmark_hash_fnv1a_32(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Hash_fnv1a_32(str)
	}
}

func Benchmark_hash_fnv1_64(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Hash_fnv1_64(str)
	}
}

func Benchmark_hash_fnv1a_64(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Hash_fnv1a_64(str)
	}
}
