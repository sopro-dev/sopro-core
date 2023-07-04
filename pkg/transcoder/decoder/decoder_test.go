package decoder

import (
	"reflect"
	"testing"
)

func TestDecodeUlaw2Lpcm(t *testing.T) {
	tests := []struct {
		logPcm  []byte
		lPcm    []byte
		wantErr bool
	}{
		{
			logPcm:  []byte{0x80, 0x80, 0x80, 0x80},
			lPcm:    []byte{0x7C, 0x7D, 0x7C, 0x7D, 0x7C, 0x7D, 0x7C, 0x7D},
			wantErr: false,
		},
		{
			logPcm:  []byte{0xFF, 0xFF, 0xFF, 0xFF},
			lPcm:    []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			wantErr: false,
		},
		{
			logPcm:  []byte{},
			lPcm:    []byte{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		lpcm, err := DecodeULawToPCM(tt.logPcm)
		if (err != nil) != tt.wantErr {
			t.Errorf("unexpected error: %v", err)
			continue
		}
		if !reflect.DeepEqual(lpcm, tt.lPcm) {
			t.Errorf("DecodeUlaw2Lpcm(%v) = %v, want %v", tt.logPcm, lpcm, tt.lPcm)
		}
	}
}
