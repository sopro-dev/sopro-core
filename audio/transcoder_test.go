package audio

import "testing"

func TestValidateAudioInfo(t *testing.T) {
	// Test case 1: Valid audio info
	info := AudioInfo{BitDepth: 16, Channels: 2, SampleRate: 44100}
	err := validateAudioInfo(info)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test case 2: Invalid bit depth
	info = AudioInfo{BitDepth: 5, Channels: 2, SampleRate: 44100}
	err = validateAudioInfo(info)
	if err == nil {
		t.Error("Expected an error for invalid bit depth")
	}
	if err != errInvalidBitDepth {
		t.Errorf("Expected error %v, got %v", errInvalidBitDepth, err)
	}

	// Test case 3: Invalid number of channels
	info = AudioInfo{BitDepth: 16, Channels: 0, SampleRate: 44100}
	err = validateAudioInfo(info)
	if err == nil {
		t.Error("Expected an error for invalid number of channels")
	}
	if err != errInvalidNumChannels {
		t.Errorf("Expected error %v, got %v", errInvalidNumChannels, err)
	}

	// Test case 4: Invalid sample rate
	info = AudioInfo{BitDepth: 16, Channels: 2, SampleRate: 0}
	err = validateAudioInfo(info)
	if err == nil {
		t.Error("Expected an error for invalid sample rate")
	}
	if err != errInvalidSampleRate {
		t.Errorf("Expected error %v, got %v", errInvalidSampleRate, err)
	}
}
