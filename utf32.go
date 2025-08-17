package stringy

import (
	"io"
)

// Utf32 is a slice of runes that represents a UTF-32 encoded string.
//
// *Utf32 implements the [io.StringWriter] and [io.Writer] interfaces.
type Utf32 []rune

var _ io.StringWriter = (*Utf32)(nil)
var _ io.Writer = (*Utf32)(nil)
var _ String = (*Utf32)(nil)

// NewUtf32 creates a new Utf32 instance from a string.
func NewUtf32(s string) Utf32 {
	return Utf32([]rune(s))
}

// AppendRunes appends a rune slice to the Utf32 slice.
func (s *Utf32) AppendRunes(r []rune) {
	*s = append(*s, r...)
}

// String returns the string representation of the Utf32 slice.
func (s Utf32) String() string {
	return string(s)
}

// Write writes a byte slice to the Utf32 slice.
// The byte slice is always assumed to be a valid UTF-8 encoded string.
// It always returns len(p) and a nil error.
func (s *Utf32) Write(p []byte) (n int, err error) {
	s.AppendRunes([]rune(string(p)))
	return len(p), nil
}

// WriteString writes a string to the Utf32 slice.
// It always returns len(str) and a nil error.
func (s *Utf32) WriteString(str string) (n int, err error) {
	s.AppendRunes([]rune(str))
	return len(str), nil
}

// CodeUnits returns the number of code units in the Utf32 slice.
func (s Utf32) CodeUnits() int {
	return len(s)
}

// CodePoints implements String.
func (s Utf32) CodePoints() int {
	return len(s)
}
