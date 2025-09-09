package helper

import (
	"math/rand"
	"strings"
)

const (
	CharacterRunes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func GenerateSerialFromString(input string) string {
	serialPrefix := input[0:6]
	randomChars := randStringRunes(6)

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
