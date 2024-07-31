package utils

import (
	"math/rand"
	"time"
)

func RandomString(n int, includeLowerCase bool, includeUpperCase bool, includeNumbers bool) string {
	lowerCaseLetters := []rune("abcdefghijklmnopqrstuvwxyz")
	upperCaseLetters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	numbers := []rune("0123456789")

	var letters []rune
	if includeLowerCase {
		letters = append(letters, lowerCaseLetters...)
	}
	if includeUpperCase {
		letters = append(letters, upperCaseLetters...)
	}
	if includeNumbers {
		letters = append(letters, numbers...)
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}
	return string(b)
}
