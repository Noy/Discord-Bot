package utils

import (
	"time"
	"math/rand"
	"strings"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789!%&*")

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}

func CaseInsensitiveContains(a, b string) bool {
	return strings.Contains(strings.ToLower(a), strings.ToLower(b))
}

func Schedule(task func(), delay time.Duration) chan bool {
	stop := make(chan bool)
	go func() {
		for {
			task()
			select {
			case <-time.After(delay):
			case <-stop:
				return
			}
		}
	}()
	return stop
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {return true}
	}
	return false
}