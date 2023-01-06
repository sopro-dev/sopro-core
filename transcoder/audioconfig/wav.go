package audioconfig

type WavConfig struct {
	SampleRate int // the sample rate in Hertz
	BitDepth   int // the bit depth (e.g. 8, 16, 24)
	Channels   int // the number of channels (e.g. 1, 2)
	Encoding   int // the encoding format (e.g. "PCM", "IEEE_FLOAT")
}
