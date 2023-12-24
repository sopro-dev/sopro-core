package audio

type AudioInfo struct {
	SampleRate  int
	Channels    int
	BitDepth    int
	FloatFormat bool
}

type AudioFormat interface {
	Decode(data []byte, info AudioInfo) []float64
	Encode(audioData []float64, info AudioInfo) []byte
}
