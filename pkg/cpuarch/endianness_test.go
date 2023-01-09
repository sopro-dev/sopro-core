package cpuarch

import (
	"testing"
)

func TestGetEndianess(t *testing.T) {
	e := GetEndianess()
	if e != LITTLE_ENDIAN && e != BIG_ENDIAN {
		t.Errorf("unexpected endianess: %d", e)
	}
}
