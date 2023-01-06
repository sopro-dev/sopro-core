package method

const (
	NOT_FILLED            = iota // Not filled
	BIT_SHIFT                    // Bit shift
	BIT_TABLE                    // Bit table, use a slice to store the values
	BIT_ADVANCED_FUNCTION        // Advanced function, use a function to calculate and return the values
)
