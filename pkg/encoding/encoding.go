package encoding

// Encoding enum:
// NOT_FILLED        = -1
// SPACE_LINEAR      = 0
// SPACE_LOGARITHMIC = 1
const (
	NOT_FILLED        = (iota - 1) // Not filled
	SPACE_LINEAR                   // PCM Linear
	SPACE_LOGARITHMIC              // PCM Logarithmic
)

// ENCODINGS is a map of encoding
var ENCODINGS = map[int]string{
	NOT_FILLED:        "Not filled",
	SPACE_LINEAR:      "Linear",
	SPACE_LOGARITHMIC: "Logarithmic",
}

// Encoding enum for variations based on existing encoding formats
const (
	SPACE_G711_1 = SPACE_LOGARITHMIC
	SPACE_ULAW   = SPACE_LOGARITHMIC
)
