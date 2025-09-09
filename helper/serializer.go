package helper

import (
	"math/rand"
	"strings"
)

const (
	CharacterRunes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	SerialLength   = 12
)

func GenerateSerialFromString(input string) string {
	serialPrefix := ""
	if len(input) >= 6 {
		serialPrefix = input[0:6]
	} else {
		serialPrefix = input
	}
	randomChars := randStringRunes(SerialLength - len(serialPrefix))

	return strings.ToUpper(serialPrefix + randomChars)
}

func randStringRunes(n int) string {
	var runes = []rune(CharacterRunes)

	b := make([]rune, n)
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}
	return string(b)
}
