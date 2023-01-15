package decimation

import "errors"

// TODO: Add different downsampling methods
// Feel free to add more downsampling methods

// Decimation downsamples the input by a factor
// by taking every factor-th sample.
// The input must be a slice of int16.
// The factor must be positive.
// The output is a slice of int16.
func Decimation[T int16 | int32 | int64 | int | byte](in []T, factor int) ([]T, error) {
	if factor <= 0 {
		return nil, errors.New("downsampling factor must be positive")
	}
	if factor == 1 {
		return in, nil
	}
	out := make([]T, len(in)/factor)
	for i := 0; i < len(out); i++ {
		out[i] = in[i*factor]
	}
	return out, nil
}
