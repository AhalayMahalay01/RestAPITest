package random

import (
	"math/rand"
	"time"
)

func NewRandomString(size int) string { // реалізація рандому, можна використати інший пакет crypto/rand для більшої безпеки
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	chars := []rune("ABCDEFGHIJKLMNPQRSTUVWXYZ" + // створюється слайс з символів для генерації
		"abcdefghijklmnpqrstyvwxyz" +
		"0123456789")

	b := make([]rune, size) //слайс буферу для символів
	for i := range b {
		b[i] = chars[rnd.Intn(len(chars))]
	}

	return string(b)
}
