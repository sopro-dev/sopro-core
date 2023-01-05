package audioconfig

const (
	// Endianness
	MULAW_BIG_ENDIAN    = "BIG"
	MULAW_LITTLE_ENDIAN = "LITTLE"

	// Encoding formats
	MULAW_ENCODING_UNSIGNED       = "UNSIGNED"
	MULAW_ENCODING_SIGNED         = "SIGNED"
	MULAW_ENCODING_FLOAT_32       = "FLOAT_32"
	MULAW_ENCODING_FLOAT_64       = "FLOAT_64"
	MULAW_ENCODING_SIGNED_8_PCM   = "SIGNED_8_PCM"
	MULAW_ENCODING_SIGNED_16_PCM  = "SIGNED_16_PCM"
	MULAW_ENCODING_SIGNED_24_PCM  = "SIGNED_24_PCM"
	MULAW_ENCODING_UNSIGNED_8_PCM = "UNSIGNED_8_PCM"
	MULAW_ENCODING_MULAW          = "ULAW"
	MULAW_ENCODING_ALAW           = "ALAW"
	// TODO: Add more encoding formats (+12)
)

type MulawConfig struct {
	SampleRate int    // the sample rate in Hertz
	BitDepth   int    // the bit depth (e.g. 8, 16, 24)
	Channels   int    // the number of channels (e.g. 1, 2)
	Endianness string // the endianness of the audio data (e.g. "BIG", "LITTLE")
	Encoding   string // the encoding format (e.g. "UNSIGNED", "SIGNED", "FLOAT")
}
