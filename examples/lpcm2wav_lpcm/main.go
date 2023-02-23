package main

import (
	"bytes"
	"math"
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

	// generate sin wave in slice of bytes
	data := []byte{}
	for i := 0; i < 10000; i++ {
		// sin wave
		data = append(data, byte(128+127*math.Sin(2*math.Pi*float64(i)/100)))
	}

	// Open the input file
	in := bytes.NewBuffer(data)

	// Create the output file
	out, err := os.Create("./internal/samples/output_lpcm.wav")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// create a transcoder
	t := &transcoder.Transcoder{
		MethodT: method.NOT_FILLED,
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
	err = t.Pcm2Wav(
		&sopro.In{
			Data: in,
			AudioFileGeneral: sopro.AudioFileGeneral{
				Format: fileformat.AUDIO_PCM,
				Config: audioconfig.PcmConfig{
					BitDepth:   8,
					Channels:   1,
					Encoding:   encoding.SPACE_LINEAR, // ulaw is logarithmic
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
