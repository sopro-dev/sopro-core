package main

import (
	"os"

	"github.com/pablodz/transcoder/transcoder/audioconfig"
	"github.com/pablodz/transcoder/transcoder/cpuarch"
	"github.com/pablodz/transcoder/transcoder/encoding"
	"github.com/pablodz/transcoder/transcoder/fileformat"
	"github.com/pablodz/transcoder/transcoder/method"
	"github.com/pablodz/transcoder/transcoder/transcoder"
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

	// Transcode the file
	err = transcoder.Mulaw2Wav(
		&transcoder.AudioFileIn{
			Data: in,
			AudioFileGeneral: transcoder.AudioFileGeneral{
				Format: fileformat.AUDIO_MULAW,
				Config: audioconfig.MulawConfig{
					BitDepth:   8,
					Channels:   1,
					Encoding:   encoding.SPACE_ULAW,
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
		&transcoder.TranscoderOneToOne{
			Method: method.BIT_TABLE,
			SourceConfigs: transcoder.TranscoderAudioConfig{
				Encoding:   encoding.SPACE_ULAW,
				Endianness: cpuarch.LITTLE_ENDIAN,
			},
			TargetConfigs: transcoder.TranscoderAudioConfig{
				Encoding:   encoding.SPACE_LINEAR,
				Endianness: cpuarch.LITTLE_ENDIAN,
			},
			BitDepth: transcoder.BIT_8,
			Verbose:  true,
		},
	)
	if err != nil {
		panic(err)
	}

}
