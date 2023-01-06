package transcoder

// func TestMulaw2Wav(t *testing.T) {
// 	// Open the input file
// 	in, err := os.Open("../../internal/samples/recording.ulaw")
// 	if err != nil {
// 		t.Fatalf("error opening input file: %v", err)
// 	}
// 	defer in.Close()

// 	// Create the output file
// 	out, err := os.Create("../../internal/samples/result_sample_ulaw_mono_8000_be.wav")
// 	if err != nil {
// 		t.Fatalf("error creating output file: %v", err)
// 	}
// 	defer out.Close()

// 	// Transcode the file
// 	err = Mulaw2Wav(
// 		&AudioFileIn{
// 			Data: in,
// 			AudioFileGeneral: AudioFileGeneral{
// 				Format: fileformat.AUDIO_MULAW,
// 				Config: audioconfig.MulawConfig{
// 					BitDepth:   8,
// 					Channels:   1,
// 					Encoding:   encoding.SPACE_ULAW,
// 					SampleRate: 8000,
// 				},
// 			},
// 		},
// 		&AudioFileOut{
// 			Data: out,
// 			AudioFileGeneral: AudioFileGeneral{
// 				Format: fileformat.AUDIO_WAV,
// 				Config: audioconfig.WavConfig{
// 					BitDepth:   8,
// 					Channels:   1,
// 					Encoding:   encoding.SPACE_LINEAR,
// 					SampleRate: 8000,
// 				},
// 			},
// 		},
// 		&TranscoderOneToOne{
// 			Method: method.BIT_TABLE,
// 			SourceConfigs: TranscoderAudioConfig{
// 				Encoding:   encoding.SPACE_ULAW,
// 				Endianness: cpuarch.LITTLE_ENDIAN,
// 			},
// 			TargetConfigs: TranscoderAudioConfig{
// 				Encoding:   encoding.SPACE_LINEAR,
// 				Endianness: cpuarch.LITTLE_ENDIAN,
// 			},
// 			BitDepth: BIT_8,
// 		},
// 	)
// 	if err != nil {
// 		t.Fatalf("error transcoding file: %v", err)
// 	}
// }
