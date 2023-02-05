package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

type char struct {
	IsPunct bool
	Value   rune
}

func Unpack(s string) (string, error) {
	// generate map
	var mapChar []char

	for key := 0; key < len(s); key++ {
		curr := rune(s[key])

		var next rune
		if len(s) > key+1 {
			next = rune(s[key+1])
		}
		if unicode.IsPunct(curr) {
			mapChar = append(mapChar, char{true, next})
			key++
			continue
		}
		mapChar = append(mapChar, char{false, curr})
	}

	// generate string
	var res = strings.Builder{}
	for keyChar := 0; keyChar < len(mapChar); keyChar++ {
		currChar := mapChar[keyChar]

		var nextChar char
		if len(mapChar) > keyChar+1 {
			nextChar = mapChar[keyChar+1]
		} else {
			res.WriteRune(currChar.Value)
			break
		}

		if unicode.IsNumber(currChar.Value) && !currChar.IsPunct {
			return "", ErrInvalidString
		}

		if unicode.IsNumber(nextChar.Value) && !nextChar.IsPunct {
			res.WriteString(strings.Repeat(string(currChar.Value), int(nextChar.Value-'0')))
			keyChar++
		} else {
			res.WriteRune(currChar.Value)
		}
	}

	return res.String(), nil
}
