package interpolation

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestBandLimitedSincInterpolation(t *testing.T) {
	//create an input slice with random values
	input := make([]byte, 1000)
	for i := range input {
		input[i] = byte(rand.Intn(256))
	}
	//define the input and output sample rates
	inputSampleRate := 44100
	outputSampleRate := 48000

	//resample the audio
	ratioOI := float64(outputSampleRate) / float64(inputSampleRate)
	output, err := BandLimitedSincInterpolation(input, ratioOI)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	//print the length of the input and output slices
	fmt.Printf("input length: %d\n", len(input))
	fmt.Printf("output length: %d\n", len(output))

	//Check that the output slice has the correct size
	if len(output) != int(float64(len(input))*float64(outputSampleRate)/float64(inputSampleRate)) {
		t.Errorf("Incorrect output slice size")
	}

	//check that the output data is not the same as the input data
	for i := range input {
		if input[i] != output[i] {
			return
		}
	}
	t.Errorf("Output data is the same as the input data")
}
