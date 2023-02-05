package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var res = strings.Builder{}

	for key, char := range s {
		var next rune
		if len(s) > key+1 {
			next = rune(s[key+1])
		}

		if unicode.IsLetter(char) {
			if unicode.IsNumber(next) {
				res.WriteString(strings.Repeat(string(char), int(next-'0')))
			} else {
				res.WriteRune(char)
			}
		}

	}

	return res.String(), nil
}
