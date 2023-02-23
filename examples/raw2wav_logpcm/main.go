package main

import (
	"os"

	"github.com/pablodz/sopro/pkg/audioconfig"
	"github.com/pablodz/sopro/pkg/cpuarch"
	"github.com/pablodz/sopro/pkg/encoding"
	"github.com/pablodz/sopro/pkg/fileformat"
	"github.com/pablodz/sopro/pkg/method"
	"github.com/pablodz/sopro/pkg/sopro"
	"github.com/pablodz/sopro/pkg/transcoder"
)

func main() {
	// Open the input file
	in, err := os.Open("./internal/samples/recording.ulaw")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	// Create the output file
	out, err := os.Create("./internal/samples/result_sample_ulaw_mono_8000_le_logpcm.wav")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// create a transcoder
	t := &transcoder.Transcoder{
		MethodT: method.BIT_LOOKUP_TABLE,
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
	err = t.Mulaw2Wav(
		&sopro.In{
			Data: in,
			AudioFileGeneral: sopro.AudioFileGeneral{
				Format: fileformat.AUDIO_MULAW,
				Config: audioconfig.MulawConfig{
					BitDepth:   8,
					Channels:   1,
					Encoding:   encoding.SPACE_LOGARITHMIC, // ulaw is logarithmic
					SampleRate: 8000,
				},
			},
		},
		&sopro.Out{
			Data: out,
			AudioFileGeneral: sopro.AudioFileGeneral{
				Format: fileformat.AUDIO_WAV,
				Config: audioconfig.WavConfig{
					BitDepth:   8,
					Channels:   1,
					Encoding:   encoding.SPACE_LOGARITHMIC,
					SampleRate: 8000,
				},
			},
		},
	)

	if err != nil {
		panic(err)
	}
}
