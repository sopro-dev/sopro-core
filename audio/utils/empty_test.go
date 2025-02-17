package utils

import (
	"bytes"
	"testing"
)

func TestAudioSilence(t *testing.T) {
	tests := []struct {
		name       string
		durationMs int
		wantLen    int // Expected total length including 44 byte header
	}{
		{
			name:       "1 second silence",
			durationMs: 1000,
			wantLen:    16044, // 44 + (8000 * 2) bytes
		},
		{
			name:       "500ms silence",
			durationMs: 500,
			wantLen:    8044, // 44 + (4000 * 2) bytes
		},
		{
			name:       "100ms silence",
			durationMs: 100,
			wantLen:    1644, // 44 + (800 * 2) bytes
		},
		{
			name:       "zero duration",
			durationMs: 0,
			wantLen:    44, // Just header
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AudioSilence(tt.durationMs)

			// Check total length
			if len(got) != tt.wantLen {
				t.Errorf("AudioSilence() length = %v, want %v", len(got), tt.wantLen)
			}

			// Verify WAV header basics
			if !bytes.Equal(got[0:4], []byte("RIFF")) {
				t.Error("AudioSilence() invalid RIFF header")
			}
			if !bytes.Equal(got[8:12], []byte("WAVE")) {
				t.Error("AudioSilence() invalid WAVE format")
			}

			// Verify silence (all bytes should be 0 after header)
			for i := 44; i < len(got); i++ {
				if got[i] != 0 {
					t.Errorf("AudioSilence() non-zero byte at position %d", i)
					break
				}
			}
		})
	}
}

func TestAudioSilenceDefault(t *testing.T) {
	got := AudioSilenceDefault()
	wantLen := 1044 // 44 + (1000 * 2) bytes

	// Check total length
	if len(got) != wantLen {
		t.Errorf("AudioSilenceDefault() length = %v, want %v", len(got), wantLen)
	}

	// Verify WAV header basics
	if !bytes.Equal(got[0:4], []byte("RIFF")) {
		t.Error("AudioSilenceDefault() invalid RIFF header")
	}
	if !bytes.Equal(got[8:12], []byte("WAVE")) {
		t.Error("AudioSilenceDefault() invalid WAVE format")
	}

	// Verify silence (all bytes should be 0 after header)
	for i := 44; i < len(got); i++ {
		if got[i] != 0 {
			t.Errorf("AudioSilenceDefault() non-zero byte at position %d", i)
			break
		}
	}
}
