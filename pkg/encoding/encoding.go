package encoding

const (
	NOT_FILLED        = (iota - 1) // Not filled
	SPACE_LINEAR                   // PCM Linear
	SPACE_LOGARITHMIC              // PCM Logarithmic
)

var ENCODINGS = map[int]string{
	SPACE_LINEAR:      "Linear",
	SPACE_LOGARITHMIC: "Logarithmic",
}

// variations
const (
	SPACE_G711_1 = SPACE_LOGARITHMIC
	SPACE_ULAW   = SPACE_LOGARITHMIC
)
