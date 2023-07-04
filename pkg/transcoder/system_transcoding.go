package transcoder

import (
	"fmt"
	"io"

	"github.com/pablodz/sopro/pkg/sopro"
	"github.com/pablodz/sopro/pkg/transcoder/decoder"
)

func TranscodeBytes(in *sopro.In, out *sopro.Out, tr *Transcoder) (int, error) {
	equalEncod := (tr.InConfigs.Encoding == tr.OutConfigs.Encoding)

	if equalEncod {
		tr.Println("Same encodings - no transcoding needed")
		bitsProcessed, err := equalSpaceEncoding(in, out, tr)
		if err != nil {
			return bitsProcessed, err
		}
	} else {
		tr.Println("Different encodings - transcoding needed")
		bitsProcessed, err := differentSpaceEncoding(in, out, tr)
		if err != nil {
			return bitsProcessed, err
		}
	}
	return 0, nil
}

func equalSpaceEncoding(in *sopro.In, out *sopro.Out, t *Transcoder) (int, error) {
	sizeBuff := 1024 // max size, more than that would be too much
	if t.SizeBuffer > 0 {
		sizeBuff = t.SizeBuffer
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

func differentSpaceEncoding(in *sopro.In, out *sopro.Out, transcoder *Transcoder) (int, error) {
	sizeBuff := 1024 // max size, more than that would be too much
	if transcoder.SizeBuffer > 0 {
		sizeBuff = transcoder.SizeBuffer
	}
	nTotal := 0
	bufIn := make([]byte, sizeBuff)    // input buffer
	bufOut := make([]byte, sizeBuff*2) //  output buffer
	for {
		n, err := in.Reader.Read(bufIn)
		if err != nil && err != io.EOF {
			return nTotal, fmt.Errorf("error reading input file: %v", err)
		}
		if n == 0 {
			break
		}
		bufIn = bufIn[:n]
		// buf2 is different size than buf
		bufOut, _ = decoder.DecodeULawToPCM(bufIn) // IMPORTANT:buf cut to n bytes
		out.Length += len(bufOut)
		if _, err = out.Writer.Write(bufOut); err != nil {
			return nTotal, fmt.Errorf("error writing output file: %v", err)
		}
		nTotal += n

		doOnceTranscoding.Do(func() {
			transcoder.Println("[Transcoder] Transcoding data - sample of the first 4 bytes (hex)")
			onlyNFirst := 4
			transcoder.Println(
				"[OLD]", fmt.Sprintf("% 2x", bufIn[:onlyNFirst]),
				"\n[NEW]", fmt.Sprintf("% 2x", bufOut[:onlyNFirst*2]),
			)
			transcoder.Println("[Transcoder] Transcoding data - sample of the first 4 bytes (decimal)")
			transcoder.Println(
				"[OLD]", fmt.Sprintf("%3d", bufIn[:onlyNFirst]),
				"\n[NEW]", fmt.Sprintf("%3d", bufOut[:onlyNFirst*2]),
			)
		})
	}
	return nTotal, nil
}
