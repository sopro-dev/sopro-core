package encoder

// func MulawToPcm(input []byte, depthBits int, signed bool) []byte {
// 	output := make([]byte, len(input))
// 	for i, b := range input {
// 		// output[i] = DecodeMulawByte(b, depthBits, signed)
// 		fmt.Printf("<%d:%d>", b, output[i])
// 	}
// 	return output
// }

// var mulawDecodeTables [MULAW_ITEMS_8BIT][MULAW_ITEMS_8BIT]byte

// const (
// 	ORIGIN_ITEMS_8BIT  = 256
// 	ORIGIN_ITEMS_16BIT = 65536
// 	// MULAW_ITEMS_24BIT = 16777216   // Not supported
// 	// MULAW_ITEMS_32BIT = 4294967296 // Not supported
// )

// func init() {
// 	for i := 0; i < 256; i++ {
// 		mulawDecodeTable[i] = make([]byte, 8)
// 		for depthBits := 1; depthBits <= 8; depthBits++ {
// 			mu := 255.0
// 			output := float64(i)
// 			output = (output / mu) * ((math.Pow(1+mu, math.Abs(output)) - 1) / mu)
// 			maxValue := math.Pow(2, float64(depthBits-1))
// 			output *= maxValue
// 			if output > maxValue {
// 				output = maxValue
// 			}
// 			if output < 0 {
// 				output = 0
// 			}
// 			mulawDecodeTable[i][depthBits-1] = byte(output)
// 		}
// 	}
// }

// func DecodeMulawByte(input byte, depthBits int, signed bool) (output byte) {

// 	if signed {
// 		return mulawDecodeTable[input][depthBits-1] - 128
// 	}
// 	return mulawDecodeTable[input][depthBits-1]
// }
