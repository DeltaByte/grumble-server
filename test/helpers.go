package test

import (
	"math/rand"
	"strings"
	"time"
)

const (
	CharsetLower   = "abcdefghijklmnopqrstuvwxyz"
	CharsetUpper   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	CharsetNumbers = "0123456789"
	CharsetDefault = CharsetLower + CharsetUpper + CharsetNumbers
)

func RandomString(length int, charSet string) string {
	rand.Seed(time.Now().Unix())

	chars := []rune(charSet)
	var output strings.Builder

	for i := 0; i < length; i++ {
			random := rand.Intn(len(chars))
			randomChar := chars[random]
			output.WriteRune(randomChar)
	}

	return output.String()
}