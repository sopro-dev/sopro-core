package resampler

func LinearInterpolation[T int16 | int32 | int64 | int | byte](data []T, ratio float64) ([]T, error) {
	// Calculate the length of the resampled data slice.
	resampledLength := int(float64(len(data)) / ratio)

	// Preallocate the resampled data slice with the correct size.
	resampledData := make([]T, resampledLength)

	// Iterate over the original data, interpolating new samples as necessary to
	// resample the data at the target sample rate.
	for i := 0; i < len(data)-1; i++ {
		// Calculate the interpolated value between the current and next samples.
		interpolatedValue := T((float64(data[i]) + float64(data[i+1])) / 2)
		resampledData[int(float64(i)/ratio)] = interpolatedValue

		// Skip the next sample if necessary.
		if ratio > 1.0 {
			i += int(ratio) - 1
		}
	}

	return resampledData, nil
}
