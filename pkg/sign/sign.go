package sign

// Sign is a type that represents the sign of the byte (signed or unsigned)
// It is used in the decoder to determine if the byte is signed or unsigned
const (
	NOT_FILLED = (iota - 1) // Not filled
	UNSIGNED                // Unsigned
	SIGNED                  // Signed
)
