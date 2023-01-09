package decoder

import "fmt"

// Decodes only raw data
// log pcm              -> linear pcm
// 8 bit log pcm (ulaw) -> 16 bit linear pcm
func DecodeFrameUlaw2Lpcm(pcm []byte) ([]byte, error) {
	if len(pcm) == 0 {
		return []byte{}, fmt.Errorf("pcm is empty")
	}
	lpcm := make([]byte, len(pcm)*2)
	for i, frame := range pcm {
		lpcmFrame := ulaw2lpcm[frame]
		copy(lpcm[i*2:], []byte{byte(lpcmFrame), byte(lpcmFrame >> 8)})
	}
	return lpcm, nil
}
