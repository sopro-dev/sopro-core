package resampler

import (
	"reflect"
	"testing"
)

func TestPCM16Downsample(t *testing.T) {
	tests := []struct {
		name     string
		in       []int16
		factor   int
		expected []int16
	}{
		{
			name:     "factor 2",
			in:       []int16{1, 2, 3, 4, 5, 6},
			factor:   2,
			expected: []int16{1, 3, 5},
		},
		{
			name:     "factor 3",
			in:       []int16{1, 2, 3, 4, 5, 6},
			factor:   3,
			expected: []int16{1, 4},
		},
		{
			name:     "factor 1",
			in:       []int16{1, 2, 3, 4, 5, 6},
			factor:   1,
			expected: []int16{1, 2, 3, 4, 5, 6},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out, err := DownsamplerDecimationByFactor(test.in, test.factor)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if !reflect.DeepEqual(out, test.expected) {
				t.Errorf("Got %v, expected %v", out, test.expected)
			}
		})
	}
}
