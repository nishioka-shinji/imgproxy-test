package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	key  := os.Getenv("IMGPROXY_KEY")
	salt := os.Getenv("IMGPROXY_SALT")

	var keyBin, saltBin []byte

	if keyBin, err = hex.DecodeString(key); err != nil {
		log.Fatal(err)
	}

	if saltBin, err = hex.DecodeString(salt); err != nil {
		log.Fatal(err)
	}

	path := "rs:fill:300:400/plain/https://storage.googleapis.com/studio-design-asset-files/projects/7kadpxLza3/s-1616x792_v-fms_webp_b5774a47-5fde-4867-987a-3a2bb4664066.webp"

	mac := hmac.New(sha256.New, keyBin)
	mac.Write(saltBin)
	mac.Write([]byte(path))
	signature := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))

	fmt.Printf("/%s%s", signature, path)
}
