package method

const (
	NOT_FILLED            = (iota - 1) // Not filled
	BIT_SHIFT                          // Bit shift
	BIT_LOOKUP_TABLE                   // Bit table, use a slice to store the values
	BIT_ADVANCED_FUNCTION              // Advanced function, use a function to calculate and return the values
)

var METHODS = map[int]string{
	NOT_FILLED:            "NOT_FILLED",
	BIT_SHIFT:             "BIT_SHIFT",
	BIT_LOOKUP_TABLE:      "BIT_LOOKUP_TABLE",
	BIT_ADVANCED_FUNCTION: "BIT_ADVANCED_FUNCTION",
}
