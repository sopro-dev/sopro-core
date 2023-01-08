package main

import (
	"os"

	"github.com/pablodz/sopro/sopro/audioconfig"
	"github.com/pablodz/sopro/sopro/cpuarch"
	"github.com/pablodz/sopro/sopro/encoding"
	"github.com/pablodz/sopro/sopro/fileformat"
	"github.com/pablodz/sopro/sopro/method"
	"github.com/pablodz/sopro/sopro/transcoder"
)

func main() {

	// Open the input file
	in, err := os.Open("./internal/samples/recording.ulaw")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	// Create the output file
	out, err := os.Create("./internal/samples/result_sample_ulaw_mono_8000_be.wav")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// create a transcoder
	t := &transcoder.Transcoder{
		Method: method.BIT_TABLE,
		SourceConfigs: transcoder.TranscoderAudioConfig{
			Endianness: cpuarch.LITTLE_ENDIAN,
		},
		TargetConfigs: transcoder.TranscoderAudioConfig{
			Endianness: cpuarch.LITTLE_ENDIAN,
		},
		SizeBufferToProcess: 1024,
		Verbose:             true,
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
