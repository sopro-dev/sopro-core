package transcoder

import (
	"fmt"
	"io"
	"sync"

	"github.com/pablodz/sopro/pkg/audioconfig"
	"github.com/pablodz/sopro/pkg/resampler"
)

var doOnceResampling sync.Once

func ResampleBytes(in *AudioFileIn, out *AudioFileOut, transcoder *Transcoder) error {
	bitsProcessed, err := differentSampleRate(in, out, transcoder)
	if err != nil {
		return err
	}
	transcoder.Println("Transcoding done:", bitsProcessed, "bits processed")

	return nil
}

func differentSampleRate(in *AudioFileIn, out *AudioFileOut, transcoder *Transcoder) (int, error) {
	sizeBuff := 1024 // max size, more than that would be too much
	if transcoder.SizeBuffer > 0 {
		sizeBuff = transcoder.SizeBuffer
	}
	nTotal := 0
	sampleRateIn := in.Config.(audioconfig.WavConfig).SampleRate
	sampleRateOut := out.Config.(audioconfig.WavConfig).SampleRate
	ratio := float64(sampleRateIn) / float64(sampleRateOut)

	bufIn := make([]byte, sizeBuff)             // input buffer
	bufOut := make([]byte, sizeBuff*int(ratio)) // output buffer
	transcoder.Println("ratio", ratio, "ratioInt", int(ratio))
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
		bufOut, _ = resampler.LinearInterpolation(bufIn, ratio) // IMPORTANT:buf cut to n bytes
		out.Length += len(bufOut)
		if _, err = out.Writer.Write(bufOut); err != nil {
			return nTotal, fmt.Errorf("error writing output file: %v", err)
		}
		nTotal += n

		doOnceResampling.Do(func() {
			transcoder.Println("[Transcoder] Transcoding data - sample of the first 4 bytes (hex)")
			onlyNFirst := 8
			transcoder.Println(
				"[OLD]", fmt.Sprintf("% 2x", bufIn[:onlyNFirst]),
				"\n[NEW]", fmt.Sprintf("% 2x", bufOut[:onlyNFirst/2]),
			)
			transcoder.Println("[Transcoder] Transcoding data - sample of the first 4 bytes (decimal)")
			transcoder.Println(
				"[OLD]", fmt.Sprintf("%3d", bufIn[:onlyNFirst]),
				"\n[NEW]", fmt.Sprintf("%3d", bufOut[:onlyNFirst/2]),
			)
		})
	}
	return nTotal, nil
}
