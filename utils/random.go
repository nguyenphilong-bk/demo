package utils

import (
	"crypto/rand"
	"encoding/base64"
)

func RandomStringCrypto(length int) (string, error) {
    bytes := make([]byte, length)
    _, err := rand.Read(bytes)
    if err != nil {
        return "", err
    }
    return base64.URLEncoding.EncodeToString(bytes)[:length], nil
}