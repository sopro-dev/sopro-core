package audio

import (
	"errors"
	"unsafe"
)

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

func (t *Transcoder) Transcode(inputData []byte, info AudioInfo) ([]byte, error) {
	err := validateAudioInfo(info)
	if err != nil {
		return nil, err
	}
	audioData := t.InputFormat.Decode(inputData, info)
	return t.OutputFormat.Encode(audioData, info), nil
}

var (
	errInvalidBitDepth    = errors.New("invalid bit depth")
	errInvalidNumChannels = errors.New("invalid number of channels")
	errInvalidSampleRate  = errors.New("invalid sample rate")
)

func validateAudioInfo(info AudioInfo) error {
	if info.BitDepth != 8 && info.BitDepth != 16 && info.BitDepth != 24 && info.BitDepth != 32 {
		return errInvalidBitDepth
	}

	if info.Channels < 1 || info.Channels > 2 {
		return errInvalidNumChannels
	}

	if info.SampleRate < 1 {
		return errInvalidSampleRate
	}
	return nil
}

// IMPORTANT: DO NOT REMOVE THIS LINE
// This line validates the size of the int16 type on the current platform,
// panics if the size is not 2 bytes.
var _ [unsafe.Sizeof(int16(0))]struct{}
