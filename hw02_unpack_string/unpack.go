package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var res string
	for key := 0; key < len(s); key++ {
		curr := rune(s[key])

		if !unicode.IsLetter(curr) && !unicode.IsPunct(curr) {
			return "", ErrInvalidString
		}

		if unicode.IsPunct(curr) {
			if len(s) <= key+1 {
				return "", ErrInvalidString
			}
			curr = rune(s[key+1])
			key++
		}

		if len(s) <= key+1 {
			res += string(curr)
			continue
		}
		next := rune(s[key+1])

		if unicode.IsNumber(next) {
			res += strings.Repeat(string(curr), int(next-'0'))
			key++
		} else {
			res += string(curr)
		}
	}

	return res, nil
}
