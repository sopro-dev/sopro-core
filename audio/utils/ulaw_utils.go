package utils

func DecodeFromULaw(uLaw byte) int16 {
	pcm := ulawDecode[uLaw]
	return int16(pcm)
}
