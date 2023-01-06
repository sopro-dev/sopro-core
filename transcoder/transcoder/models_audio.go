package transcoder

import (
	"bufio"
	"io"
	"os"
)

type AudioFileIn struct {
	Data             io.Reader     // the audio data
	Reader           *bufio.Reader //  reader
	AudioFileGeneral               // general information about the audio file
}

type AudioFileOut struct {
	Data             io.Writer     // the audio data
	Writer           *bufio.Writer // writer
	AudioFileGeneral               // general information about the audio file
}

type AudioFileGeneral struct {
	File   *os.File    // the audio file
	Format string      // the audio file format (e.g. "mp3", "ogg", "wav", etc.)
	Length int         // the length of the audio in seconds
	Config interface{} // the specific configuration options for the audio file format
}
