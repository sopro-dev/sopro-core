package models

import (
	"bufio"
	"io"
	"os"
)

type Transcoder struct {
	Method                int                   // the method of transcoding (e.g. 1, 2, 3, etc.)
	MethodAdvancedConfigs interface{}           // the specific configuration options for the transcoding method
	SizeBufferToProcess   int                   // the size of the buffer to read from the input file. Default is 1024
	SourceConfigs         TranscoderAudioConfig // the source configuration
	TargetConfigs         TranscoderAudioConfig // the target configuration
	BitDepth              int                   // the bit depth (e.g. 8, 16, 24) Needs to be equal for source and target
	Verbose               bool                  // if true, the transcoding process will be verbose
}

type TranscoderAudioConfig struct {
	Encoding   int // the encoding format (e.g. "UNSIGNED", "SIGNED", "FLOAT", "PCM", "IEEE_FLOAT", "MULAW")
	Endianness int // the endianness of the audio data (e.g. "BIG", "LITTLE")
}

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
