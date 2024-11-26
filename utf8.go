package stringy

type Utf8 string

func (s Utf8) String() string {
	return string(s)
}
