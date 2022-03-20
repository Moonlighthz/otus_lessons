package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var strBuilder strings.Builder
	runeStr := []rune(str)
	validate := false
	var slash bool

	for i, char := range runeStr {
		if (i == 0 && unicode.IsDigit(char)) || (validate && unicode.IsDigit(char)) {
			return "", ErrInvalidString
		}

		if char == '\\' && !slash {
			slash = true
			continue
		}

		if slash {
			strBuilder.WriteString(string(char))
			slash = false
			continue
		}

		if unicode.IsDigit(char) {
			number := int(char - '0')
			validate = true

			if number == 0 {
				newStr := strBuilder.String()[:len(strBuilder.String())-1]
				strBuilder.Reset()
				strBuilder.WriteString(newStr)
				continue
			}

			strBuilder.WriteString(strings.Repeat(string(runeStr[i-1]), number-1))
			continue
		}

		strBuilder.WriteString(string(char))
		validate = false
	}

	return strBuilder.String(), nil
}
