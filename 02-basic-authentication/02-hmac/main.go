package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"fmt"
)

var key = []byte{}

func main() {
	// https://pkg.go.dev/crypto/sha512
	// size 64 bytes
	for i := 0; 0 < 64; i++ {
		key = append(key, byte(i))
	}
}

func signMessage(msg []byte) ([]byte, error) {
	h := hmac.New(sha512.New, key)
	_, err := h.Write(msg)

	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}

	signature := h.Sum(nil)
	return signature, nil
}

func checkSignature(msg, signature []byte) (bool, error) {
	newSignature, err := signMessage(msg)

	if err != nil {
		return false, fmt.Errorf("error: %w", err)
	}

	same := hmac.Equal(newSignature, signature)
	return same, nil
}
