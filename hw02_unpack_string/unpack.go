package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	data := []rune(input)
	result := ""

	for key, char := range data {
		if key == 0 && unicode.IsDigit(char) {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(char) && unicode.IsDigit(data[key-1]) {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(char) {
			count := int(char - '0')

			if count == 0 {
				result = result[:len(result)-1]
				continue
			}

			result += strings.Repeat(string(data[key-1]), count-1)
			continue
		}

		result += string(char)
	}

	return result, nil
}
