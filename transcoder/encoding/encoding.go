package encoding

func Mulaw2Linear(in []byte) []byte {
	out := make([]byte, len(in))
	for i, b := range in {
		out[i] = mulawDecodeTable[b]
	}
	return out
}

var mulawDecodeTable = GenerateMulawDecodeTable()

func GenerateMulawDecodeTable() [256]byte {
	var table [256]byte
	for i := 0; i < 256; i++ {
		b := byte(i)
		sign := b & 0x80
		segment := (b >> 4) & 0x07
		factor := b & 0x0f
		if sign == 0 {
			table[i] = (byte(255-factor) << 4) + factor
		} else {
			table[i] = (byte(factor+1) << 4) + (15 - factor)
		}
		if segment > 1 {
			table[i] = table[i] << uint(segment-1)
		} else if segment == 1 {
			table[i] = byte(int(table[i]) + 128)
		}
	}
	return table
}
