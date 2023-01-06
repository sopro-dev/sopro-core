package transcoder

import (
	"fmt"

	"github.com/pablodz/transcoder/transcoder/encoding"
	"github.com/pablodz/transcoder/transcoder/method"
)

// Function that find the ideal method for transcoding your samples
func bitTranscoder(b []byte, transcoder *TranscoderOneToOne) func() ([]byte, error) {
	return func() ([]byte, error) {
		switch transcoder.Method {
		case method.BIT_TABLE:
			// Bit table: use a slice to store the values
			switch transcoder.BitDepth {
			case BIT_8:
				// Only 8 bit depth
				switch {
				case transcoder.SourceConfigs.Encoding == encoding.SPACE_ULAW && transcoder.TargetConfigs.Encoding == encoding.SPACE_LINEAR:
					// ulaw -> linear
					// return g711.DecodeUlaw(b), nil
					// transcoder.Println("[debug transcoding] ulaw -> linear (8 bit depth)")
					return table_ulaw_8_linear(
						b,
						transcoder.SourceConfigs.Encoding,
						transcoder.TargetConfigs.Encoding,
					), nil
				}

			}
		}

		transcoder.Println("[error transcoding] method not found")
		return []byte{}, fmt.Errorf(
			"[error transcoding]%v[source]%v[target]%v[bitDepth]%v",
			"method not found",
			transcoder.SourceConfigs.Encoding,
			transcoder.TargetConfigs.Encoding,
			transcoder.BitDepth,
		)
	}
}
