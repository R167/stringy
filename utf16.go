package stringy

import "unicode/utf16"

const (
	maxCodeUnit = 0xFFFF
)

type Utf16 []uint16

// AppendRunes appends a rune slice to the Utf16 slice.
func (s *Utf16) AppendRunes(r []rune) {
	*s = append(*s, utf16.Encode(r)...)
}

// String returns the string representation of the Utf16 slice.
func (s Utf16) String() string {
	return string(utf16.Decode(s))
}

// Write writes a byte slice to the Utf16 slice.
// The byte slice is always assumed to be a valid UTF-8 encoded string.
// It always returns len(p) and a nil error.
func (s *Utf16) Write(p []byte) (n int, err error) {
	*s = utf16.Encode([]rune(string(p)))
	return len(p), nil
}

// WriteString writes a string to the Utf16 slice.
// It always returns len(str) and a nil error.
func (s *Utf16) WriteString(str string) (n int, err error) {
	*s = utf16.Encode([]rune(str))
	return len(str), nil
}

// CodeUnits returns the number of code units in the Utf16 slice.
func (s Utf16) CodeUnits() int {
	return len(s)
}

// Utf16CodeUnits returns the number of code units in the UTF-16 encoding of the string.
func Utf16CodeUnits(str string) int {
	n := len(str)
	for _, r := range str {
		if r > maxCodeUnit {
			n++
		}
	}
	return n
}
