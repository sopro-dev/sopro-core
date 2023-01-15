package interpolation

import "testing"

func TestLinearInterpolation(t *testing.T) {
	data := []int16{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int16{1, 3, 5, 7, 9}
	ratioIO := float64(16000) / float64(8000)
	resampledData, err := LinearInterpolation(data, ratioIO)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !testEq(resampledData, expected) {
		t.Errorf("Expected %v, got %v", expected, resampledData)
	}
}

func TestLinearInterpolation2(t *testing.T) {
	data := []int16{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int16{1, 0, 2, 0, 3, 0, 4, 0, 5, 0, 6, 0, 7, 0, 8, 0, 9, 0, 0, 0}
	ratioIO := float64(8000) / float64(16000)
	resampledData, err := LinearInterpolation(data, ratioIO)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !testEq(resampledData, expected) {
		t.Errorf("Expected %v, got %v", expected, resampledData)
	}
}

// testEq is a helper function to compare two slices of int16 values.
func testEq[T int16 | int32 | int64 | int](a, b []T) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
