package helpers

import (
	"crypto/rand"
	"math/big"
)

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		b[i] = charset[index.Int64()]
	}
	return string(b)
}

func RandomNumber(length int) string {
	return stringWithCharset(length, "1234567890")
}

func RandomAlphabet(length int) string {
	return stringWithCharset(length, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
}

func RandomAlphaNum(length int) string {
	return stringWithCharset(length, "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
}

func RandomLowerAlphaNum(length int) string {
	return stringWithCharset(length, "abcdefghijklmnopqrstuvwxyz0123456789")
}

// without 0 to reduce O and 0 confusion
func RandomAlphaNumID(length int) string {
	return stringWithCharset(length, "ABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")
}
