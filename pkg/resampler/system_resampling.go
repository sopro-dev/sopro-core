package resampler

import (
	"fmt"
	"io"

	"github.com/pablodz/sopro/pkg/audioconfig"
	"github.com/pablodz/sopro/pkg/resampler/interpolation"
	"github.com/pablodz/sopro/pkg/sopro"
)

func ResampleBytes(in *sopro.In, out *sopro.Out, rs *Resampler) (int, error) {

	switch rs.MethodR {
	case LINEAR_INTERPOLATION:
		return linear_interpolation(in, out, rs)
	case BAND_LIMITED_INTERPOLATION:
		return band_limited_interpolation(in, out, rs)
	case FRACTIONAL_DELAY_FILTER:
		return fractional_delay_filter(in, out, rs)
	}

	return 0, nil
}

func linear_interpolation(in *sopro.In, out *sopro.Out, rs *Resampler) (int, error) {
	sizeBuff := 1024 // max size, more than that would be too much
	if rs.SizeBuffer > 0 {
		sizeBuff = rs.SizeBuffer
	}
	nTotal := 0
	sampleRateIn := in.Config.(audioconfig.WavConfig).SampleRate
	sampleRateOut := out.Config.(audioconfig.WavConfig).SampleRate
	ratioIO := float64(sampleRateIn) / float64(sampleRateOut)
	// ratioOI := float64(sampleRateOut) / float64(sampleRateIn)

	bufIn := make([]byte, sizeBuff)               // input buffer
	bufOut := make([]byte, sizeBuff*int(ratioIO)) // output buffer
	rs.Println("ratio", ratioIO, "ratioInt", int(ratioIO))
	for {
		n, err := in.Reader.Read(bufIn)
		if err != nil && err != io.EOF {
			return nTotal, fmt.Errorf("error reading input file: %v", err)
		}
		if n == 0 {
			break
		}
		bufIn = bufIn[:n] // buf2 is different size than buf

		bufOut, _ = interpolation.LinearInterpolation(bufIn, ratioIO) // IMPORTANT:buf cut to n bytes
		out.Length += len(bufOut)
		if _, err = out.Writer.Write(bufOut); err != nil {
			return nTotal, fmt.Errorf("error writing output file: %v", err)
		}
		nTotal += n

		doOnceResampling.Do(func() {
			rs.Println("[Transcoder] Transcoding data - sample of the first 4 bytes (hex)")
			onlyNFirst := 8
			rs.Println(
				"[OLD]", fmt.Sprintf("% 2x", bufIn[:onlyNFirst]),
				"\n[NEW]", fmt.Sprintf("% 2x", bufOut[:onlyNFirst/2]),
			)
			rs.Println("[Transcoder] Transcoding data - sample of the first 4 bytes (decimal)")
			rs.Println(
				"[OLD]", fmt.Sprintf("%3d", bufIn[:onlyNFirst]),
				"\n[NEW]", fmt.Sprintf("%3d", bufOut[:onlyNFirst/2]),
			)
		})
	}
	return nTotal, nil
}

func band_limited_interpolation(in *sopro.In, out *sopro.Out, r *Resampler) (int, error) {
	sizeBuff := 1024 // max size, more than that would be too much
	if r.SizeBuffer > 0 {
		sizeBuff = r.SizeBuffer
	}
	nTotal := 0
	sampleRateIn := in.Config.(audioconfig.WavConfig).SampleRate
	sampleRateOut := out.Config.(audioconfig.WavConfig).SampleRate
	ratioIO := float64(sampleRateIn) / float64(sampleRateOut)

	sizeOutBuff := int(float64(sizeBuff) * ratioIO)
	bufIn := make([]byte, sizeBuff)     // input buffer
	bufOut := make([]byte, sizeOutBuff) // output buffer
	r.Println("ratio", ratioIO, "ratioInt", int(ratioIO), "sizeBuff", sizeBuff, "sizeOutBuff", sizeOutBuff)
	for {
		n, err := in.Reader.Read(bufIn)
		if err != nil && err != io.EOF {
			return nTotal, fmt.Errorf("error reading input file: %v", err)
		}
		if n == 0 {
			break
		}
		bufIn = bufIn[:n] // buf2 is different size than buf

		bufOut, _ = interpolation.BandLimitedSincInterpolation(bufIn, ratioIO) // IMPORTANT:buf cut to n bytes
		out.Length += len(bufOut)
		if _, err = out.Writer.Write(bufOut); err != nil {
			return nTotal, fmt.Errorf("error writing output file: %v", err)
		}
		nTotal += n

		doOnceResampling.Do(func() {
			r.Println("[Transcoder] Transcoding data - sample of the first 4 bytes (hex)")
			onlyNFirst := 8
			r.Println(
				"[OLD]", fmt.Sprintf("% 2x", bufIn[:onlyNFirst]),
				"\n[NEW]", fmt.Sprintf("% 2x", bufOut[:onlyNFirst/2]),
			)
			r.Println("[Transcoder] Transcoding data - sample of the first 4 bytes (decimal)")
			r.Println(
				"[OLD]", fmt.Sprintf("%3d", bufIn[:onlyNFirst]),
				"\n[NEW]", fmt.Sprintf("%3d", bufOut[:onlyNFirst/2]),
			)
		})
	}
	return nTotal, nil
}

func fractional_delay_filter(in *sopro.In, out *sopro.Out, r *Resampler) (int, error) {
	sizeBuff := 1024 // max size, more than that would be too much
	if r.SizeBuffer > 0 {
		sizeBuff = r.SizeBuffer
	}
	nTotal := 0
	sampleRateIn := in.Config.(audioconfig.WavConfig).SampleRate
	sampleRateOut := out.Config.(audioconfig.WavConfig).SampleRate
	ratioIO := float64(sampleRateIn) / float64(sampleRateOut)

	sizeOutBuff := int(float64(sizeBuff) * ratioIO)
	bufIn := make([]byte, sizeBuff)     // input buffer
	bufOut := make([]byte, sizeOutBuff) // output buffer
	r.Println("ratio", ratioIO, "ratioInt", int(ratioIO), "sizeBuff", sizeBuff, "sizeOutBuff", sizeOutBuff)
	for {
		n, err := in.Reader.Read(bufIn)
		if err != nil && err != io.EOF {
			return nTotal, fmt.Errorf("error reading input file: %v", err)
		}
		if n == 0 {
			break
		}
		bufIn = bufIn[:n] // buf2 is different size than buf

		// bufOut, _ = filter.FractionalDelayFilter(bufIn, ratioIO) // IMPORTANT:buf cut to n bytes
		out.Length += len(bufOut)
		if _, err = out.Writer.Write(bufOut); err != nil {
			return nTotal, fmt.Errorf("error writing output file: %v", err)
		}
		nTotal += n

		doOnceResampling.Do(func() {
			r.Println("[Transcoder] Transcoding data - sample of the first 4 bytes (hex)")
			onlyNFirst := 8
			r.Println(
				"[OLD]", fmt.Sprintf("% 2x", bufIn[:onlyNFirst]),
				"\n[NEW]", fmt.Sprintf("% 2x", bufOut[:onlyNFirst/2]),
			)
			r.Println("[Transcoder] Transcoding data - sample of the first 4 bytes (decimal)")
			r.Println(
				"[OLD]", fmt.Sprintf("%3d", bufIn[:onlyNFirst]),
				"\n[NEW]", fmt.Sprintf("%3d", bufOut[:onlyNFirst/2]),
			)
		})
	}
	return nTotal, nil
}
