package utils

import (
	"time"
	"math/rand"
	"strings"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func CaseInsensitiveContains(a, b string) bool {
	return strings.Contains(strings.ToLower(a), strings.ToLower(b))
}