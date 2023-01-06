package transcoder

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/guptarohit/asciigraph"
	"github.com/pablodz/transcoder/transcoder/audioconfig"
	"github.com/pablodz/transcoder/transcoder/cpuarch"
)

// TODO: split functions for different sizes of files
// Transcode an ulaw file to a wav file (large files supported)
// https://raw.githubusercontent.com/corkami/pics/master/binary/WAV.png
// http://www.topherlee.com/software/pcm-tut-wavformat.html
func Mulaw2Wav(in *AudioFileIn, out *AudioFileOut, transcoder *TranscoderOneToOne) (err error) {

	// read all the file
	if transcoder.Verbose {
		f := in.Data.(*os.File)
		bytes, err := io.ReadAll(f)
		if err != nil {
			panic(err)
		}

		stringSlice := make([]float64, len(bytes))
		for i, val := range bytes {
			stringSlice[i] = float64(val)
		}
		fmt.Println(asciigraph.Plot(
			stringSlice,
			asciigraph.Height(10),
			asciigraph.Width(80),
			asciigraph.Caption("Graph for input ulaw file"),
		))
		defer f.Close()
	}

	// Get the WAV file configuration
	channels := out.Config.(audioconfig.WavConfig).Channels
	sampleRate := out.Config.(audioconfig.WavConfig).SampleRate
	bitsPerSample := out.Config.(audioconfig.WavConfig).BitDepth

	// Create a buffered reader and writer
	in.Reader = bufio.NewReader(in.Data)
	out.Writer = bufio.NewWriter(out.Data)
	out.Length = 0
	transcoder.SourceConfigs.Encoding = in.Config.(audioconfig.MulawConfig).Encoding
	transcoder.TargetConfigs.Encoding = out.Config.(audioconfig.WavConfig).Encoding
	transcoder.BitDepth = bitsPerSample

	// Riff header
	out.Writer.Write([]byte{'R', 'I', 'F', 'F'})
	out.Writer.Write([]byte{0, 0, 0, 0}) // placeholder for file size
	out.Writer.Write([]byte{'W', 'A', 'V', 'E'})
	// Wave header
	out.Writer.Write([]byte{'f', 'm', 't', ' '})
	out.Writer.Write([]byte{16, 0, 0, 0})       // size of fmt chunk
	out.Writer.Write([]byte{1, 0})              // audio format (PCM)
	out.Writer.Write([]byte{byte(channels), 0}) // number of channels

	// Set the sample rate and byte rate
	sampleRateBytes := []byte{
		byte(sampleRate & 0xff),
		byte(sampleRate >> 8 & 0xff),
		byte(sampleRate >> 16 & 0xff),
		byte(sampleRate >> 24 & 0xff),
	}
	transcoder.Println("Sample rate:", fmt.Sprintf("% 02x", sampleRateBytes))

	out.Writer.Write(sampleRateBytes)
	byteRate := []byte{
		byte(sampleRate * channels * (bitsPerSample / 8) & 0xff),
		byte(sampleRate * channels * (bitsPerSample / 8) >> 8 & 0xff),
		byte(sampleRate * channels * (bitsPerSample / 8) >> 16 & 0xff),
		byte(sampleRate * channels * (bitsPerSample / 8) >> 24 & 0xff),
	}
	transcoder.Println("Byte rate:", fmt.Sprintf("% 02x", byteRate))

	out.Writer.Write(byteRate)
	out.Writer.Write([]byte{byte(channels * (bitsPerSample / 8)), 0}) // block align
	out.Writer.Write([]byte{byte(bitsPerSample), 0})                  // bits per sample

	// Write the data chunk header
	out.Writer.Write([]byte{'d', 'a', 't', 'a'})
	out.Writer.Write([]byte{0, 0, 0, 0}) // placeholder for data size

	out.Length += 44
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

	// Update the file size and data size fields
	fileSize := []byte{
		byte((out.Length - 8) & 0xff),
		byte((out.Length - 8) >> 8 & 0xff),
		byte((out.Length - 8) >> 16 & 0xff),
		byte((out.Length - 8) >> 24 & 0xff),
	}
	r, err := out.File.Seek(4, io.SeekStart)
	if err != nil {
		return fmt.Errorf("[1]error seeking file: %v", err)
	}
	transcoder.Println("Seeked to:", r)

	n, err := out.File.Write(fileSize)
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
	r, err = out.File.Seek(40, io.SeekStart)
	if err != nil {
		return fmt.Errorf("[2]error seeking file: %v", err)
	}
	transcoder.Println("Seeked to:", r)
	n, err = out.File.Write(dataSize)
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

	buf := make([]byte, 10) // read and write in chunks of 1024 bytes
	n := 0
	err := error(nil)

	tBestFunc := bitTranscoder(buf, transcoder)

	if equalEncoding {
		log.Println("Same encodings")
		for {
			// Read a chunk from the input file
			n, err = in.Reader.Read(buf)
			if err != nil && err != io.EOF {
				return -1, fmt.Errorf("error reading input file: %v", err)
			}
			if n == 0 {
				break
			}
			out.Length += len(buf)

			// Write the chunk to the output file
			if _, err := out.Writer.Write(buf[:n]); err != nil {
				return -1, fmt.Errorf("error writing output file: %v", err)
			}
		}
	} else {
		log.Println("Different encodings")
		for {
			// Read a chunk from the input file
			n, err = in.Reader.Read(buf)
			if err != nil && err != io.EOF {
				return -1, fmt.Errorf("error reading input file: %v", err)
			}
			if n == 0 {
				break
			}

			// Transcode the chunk
			buf2, err := tBestFunc()
			if err != nil {
				return -1, fmt.Errorf("error getting best transcoder: %v", err)
			}
			out.Length += len(buf2)

			doOnceTD.Do(func() {
				transcoder.Println(strings.Repeat("-", 80))
				transcoder.Println("OLD:", fmt.Sprintf("% 02x", buf[:n]), "NEW:", fmt.Sprintf("% 02x", buf2[:n]), "LEN:", n)
				transcoder.Println(strings.Repeat("-", 80))
			})

			// Write the chunk to the output file
			if _, err := out.Writer.Write(buf2[:n]); err != nil {
				return -1, fmt.Errorf("error writing output file: %v", err)
			}
		}

	}

	return out.Length, nil
}
