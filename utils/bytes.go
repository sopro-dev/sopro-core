package utils

import "bytes"

// MergeSliceOfBytes merges multiple slices of bytes into a single slice of bytes
// Mostly used to generate good spacing for documentation
// Example:
//
//	slice1 := []byte{1, 2, 3}
//	slice2 := []byte{4, 5, 6}
//	slice3 := []byte{7, 8, 9}
//	slice := mergeSliceOfBytes(slice1, slice2, slice3)
//	fmt.Println(slice)
//
// Output:
//
//	[1 2 3 4 5 6 7 8 9]
func MergeSliceOfBytes(slices ...[]byte) []byte {
	var buffer bytes.Buffer
	for _, slice := range slices {
		buffer.Write(slice)
	}
	return buffer.Bytes()
}
