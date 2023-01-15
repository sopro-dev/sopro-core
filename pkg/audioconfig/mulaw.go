package audioconfig

// MulawConfig is a struct that contains the configuration of a MULAW (RAW) audio file.
type MulawConfig struct {
	SampleRate int // the sample rate in Hertz
	BitDepth   int // the bit depth (e.g. 8, 16, 24)
	Channels   int // the number of channels (e.g. 1, 2)
	Encoding   int // the encoding format (e.g. "UNSIGNED", "SIGNED", "FLOAT")
}
