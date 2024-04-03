package helpers

import (
	"math/rand"
	"strings"
)

func GenerateShortURL(urlLen int) string {
	var alphabet = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	var alphabetLen = len(alphabet)

	//  62^7 options should be plenty
	sb := strings.Builder{}

	for range urlLen {
		sb.WriteByte(alphabet[rand.Intn(alphabetLen)])
	}

	return sb.String()
}
