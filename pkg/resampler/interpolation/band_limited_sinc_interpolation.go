package interpolation

import (
	"errors"
	"math"
)

// BandLimitedSincInterpolation resamples the input slice using
// the band-limited sinc interpolation method.
//
// Warning: size of the input slice must be greater than 0 and power of 2.
// Warning: For high-quality audio resampling, a byte size of 8192 or higher might be suitable.
// Warning: For lower quality audio resampling, a byte size of 4096 or lower might be acceptable.
// Warning: For speech recognition and other similar applications, a byte size of 2048 or lower may be sufficient.
// Warning: For real-time audio processing, a byte size of 1024 or lower may be necessary to meet the processing requirements.
<<<<<<< HEAD
func BandLimitedSincInterpolation[T int16 | int32 | int64 | int | byte](input []T, ratioIO float64) ([]T, error) {
	if ratioIO <= 0 {
		return input, errors.New("ratio must be positive")
	}
	//calculate the number of samples in the input slice
	inputSamples := len(input)

	//calculate the number of samples in the output slice
	outputSamples := int(math.Ceil(float64(inputSamples) / ratioIO))

	//allocate the output slice with the correct size
	output := make([]T, outputSamples)

	//resample the audio
	for i := 0; i < outputSamples; i++ {
		//calculate the corresponding sample index in the input slice
		inputIndex := float64(i) * ratioIO
		//calculate the fractional part of the index
		alpha := inputIndex - math.Floor(inputIndex)
		_ = alpha // TODO: use alpha
		//initialize the sum
		sum := 0.0
		//apply the band-limited sinc interpolation
		for j := -8; j <= 8; j++ {
			//calculate the index of the sample in the input slice
			sampleIndex := int(inputIndex) + j
			//check if the index is in range
			if sampleIndex < 0 || sampleIndex >= inputSamples {
				continue
			}
			//calculate the sinc function
			// denom := alpha + float64(j)
			// // sinc := utils.Sinc(denom)

			// //multiply the sample value by the sinc function
			// sum += float64(input[sampleIndex]) * sinc
		}
		//store the float value of the sum in the output slice
		output[i] = T(sum)
=======
func BandLimitedSincInterpolation[T int16 | int32 | int64 | int | byte](input []T, ratioOI float64) ([]T, error) {
	if ratioOI <= 0 {
		return input, errors.New("ratio must be positive")
	}
	// calculate the number of samples in the input slice
	inputSamples := len(input)

	// calculate the number of samples in the output slice
	outputSamples := int(float64(inputSamples) * ratioOI)

	// allocate the output slice with the correct size
	output := make([]T, outputSamples)

	// resample the audio
	for i := 0; i < outputSamples; i++ {
		// calculate the corresponding sample index in the input slice
		inputIndex := float64(i) / ratioOI
		// calculate the fractional part of the index
		alpha := inputIndex - math.Floor(inputIndex)
		// initialize the sum
		sum := 0.0
		// apply the band-limited sinc interpolation
		for j := -4; j <= 4; j++ {
			// calculate the index of the sample in the input slice
			sampleIndex := int(math.Floor(inputIndex)) + j
			// check if the index is in range
			if sampleIndex < 0 || sampleIndex >= inputSamples {
				continue
			}
			// calculate the sinc function
			sinc := math.Sin(math.Pi*(alpha+float64(j))) / (math.Pi * (alpha + float64(j)))
			// multiply the sample value by the sinc function
			sum += float64(input[sampleIndex]) * sinc
		}
		// round the result and store it in the output slice
		output[i] = T(math.Round(sum))
>>>>>>> d598982 (refactor, new resampler, new transcoder, sopro models, sinc interpolation and more examples)
	}
	return output, nil
}
