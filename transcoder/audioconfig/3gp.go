package audioconfig

type ThreeGpConfig struct {
	AudioCodec string // the audio codec used (e.g. "AMR-NB", "AMR-WB", "AAC")
	SampleRate int    // the sample rate in Hertz
	BitRate    int    // the bit rate in bits per second
	Channels   int    // the number of channels (1 for mono, 2 for stereo)
}
