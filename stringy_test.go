package stringy

import (
	"testing"
	"unicode/utf16"
	"unicode/utf8"
)

func TestUtf16_AppendRunesAndString(t *testing.T) {
	runes := []rune("hello 𝄞 world")
	var s Utf16
	s.AppendRunes(runes)
	if got := s.String(); got != string(runes) {
		t.Errorf("Utf16.String() = %q, want %q", got, string(runes))
	}
}

func TestUtf16_WriteAndWriteString(t *testing.T) {
	var s Utf16
	str := "foo 𝄞 bar"
	n, err := s.Write([]byte(str))
	if err != nil || n != len(str) {
		t.Errorf("Utf16.Write() = (%d, %v), want (%d, nil)", n, err, len(str))
	}
	n, err = s.WriteString(str)
	if err != nil || n != len(str) {
		t.Errorf("Utf16.WriteString() = (%d, %v), want (%d, nil)", n, err, len(str))
	}
	if got := s.String(); got != str+str {
		t.Errorf("Utf16.String() = %q, want %q", got, str+str)
	}
}

func TestUtf16_CodeUnitsAndCodePoints(t *testing.T) {
	str := "a𝄞b"
	utf16s := utf16.Encode([]rune(str))
	s := Utf16(utf16s)
	if got := s.CodeUnits(); got != len(utf16s) {
		t.Errorf("Utf16.CodeUnits() = %d, want %d", got, len(utf16s))
	}
	if got := s.CodePoints(); got != utf8.RuneCountInString(str) {
		t.Errorf("Utf16.CodePoints() = %d, want %d", got, utf8.RuneCountInString(str))
	}
}

func TestUtf16CodeUnitsFunc(t *testing.T) {
	str := "a𝄞b"
	want := len(utf16.Encode([]rune(str)))
	if got := Utf16CodeUnits(str); got != want {
		t.Errorf("Utf16CodeUnits(%q) = %d, want %d", str, got, want)
	}
}

func TestUtf32_AppendRunesAndString(t *testing.T) {
	runes := []rune("hello 𝄞 world")
	var s Utf32
	s.AppendRunes(runes)
	if got := s.String(); got != string(runes) {
		t.Errorf("Utf32.String() = %q, want %q", got, string(runes))
	}
}

func TestUtf32_WriteAndWriteString(t *testing.T) {
	var s Utf32
	str := "foo 𝄞 bar"
	n, err := s.Write([]byte(str))
	if err != nil || n != len(str) {
		t.Errorf("Utf32.Write() = (%d, %v), want (%d, nil)", n, err, len(str))
	}
	n, err = s.WriteString(str)
	if err != nil || n != len(str) {
		t.Errorf("Utf32.WriteString() = (%d, %v), want (%d, nil)", n, err, len(str))
	}
	if got := s.String(); got != str+str {
		t.Errorf("Utf32.String() = %q, want %q", got, str+str)
	}
}

func TestUtf32_CodeUnitsAndCodePoints(t *testing.T) {
	str := "a𝄞b"
	s := Utf32([]rune(str))
	if got := s.CodeUnits(); got != len([]rune(str)) {
		t.Errorf("Utf32.CodeUnits() = %d, want %d", got, len([]rune(str)))
	}
	if got := s.CodePoints(); got != len([]rune(str)) {
		t.Errorf("Utf32.CodePoints() = %d, want %d", got, len([]rune(str)))
	}
}

func TestUtf8_CodeUnitsAndCodePoints(t *testing.T) {
	str := Utf8("a𝄞b")
	if got := str.CodeUnits(); got != len("a𝄞b") {
		t.Errorf("Utf8.CodeUnits() = %d, want %d", got, len("a𝄞b"))
	}
	if got := str.CodePoints(); got != utf8.RuneCountInString("a𝄞b") {
		t.Errorf("Utf8.CodePoints() = %d, want %d", got, utf8.RuneCountInString("a𝄞b"))
	}
}
