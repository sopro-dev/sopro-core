package transcoder

func (t *Transcoder) Mulaw2Wav(in *AudioFileIn, out *AudioFileOut) error {
	return mulaw2Wav(in, out, t)
}
