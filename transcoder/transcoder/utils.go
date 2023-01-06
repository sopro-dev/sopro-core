package transcoder

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

func (t *TranscoderOneToOne) Println(items ...any) {
	if t.Verbose {
		log.Println(items...)
	}
}

func printTableInt8(items []int8, numPerRow int) {
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
