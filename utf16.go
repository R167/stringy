package stringy

import (
	"io"
	"unicode/utf16"
)

// constants for utf16 encoding
const (
	maxCodeUnit = 0xFFFF

	surr1 = 0xd800
	surr2 = 0xdc00
	surr3 = 0xe000
)

// Utf16 is a slice of uint16 values that represents a UTF-16 encoded string.
//
// *Utf16 implements the [io.StringWriter] and [io.Writer] interfaces.
type Utf16 []uint16

var _ io.StringWriter = (*Utf16)(nil)
var _ io.Writer = (*Utf16)(nil)
var _ String = (*Utf16)(nil)

func NewUtf16(s string) Utf16 {
	return Utf16(utf16.Encode([]rune(s)))
}

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
	s.AppendRunes([]rune(string(p)))
	return len(p), nil
}

// WriteString writes a string to the Utf16 slice.
// It always returns len(str) and a nil error.
func (s *Utf16) WriteString(str string) (n int, err error) {
	s.AppendRunes([]rune(str))
	return len(str), nil
}

// CodeUnits returns the number of code units in the Utf16 slice.
func (s Utf16) CodeUnits() int {
	return len(s)
}

// CodePoints returns the number of code points in the Utf16 slice.
func (s Utf16) CodePoints() int {
	length := 0
	for i := 0; i < len(s); i++ {
		length++
		r := s[i]
		// if we're in the middle of a valid surrogate pair, skip
		if surr1 <= r && r < surr2 && i+1 < len(s) && surr2 <= s[i+1] && s[i+1] < surr3 {
			i++
		}
	}
	return length
}

// Utf16CodeUnits returns the number of code units in the UTF-16 encoding of the string.
func Utf16CodeUnits(str string) int {
	n := 0
	for _, r := range str {
		n++
		if r > maxCodeUnit {
			n++
		}
	}
	return n
}
