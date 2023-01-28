package utils

import (
	"fmt"
	"sync"

	"github.com/pablodz/sopro/pkg/audioconfig"
)

// Generate a pool to only allocate one time
var bytePool = sync.Pool{
	New: func() interface{} {
		// Return a new slice of 44 bytes for wav headers
		return make([]byte, 44)
	},
}

func GenerateSpaceForWavHeader() []byte {
	// Get a new slice from the pool
	buffer := bytePool.Get().([]byte)
	// Release the slice back to the pool
	bytePool.Put(&buffer)
	return buffer
}

// GenerateWavHeadersWithSize generates the headers for a wav file without the size fields
func GenerateWavHeadersWithSize(config audioconfig.WavConfig, length int) []byte {
	if config.WaveFormat == 0 {
		fmt.Println("WARNING: WaveFormat not set, defaulting to PCM (1)")
		config.WaveFormat = 1
	}

	totalLength := length - 8
	if totalLength < 0 {
		totalLength = 0
	}
	subChunk2Size := length - 44
	if subChunk2Size < 0 {
		subChunk2Size = 0
	}

	return MergeSliceOfBytes(
		[]byte{'R', 'I', 'F', 'F'}, // Chunk ID
		[]byte{
			byte((length - 8) & 0xff),       // 1st byte of length
			byte((length - 8) >> 8 & 0xff),  // 2nd byte of length
			byte((length - 8) >> 16 & 0xff), // 3rd byte of length
			byte((length - 8) >> 24 & 0xff), // 4th byte of length
		}, // Chunk size
		[]byte{'W', 'A', 'V', 'E'},                                                             // Format
		[]byte{'f', 'm', 't', ' '},                                                             // Sub-chunk 1 ID
		[]byte{16, 0, 0, 0},                                                                    // Sub-chunk 1 size
		[]byte{config.WaveFormat, 0},                                                           // Audio format (PCM)
		[]byte{byte(config.Channels), 0},                                                       // Number of channels
		[]byte{byte(config.SampleRate & 0xFF)},                                                 // sample rate (low)
		[]byte{byte(config.SampleRate >> 8 & 0xFF)},                                            // sample rate (mid)
		[]byte{byte(config.SampleRate >> 16 & 0xFF)},                                           // sample rate (high)
		[]byte{byte(config.SampleRate >> 24 & 0xFF)},                                           // sample rate (high)
		[]byte{byte(config.SampleRate * config.Channels * (config.BitDepth / 8) & 0xFF)},       // byte rate (low)
		[]byte{byte(config.SampleRate * config.Channels * (config.BitDepth / 8) >> 8 & 0xFF)},  // byte rate (mid)
		[]byte{byte(config.SampleRate * config.Channels * (config.BitDepth / 8) >> 16 & 0xFF)}, // byte rate (high)
		[]byte{byte(config.SampleRate * config.Channels * (config.BitDepth / 8) >> 24 & 0xFF)}, // byte rate (high)
		[]byte{byte(config.Channels * (config.BitDepth / 8)), 0},                               // block align
		[]byte{byte(config.BitDepth), 0},                                                       // bits per sample
		[]byte{'d', 'a', 't', 'a'},                                                             // Sub-chunk 2 ID
		[]byte{
			byte((length - 44) & 0xff),       // 1st byte of sub-chunk 2 size
			byte((length - 44) >> 8 & 0xff),  // 2nd byte of sub-chunk 2 size
			byte((length - 44) >> 16 & 0xff), // 3rd byte of sub-chunk 2 size
			byte((length - 44) >> 24 & 0xff), // 4th byte of sub-chunk 2 size
		}, // Sub-chunk 2 size
	)
}
