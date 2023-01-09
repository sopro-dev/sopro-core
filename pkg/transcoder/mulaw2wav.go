package transcoder

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"

	"github.com/guptarohit/asciigraph"
	"github.com/pablodz/sopro/pkg/audioconfig"
	"github.com/pablodz/sopro/pkg/cpuarch"
	"github.com/pablodz/sopro/pkg/encoding"
	"golang.org/x/term"
)

var WIDTH_TERMINAL = 80
var HEIGHT_TERMINAL = 30

func init() {

	err := error(nil)
	WIDTH_TERMINAL, HEIGHT_TERMINAL, err = term.GetSize(0)
	if err != nil {
		log.Fatal(err)
	}
}

// TODO: split functions for different sizes of files
// Transcode an ulaw file to a wav file (large files supported)
// https://raw.githubusercontent.com/corkami/pics/master/binary/WAV.png
// http://www.topherlee.com/software/pcm-tut-wavformat.html
func mulaw2Wav(in *AudioFileIn, out *AudioFileOut, transcoder *Transcoder) (err error) {

	// read all the file
	if transcoder.Verbose {
		graphIn(in)
	}

	// Get the WAV file configuration
	channels := out.Config.(audioconfig.WavConfig).Channels
	sampleRate := out.Config.(audioconfig.WavConfig).SampleRate
	bitsPerSample := out.Config.(audioconfig.WavConfig).BitDepth
	transcoder.SourceConfigs.Encoding = in.Config.(audioconfig.MulawConfig).Encoding
	transcoder.TargetConfigs.Encoding = out.Config.(audioconfig.WavConfig).Encoding
	transcoder.BitDepth = bitsPerSample

	if transcoder.SourceConfigs.Endianness == cpuarch.NOT_FILLED && transcoder.TargetConfigs.Endianness == cpuarch.NOT_FILLED {
		transcoder.SourceConfigs.Endianness = cpuarch.LITTLE_ENDIAN // replace with cpuarch.GetEndianess()
		transcoder.TargetConfigs.Endianness = cpuarch.LITTLE_ENDIAN
	}

	transcoder.Println(
		"\n[Format]                      ", in.Format, "=>", out.Format,
		"\n[Encoding]                    ", encoding.ENCODINGS[in.Config.(audioconfig.MulawConfig).Encoding], "=>", encoding.ENCODINGS[out.Config.(audioconfig.WavConfig).Encoding],
		"\n[Channels]                    ", in.Config.(audioconfig.MulawConfig).Channels, "=>", channels,
		"\n[SampleRate]                  ", in.Config.(audioconfig.MulawConfig).SampleRate, "=>", sampleRate, "kHz",
		"\n[BitDepth]                    ", in.Config.(audioconfig.MulawConfig).BitDepth, "=>", bitsPerSample, "bytes",
		"\n[Transcoder][Source][Encoding]", encoding.ENCODINGS[transcoder.SourceConfigs.Encoding],
		"\n[Transcoder][Target][Encoding]", encoding.ENCODINGS[transcoder.TargetConfigs.Encoding],
		"\n[Transcoder][BitDepth]        ", transcoder.BitDepth,
		"\n[Transcoder][Endianness]      ", cpuarch.ENDIANESSES[cpuarch.GetEndianess()],
	)

	// Create a buffered reader and writer
	in.Reader = bufio.NewReader(in.Data)
	out.Writer = bufio.NewWriter(out.Data)
	out.Length = 0

	headersWav := []byte{
		'R', 'I', 'F', 'F', // Chunk ID
		0, 0, 0, 0, // Chunk size
		'W', 'A', 'V', 'E', // Format
		'f', 'm', 't', ' ', // Sub-chunk 1 ID
		16, 0, 0, 0, // Sub-chunk 1 size
		1, 0, // Audio format (PCM)
		byte(channels), 0, // Number of channels
		byte(sampleRate & 0xFF),                                        // sample rate (low)
		byte(sampleRate >> 8 & 0xFF),                                   // sample rate (mid)
		byte(sampleRate >> 16 & 0xFF),                                  // sample rate (high)
		byte(sampleRate >> 24 & 0xFF),                                  // sample rate (high)
		byte(sampleRate * channels * (bitsPerSample / 8) & 0xFF),       // byte rate (low)
		byte(sampleRate * channels * (bitsPerSample / 8) >> 8 & 0xFF),  // byte rate (mid)
		byte(sampleRate * channels * (bitsPerSample / 8) >> 16 & 0xFF), // byte rate (high)
		byte(sampleRate * channels * (bitsPerSample / 8) >> 24 & 0xFF), // byte rate (high)
		byte(channels * (bitsPerSample / 8)), 0,                        // block align
		byte(bitsPerSample), 0, // bits per sample
		'd', 'a', 't', 'a',
		0, 0, 0, 0,
	}
	out.Writer.Write(headersWav)
	out.Length += len(headersWav)

	if transcoder.Verbose {
		audioconfig.PrintWavHeaders(headersWav)
	}

	// Copy the data from the input file to the output file in chunks
	if err = TranscodeBytes(in, out, transcoder); err != nil {
		return fmt.Errorf("error converting bytes: %v", err)
	}

	// Flush the output file
	if err := out.Writer.Flush(); err != nil {
		return fmt.Errorf("error flushing output file: %v", err)
	}
	transcoder.Println("Wrote", out.Length, "bytes to output file")

	// Update the file size and data size fields
	fileFixer := out.Data.(*os.File)
	r, err := fileFixer.Seek(4, io.SeekStart)
	if err != nil {
		return fmt.Errorf("error seeking file: %v", err)
	}
	transcoder.Println("Seeked to:", r)
	fileSize := []byte{
		byte((out.Length - 8) & 0xff),
		byte((out.Length - 8) >> 8 & 0xff),
		byte((out.Length - 8) >> 16 & 0xff),
		byte((out.Length - 8) >> 24 & 0xff),
	}
	n, err := fileFixer.Write(fileSize)
	if err != nil {
		return fmt.Errorf("error writing file size: %v", err)
	}
	transcoder.Println("File size:", fmt.Sprintf("% 02x", fileSize), "bytes written:", n)
	dataSize := []byte{
		byte((out.Length - 44) & 0xff),
		byte((out.Length - 44) >> 8 & 0xff),
		byte((out.Length - 44) >> 16 & 0xff),
		byte((out.Length - 44) >> 24 & 0xff),
	}
	r, err = fileFixer.Seek(40, io.SeekStart)
	if err != nil {
		return fmt.Errorf("[2]error seeking file: %v", err)
	}
	transcoder.Println("Seeked to:", r)
	n, err = fileFixer.Write(dataSize)
	if err != nil {
		return fmt.Errorf("error writing data size: %v", err)
	}
	transcoder.Println("Data size:", fmt.Sprintf("% 02x", dataSize), "bytes written:", n)

	if transcoder.Verbose {
		graphOut(in, out)
	}

	return nil

}

func graphIn(in *AudioFileIn) {
	log.Println("[WARNING] Reading the whole file into memory. This may take a while...")
	// make an independent copy of the file
	file := in.Data.(*os.File)
	f, err := os.Open(file.Name())
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	values := make([]float64, len(data))
	for i, val := range data {
		values[i] = float64(val)
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
}

func graphOut(in *AudioFileIn, out *AudioFileOut) {
	log.Println("[WARNING] Reading the whole file into memory. This may take a while...")
	file := in.Data.(*os.File)
	f, err := os.Open(file.Name())
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	input := make([]float64, len(data)*2)
	for i, val := range data {
		input[i*2] = float64(val)
		input[i*2+1] = float64(val)
	}

	outFile := out.Data.(*os.File)
	fOut, err := os.Open(outFile.Name())
	if err != nil {
		log.Fatal(err)
	}
	defer fOut.Close()
	outData, err := io.ReadAll(fOut)
	if err != nil {
		log.Fatal(err)
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
}
