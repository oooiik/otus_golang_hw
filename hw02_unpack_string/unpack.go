package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	st := ""
	if 0 >= len(s) {
		return st, nil
	}
	var lastV int32
	for _, valS := range s {
		if unicode.IsNumber(valS) {

			num, err := strconv.Atoi(string(valS))
			if err != nil {
				return "", ErrInvalidString
			}

			if lastV == 0 {
				return "", ErrInvalidString
			}

			if num == 0 {
				st = st[:len(st)-1]
			} else {
				st = st + strings.Repeat(string(lastV), num-1)
			}
			lastV = 0
		} else {
			lastV = valS
			st += string(lastV)
		}
	}

	return st, nil
}
