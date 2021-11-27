package main

import (
	"crypto/hmac"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/RyuaNerin/go-krypto/lsh256"
)

func main() {
	keyHex := flag.String("key", "", "Secret key")
	flag.Parse()
	key, err := hex.DecodeString(*keyHex)
	if err != nil {
		panic(err)
	}
	h := hmac.New(lsh256.New, key)
	if _, err = io.Copy(h, os.Stdin); err != nil {
		panic(err)
	}
	fmt.Println("MAC:", hex.EncodeToString(h.Sum(nil)))
}
