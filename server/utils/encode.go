package utils

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

const (
	baseURL = "http://localhost:8080"
	// Offset ID so it appears like a 7 digit alphanumeric
	// todo: look into mapping sequential ID to a non-sequential number
	offset       = 107253422234
	characterSet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

// We are going to take ID and multiply it to get a 7 digit base62 number
// Todo: use another method to avoid sequential urls
func EncodeURL(num uint64) string {
	return fmt.Sprintf("%s/%s", baseURL, ToBase62(num+offset))
}

// DecodeURL
func DecodeURL(shortURL string) (uint64, error) {
	num, err := FromBase62(shortURL)
	if err != nil {
		return 0, err
	}
	id := uint64(num) - 107253422234
	return id, err
}

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
