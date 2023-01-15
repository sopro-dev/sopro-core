package transcoder

import "github.com/pablodz/sopro/pkg/sopro"

// TODO: add functions to validate the input and output configuration
// ex. validate formats
// ex. validate sample rate
// ex. validate bit depth
// ex. validate channels
// ex. validate endianness
// ex. validate encoding
// ex. validate file size
// validate not 1-1 transcoding inputs as 8k output as 16k

func ValidateInput(in *sopro.In) error {
	return nil
}

func ValidateOutput(out *sopro.Out) error {
	return nil
}

func ValidateTranscoder(tr *Transcoder) error {
	return nil
}
