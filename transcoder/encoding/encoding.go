package encoding

const (
	NOT_FILLED   = iota // Not filled
	SPACE_LINEAR        // PCM Linear
	SPACE_G711_1        // G711.1
)

// variations
const (
	SPACE_ULAW = SPACE_G711_1
)
