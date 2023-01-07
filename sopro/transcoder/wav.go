package transcoder

import (
	"fmt"
	"log"
	"strings"
)

func printWavHeaders(headersWav []byte) {
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
