package transcoder

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"
	"sync"

	"github.com/guptarohit/asciigraph"
	"github.com/pablodz/transcoder/transcoder/audioconfig"
	"github.com/pablodz/transcoder/transcoder/cpuarch"
	"github.com/zaf/g711"
	"golang.org/x/term"
)

var WIDTH_TERMINAL = 80
var HEIGHT_TERMINAL = 10

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
func Mulaw2Wav(in *AudioFileIn, out *AudioFileOut, transcoder *TranscoderOneToOne) (err error) {

	// read all the file
	if transcoder.Verbose {
		transcoder.Println("[WARNING] Reading the whole file into memory. This may take a while...")
		f := in.Data.(*os.File)
		// make an independent copy of the file
		f2, err := os.Open(f.Name())
		if err != nil {
			log.Fatal(err)
		}
		bytes, err := io.ReadAll(f2)
		if err != nil {
			log.Fatal(err)
		}

		stringSlice := make([]float64, len(bytes))
		for i, val := range bytes {
			stringSlice[i] = float64(val)
		}
		fmt.Println(asciigraph.Plot(
			stringSlice,
			asciigraph.Height(HEIGHT_TERMINAL/3),
			asciigraph.Width(WIDTH_TERMINAL-10),
			asciigraph.Caption("Graph for input ulaw file"),
		))
		defer f2.Close()
	}

	// Get the WAV file configuration
	channels := out.Config.(audioconfig.WavConfig).Channels
	sampleRate := out.Config.(audioconfig.WavConfig).SampleRate
	bitsPerSample := out.Config.(audioconfig.WavConfig).BitDepth

	transcoder.Println("Channels:", channels)
	transcoder.Println("Sample rate:", sampleRate)
	transcoder.Println("Bits per sample:", bitsPerSample)

	// Create a buffered reader and writer
	in.Reader = bufio.NewReaderSize(in.Data, 64)
	out.Writer = bufio.NewWriter(out.Data)
	out.Length = 0
	transcoder.SourceConfigs.Encoding = in.Config.(audioconfig.MulawConfig).Encoding
	transcoder.TargetConfigs.Encoding = out.Config.(audioconfig.WavConfig).Encoding
	transcoder.BitDepth = bitsPerSample

	headersWav := []byte{
		'R', 'I', 'F', 'F', // Chunk ID
		0, 0, 0, 0, // Chunk size
		'W', 'A', 'V', 'E', // Format
		'f', 'm', 't', ' ', // Sub-chunk 1 ID
		16, 0, 0, 0, // Sub-chunk 1 size
		1, 0, // Audio format (PCM)
		byte(channels), 0, // Number of channels
		byte(sampleRate & 0xff),                                        // sample rate (low)
		byte(sampleRate >> 8 & 0xff),                                   // sample rate (mid)
		byte(sampleRate >> 16 & 0xff),                                  // sample rate (high)
		byte(sampleRate >> 24 & 0xff),                                  // sample rate (high)
		byte(sampleRate * channels * (bitsPerSample / 8) & 0xff),       // byte rate (low)
		byte(sampleRate * channels * (bitsPerSample / 8) >> 8 & 0xff),  // byte rate (mid)
		byte(sampleRate * channels * (bitsPerSample / 8) >> 16 & 0xff), // byte rate (high)
		byte(sampleRate * channels * (bitsPerSample / 8) >> 24 & 0xff), // byte rate (high)
		byte(channels * (bitsPerSample / 8)), 0,                        // block align
		byte(bitsPerSample), 0, // bits per sample
		'd', 'a', 't', 'a',
		0, 0, 0, 0,
	}
	// Riff header
	out.Writer.Write(headersWav)
	for i := 0; i < 44; i += 4 {
		transcoder.Println("Header:", fmt.Sprintf("% 02x", headersWav[i:i+4]))
	}

	out.Length += len(headersWav)
	transcoder.Println("Length:", out.Length)

	// Copy the data from the input file to the output file in chunks
	out.Length, err = TranscodeBytes(in, out, transcoder)
	if err != nil {
		return fmt.Errorf("error converting bytes: %v", err)
	}
	if err := out.Writer.Flush(); err != nil {
		return fmt.Errorf("error flushing output file: %v", err)
	}
	transcoder.Println("Wrote", out.Length, "bytes to output file")

	if transcoder.Verbose {
		transcoder.Println("[WARNING] Reading the whole file into memory. This may take a while...")
		f := in.Data.(*os.File)
		// make an independent copy of the file
		f2, err := os.Open(f.Name())
		if err != nil {
			log.Fatal(err)
		}
		bytes, err := io.ReadAll(f2)
		if err != nil {
			log.Fatal(err)
		}

		stringSlice1 := make([]float64, len(bytes)*2)
		for i, val := range bytes {
			stringSlice1[i*2] = float64(val)
			stringSlice1[i*2+1] = float64(val)
		}
		defer f2.Close()
		// graph of the output file
		fOut := out.Data.(*os.File)
		// make an independent copy of the file
		f3, err := os.Open(fOut.Name())
		if err != nil {
			log.Fatal(err)
		}
		bytes, err = io.ReadAll(f3)
		if err != nil {
			log.Fatal(err)
		}

		stringSlice2 := make([]float64, len(bytes))
		for i, val := range bytes {
			if i%2 == 0 {
				continue
			}
			stringSlice2[i] = float64(val)
		}

		stringSlice3 := make([]float64, len(bytes))
		for i := range bytes {
			stringSlice3[i] = float64(math.MaxInt8)
		}
		l1 := stringSlice3[44:] // 0
		l2 := stringSlice1      // ulaw
		l3 := stringSlice2[44:] // wav

		// define limits forced
		// l1[0] = 0
		// l1[len(l1)-1] = float64(math.MaxUint8)
		l2[0] = 0
		l2[len(l2)-1] = math.Round(float64(math.MaxUint8) / 2)
		l3[0] = 0
		l3[len(l3)-1] = math.Round(float64(math.MaxUint8) / 2)

		log.Println("Sample of the input file (ulaw) (first 100 samples of n)")

		// resample to double the size
		fmt.Println(asciigraph.PlotMany(
			[][]float64{l1, l2},
			asciigraph.Height(HEIGHT_TERMINAL/3),
			asciigraph.Width(WIDTH_TERMINAL-10),
			asciigraph.Caption("Graph for input ulaw file"),
			asciigraph.SeriesColors(
				asciigraph.Blue,
				asciigraph.Red,
				asciigraph.Green,
			),
		))

		log.Println("len l1", len(l1), "len l2", len(l2), "len l3", len(l3))
		fmt.Println(asciigraph.PlotMany(
			[][]float64{l1, l2, l3},
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
		defer f2.Close()
	}

	// Update the file size and data size fields
	fileSize := []byte{
		byte((out.Length - 8) & 0xff),
		byte((out.Length - 8) >> 8 & 0xff),
		byte((out.Length - 8) >> 16 & 0xff),
		byte((out.Length - 8) >> 24 & 0xff),
	}
	newF := out.Data.(*os.File)
	r, err := newF.Seek(4, io.SeekStart)
	if err != nil {
		return fmt.Errorf("[1]error seeking file: %v", err)
	}
	transcoder.Println("Seeked to:", r)

	n, err := newF.Write(fileSize)
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
	r, err = newF.Seek(40, io.SeekStart)
	if err != nil {
		return fmt.Errorf("[2]error seeking file: %v", err)
	}
	transcoder.Println("Seeked to:", r)
	n, err = newF.Write(dataSize)
	if err != nil {
		return fmt.Errorf("error writing data size: %v", err)
	}
	transcoder.Println("Data size:", fmt.Sprintf("% 02x", dataSize), "bytes written:", n)

	return nil

}

var doOnceTD sync.Once

func TranscodeBytes(in *AudioFileIn, out *AudioFileOut, transcoder *TranscoderOneToOne) (int, error) {

	equalEncoding := (transcoder.SourceConfigs.Encoding == transcoder.TargetConfigs.Encoding)
	if transcoder.SourceConfigs.Endianness == cpuarch.NOT_FILLED && transcoder.TargetConfigs.Endianness == cpuarch.NOT_FILLED {
		transcoder.SourceConfigs.Endianness = cpuarch.LITTLE_ENDIAN
		transcoder.TargetConfigs.Endianness = cpuarch.LITTLE_ENDIAN
	}

	if equalEncoding {
		log.Println("Same encodings")
		for {
			break
			// Read a chunk from the input file
			// n, err := in.Reader.Read(buf)
			// if err != nil && err != io.EOF {
			// 	return -1, fmt.Errorf("error reading input file: %v", err)
			// }
			// if n == 0 {
			// 	break
			// }
			// // out.Length += len(buf)

			// // // Write the chunk to the output file
			// // _, err := out.Writer.Write(buf[:n])
			// // if err != nil {
			// // 	return -1, fmt.Errorf("error writing output file: %v", err)
			// // }
		}
	} else {
		log.Println("Different encodings")
		m := 32 // 32 no more buffer
		buf := make([]byte, 2*m)
		for {
			n, err := in.Reader.Read(buf)
			if err != nil && err != io.EOF {
				return -1, fmt.Errorf("error reading input file: %v", err)
			}
			if n == 0 {
				break
			}
			buf2 := g711.DecodeUlaw(buf)
			out.Length += len(buf2)

			if _, err := out.Writer.Write(buf2); err != nil {
				return -1, fmt.Errorf("error writing output file: %v", err)
			}
			out.Writer.Flush()
			doOnceTD.Do(func() {
				transcoder.Println(strings.Repeat("-", 80))
				transcoder.Println("Transcoding bytes")
				onlyNFirst := 4
				transcoder.Println(
					"|OLD|", fmt.Sprintf("% 02x", buf[:onlyNFirst]),
					"|NEW|", fmt.Sprintf("% 02x", buf2[:onlyNFirst]),
					"|LEN|", onlyNFirst,
				)
				transcoder.Println(
					"|OLD|", fmt.Sprintf("%03d", buf[:onlyNFirst]),
					"|NEW|", fmt.Sprintf("%03d", buf2[:onlyNFirst]),
					"|LEN|", onlyNFirst,
				)
				transcoder.Println(strings.Repeat("-", 80))
			})

		}

	}

	return out.Length, nil
}
