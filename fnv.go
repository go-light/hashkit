package hashkit

const (
	_FNV_64_OFFSET_BASIS = uint64(0xcbf29ce484222325)
	_FNV_64_PRIME = uint64(0x100000001b3)
	_FNV_32_OFFSET_BASIS = uint32(2166136261)
	_FNV_32_PRIME = uint32(16777619)
)

func Hash_fnv1_64(key string) uint64 {
	hash1 := _FNV_64_OFFSET_BASIS
	for i := 0; i < len(key); i++ {
		hash1 *= _FNV_64_PRIME
		hash1 ^= uint64(key[i])
	}

	return hash1
}

func Hash_fnv1a_64(key string) uint64 {
	hash1 := _FNV_64_OFFSET_BASIS
	for i := 0; i < len(key); i++ {
		val := uint64(key[i])
		hash1 ^= val
		hash1 *= _FNV_64_PRIME
	}

	return hash1
}


func Hash_fnv1_32(key string) uint32 {
	hash1 := _FNV_32_OFFSET_BASIS
	for i := 0; i < len(key); i++ {
		val := uint32(key[i])
		hash1 *= _FNV_32_PRIME
		hash1 ^= val
	}

	return hash1
}

func Hash_fnv1a_32(key string) uint32 {
	hash1 := _FNV_32_OFFSET_BASIS
	for i := 0; i < len(key); i++ {
		val := uint32(key[i])
		hash1 ^= val
		hash1 *= _FNV_32_PRIME
	}

	return hash1
}