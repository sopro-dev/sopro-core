package main

import (
	"io"
	"os"

	"github.com/sopro-dev/sopro-core/audio"
	"github.com/sopro-dev/sopro-core/audio/formats/pcm"
	mulaw "github.com/sopro-dev/sopro-core/audio/formats/ulaw"
	"github.com/sopro-dev/sopro-core/audio/utils"
)

func main() {
	// read file on "internal/samples/sample.ul"
	f, err := os.Open("internal/samples/sample.ul")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// copy all to an array
	inputData, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	audioInfo := audio.AudioInfo{
		SampleRate:  8000,
		Channels:    1,
		BitDepth:    8,
		FloatFormat: false,
		Verbose:     false,
	}

	transcoder := audio.NewTranscoder(&mulaw.MuLawFormat{}, &pcm.PCMFormat{})
	outputData, err := transcoder.Transcode(inputData, audioInfo)
	if err != nil {
		panic(err)
	}

	// Store the output data to a file...
	f, err = os.Create("output_converted.wav")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// create headers
	headers := utils.GenerateWavHeadersWithConfig(&utils.WavHeader{
		Length:     uint32(len(outputData) + 44),
		WaveFormat: utils.WAVE_FORMAT_PCM,
		Channels:   1,
		SampleRate: 8000,
		BitDepth:   16,
		Verbose:    audioInfo.Verbose,
	})

	f.Write(headers)
	f.Seek(44, 0)
	f.Write(outputData)
}
