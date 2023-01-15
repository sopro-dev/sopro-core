package sopro

import (
	"bufio"
	"io"
	"os"
)

// In is a struct that contains the input audio file or stream.
type In struct {
	Data             io.Reader     // the audio data
	Reader           *bufio.Reader //  reader
	AudioFileGeneral               // general information about the audio file
}

// Out is a struct that contains the output audio file or stream.
type Out struct {
	Data             io.Writer     // the audio data
	Writer           *bufio.Writer // writer
	AudioFileGeneral               // general information about the audio file
}

// AudioFileGeneral is a struct that contains general information about the audio file or stream.
type AudioFileGeneral struct {
	File   *os.File    // the audio file
	Format string      // the audio file format (e.g. "mp3", "ogg", "wav", etc.)
	Length int         // the length of the audio in seconds
	Config interface{} // the specific configuration options for the audio file format
}

// AudioConfig is a struct that contains the configuration
// of the source and target audio files or streams.
type AudioConfig struct {
	Encoding   int // the encoding format (e.g. "UNSIGNED", "SIGNED", "FLOAT", "PCM", "IEEE_FLOAT", "MULAW")
	Endianness int // the endianness of the audio data (e.g. "BIG", "LITTLE")
}

func (a *AudioFileGeneral) NewConfig(config interface{}) {
	a.Config = config
}
