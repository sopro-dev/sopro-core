package utils

func DecodeFromULaw(uLaw byte) int16 {
	uLaw = ^uLaw
	sign := int16(1)
	if uLaw&0x80 != 0 {
		sign = -1
	}
	exponent := int((uLaw >> 4) & 0x07)
	mantissa := int(uLaw&0x0F) + 16
	return sign * int16(mantissa) << uint(exponent+3)
}

func EncodeToULaw(sample int16) byte {
	sign := byte(0)
	if sample < 0 {
		sign = 0x80
		sample = -sample
	}

	sample = sample + 33
	if sample > 0x7FFF {
		sample = 0x7FFF
	}

	exponent := byte(7)
	for (sample & 0x4000) == 0 {
		exponent--
		sample <<= 1
	}

	mantissa := byte(sample >> 6)
	return ^byte(sign | (exponent << 4) | mantissa)
}
