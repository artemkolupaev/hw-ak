package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var res strings.Builder
	var lastRune rune
	var esc bool

	for i, curRune := range s {
		if unicode.IsDigit(curRune) && i == 0 {
			return "", ErrInvalidString
		}
		if unicode.IsDigit(curRune) && unicode.IsDigit(lastRune) {
			return "", ErrInvalidString
		}
		if unicode.IsLetter(curRune) || unicode.IsSpace(curRune) || unicode.IsSymbol(curRune) {
			res.WriteRune(curRune)
		}
		if unicode.IsDigit(curRune) && esc {
			res.WriteRune(curRune)
		}
		if (unicode.IsDigit(curRune) && !esc) && (unicode.IsLetter(lastRune) || unicode.IsSpace(lastRune) || unicode.IsSymbol(lastRune)) {
			runeInt := int(curRune - '0')
			if runeInt == 0 {
				s = strings.TrimSuffix(res.String(), string(lastRune))
				res.Reset()
				res.WriteString(s)
			} else {
				res.WriteString(strings.Repeat(string(lastRune), runeInt-1))
			}
		}
		if string(curRune) == "\\" {
			esc = true
		} else {
			esc = false
		}
		lastRune = curRune
	}

	return res.String(), nil
}
