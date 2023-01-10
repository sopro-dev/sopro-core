package resampler

import "errors"

// TODO: Add different downsampling methods

// Decimation by factor
func DownsamplerDecimationByFactor[T int16](in []T, factor int) ([]T, error) {
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
