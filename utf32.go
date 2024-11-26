package stringy

type Utf32 []rune

func (s Utf32) String() string {
	return string(s)
}
