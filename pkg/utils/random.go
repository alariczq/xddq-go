package utils

import (
	"crypto/rand"
)

func RandomString(n int) string {
	chars := "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charsLen := len(chars)
	result := make([]byte, n)
	rand.Read(result)
	for i, b := range result {
		result[i] = chars[b%byte(charsLen)]
	}
	return string(result)
}
