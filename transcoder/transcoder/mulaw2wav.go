package transcoder

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/pablodz/transcoder/transcoder/audioconfig"
	"github.com/pablodz/transcoder/transcoder/encoding"
)

// TODO: split functions for different sizes of files
// Transcode an ulaw file to a wav file (large files supported)
func Mulaw2Wav(in *AudioFileIn, out *AudioFileOut) (err error) {

	// Get the WAV file configuration
	channels := out.Config.(audioconfig.WavConfig).Channels
	sampleRate := out.Config.(audioconfig.WavConfig).SampleRate
	bitsPerSample := out.Config.(audioconfig.WavConfig).BitDepth

	// Create a buffered reader and writer
	in.Reader = bufio.NewReader(in.Data)
	out.Writer = bufio.NewWriter(out.Data)

	// Set the WAV file header
	out.Writer.Write([]byte{'R', 'I', 'F', 'F'})
	out.Writer.Write([]byte{0, 0, 0, 0}) // placeholder for file size
	out.Writer.Write([]byte{'W', 'A', 'V', 'E'})
	out.Writer.Write([]byte{'f', 'm', 't', ' '})
	out.Writer.Write([]byte{16, 0, 0, 0})       // size of fmt chunk
	out.Writer.Write([]byte{1, 0})              // audio format (PCM)
	out.Writer.Write([]byte{byte(channels), 0}) // number of channels

	// Set the sample rate and byte rate
	sampleRateBytes := []byte{byte(sampleRate & 0xff), byte(sampleRate >> 8 & 0xff), byte(sampleRate >> 16 & 0xff), byte(sampleRate >> 24 & 0xff)}
	out.Writer.Write(sampleRateBytes)
	byteRate := []byte{byte(sampleRate * channels * (bitsPerSample / 8) & 0xff), byte(sampleRate * channels * (bitsPerSample / 8) >> 8 & 0xff), byte(sampleRate * channels * (bitsPerSample / 8) >> 16 & 0xff), byte(sampleRate * channels * (bitsPerSample / 8) >> 24 & 0xff)}
	out.Writer.Write(byteRate)
	out.Writer.Write([]byte{byte(channels * (bitsPerSample / 8)), 0}) // block align
	out.Writer.Write([]byte{byte(bitsPerSample), 0})                  // bits per sample

	// Write the data chunk header
	out.Writer.Write([]byte{'d', 'a', 't', 'a'})
	out.Writer.Write([]byte{0, 0, 0, 0}) // placeholder for data size

	// Copy the data from the input file to the output file in chunks
	if err = ConvertBytes(in, out); err != nil {
		return err
	}

	// Update the file size and data size fields
	out.File.Seek(4, 0)
	fileInfo, err := out.File.Stat()
	if err != nil {
		return fmt.Errorf("error getting file info: %v", err)
	}
	fileSize := []byte{byte(fileInfo.Size() - 8&0xff), byte(fileInfo.Size() - 8>>8&0xff), byte(fileInfo.Size() - 8>>16&0xff), byte(fileInfo.Size() - 8>>24&0xff)}
	out.File.Write(fileSize)
	out.File.Seek(40, 0)
	dataSize := []byte{byte((fileInfo.Size() - 44) & 0xff), byte((fileInfo.Size() - 44) >> 8 & 0xff), byte((fileInfo.Size() - 44) >> 16 & 0xff), byte((fileInfo.Size() - 44) >> 24 & 0xff)}
	out.File.Write(dataSize)
	return nil
}

func ConvertBytes(in *AudioFileIn, out *AudioFileOut) error {

	buf := make([]byte, 1024) // read and write in chunks of 1024 bytes
	dif_encoding := strings.EqualFold(
		in.Config.(audioconfig.MulawConfig).Encoding,
		out.Config.(audioconfig.WavConfig).Encoding,
	)
	if !dif_encoding {
		for {
			// Read a chunk from the input file
			n, err := in.Reader.Read(buf)
			if err != nil && err != io.EOF {
				return fmt.Errorf("error reading input file: %v", err)
			}
			if n == 0 {
				break
			}

			// Write the chunk to the output file
			if _, err := out.Writer.Write(buf[:n]); err != nil {
				return fmt.Errorf("error writing output file: %v", err)
			}
		}
	} else {
		for {
			// Read a chunk from the input file
			n, err := in.Reader.Read(buf)
			if err != nil && err != io.EOF {
				return fmt.Errorf("error reading input file: %v", err)
			}
			if n == 0 {
				break
			}

			switch {
			case in.Config.(audioconfig.MulawConfig).Encoding == audioconfig.MULAW_ENCODING_MULAW &&
				out.Config.(audioconfig.WavConfig).Encoding == audioconfig.PCM_ENCODING_SIGNED_16_PCM:
				buf = encoding.Mulaw2Linear(buf)
			}

			// Write the chunk to the output file
			if _, err := out.Writer.Write(buf[:n]); err != nil {
				return fmt.Errorf("error writing output file: %v", err)
			}
		}
	}

	if err := out.Writer.Flush(); err != nil {
		return fmt.Errorf("error flushing output file: %v", err)
	}

	return nil
}
