package stringy

type String interface {
	// String returns the string representation of the string.
	String() string
	// CodeUnits returns the number of code units in the string.
	CodeUnits() int
	// CodePoints returns the number of code points in the string.
	// This is the number of runes in the string and what you would use for character-based operations.
	CodePoints() int
}
