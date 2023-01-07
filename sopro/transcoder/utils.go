package transcoder

import (
	"fmt"
	"os"
	"strconv"

	"github.com/guptarohit/asciigraph"
	"github.com/olekukonko/tablewriter"
)

func (t *Transcoder) Println(items ...any) {
	if t.Verbose {
		fmt.Println(items...)
	}
}

func printGraphInt16(items []int16) {
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

func printTableInt16(items []int16, numPerRow int) {
	// Convert the slice of int16 values to a slice of strings
	stringSlice := make([]string, len(items))
	for i, val := range items {
		stringSlice[i] = strconv.FormatInt(int64(val), 10)
	}
	printTable(stringSlice, numPerRow)
}

func printTable(items []string, numPerRow int) {
	// Create a new table
	table := tablewriter.NewWriter(os.Stdout)
	// table.SetColWidth(80)
	table.SetColumnSeparator("|")
	table.SetRowSeparator("-")
	table.SetCenterSeparator(".")

	cols := []string{}
	for i := 0; i < numPerRow; i++ {
		cols = append(cols, fmt.Sprintf("%v", i))
	}
	table.SetHeader(cols)

	nRows := len(items) / numPerRow

	data := make([][]string, nRows)
	for i := 0; i < nRows; i++ {
		data[i] = make([]string, numPerRow)
	}

	for i := 0; i < nRows; i++ {
		for j := 0; j < numPerRow; j++ {
			data[i][j] = items[i*numPerRow+j]
		}
	}

	for _, v := range data {
		table.Append(v)
	}

	table.Render()
}

func LimitsSliceFloat64(items []float64) (float64, float64) {
	min := items[0]
	max := items[0]
	for _, val := range items {
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}
	return min, max
}
