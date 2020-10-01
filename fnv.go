package hashkit

const (
	Fnv64OffsetBasis = uint64(0xcbf29ce484222325)
	Fnv64Prime       = uint64(0x100000001b3)
	Fnv32OffsetBasis = uint32(2166136261)
	Fnv32Prime       = uint32(16777619)
)

// Fnv64 return FNV-1 64 bit hash value
func Fnv64(key string) uint64 {
	hash1 := Fnv64OffsetBasis
	for i := 0; i < len(key); i++ {
		hash1 *= Fnv64Prime
		hash1 ^= uint64(key[i])
	}

	return hash1
}

// Fnv64a return FNV-1 64 bit hash value
func Fnv64a(key string) uint64 {
	hash1 := Fnv64OffsetBasis
	for i := 0; i < len(key); i++ {
		val := uint64(key[i])
		hash1 ^= val
		hash1 *= Fnv64Prime
	}

	return hash1
}

// Fnv32 return FNV-1 32 bit hash value
func Fnv32(key string) uint32 {
	hash1 := Fnv32OffsetBasis
	for i := 0; i < len(key); i++ {
		val := uint32(key[i])
		hash1 *= Fnv32Prime
		hash1 ^= val
	}

	return hash1
}

// Fnv32a return FNV-1a 32 bit hash value
func Fnv32a(key string) uint32 {
	hash1 := Fnv32OffsetBasis
	for i := 0; i < len(key); i++ {
		val := uint32(key[i])
		hash1 ^= val
		hash1 *= Fnv32Prime
	}

	return hash1
}