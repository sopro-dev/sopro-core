package transcoder

import (
	"sync"

	"github.com/pablodz/sopro/pkg/sopro"
)

// Bit enum to convert standard bit sizes
const (
	BIT_8    = 8
	BIT_16   = 16
	BIT_24   = 24
	BIT_32   = 32
	BIT_64   = 64
	BIT_128  = 128
	BIT_256  = 256
	BIT_512  = 512
	BIT_1024 = 1024
)

// Transcoder is the main struct for the transcoding process
type Transcoder struct {
	MethodT                int               // MethodT is the method of transcoding if needed. Default is 0 (no transcoding)
	MethodTAdvancedConfigs interface{}       // MethodAdvancedConfigs is the method of advanced configs if needed
	SizeBuffer             int               // SizeBuffer is the size of the buffer used in the transcoding process.Default is 1024
	InConfigs              sopro.AudioConfig // the source configuration
	OutConfigs             sopro.AudioConfig // the target configuration
	BitDepth               int               // the bit depth (e.g. 8, 16, 24) Needs to be equal for source and target
	Verbose                bool              // if true, the transcoding process will be verbose
}

var doOnceTranscoding sync.Once // doOnceTranscoding is used to initialize the transcoding methods only once if verbose is true
