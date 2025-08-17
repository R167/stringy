# stringy

Unicode string encoding handling for Go. Provides helpers for working with UTF-8, UTF-16, and UTF-32 string representations, including code unit and code point counting.

## Why does this library exist?

Prior to Go 1.23, there was no built-in way to count UTF-16 code units in a string. This is important for interoperability with systems and protocols (such as JavaScript, Windows APIs, and some network protocols) that use UTF-16 as their string encoding. Go's built-in string type is UTF-8, and while it is easy to work with runes (Unicode code points), handling UTF-16 code units correctly is non-trivial, especially for characters outside the Basic Multilingual Plane (BMP) that require surrogate pairs.

This library provides types and helpers that "just work" for handling UTF-16, UTF-32, and UTF-8 string representations, including correct code unit and code point counting, and conversion between representations.

## How it works

- `Utf16` is a slice of `uint16` representing a UTF-16 encoded string. It implements methods for appending, writing, and converting to/from Go strings, as well as counting code units and code points.
- `Utf32` is a slice of `rune` (Unicode code points), with similar methods for manipulation and conversion.
- `Utf8` is a wrapper around Go's built-in string type, with helpers for code unit and code point counting.

All types implement the `String` interface, which provides methods for getting the string representation, code unit count, and code point count.

## Example usage

```go
import "github.com/R167/stringy"

var s stringy.Utf16
s.WriteString("hello 𝄞 world")
fmt.Println(s.String()) // prints: hello 𝄞 world
fmt.Println(s.CodeUnits()) // prints: number of UTF-16 code units
fmt.Println(s.CodePoints()) // prints: number of Unicode code points (runes)
```
