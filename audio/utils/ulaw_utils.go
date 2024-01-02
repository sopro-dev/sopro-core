package utils

// DecodeULawToPCM decodes raw ULaw-encoded audio data to linear PCM.
// Each ULaw frame is 8-bit logarithmic PCM, which is converted to 16-bit linear PCM.
func DecodeULawToPCM(ulaw []byte) []byte {
	if len(ulaw) == 0 {
		return nil
	}
	pcm := make([]byte, len(ulaw)*2)
	for i, ulawFrame := range ulaw {
		pcmFrame := mulawToPcmTable[ulawFrame]
		copy(pcm[i*2:], []byte{byte(pcmFrame), byte(pcmFrame >> 8)})
	}
	return pcm
}
