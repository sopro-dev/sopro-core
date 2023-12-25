package utils

import (
	"encoding/binary"
)

// MergeSliceOfBytes merges multiple slices of bytes into a single slice
func MergeSliceOfBytes(slices ...[]byte) []byte {
	var result []byte
	for _, s := range slices {
		result = append(result, s...)
	}
	return result
}

// IntToBytes converts an unsigned integer to a little-endian byte slice
func IntToBytes(i interface{}) []byte {
	switch v := i.(type) {
	case uint32:
		buf := make([]byte, 4)
		binary.LittleEndian.PutUint32(buf, v)
		return buf
	case uint16:
		buf := make([]byte, 2)
		binary.LittleEndian.PutUint16(buf, v)
		return buf
	default:

	}
	return nil
}
