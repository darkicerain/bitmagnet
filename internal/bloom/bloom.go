package bloom

import (
	"bitmagnet-io/bitmagnet/internal/protocol/dht"
	"encoding/binary"
	"github.com/bits-and-blooms/bloom/v3"
)

type Filter = bloom.BloomFilter

const (
	size     = 32
	byteSize = size * 8
	M        = byteSize * 8
	K        = 2
)

var NewWithEstimates = bloom.NewWithEstimates

func FromScrape(f dht.ScrapeBloomFilter) Filter {
	return *bloom.FromWithM(convertBytes(f), M, K)
}

func convertBytes(b [byteSize]byte) []uint64 {
	ret := make([]uint64, size)
	for i := 0; i < size; i++ {
		startPos := i * 8
		ret[i] = binary.BigEndian.Uint64(b[startPos : startPos+8])
	}
	return ret
}
