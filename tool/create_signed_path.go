package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

func main() {
	key  := os.Getenv("IMGPROXY_KEY")
	salt := os.Getenv("IMGPROXY_SALT")

	var keyBin, saltBin []byte
	var err error

	if keyBin, err = hex.DecodeString(key); err != nil {
		log.Fatal(err)
	}

	if saltBin, err = hex.DecodeString(salt); err != nil {
		log.Fatal(err)
	}

	processingOptions := "thumb"
	sourceUrl := "https://storage.googleapis.com/studio-design-asset-files/projects/7kadpxLza3/s-1616x792_v-fms_webp_b5774a47-5fde-4867-987a-3a2bb4664066.webp"
	path := fmt.Sprintf("/%s/plain/%s", processingOptions, sourceUrl)

	mac := hmac.New(sha256.New, keyBin)
	mac.Write(saltBin)
	mac.Write([]byte(path))
	signature := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))

	fmt.Printf("http://localhost:10000/%s%s", signature, path)
}
