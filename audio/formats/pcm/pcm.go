// audio/formats/pcm/pcm.go
package pcm

import (
	"math"

	"github.com/sopro-dev/sopro-core/audio"
)

type PCMFormat struct{}

func (f *PCMFormat) Decode(data []byte, info audio.AudioInfo) []float64 {
	sampleSize := info.BitDepth / 8
	numSamples := len(data) / (sampleSize * info.Channels)
	pcm := make([]float64, numSamples*info.Channels)

	for i := 0; i < numSamples; i++ {
		for ch := 0; ch < info.Channels; ch++ {
			startIndex := (i*info.Channels + ch) * sampleSize
			endIndex := startIndex + sampleSize
			sample := int64(0)

			for j := startIndex; j < endIndex; j++ {
				sample |= int64(data[j]) << ((j - startIndex) * 8)
			}

			divider := math.Pow(2, float64(info.BitDepth-1))
			if info.FloatFormat {
				divider = math.Pow(2, float64(info.BitDepth))
			}

			pcm[i*info.Channels+ch] = float64(sample) / divider
		}
	}

	return pcm
}

func (f *PCMFormat) Encode(audioData []float64, info audio.AudioInfo) []byte {
	sampleSize := info.BitDepth / 8
	numSamples := len(audioData) / info.Channels
	encoded := make([]byte, numSamples*info.Channels*sampleSize)

	for i := 0; i < numSamples; i++ {
		for ch := 0; ch < info.Channels; ch++ {
			sampleIndex := i*info.Channels + ch
			sample := audioData[sampleIndex]

			sample = math.Max(-1.0, math.Min(sample, 1.0))

			if !info.FloatFormat {
				sample *= 0.5
			}
			sampleInt := int64((1 << (info.BitDepth - 1)) * int(sample))
			for j := 0; j < sampleSize; j++ {
				shift := uint(j * 8)
				encoded[i*info.Channels*sampleSize+ch*sampleSize+j] = byte((sampleInt >> shift) & 0xFF)
			}
		}
	}

	return encoded
}
