package transcoder

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/pablodz/sopro/pkg/audioconfig"
	"github.com/pablodz/sopro/pkg/cpuarch"
	"github.com/pablodz/sopro/pkg/encoding"
	"golang.org/x/term"
)

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
func mulaw2WavLogpcm(in *AudioFileIn, out *AudioFileOut, transcoder *Transcoder) (err error) {

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
		audioconfig.WAVE_FORMAT_MULAW, 0, // Audio format (1 = PCM)
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
