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
		bitsProcessed, err = equalSpaceEncoding(in, out)
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

func equalSpaceEncoding(in *AudioFileIn, out *AudioFileOut) (int, error) {
	nTotal, err := 0, error(nil)
	buf := make([]byte, 32) // read and write in chunks of 1024 byte
	for err == nil {
		n, err := in.Reader.Read(buf)
		if err != nil && err != io.EOF {
			return nTotal, fmt.Errorf("error reading input file: %v", err)
		}
		if n == 0 {
			break
		}
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
	nTotal, err := 0, error(nil)
	buf := make([]byte, 32) // read and write in chunks of 1024 byte
	for err == nil {
		n, err := in.Reader.Read(buf)
		if err != nil && err != io.EOF {
			return nTotal, fmt.Errorf("error reading input file: %v", err)
		}
		if n == 0 {
			break
		}
		// buf2 is different size than buf
		buf2, _ := decoder.DecodeFrameUlaw2Lpcm(buf)
		out.Length += len(buf2)
		nTotal += n

		if _, err = out.Writer.Write(buf); err != nil {
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
