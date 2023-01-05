package transcoder

import (
	"bufio"
	"fmt"
	"io"

	"github.com/pablodz/transcoder/transcoder/audioconfig"
)

// Pcm2Wav converts a PCM file to a WAV file
func Pcm2Wav(in AudioFileIn, out AudioFileOut) (err error) {
	// Get the WAV file configuration
	channels := out.Config.(audioconfig.WavConfig).Channels
	sampleRate := out.Config.(audioconfig.WavConfig).SampleRate
	bitsPerSample := out.Config.(audioconfig.WavConfig).BitDepth

	// Create a buffered reader and writer
	inputReader := bufio.NewReader(in.Data)
	outputWriter := bufio.NewWriter(out.Data)

	// Set the WAV file header
	outputWriter.Write([]byte{'R', 'I', 'F', 'F'})
	outputWriter.Write([]byte{0, 0, 0, 0}) // placeholder for file size
	outputWriter.Write([]byte{'W', 'A', 'V', 'E'})
	outputWriter.Write([]byte{'f', 'm', 't', ' '})
	outputWriter.Write([]byte{16, 0, 0, 0})       // size of fmt chunk
	outputWriter.Write([]byte{1, 0})              // audio format (PCM)
	outputWriter.Write([]byte{byte(channels), 0}) // number of channels

	// Set the sample rate and byte rate
	// sampleRateBytes := []byte{byte(sampleRate & 0xff), byte(sampleRate >> 8 & 0xff),

	// Set the sample rate and byte rate
	sampleRateBytes := []byte{byte(sampleRate & 0xff), byte(sampleRate >> 8 & 0xff), byte(sampleRate >> 16 & 0xff), byte(sampleRate >> 24 & 0xff)}
	outputWriter.Write(sampleRateBytes)
	byteRate := []byte{byte(sampleRate * channels * (bitsPerSample / 8) & 0xff), byte(sampleRate * channels * (bitsPerSample / 8) >> 8 & 0xff), byte(sampleRate * channels * (bitsPerSample / 8) >> 16 & 0xff), byte(sampleRate * channels * (bitsPerSample / 8) >> 24 & 0xff)}
	outputWriter.Write(byteRate)
	outputWriter.Write([]byte{byte(channels * (bitsPerSample / 8)), 0}) // block align
	outputWriter.Write([]byte{byte(bitsPerSample), 0})                  // bits per sample
	// Write the data chunk header
	outputWriter.Write([]byte{'d', 'a', 't', 'a'})
	outputWriter.Write([]byte{0, 0, 0, 0}) // placeholder for data size

	// Copy the data from the input file to the output file in chunks
	buf := make([]byte, 1024) // read and write in chunks of 1024 bytes
	for {
		// Read a chunk from the input file
		n, err := inputReader.Read(buf)
		if err != nil && err != io.EOF {
			return fmt.Errorf("error reading input file: %v", err)
		}
		if n == 0 {
			break
		}
		// Write the chunk to the output file
		if _, err := outputWriter.Write(buf[:n]); err != nil {
			return fmt.Errorf("error writing output file: %v", err)
		}
	}
	if err = outputWriter.Flush(); err != nil {
		return fmt.Errorf("error flushing output file: %v", err)
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
