package transcoder

type Transcoder struct {
	Method          int                   // the method of transcoding (e.g. 1, 2, 3, etc.)
	AdvancedConfigs interface{}           // the specific configuration options for the transcoding method
	SourceConfigs   TranscoderAudioConfig // the source configuration
	TargetConfigs   TranscoderAudioConfig // the target configuration
	BitDepth        int                   // the bit depth (e.g. 8, 16, 24) Needs to be equal for source and target
	Verbose         bool                  // if true, the transcoding process will be verbose
}

type TranscoderAudioConfig struct {
	Encoding   int // the encoding format (e.g. "UNSIGNED", "SIGNED", "FLOAT", "PCM", "IEEE_FLOAT", "MULAW")
	Endianness int // the endianness of the audio data (e.g. "BIG", "LITTLE")
}
