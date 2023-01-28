package sopro

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math"
	"os"

	"github.com/guptarohit/asciigraph"
	"golang.org/x/term"
)

// Default terminal size
var (
	WIDTH_TERMINAL  = 80
	HEIGHT_TERMINAL = 30
)

// GraphIn graphs the input file to the terminal
func GraphIn(in *In) error {
	log.Println("[WARNING] Reading the whole file into memory. This may take a while...")
	// check if in is *bytes.Buffer
	if _, ok := in.Data.(*bytes.Buffer); ok {
		log.Println("Input file is a bytes.Buffer")
		return nil
	}

	// make an independent copy of the file
	f, err := os.Open(in.Data.(*os.File).Name())
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	values := make([]float64, len(data))
	for i, val := range data {
		values[i] = float64(val)
	}

	WIDTH_TERMINAL, HEIGHT_TERMINAL, err = term.GetSize(0)
	if err != nil {
		WIDTH_TERMINAL = 80
		HEIGHT_TERMINAL = 24
		log.Println("Error getting terminal size, using default values (80x24) instead")
	}

	fmt.Println(asciigraph.Plot(
		values,
		asciigraph.Height(HEIGHT_TERMINAL/3),
		asciigraph.Width(WIDTH_TERMINAL-10),
		asciigraph.Caption("Graph for input ulaw file"),
		asciigraph.SeriesColors(
			asciigraph.Red,
		),
	))

	return nil
}

// GraphOut graphs the output file to the terminal
func GraphOut(in *In, out *Out) error {
	log.Println("[WARNING] Reading the whole file into memory. This may take a while...")
	// check if in is *bytes.Buffer
	if _, ok := in.Data.(*bytes.Buffer); ok {
		log.Println("Input file is a bytes.Buffer")
		return nil
	}
	f, err := os.Open(in.Data.(*os.File).Name())
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}
	input := make([]float64, len(data)*2)
	for i, val := range data {
		input[i*2] = float64(val)
		input[i*2+1] = float64(val)
	}

	fOut, err := os.Open(out.Data.(*os.File).Name())
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer fOut.Close()
	outData, err := io.ReadAll(fOut)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}
	output := make([]float64, len(outData))
	for i, val := range outData {
		if i%2 == 0 {
			continue
		}
		output[i] = float64(val)
	}

	maxData := make([]float64, len(outData))
	for i := range outData {
		maxData[i] = float64(math.MaxInt8)
	}
	linesMiddle := maxData[44:]
	lineInput := input
	lineOutput := output[44:]

	lineInput[0] = 0
	lineInput[len(lineInput)-1] = math.Round(float64(math.MaxUint8) / 2)
	lineOutput[0] = 0
	lineOutput[len(lineOutput)-1] = math.Round(float64(math.MaxUint8) / 2)

	WIDTH_TERMINAL, HEIGHT_TERMINAL, err = term.GetSize(0)
	if err != nil {
		WIDTH_TERMINAL = 80
		HEIGHT_TERMINAL = 24
		log.Println("Error getting terminal size, using default values (80x24) instead")
	}
	log.Println("Sample of the input file (ulaw) (first 100 samples of n)")
	fmt.Println(asciigraph.PlotMany(
		[][]float64{linesMiddle, lineInput},
		asciigraph.Height(HEIGHT_TERMINAL/3),
		asciigraph.Width(WIDTH_TERMINAL-10),
		asciigraph.Caption("Graph for input ulaw file"),
		asciigraph.SeriesColors(
			asciigraph.Blue,
			asciigraph.Red,
			asciigraph.Green,
		),
	))

	log.Println("Length Zeros", len(linesMiddle), "Length Input", len(lineInput), "Length Output[44:]", len(lineOutput))
	fmt.Println(asciigraph.PlotMany(
		[][]float64{linesMiddle, lineInput, lineOutput},
		asciigraph.Height(HEIGHT_TERMINAL/3),
		asciigraph.Width(WIDTH_TERMINAL-10),
		asciigraph.Caption("Graph for output wav data file"),
		asciigraph.SeriesColors(
			asciigraph.Blue,
			asciigraph.Red,
			asciigraph.Green,
		),
	))
	fmt.Println("*First and last byte are not representative")

	return nil
}
