package utils

import (
	"errors"
	"math"
	"strings"
)

const (
	characterSet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

func ToBase62(num uint64) string {
	encoded := ""
	for num > 0 {
		r := num % 62

		num /= 62
		encoded = string(characterSet[r]) + encoded
	}
	return encoded
}

func FromBase62(encoded string) (uint64, error) {
	var val uint64
	for index, char := range encoded {
		pow := len(encoded) - (index + 1)
		pos := strings.IndexRune(characterSet, char)
		if pos == -1 {
			return 0, errors.New("invalid character: " + string(char))
		}

		val += uint64(pos) * uint64(math.Pow(float64(62), float64(pow)))
	}

	return val, nil
}
