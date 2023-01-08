package transcoder

import (
	"fmt"
	"io"
	"log"
	"sync"

	"github.com/pablodz/sopro/sopro/decoder"
)

var doOnceTranscoding sync.Once

func TranscodeBytes(in *AudioFileIn, out *AudioFileOut, transcoder *Transcoder) error {

	equalEncod := (transcoder.SourceConfigs.Encoding == transcoder.TargetConfigs.Encoding)
	bitsProcessed := 0
	err := error(nil)

	if equalEncod {

		log.Println("Same encodings - no transcoding needed")
		bitsProcessed, err = equalSpaceEncoding(in, out, transcoder)
		if err != nil {
			return err
		}

	} else {

		transcoder.Println("Different encodings - transcoding needed")
		bitsProcessed, err = differentSpaceEncoding(in, out, transcoder)
		if err != nil {
			return err
		}

	}

	log.Println("Transcoding done:", bitsProcessed, "bits processed")

	return nil
}

func equalSpaceEncoding(in *AudioFileIn, out *AudioFileOut, transcoder *Transcoder) (int, error) {
	sizeBuff := 1024 // max size, more than that would be too much
	if transcoder.SizeBufferToProcess > 0 {
		sizeBuff = transcoder.SizeBufferToProcess
	}
	nTotal := 0
	buf := make([]byte, sizeBuff) // read and write in chunks of 1024 byte
	for {
		n, err := in.Reader.Read(buf)
		if err != nil && err != io.EOF {
			return nTotal, fmt.Errorf("error reading input file: %v", err)
		}
		if n == 0 {
			break
		}
		buf = buf[:n]
		// TODO: add some function here
		out.Length += len(buf)
		nTotal += n

		if _, err = out.Writer.Write(buf); err != nil {
			return nTotal, fmt.Errorf("error writing output file: %v", err)
		}
	}
	return nTotal, nil
}

func differentSpaceEncoding(in *AudioFileIn, out *AudioFileOut, transcoder *Transcoder) (int, error) {
	sizeBuff := 1024 // max size, more than that would be too much
	if transcoder.SizeBufferToProcess > 0 {
		sizeBuff = transcoder.SizeBufferToProcess
	}
	nTotal := 0
	buf := make([]byte, sizeBuff)    // input buffer
	buf2 := make([]byte, sizeBuff*2) //  output buffer
	for {
		n, err := in.Reader.Read(buf)
		if err != nil && err != io.EOF {
			return nTotal, fmt.Errorf("error reading input file: %v", err)
		}
		if n == 0 {
			break
		}
		buf = buf[:n]
		// buf2 is different size than buf
		buf2, _ = decoder.DecodeFrameUlaw2Lpcm(buf) // IMPORTANT:buf cut to n bytes
		out.Length += len(buf2)
		if _, err = out.Writer.Write(buf2); err != nil {
			return nTotal, fmt.Errorf("error writing output file: %v", err)
		}

		doOnceTranscoding.Do(func() {
			transcoder.Println("[Transcoder] Transcoding data - sample of the first 4 bytes (hex)")
			onlyNFirst := 4
			transcoder.Println(
				"[OLD]", fmt.Sprintf("% 2x", buf[:onlyNFirst]),
				"\n[NEW]", fmt.Sprintf("% 2x", buf2[:onlyNFirst*2]),
			)
			transcoder.Println("[Transcoder] Transcoding data - sample of the first 4 bytes (decimal)")
			transcoder.Println(
				"[OLD]", fmt.Sprintf("%3d", buf[:onlyNFirst]),
				"\n[NEW]", fmt.Sprintf("%3d", buf2[:onlyNFirst*2]),
			)
		})
	}
	return nTotal, nil
}
