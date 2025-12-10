package helper

import (
	"math/rand"
	"strings"
)

const (
	CharacterRunes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	SerialLength   = 12
)

func GenerateSerialFromString(serialPrefix, input string) (serial string) {
	stringPrefix := ""
	if len(input) >= 6 {
		stringPrefix = input[0:6]
	} else {
		stringPrefix = input
	}

	randomChars := randStringRunes(SerialLength - len(stringPrefix))
	serial = strings.ToUpper(stringPrefix + randomChars)

	if serialPrefix != "" {
		serial = serialPrefix + "-" + serial
	}

	return serial
}

func randStringRunes(n int) string {
	var runes = []rune(CharacterRunes)

	b := make([]rune, n)
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}
	return string(b)
}
