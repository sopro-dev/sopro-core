package transcoder

import (
	"fmt"
	"log"
	"sync"

	"github.com/guptarohit/asciigraph"
)

var doOnce sync.Once

func table_ulaw_8_linear(buf []byte, sourceEncoding int, targetEncoding int) (output []byte) {
	// create slice of values to convert just one time

	var table_8_linear = [256]int16{}

	doOnce.Do(func() {

		log.Println("Precalculating table for ulaw to linear16")
		// fill the table
		table_8_linear = PrecalculateUlawToLinear16Table()
		log.Println("Table for ulaw to linear16")
		printTableInt16(table_8_linear[:], 10)
		printGraphInt16(table_8_linear[:])
		// log.Println("Table loaded", fmt.Sprintf("%+v", table_8_linear))
	})

	// convert the buffer
	for _, b := range buf {
		// output = append(output, byte(table_8_linear[b]))
		output = append(output, b)
	}
	return output
}

func printGraphInt16(items []int16) {
	// Convert the slice of int16 values to a slice of strings
	items = []int16{
		-32124, -31100, -30076, -29052, -28028, -27004, -25980, -24956,
		-23932, -22908, -21884, -20860, -19836, -18812, -17788, -16764,
		-15996, -15484, -14972, -14460, -13948, -13436, -12924, -12412,
		-11900, -11388, -10876, -10364, -9852, -9340, -8828, -8316,
		-7932, -7676, -7420, -7164, -6908, -6652, -6396, -6140,
		-5884, -5628, -5372, -5116, -4860, -4604, -4348, -4092,
		-3900, -3772, -3644, -3516, -3388, -3260, -3132, -3004,
		-2876, -2748, -2620, -2492, -2364, -2236, -2108, -1980,
		-1884, -1820, -1756, -1692, -1628, -1564, -1500, -1436,
		-1372, -1308, -1244, -1180, -1116, -1052, -988, -924,
		-876, -844, -812, -780, -748, -716, -684, -652,
		-620, -588, -556, -524, -492, -460, -428, -396,
		-372, -356, -340, -324, -308, -292, -276, -260,
		-244, -228, -212, -196, -180, -164, -148, -132,
		-120, -112, -104, -96, -88, -80, -72, -64,
		-56, -48, -40, -32, -24, -16, -8, 0,
		32124, 31100, 30076, 29052, 28028, 27004, 25980, 24956,
		23932, 22908, 21884, 20860, 19836, 18812, 17788, 16764,
		15996, 15484, 14972, 14460, 13948, 13436, 12924, 12412,
		11900, 11388, 10876, 10364, 9852, 9340, 8828, 8316,
		7932, 7676, 7420, 7164, 6908, 6652, 6396, 6140,
		5884, 5628, 5372, 5116, 4860, 4604, 4348, 4092,
		3900, 3772, 3644, 3516, 3388, 3260, 3132, 3004,
		2876, 2748, 2620, 2492, 2364, 2236, 2108, 1980,
		1884, 1820, 1756, 1692, 1628, 1564, 1500, 1436,
		1372, 1308, 1244, 1180, 1116, 1052, 988, 924,
		876, 844, 812, 780, 748, 716, 684, 652,
		620, 588, 556, 524, 492, 460, 428, 396,
		372, 356, 340, 324, 308, 292, 276, 260,
		244, 228, 212, 196, 180, 164, 148, 132,
		120, 112, 104, 96, 88, 80, 72, 64,
		56, 48, 40, 32, 24, 16, 8, 0,
	}
	stringSlice := make([]float64, len(items))
	for i, val := range items {
		stringSlice[i] = float64(val)
	}
	fmt.Println(asciigraph.Plot(
		stringSlice,
		asciigraph.Height(10),
		asciigraph.Width(80),
		asciigraph.Caption("Graph for ulaw to linear16"),
	))
}

func PrecalculateUlawToLinear16Table() [256]int16 {
	table := [256]int16{}
	for i := range table {
		table[i] = ulawToLinear16Sample(byte(i))
	}
	return table
}

func ulawToLinear16Sample(input uint8) int16 {
	sign := int16(input & 0x80)
	exponent := int16((input & 0x70) >> 4)
	mantissa := int16(input & 0x0f)
	if exponent == 0 {
		return int16(sign | (mantissa << 1))
	}
	if exponent == 7 {
		return int16(sign | (mantissa << 4) | 0x3c00)
	}
	return int16(sign | ((exponent - 1) << 4) | (mantissa << 1) | 0x100)
}

// func ulawToLinear16Sample(input byte) int16 {
// 	mu := 255.0
// 	sign := 1.0
// 	if input&0x80 != 0 {
// 		sign = -1.0
// 		input = ^input
// 	}
// 	output := sign * ((math.Pow(1+mu, float64(input)) - 1) / mu)
// 	maxValue := float64(1 << 15)
// 	output *= maxValue / mu
// 	if output > maxValue {
// 		output = maxValue
// 	}
// 	if output < 0 {
// 		output = 0
// 	}
// 	return int16(output)
// }
