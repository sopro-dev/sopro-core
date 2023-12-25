package mulaw

import (
	"log"

	"github.com/sopro-dev/sopro-core/audio"
	"github.com/sopro-dev/sopro-core/audio/utils"
)

type MuLawFormat struct{}

func (f *MuLawFormat) Decode(data []byte, info audio.AudioInfo) []float64 {

	pcmData := make([]float64, len(data))
	for i := 0; i < len(data); i++ {
		pcmData[i] = float64(utils.DecodeFromULaw(data[i]))
	}

	log.Println("[Bytes][Org]", data[0:100])
	log.Println("[Bytes][Pcm]", pcmData[0:100])

	return pcmData
}

func (f *MuLawFormat) Encode(audioData []float64, info audio.AudioInfo) []byte {
	log.Printf("Not implemented")
	return nil
}
