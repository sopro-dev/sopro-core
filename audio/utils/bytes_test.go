package utils

import (
	"bytes"
	"testing"
)

func TestMergeSliceOfBytes(t *testing.T) {
	slices := [][]byte{
		[]byte("hello"),
		[]byte(" "),
		[]byte("world"),
		[]byte(""),
	}
	expected := []byte("hello world")

	result := MergeSliceOfBytes(slices...)

	if !bytes.Equal(result, expected) {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestIntToBytes(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected []byte
	}{
		{uint32(123456), []byte{64, 226, 1, 0}}, // 123456 in little-endian bytes
		{uint16(0), []byte{0, 0}},               // 0 in little-endian bytes
		{uint16(65535), []byte{255, 255}},       // 65535 in little-endian bytes
		{uint16(1), []byte{1, 0}},               // 65536 in little-endian bytes
		{uint32(2084), []byte{24, 8, 0, 0}},     // 65536 in little-endian bytes
		// Add more test cases as needed
	}

	for _, test := range tests {
		result := IntToBytes(test.input)
		if !bytesEqual(result, test.expected) {
			t.Errorf("For input %v, expected %v, but got %v", test.input, test.expected, result)
		}
	}
}

// bytesEqual compares two byte slices for equality
func bytesEqual(a, b []byte) bool {
	return len(a) == len(b) && (len(a) == 0 || (a[0] == b[0] && bytesEqual(a[1:], b[1:])))
}
