package cpuarch

import (
	"encoding/binary"
)

const (
	NOT_FILLED    = (iota - 1) // Not filled
	LITTLE_ENDIAN              // Little Endian
	BIG_ENDIAN                 // Big Endian
)

var ENDIANESSES = map[int]string{
	LITTLE_ENDIAN: "Little Endian",
	BIG_ENDIAN:    "Big Endian",
}

func GetEndianess() int {
	if binary.LittleEndian.Uint16([]byte{0x01, 0x00}) == 1 {
		return LITTLE_ENDIAN
	} else {
		return BIG_ENDIAN
	}
}
