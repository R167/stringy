package stringy

import "unicode/utf8"

// Utf8 is a string type that supports UTF-8 encoding.
//
// It is implemented as a string since Go's built-in string type is already UTF-8 encoded.
// Because of this, Utf8 does not implement the same interfaces for updating as the other types.
type Utf8 string

var _ String = (*Utf8)(nil)

// NewUtf8 creates a new Utf8 wrapper from a string.
func NewUtf8(s string) Utf8 {
	return Utf8(s)
}

// String returns the string underlying string from the Utf8 wrapper.
func (s Utf8) String() string {
	return string(s)
}

// CodeUnits returns the number of code units in the Utf8 string.
// For utf8, this is the number of bytes.
func (s Utf8) CodeUnits() int {
	return len(s)
}

// CodePoints returns the unicode charachter length of this string.
func (s Utf8) CodePoints() int {
	return utf8.RuneCountInString(string(s))
}
