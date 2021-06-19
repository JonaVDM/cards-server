package utils

import (
	"math/rand"
	"time"
)

func RandomString(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")

	rand.Seed(time.Now().Unix())

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
