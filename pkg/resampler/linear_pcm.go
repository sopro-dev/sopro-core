package resampler

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

<<<<<<< HEAD
func linearPcm(in *sopro.In, out *sopro.Out, rs *Resampler) error {
	// read all the file
	if rs.Verbose {
=======
func linearPcm(in *sopro.In, out *sopro.Out, resampler *Resampler) error {
	// read all the file
	if resampler.Verbose {
>>>>>>> d598982 (refactor, new resampler, new transcoder, sopro models, sinc interpolation and more examples)
		sopro.GraphIn(in)
	}

	// Get the WAV file configuration
	channels := out.Config.(audioconfig.WavConfig).Channels
	sampleRate := out.Config.(audioconfig.WavConfig).SampleRate
	bitsPerSample := out.Config.(audioconfig.WavConfig).BitDepth
<<<<<<< HEAD
	rs.InConfigs.Encoding = in.Config.(audioconfig.WavConfig).Encoding
	rs.OutConfigs.Encoding = out.Config.(audioconfig.WavConfig).Encoding

	if rs.InConfigs.Endianness == cpuarch.NOT_FILLED && rs.OutConfigs.Endianness == cpuarch.NOT_FILLED {
		rs.InConfigs.Endianness = cpuarch.LITTLE_ENDIAN // replace with cpuarch.GetEndianess()
		rs.OutConfigs.Endianness = cpuarch.LITTLE_ENDIAN
	}

	rs.Println(
=======
	resampler.InConfigs.Encoding = in.Config.(audioconfig.WavConfig).Encoding
	resampler.OutConfigs.Encoding = out.Config.(audioconfig.WavConfig).Encoding

	if resampler.InConfigs.Endianness == cpuarch.NOT_FILLED && resampler.OutConfigs.Endianness == cpuarch.NOT_FILLED {
		resampler.InConfigs.Endianness = cpuarch.LITTLE_ENDIAN // replace with cpuarch.GetEndianess()
		resampler.OutConfigs.Endianness = cpuarch.LITTLE_ENDIAN
	}

	resampler.Println(
>>>>>>> d598982 (refactor, new resampler, new transcoder, sopro models, sinc interpolation and more examples)
		"\n[Format]                      ", in.Format, "=>", out.Format,
		"\n[Encoding]                    ", encoding.ENCODINGS[in.Config.(audioconfig.WavConfig).Encoding], "=>", encoding.ENCODINGS[out.Config.(audioconfig.WavConfig).Encoding],
		"\n[Channels]                    ", in.Config.(audioconfig.WavConfig).Channels, "=>", channels,
		"\n[SampleRate]                  ", in.Config.(audioconfig.WavConfig).SampleRate, "=>", sampleRate, "Hz",
		"\n[BitDepth]                    ", in.Config.(audioconfig.WavConfig).BitDepth, "=>", bitsPerSample, "bytes",
<<<<<<< HEAD
		"\n[Transcoder][Source][Encoding]", encoding.ENCODINGS[rs.InConfigs.Encoding],
		"\n[Transcoder][Target][Encoding]", encoding.ENCODINGS[rs.OutConfigs.Encoding],
=======
		"\n[Transcoder][Source][Encoding]", encoding.ENCODINGS[resampler.InConfigs.Encoding],
		"\n[Transcoder][Target][Encoding]", encoding.ENCODINGS[resampler.OutConfigs.Encoding],
>>>>>>> d598982 (refactor, new resampler, new transcoder, sopro models, sinc interpolation and more examples)
		"\n[Transcoder][Endianness]      ", cpuarch.ENDIANESSES[cpuarch.GetEndianess()],
	)

	// Create a buffered reader and writer
	in.Reader = bufio.NewReader(in.Data)
	out.Writer = bufio.NewWriter(out.Data)
	out.Length = 0
	headerWavZero := utils.GenerateSpaceForWavHeader()

	out.Writer.Write(headerWavZero)
	out.Length += 44

	in.Reader.Discard(44) // avoid first 44 bytes of in

	// Copy the data from the input file to the output file in chunks
<<<<<<< HEAD
	if _, err := ResampleBytes(in, out, rs); err != nil {
		return fmt.Errorf("error converting bytes: %v", err)
	}
=======
	if _, err := ResampleBytes(in, out, resampler); err != nil {
		return fmt.Errorf("error converting bytes: %v", err)
	}

>>>>>>> d598982 (refactor, new resampler, new transcoder, sopro models, sinc interpolation and more examples)
	// Flush the output file
	err := out.Writer.Flush()
	if err != nil {
		return fmt.Errorf("error flushing output file: %v", err)
	}
<<<<<<< HEAD
	rs.Println("Wrote", out.Length, "bytes to output file")
=======
	resampler.Println("Wrote", out.Length, "bytes to output file")
>>>>>>> d598982 (refactor, new resampler, new transcoder, sopro models, sinc interpolation and more examples)

	// Update the file size and data size fields
	fileFixer := out.Data.(*os.File)
	r, err := fileFixer.Seek(0, io.SeekStart)
	if err != nil {
		return fmt.Errorf("error seeking file: %v", err)
	}
<<<<<<< HEAD
	rs.Println("Seeked to:", r)
=======
	resampler.Println("Seeked to:", r)
>>>>>>> d598982 (refactor, new resampler, new transcoder, sopro models, sinc interpolation and more examples)

	out.NewConfig(audioconfig.WavConfig{
		BitDepth:   bitsPerSample,
		Channels:   channels,
<<<<<<< HEAD
		Encoding:   rs.OutConfigs.Encoding,
=======
		Encoding:   resampler.OutConfigs.Encoding,
>>>>>>> d598982 (refactor, new resampler, new transcoder, sopro models, sinc interpolation and more examples)
		SampleRate: sampleRate,
		WaveFormat: audioconfig.WAVE_FORMAT_PCM,
	})
	headersWav := utils.GenerateWavHeadersWithSize(out.Config.(audioconfig.WavConfig), out.Length)

<<<<<<< HEAD
	if rs.Verbose {
=======
	if resampler.Verbose {
>>>>>>> d598982 (refactor, new resampler, new transcoder, sopro models, sinc interpolation and more examples)
		audioconfig.PrintWavHeaders(headersWav)
	}

	n, err := fileFixer.Write(headersWav)
	if err != nil {
		return fmt.Errorf("error writing file size: %v", err)
	}
<<<<<<< HEAD
	rs.Println("File size:", fmt.Sprintf("% 02x", out.Length-8), "bytes written:", n)
	rs.Println("Data size:", fmt.Sprintf("% 02x", out.Length-44), "bytes written:", n)

	if rs.Verbose {
=======
	resampler.Println("File size:", fmt.Sprintf("% 02x", out.Length-8), "bytes written:", n)
	resampler.Println("Data size:", fmt.Sprintf("% 02x", out.Length-44), "bytes written:", n)

	if resampler.Verbose {
>>>>>>> d598982 (refactor, new resampler, new transcoder, sopro models, sinc interpolation and more examples)
		sopro.GraphOut(in, out)
	}

	return nil
}
