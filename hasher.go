package hasher

import "github.com/ninepeach/go-hasher/xxhash"

func Sum64(key []byte) uint64 {
	return xxhash.Sum64(key)
}

