package main

import (
	"bytes"
	"os"

	"github.com/pablodz/sopro/pkg/audioconfig"
	"github.com/pablodz/sopro/pkg/cpuarch"
	"github.com/pablodz/sopro/pkg/encoding"
	"github.com/pablodz/sopro/pkg/fileformat"
	"github.com/pablodz/sopro/pkg/method"
	"github.com/pablodz/sopro/pkg/transcoder"
)

func main() {
	data := []byte{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 4, 2, 3, 3, 5, 1, 7, 8, 1, 4, 0}
	// Open the input file
	in := bytes.NewBuffer(data)

	// Create the output file
	out, err := os.Create("./internal/samples/output.wav")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// create a transcoder
	t := &transcoder.Transcoder{
		MethodT: method.BIT_LOOKUP_TABLE,
		SourceConfigs: transcoder.TranscoderAudioConfig{
			Endianness: cpuarch.LITTLE_ENDIAN,
		},
		TargetConfigs: transcoder.TranscoderAudioConfig{
			Endianness: cpuarch.LITTLE_ENDIAN,
		},
		SizeBuffer: 1024,
		Verbose:    true,
	}

	// Transcode the file
	err = t.Mulaw2Wav(
		&transcoder.AudioFileIn{
			Data: in,
			AudioFileGeneral: transcoder.AudioFileGeneral{
				Format: fileformat.AUDIO_MULAW,
				Config: audioconfig.MulawConfig{
					BitDepth:   8,
					Channels:   1,
					Encoding:   encoding.SPACE_LOGARITHMIC, // ulaw is logarithmic
					SampleRate: 8000,
				},
			},
		},
		&transcoder.AudioFileOut{
			Data: out,
			AudioFileGeneral: transcoder.AudioFileGeneral{
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
