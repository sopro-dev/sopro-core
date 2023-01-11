package transcoder

import (
	"fmt"

	"github.com/pablodz/sopro/pkg/audioconfig"
	"github.com/pablodz/sopro/pkg/encoding"
	"github.com/pablodz/sopro/pkg/method"
	"github.com/pablodz/sopro/pkg/resampler"
)

var WIDTH_TERMINAL = 80
var HEIGHT_TERMINAL = 30

const ErrUnsupportedConversion = "unsupported conversion"

func (t *Transcoder) Mulaw2Wav(in *AudioFileIn, out *AudioFileOut) error {

	inSpace := in.Config.(audioconfig.MulawConfig).Encoding
	outSpace := out.Config.(audioconfig.WavConfig).Encoding

	switch {
	case t.MethodT == method.BIT_LOOKUP_TABLE &&
		inSpace == encoding.SPACE_LOGARITHMIC &&
		outSpace == encoding.SPACE_LINEAR:
		return mulaw2WavLpcm(in, out, t)
	case t.MethodT == method.BIT_LOOKUP_TABLE &&
		inSpace == encoding.SPACE_LOGARITHMIC &&
		outSpace == encoding.SPACE_LOGARITHMIC:
		return mulaw2WavLogpcm(in, out, t)
	default:
		return fmt.Errorf(
			"[%s] %s: %s -> %s",
			method.METHODS[t.MethodT],
			ErrUnsupportedConversion,
			encoding.ENCODINGS[inSpace],
			encoding.ENCODINGS[outSpace],
		)

	}
}

func (t *Transcoder) Wav2Wav(in *AudioFileIn, out *AudioFileOut) error {

	inSpace := in.Config.(audioconfig.WavConfig).Encoding
	outSpace := out.Config.(audioconfig.WavConfig).Encoding

	switch {
	case t.MethodR == resampler.LINEAR_INTERPOLATION &&
		inSpace == encoding.SPACE_LINEAR &&
		outSpace == encoding.SPACE_LINEAR:
		return wavLpcm2wavLpcm(in, out, t)
	case t.MethodR == resampler.LINEAR_INTERPOLATION &&
		inSpace == encoding.SPACE_LOGARITHMIC &&
		outSpace == encoding.SPACE_LOGARITHMIC:
		fallthrough
	default:
		return fmt.Errorf(
			"%s: %s -> %s",
			ErrUnsupportedConversion,
			encoding.ENCODINGS[inSpace],
			encoding.ENCODINGS[outSpace],
		)

	}
}
