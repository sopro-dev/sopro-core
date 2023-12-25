package audio

type AudioInfo struct {
	SampleRate  int
	Channels    int
	BitDepth    int
	FloatFormat bool
	Verbose     bool
}

type AudioFormat interface {
	Decode(data []byte, info AudioInfo) []byte
	Encode(audioData []byte, info AudioInfo) []byte
}
