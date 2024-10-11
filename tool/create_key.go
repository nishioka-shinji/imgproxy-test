package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func generateHex(size int) (string, error) {
	bytes := make([]byte, size)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func main() {
	key, err := generateHex(32)
	if err != nil {
		fmt.Println("Error generating key:", err)
		return
	}

	salt, err := generateHex(32)
	if err != nil {
		fmt.Println("Error generating salt:", err)
		return
	}

	fmt.Println("Key:")
	fmt.Println(key)
	fmt.Println("Salt:")
	fmt.Println(salt)
}
