package utils

import "encoding/binary"

type ENDIANNESS int

const (
	LITTLE_ENDIAN ENDIANNESS = iota
	BIG_ENDIAN
)

// GetEndianess returns the endianness of the CPU safety
func GetEndianess() ENDIANNESS {
	if binary.LittleEndian.Uint16([]byte{0x01, 0x00}) == 1 {
		return LITTLE_ENDIAN
	} else {
		return BIG_ENDIAN
	}
}
