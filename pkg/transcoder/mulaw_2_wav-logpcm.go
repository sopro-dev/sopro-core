package transcoder

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/pablodz/sopro/pkg/audioconfig"
	"github.com/pablodz/sopro/pkg/cpuarch"
	"github.com/pablodz/sopro/pkg/encoding"
	"github.com/pablodz/sopro/pkg/sopro"
	"github.com/pablodz/sopro/utils"
)

// Transcode an ulaw file to a wav file (large files supported)
// https://raw.githubusercontent.com/corkami/pics/master/binary/WAV.png
// http://www.topherlee.com/software/pcm-tut-wavformat.html
func mulaw2WavLogpcm(in *sopro.In, out *sopro.Out, tr *Transcoder) (err error) {
	// read all the file
	if tr.Verbose {
		sopro.GraphIn(in)
	}

	// Get the WAV file configuration
	channels := out.Config.(audioconfig.WavConfig).Channels
	sampleRate := out.Config.(audioconfig.WavConfig).SampleRate
	bitsPerSample := out.Config.(audioconfig.WavConfig).BitDepth
	tr.InConfigs.Encoding = in.Config.(audioconfig.MulawConfig).Encoding
	tr.OutConfigs.Encoding = out.Config.(audioconfig.WavConfig).Encoding
	tr.BitDepth = bitsPerSample

	if tr.InConfigs.Endianness == cpuarch.NOT_FILLED && tr.OutConfigs.Endianness == cpuarch.NOT_FILLED {
		tr.InConfigs.Endianness = cpuarch.LITTLE_ENDIAN // replace with cpuarch.GetEndianess()
		tr.OutConfigs.Endianness = cpuarch.LITTLE_ENDIAN
	}

	tr.Println(
		"\n[Format]                      ", in.Format, "=>", out.Format,
		"\n[Encoding]                    ", encoding.ENCODINGS[in.Config.(audioconfig.MulawConfig).Encoding], "=>", encoding.ENCODINGS[out.Config.(audioconfig.WavConfig).Encoding],
		"\n[Channels]                    ", in.Config.(audioconfig.MulawConfig).Channels, "=>", channels,
		"\n[SampleRate]                  ", in.Config.(audioconfig.MulawConfig).SampleRate, "=>", sampleRate, "Hz",
		"\n[BitDepth]                    ", in.Config.(audioconfig.MulawConfig).BitDepth, "=>", bitsPerSample, "bytes",
		"\n[Transcoder][Source][Encoding]", encoding.ENCODINGS[tr.InConfigs.Encoding],
		"\n[Transcoder][Target][Encoding]", encoding.ENCODINGS[tr.OutConfigs.Encoding],
		"\n[Transcoder][BitDepth]        ", tr.BitDepth,
		"\n[Transcoder][Endianness]      ", cpuarch.ENDIANESSES[cpuarch.GetEndianess()],
	)

	// Create a buffered reader and writer
	in.Reader = bufio.NewReader(in.Data)
	out.Writer = bufio.NewWriter(out.Data)
	out.Length = 0
	headerWavZero := utils.GenerateSpaceForWavHeader()

	out.Writer.Write(headerWavZero)
	out.Length += 44

	// Copy the data from the input file to the output file in chunks
	if _, err := TranscodeBytes(in, out, tr); err != nil {
		return fmt.Errorf("error converting bytes: %v", err)
	}

	// Flush the output file
	err = out.Writer.Flush()
	if err != nil {
		return fmt.Errorf("error flushing output file: %v", err)
	}
	tr.Println("Wrote", out.Length, "bytes to output file")

	// Update the file size and data size fields
	fileFixer := out.Data.(*os.File)
	r, err := fileFixer.Seek(0, io.SeekStart)
	if err != nil {
		return fmt.Errorf("error seeking file: %v", err)
	}
	tr.Println("Seeked to:", r)

	out.NewConfig(audioconfig.WavConfig{
		Channels:   channels,
		SampleRate: sampleRate,
		BitDepth:   bitsPerSample,
		Encoding:   tr.OutConfigs.Encoding,
		WaveFormat: audioconfig.WAVE_FORMAT_MULAW,
	})
	headersWav := utils.GenerateWavHeadersWithSize(out.Config.(audioconfig.WavConfig), out.Length)

	if tr.Verbose {
		audioconfig.PrintWavHeaders(headersWav)
	}

	n, err := fileFixer.Write(headersWav)
	if err != nil {
		return fmt.Errorf("error writing file size: %v", err)
	}
	tr.Println("File size:", fmt.Sprintf("% 02x", out.Length-8), "bytes written:", n)
	tr.Println("Data size:", fmt.Sprintf("% 02x", out.Length-44), "bytes written:", n)

	if tr.Verbose {
		sopro.GraphOut(in, out)
	}

	return nil
}
