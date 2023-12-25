// audio/formats/pcm/pcm.go
package pcm

import (
	"log"

	"github.com/sopro-dev/sopro-core/audio"
)

type PCMFormat struct{}

func (f *PCMFormat) Decode(data []byte, info audio.AudioInfo) []float64 {
	log.Printf("Not implemented")
	return nil
}

func (f *PCMFormat) Encode(audioData []float64, info audio.AudioInfo) []byte {

	// convert float64 to byte
	data := make([]byte, len(audioData))
	for i := 0; i < len(audioData); i++ {
		data[i] = byte(audioData[i])
	}

	return data
}
