package mulaw

import (
	"log"

	"github.com/sopro-dev/sopro-core/audio"
	"github.com/sopro-dev/sopro-core/audio/utils"
)

type MuLawFormat struct{}

func (f *MuLawFormat) Decode(data []byte, info audio.AudioInfo) []byte {
	pcmData := make([]byte, len(data)*2)
	// remember that each byte need to be converted to two bytes
	for i := 0; i < len(data); i++ {

		newFrame := utils.DecodeULawToPCM(data[i : i+1])
		pcmData[i*2] = newFrame[0]
		pcmData[i*2+1] = newFrame[1]

	}
	if info.Verbose {
		log.Println("[MuLaw][Decode] Decoded to PCM")
		log.Println("[MuLaw][Decode] PCM Data [0  :100]", pcmData[0:100])
		log.Println("[MuLaw][Decode] PCM Data [100:200]", pcmData[100:200])
	}

	return pcmData
}

func (f *MuLawFormat) Encode(audioData []byte, info audio.AudioInfo) []byte {
	log.Printf("Not implemented")
	return nil
}
