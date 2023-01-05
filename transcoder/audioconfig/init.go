package audioconfig

type AudioFile struct {
	Format string      // the audio file format (e.g. "mp3", "ogg", "wav", etc)
	Data   []byte      // the audio data
	Length float64     // the length of the audio in seconds
	Config interface{} // the specific configuration options for the audio file format
}
