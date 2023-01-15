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
	}
	expected := []byte("hello world")

	result := MergeSliceOfBytes(slices...)

	if !bytes.Equal(result, expected) {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}
