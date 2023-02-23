package transcoder

import (
	"fmt"

	"github.com/pablodz/sopro/pkg/audioconfig"
	"github.com/pablodz/sopro/pkg/encoding"
	"github.com/pablodz/sopro/pkg/method"
	"github.com/pablodz/sopro/pkg/sopro"
)

func (t *Transcoder) Mulaw2Wav(in *sopro.In, out *sopro.Out) error {
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
			sopro.ErrUnsupportedTranscoding,
			encoding.ENCODINGS[inSpace],
			encoding.ENCODINGS[outSpace],
		)

	}
}

func (t *Transcoder) Pcm2Wav(in *sopro.In, out *sopro.Out) error {
	inSpace := in.Config.(audioconfig.PcmConfig).Encoding
	outSpace := out.Config.(audioconfig.WavConfig).Encoding

	switch {
	case t.MethodT == method.NOT_FILLED &&
		inSpace == encoding.SPACE_LINEAR &&
		outSpace == encoding.SPACE_LINEAR:
		return lpcm2WavLpcm(in, out, t)
	default:
		return fmt.Errorf(
			"[%s] %s: %s -> %s",
			method.METHODS[t.MethodT],
			sopro.ErrUnsupportedTranscoding,
			encoding.ENCODINGS[inSpace],
			encoding.ENCODINGS[outSpace],
		)

	}
}
