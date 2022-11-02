package util

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"math/rand"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func GenerateKey(length int) (apiKey string, hash string, error error) {
	key := getApiKey(length)
	h := sha256.New()
	_, err := io.WriteString(h, key)
	if err != nil {
		return "", "", err
	}

	return key, hex.EncodeToString(h.Sum(nil)), nil
}

func getApiKey(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
