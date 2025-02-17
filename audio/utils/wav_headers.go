package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"sync"
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

// WavHeader represents the WAV file header
type WavHeader struct {
	Length     uint32     `json:"length"`
	WaveFormat WaveFormat `json:"wave_format"`
	Channels   uint16     `json:"channels"`
	SampleRate uint32     `json:"sample_rate"`
	BitDepth   uint16     `json:"bit_depth"`
	Verbose    bool       `json:"verbose"`
}

type WaveFormat uint16

// WAV FORMAT CODES
const (
	WAVE_FORMAT_UNKNOWN            = 0x0000                // Microsoft Unknown Wave Format
	WAVE_FORMAT_PCM                = 0x0001                // Microsoft PCM Format
	WAVE_FORMAT_ADPCM              = 0x0002                // Microsoft ADPCM Format
	WAVE_FORMAT_IEEE_FLOAT         = 0x0003                // IEEE float
	WAVE_FORMAT_VSELP              = 0x0004                // Compaq Computer's VSELP
	WAVE_FORMAT_IBM_CVSD           = 0x0005                // IBM CVSD
	WAVE_FORMAT_ALAW               = 0x0006                // ALAW
	WAVE_FORMAT_MULAW              = 0x0007                // MULAW
	WAVE_FORMAT_DTS                = 0x0008                // Digital Theater Systems DTS
	WAVE_FORMAT_DRM                = 0x0009                // Microsoft Corporation
	WAVE_FORMAT_WMAVOICE9          = 0x000A                // Microsoft Corporation
	WAVE_FORMAT_WMAVOICE10         = 0x000B                // Microsoft Corporation
	WAVE_FORMAT_OKI_ADPCM          = 0x0010                // OKI ADPCM
	WAVE_FORMAT_DVI_ADPCM          = 0x0011                // Intel's DVI ADPCM
	WAVE_FORMAT_IMA_ADPCM          = WAVE_FORMAT_DVI_ADPCM // Intel's DVI ADPCM
	WAVE_FORMAT_MEDIASPACE_ADPCM   = 0x0012                // Videologic's MediaSpace ADPCM
	WAVE_FORMAT_SIERRA_ADPCM       = 0x0013                // Sierra ADPCM
	WAVE_FORMAT_G723_ADPCM         = 0x0014                // G.723 ADPCM
	WAVE_FORMAT_DIGISTD            = 0x0015                // DSP Solution's DIGISTD
	WAVE_FORMAT_DIGIFIX            = 0x0016                // DSP Solution's DIGIFIX
	WAVE_FORMAT_DIALOGIC_OKI_ADPCM = 0x0017                // Dialogic Corporation
	WAVE_FORMAT_MEDIAVISION_ADPCM  = 0x0018                // Media Vision ADPCM
)

// GenerateWavHeadersWithConfig generates the headers for a WAV file without the size fields
func GenerateWavHeadersWithConfig(head *WavHeader) []byte {
	if head.WaveFormat == 0 {
		fmt.Println("WARNING: WaveFormat not set, defaulting to PCM (1)")
		head.WaveFormat = 1
	}

	totalLength := head.Length - 8
	subChunk2Size := head.Length - 44

	chunkID := []byte{'R', 'I', 'F', 'F'}
	format := []byte{'W', 'A', 'V', 'E'}
	subChunk1ID := []byte{'f', 'm', 't', ' '}
	subChunk1Size := uint32(16)            // Update subChunk1Size to 16
	audioFormat := uint16(head.WaveFormat) // Update audioFormat to uint16
	numChannels := uint16(head.Channels)   // Update numChannels to uint16

	sampleRate := uint32(head.SampleRate)
	byteRate := sampleRate * uint32(head.Channels) * uint32(head.BitDepth/8)
	blockAlign := uint16(head.Channels * (head.BitDepth / 8))
	bitsPerSample := uint16(head.BitDepth)

	subChunk2ID := []byte{'d', 'a', 't', 'a'}

	if head.Verbose {

		b, err := json.Marshal(head)
		if err != nil {
			fmt.Println("Error: ", err)
		}

		fmt.Println("==== Audio File Config ====")
		fmt.Println(string(b))
		fmt.Println("==== Audio File Header ====")

		fmt.Printf("[ 1- 4] Chunk ID:          <%s>\n", string(chunkID))
		fmt.Printf("[ 5- 8] Chunk Size:        <%d> =%d bytes\n", totalLength, totalLength)
		fmt.Printf("[ 9-12] Format:            <%s>\n", string(format))
		fmt.Printf("[13-16] SubChunk1 ID:      <%s>\n", string(subChunk1ID))
		fmt.Printf("[17-20] SubChunk1 Size:    <%d>\n", subChunk1Size)
		fmt.Printf("[21-22] Audio Format:      <%d>\n", audioFormat)
		fmt.Printf("[23-24] Num Channels:      <%d>\n", numChannels)
		fmt.Printf("[25-28] Sample Rate:       <%d> =%d Hz\n", sampleRate, head.SampleRate)
		fmt.Printf("[29-32] Byte Rate:         <%d> =%d bytes/sec\n", byteRate, byteRate)
		fmt.Printf("[33-34] Block Align:       <%d>\n", blockAlign)
		fmt.Printf("[35-36] Bits Per Sample:   <%d>\n", bitsPerSample)
		fmt.Printf("[37-40] SubChunk2 ID:      <%s>\n", string(subChunk2ID))
		fmt.Printf("[41-44] SubChunk2 Size:    <%d>\n", subChunk2Size)
		fmt.Println("==========================")
	}

	return MergeSliceOfBytes(
		chunkID,
		IntToBytes(totalLength),
		format,
		subChunk1ID,
		IntToBytes(subChunk1Size),
		IntToBytes(audioFormat),
		IntToBytes(numChannels),
		IntToBytes(sampleRate),
		IntToBytes(byteRate),
		IntToBytes(blockAlign),
		IntToBytes(bitsPerSample),
		subChunk2ID,
		IntToBytes(subChunk2Size),
	)
}

// PrintWavHeaders prints the headers of a WAV file
// first 44 bytes of a WAV file
func PrintWavHeaders(headersWav []byte) {
	if len(headersWav) != 44 {
		log.Println("[ERROR] Headers are not 44 bytes long")
		return
	}
	fmt.Println("Headers (WAV):")
	comments := []string{
		"(4) Chunk ID [RIFF]",
		"(4) Chunk size",
		"(4) Format [WAVE]",
		"(4) Sub-chunk 1 ID [fmt ]",
		"(4) Sub-chunk 1 size",
		"(2) Audio format (PCM) & (2) Number of channels",
		"(4) Sample rate",
		"(4) Byte rate",
		"(2) Block align & (2) Bits per sample",
		"(4) Sub-chunk 2 ID [data]",
		"(4) Sub-chunk 2 size",
	}
	for i := 0; i < 44; i += 4 {
		fmt.Println(
			fmt.Sprintf("[%2d,%2d]", i, i+4),
			fmt.Sprintf("% 2x", headersWav[i:i+4]),
			"\t<"+strings.ToUpper(string(headersWav[i:i+4]))+">\t",
			comments[i/4],
		)
	}
}
