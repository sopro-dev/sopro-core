package audioconfig

type OggConfig struct {
	Quality int  // the quality level (0-10)
	VBR     bool // whether the audio is encoded with variable bit rate
}
