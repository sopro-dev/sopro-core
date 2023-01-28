package cpuarch

import (
	"encoding/binary"
)

// Endianness enum:
// NOT_FILLED    = -1
// LITTLE_ENDIAN = 0
// BIG_ENDIAN    = 1
const (
	NOT_FILLED    = (iota - 1) // Not filled
	LITTLE_ENDIAN              // Little Endian
	BIG_ENDIAN                 // Big Endian
)

// ENDIANESSES is a map of endianness
var ENDIANESSES = map[int]string{
	NOT_FILLED:    "Not Filled",
	LITTLE_ENDIAN: "Little Endian",
	BIG_ENDIAN:    "Big Endian",
}

// GetEndianess returns the endianness of the CPU safety
func GetEndianess() int {
	if binary.LittleEndian.Uint16([]byte{0x01, 0x00}) == 1 {
		return LITTLE_ENDIAN
	} else {
		return BIG_ENDIAN
	}
}
