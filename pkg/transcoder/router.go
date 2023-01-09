package transcoder

import (
	"fmt"

	"github.com/pablodz/sopro/pkg/audioconfig"
	"github.com/pablodz/sopro/pkg/encoding"
)

const ErrUnsupportedConversion = "unsupported conversion"

func (t *Transcoder) Mulaw2Wav(in *AudioFileIn, out *AudioFileOut) error {

	inSpace := in.Config.(audioconfig.MulawConfig).Encoding
	outSpace := out.Config.(audioconfig.WavConfig).Encoding

	switch {
	case inSpace == encoding.SPACE_LOGARITHMIC && outSpace == encoding.SPACE_LINEAR:
		return mulaw2WavLpcm(in, out, t)
	case inSpace == encoding.SPACE_LOGARITHMIC && outSpace == encoding.SPACE_LOGARITHMIC:
		return mulaw2WavLogpcm(in, out, t)
	default:
		return fmt.Errorf(
			"%s: %s -> %s",
			ErrUnsupportedConversion,
			encoding.ENCODINGS[inSpace],
			encoding.ENCODINGS[outSpace],
		)

	}
}
