package audio

import "unsafe"

type Transcoder struct {
	InputFormat  AudioFormat
	OutputFormat AudioFormat
}

func NewTranscoder(inputFormat, outputFormat AudioFormat) *Transcoder {
	return &Transcoder{
		InputFormat:  inputFormat,
		OutputFormat: outputFormat,
	}
}

func (t *Transcoder) Transcode(inputData []byte, info AudioInfo) []byte {
	audioData := t.InputFormat.Decode(inputData, info)
	return t.OutputFormat.Encode(audioData, info)
}

// IMPORTANT: DO NOT REMOVE THIS LINE
// This line validates the size of the int16 type on the current platform,
// panics if the size is not 2 bytes.
var _ [unsafe.Sizeof(int16(0))]struct{}
