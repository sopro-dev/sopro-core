package resampler

import (
	"fmt"

	"github.com/pablodz/sopro/pkg/audioconfig"
	"github.com/pablodz/sopro/pkg/encoding"
	"github.com/pablodz/sopro/pkg/sopro"
)

<<<<<<< HEAD
func (rs *Resampler) Wav(in *sopro.In, out *sopro.Out) error {
=======
func (r *Resampler) Wav(in *sopro.In, out *sopro.Out) error {
>>>>>>> d598982 (refactor, new resampler, new transcoder, sopro models, sinc interpolation and more examples)
	inSpace := in.Config.(audioconfig.WavConfig).Encoding
	outSpace := out.Config.(audioconfig.WavConfig).Encoding

	switch {
<<<<<<< HEAD
	case rs.MethodR == LINEAR_INTERPOLATION &&
		inSpace == encoding.SPACE_LINEAR &&
		outSpace == encoding.SPACE_LINEAR:
		return linearPcm(in, out, rs)
	case rs.MethodR == BAND_LIMITED_INTERPOLATION &&
		inSpace == encoding.SPACE_LINEAR &&
		outSpace == encoding.SPACE_LINEAR:
		return linearPcm(in, out, rs)
	case rs.MethodR == FRACTIONAL_DELAY_FILTER &&
		inSpace == encoding.SPACE_LINEAR &&
		outSpace == encoding.SPACE_LINEAR:
		return linearPcm(in, out, rs)
	case rs.MethodR == LINEAR_INTERPOLATION &&
=======
	case r.MethodR == LINEAR_INTERPOLATION &&
		inSpace == encoding.SPACE_LINEAR &&
		outSpace == encoding.SPACE_LINEAR:
		return linearPcm(in, out, r)
	case r.MethodR == BAND_LIMITED_INTERPOLATION &&
		inSpace == encoding.SPACE_LINEAR &&
		outSpace == encoding.SPACE_LINEAR:
		return linearPcm(in, out, r)
	case r.MethodR == LINEAR_INTERPOLATION &&
>>>>>>> d598982 (refactor, new resampler, new transcoder, sopro models, sinc interpolation and more examples)
		inSpace == encoding.SPACE_LOGARITHMIC &&
		outSpace == encoding.SPACE_LOGARITHMIC:
		fallthrough
	default:
		return fmt.Errorf(
			"%s: %s -> %s",
			sopro.ErrUnsupportedTranscoding,
			encoding.ENCODINGS[inSpace],
			encoding.ENCODINGS[outSpace],
		)

	}
}
