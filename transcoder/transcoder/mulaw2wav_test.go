package transcoder

import (
	"os"
	"testing"

	"github.com/pablodz/transcoder/transcoder/audioconfig"
	"github.com/pablodz/transcoder/transcoder/fileformat"
)

func TestMulaw2Wav(t *testing.T) {
	// Open the input file
	in, err := os.Open("../../internal/samples/sample_ulaw_mono_8000_be.ulaw")
	if err != nil {
		t.Fatalf("error opening input file: %v", err)
	}
	defer in.Close()

	// Create the output file
	out, err := os.Create("../../internal/samples/result_sample_ulaw_mono_8000_be.wav")
	if err != nil {
		t.Fatalf("error creating output file: %v", err)
	}
	defer out.Close()

	// Transcode the file
	err = Mulaw2Wav(
		&AudioFileIn{
			Data: in,
			AudioFileGeneral: AudioFileGeneral{
				Format: fileformat.AUDIO_MULAW,
				Config: audioconfig.MulawConfig{
					SampleRate: 8000,
					BitDepth:   8,
					Channels:   1,
					Endianness: audioconfig.MULAW_BIG_ENDIAN,
					Encoding:   audioconfig.MULAW_ENCODING_MULAW,
				},
			},
		},
		&AudioFileOut{
			Data: out,
			AudioFileGeneral: AudioFileGeneral{
				Format: fileformat.AUDIO_WAV,
				Config: audioconfig.WavConfig{
					Channels:   1,
					SampleRate: 8000,
					BitDepth:   16,
					Encoding:   audioconfig.PCM_ENCODING_SIGNED_16_PCM,
				},
			},
		})
	if err != nil {
		t.Fatalf("error transcoding file: %v", err)
	}
}
