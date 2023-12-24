package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/sopro-dev/sopro-core/audio"
	"github.com/sopro-dev/sopro-core/audio/formats/pcm"
	"github.com/sopro-dev/sopro-core/audio/formats/ulaw"
)

func main() {
	numSamples := 80000
	numChannels := 2
	samplesPerSecond := 44100

	inputData := make([]byte, numSamples*numChannels)
	for i := 0; i < numSamples; i++ {
		// Calculate the current second
		currentSecond := i / samplesPerSecond

		// Generate noise for the first second, then silence for the second second
		if currentSecond%2 == 0 {
			sample := byte(127 * (2.0*rand.Float64() - 1.0)) // Generate random noise
			inputData[i*numChannels] = sample
			inputData[i*numChannels+1] = sample
		} else {
			inputData[i*numChannels] = 0   // Silence
			inputData[i*numChannels+1] = 0 // Silence
		}
	}

	audioInfo := audio.AudioInfo{
		SampleRate:  44100,
		Channels:    numChannels,
		BitDepth:    16,
		FloatFormat: false,
	}

	transcoder := audio.NewTranscoder(&pcm.PCMFormat{}, &ulaw.ULawFormat{})
	outputData := transcoder.Transcode(inputData, audioInfo)

	// Process the output data as needed...
	fmt.Println(outputData[:200])

	// Store the output data to a file...
	f, err := os.Create("output.ul")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.Write(outputData)
	if err != nil {
		panic(err)
	}
}
