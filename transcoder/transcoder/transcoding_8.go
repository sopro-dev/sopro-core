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

	var table_8_linear = [256]int8{}

	doOnce.Do(func() {

		log.Println("Precalculating table for ulaw to linear16")
		// fill the table
		// table_8_linear = PrecalculateUlawToLinear16Table()
		table_8_linear = table_ulaw_8_linear_8
		log.Println("Table for ulaw to linear16")
		printTableInt8(table_8_linear[:], 10)
		log.Println("Graph for ulaw to linear16")
		printGraphInt8(table_8_linear[:])
		// log.Println("Table loaded", fmt.Sprintf("%+v", table_8_linear))
	})

	// convert the buffer
	for _, b := range buf {
		output = append(output, byte(table_ulaw_8_linear_8[b]))
		// output = append(output, b)
	}
	return output
}

func printGraphInt8(items []int8) {
	// Convert the slice of int16 values to a slice of strings

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
	minY, maxY := LimitsSliceFloat64(stringSlice)
	fmt.Println("MinY:   ", minY)
	fmt.Println("MaxY:   ", maxY)
	fmt.Println("DeltaY: ", (maxY - minY))
	fmt.Println("MinX:   ", 0)
	fmt.Println("MaxX:   ", len(stringSlice))
	fmt.Println("DeltaX: ", (len(stringSlice) - 0))
	fmt.Println()

}

// func PrecalculateUlawToLinear16Table() [256]int16 {
// 	table := [256]int16{}
// 	for i := range table {
// 		table[i] = ulawToLinear16Sample(byte(i))
// 	}
// 	return table
// }

// func ulawToLinear16Sample(input uint8) int16 {
// 	sign := int16(input & 0x80)
// 	exponent := int16((input & 0x70) >> 4)
// 	mantissa := int16(input & 0x0f)
// 	if exponent == 0 {
// 		return int16(sign | (mantissa << 1))
// 	}
// 	if exponent == 7 {
// 		return int16(sign | (mantissa << 4) | 0x3c00)
// 	}
// 	return int16(sign | ((exponent - 1) << 4) | (mantissa << 1) | 0x100)
// }

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
