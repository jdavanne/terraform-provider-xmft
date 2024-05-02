package tools

import (
	"math/rand"
	"time"
)

const (
	shortIDLength = 20
	charset       = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func _stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// GenerateID generate a random ID string (20 characters)
func GenerateID(n int) string {
	return _stringWithCharset(n, charset)
}

// GenerateShortID generate a 128bit random ID string (22 characters)
func GenerateShortID() string {
	return _stringWithCharset(shortIDLength, charset)
}

// GenerateShortID generate a 128bit random ID string (n characters)
func GenerateShortIDn(n int) string {
	return _stringWithCharset(n, charset)
}
