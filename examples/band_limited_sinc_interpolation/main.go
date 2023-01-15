package main

import (
	"os"

	"github.com/pablodz/sopro/pkg/audioconfig"
	"github.com/pablodz/sopro/pkg/cpuarch"
	"github.com/pablodz/sopro/pkg/encoding"
	"github.com/pablodz/sopro/pkg/fileformat"
	"github.com/pablodz/sopro/pkg/resampler"
	"github.com/pablodz/sopro/pkg/sopro"
)

func main() {
	// Open the input file
	in, err := os.Open("./internal/samples/v1_16b_16000.wav")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	// Create the output file
	out, err := os.Create("./internal/samples/v1_16b_8000_2.wav")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// create a transcoder
	r := &resampler.Resampler{
		MethodR: resampler.BAND_LIMITED_INTERPOLATION,
		InConfigs: sopro.AudioConfig{
			Endianness: cpuarch.LITTLE_ENDIAN,
		},
		OutConfigs: sopro.AudioConfig{
			Endianness: cpuarch.LITTLE_ENDIAN,
		},
		SizeBuffer: 1024,
		Verbose:    true,
	}

	// Transcode the file
	err = r.Wav(
		&sopro.In{
			Data: in,
			AudioFileGeneral: sopro.AudioFileGeneral{
				Format: fileformat.AUDIO_WAV,
				Config: audioconfig.WavConfig{
					BitDepth:   16,
					Channels:   1,
					Encoding:   encoding.SPACE_LINEAR, // ulaw is logarithmic
					SampleRate: 16000,
				},
			},
		},
		&sopro.Out{
			Data: out,
			AudioFileGeneral: sopro.AudioFileGeneral{
				Format: fileformat.AUDIO_WAV,
				Config: audioconfig.WavConfig{
					BitDepth:   16,
					Channels:   1,
					Encoding:   encoding.SPACE_LINEAR,
					SampleRate: 8000,
				},
			},
		},
	)

	if err != nil {
		panic(err)
	}
}
