package utils

// AudioSilenceDefault creates a WAV silence buffer with 1000 samples
func AudioSilenceDefault() []byte {
	// Create 1000 samples of silence (2000 bytes for 16-bit)
	empty := make([]byte, 1000)

	headers := GenerateWavHeadersWithConfig(&WavHeader{
		Length:     uint32(len(empty) + 44), // 44 is WAV header size
		WaveFormat: WAVE_FORMAT_PCM,
		Channels:   1,
		SampleRate: 8000,
		BitDepth:   16,
		Verbose:    false,
	})

	// Combine headers and empty samples
	result := make([]byte, len(headers)+len(empty))
	copy(result, headers)
	copy(result[len(headers):], empty)
	return result
}

// AudioSilence creates a WAV silence buffer of specified duration in milliseconds
func AudioSilence(durationMs int) []byte {
	// Calculate number of samples needed (8000Hz * (durationMs/1000))
	numSamples := (8000 * durationMs) / 1000

	// Pre-allocate the exact buffer size needed (2 bytes per sample for 16-bit)
	empty := make([]byte, numSamples*2)

	headers := GenerateWavHeadersWithConfig(&WavHeader{
		Length:     uint32(len(empty) + 44), // 44 is WAV header size
		WaveFormat: WAVE_FORMAT_PCM,
		Channels:   1,
		SampleRate: 8000,
		BitDepth:   16,
		Verbose:    false, // Set to false unless debugging
	})

	// Pre-allocate the final buffer to avoid reallocation
	result := make([]byte, len(headers)+len(empty))
	copy(result, headers)
	copy(result[len(headers):], empty)
	return result
}
