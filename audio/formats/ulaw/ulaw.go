package ulaw

import (
	"math"

	"github.com/sopro-dev/sopro-core/audio"
	"github.com/sopro-dev/sopro-core/audio/utils"
)

type ULawFormat struct{}

func (f *ULawFormat) Decode(data []byte, info audio.AudioInfo) []float64 {
	numSamples := len(data)
	pcm := make([]float64, numSamples/info.Channels)

	for i := 0; i < numSamples; i += info.Channels {
		// Process each channel separately
		for ch := 0; ch < info.Channels; ch++ {
			pcm[i/info.Channels] = float64(utils.DecodeFromULaw(data[i+ch])) / 32768.0
		}
	}

	return pcm
}

func (f *ULawFormat) Encode(audioData []float64, info audio.AudioInfo) []byte {
	encoded := make([]byte, len(audioData)*info.Channels)

	for i, sample := range audioData {
		sampleInt := int16(math.Max(-1.0, math.Min(sample, 1.0)) * 32768.0)
		for ch := 0; ch < info.Channels; ch++ {
			encoded[i*info.Channels+ch] = utils.EncodeToULaw(sampleInt)
		}
	}

	return encoded
}
