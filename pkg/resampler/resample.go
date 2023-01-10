package resampler

const (
	NOT_FILLED               = (iota - 1) // Not filled
	LINEAR_INTERPOLATION                  // Linear interpolation
	POLYNOMIAL_INTERPOLATION              // Polynomial interpolation
)

var RESAMPLER_METHODS = map[int]string{
	NOT_FILLED:               "NOT_FILLED",
	LINEAR_INTERPOLATION:     "LINEAR_INTERPOLATION",
	POLYNOMIAL_INTERPOLATION: "POLYNOMIAL_INTERPOLATION",
}
