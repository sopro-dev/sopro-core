package main

import (
	"os"

	"github.com/pablodz/sopro/pkg/audioconfig"
	"github.com/pablodz/sopro/pkg/cpuarch"
	"github.com/pablodz/sopro/pkg/encoding"
	"github.com/pablodz/sopro/pkg/fileformat"
	"github.com/pablodz/sopro/pkg/resampler"
	"github.com/pablodz/sopro/pkg/transcoder"
)

func main() {

	// Open the input file
	in, err := os.Open("./internal/samples/v1_16b_16000.wav")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	// Create the output file
	out, err := os.Create("./internal/samples/v1_16b_8000.wav")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// create a transcoder
	t := &transcoder.Transcoder{
		MethodR: resampler.LINEAR_INTERPOLATION,
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
	err = t.Wav2Wav(
		&transcoder.AudioFileIn{
			Data: in,
			AudioFileGeneral: transcoder.AudioFileGeneral{
				Format: fileformat.AUDIO_WAV,
				Config: audioconfig.WavConfig{
					BitDepth:   16,
					Channels:   1,
					Encoding:   encoding.SPACE_LINEAR, // ulaw is logarithmic
					SampleRate: 16000,
				},
			},
		},
		&transcoder.AudioFileOut{
			Data: out,
			AudioFileGeneral: transcoder.AudioFileGeneral{
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
