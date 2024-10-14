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
	// sourceUrl := "s3://imgproxy-test-nishioka/anime.gif"
	sourceUrl := "s3://imgproxy-test-nishioka/jpeg1.jpg"
	// sourceUrl := "s3://imgproxy-test-nishioka/png4.png"
	encodedUrl := base64.URLEncoding.EncodeToString([]byte(sourceUrl))
	path := fmt.Sprintf("/%s/%s", processingOptions, encodedUrl)

	mac := hmac.New(sha256.New, keyBin)
	mac.Write(saltBin)
	mac.Write([]byte(path))
	signature := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))

	fmt.Printf("http://localhost:10000/%s%s", signature, path)
}
