package audioconfig

const (
	PCM_ENCODING_SIGNED_8_PCM  = "SIGNED_8_PCM"
	PCM_ENCODING_SIGNED_16_PCM = "SIGNED_16_PCM"
	PCM_ENCODING_SIGNED_24_PCM = "SIGNED_24_PCM"
)

type WavConfig struct {
	SampleRate int    // the sample rate in Hertz
	BitDepth   int    // the bit depth (e.g. 8, 16, 24)
	Encoding   string // the encoding format (e.g. "PCM", "IEEE_FLOAT")
	Channels   int    // the number of channels (e.g. 1, 2)
}
