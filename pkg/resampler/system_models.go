package resampler

import (
	"sync"

	"github.com/pablodz/sopro/pkg/sopro"
)

// Resampler struct to hold resampler information only
const (
	NOT_FILLED                 = (iota - 1) // Not filled
	LINEAR_INTERPOLATION                    // Linear interpolation
	POLYNOMIAL_INTERPOLATION                // Polynomial interpolation
	BAND_LIMITED_INTERPOLATION              // Band limited interpolation
	FRACTIONAL_DELAY_FILTER
)

// RESAMPLER_METHODS is a map of resampler methods
var RESAMPLER_METHODS = map[int]string{
	NOT_FILLED:               "NOT_FILLED",
	LINEAR_INTERPOLATION:     "LINEAR_INTERPOLATION",
	POLYNOMIAL_INTERPOLATION: "POLYNOMIAL_INTERPOLATION",
}

type Resampler struct {
	MethodR                int               // MethodR is the method of resampling if needed. Default is 0 (no resampling)
	MethodRAdvancedConfigs interface{}       // MethodAdvancedConfigs is the method of advanced configs if needed
	SizeBuffer             int               // SizeBuffer is the size of the buffer used in the resampling process.Default is 1024
	InConfigs              sopro.AudioConfig // the source configuration
	OutConfigs             sopro.AudioConfig // the target configuration
	Verbose                bool              // if true, the resampling process will be verbose
}

var doOnceResampling sync.Once // doOnceResampling is used to initialize the resampling methods only once if verbose is true
