package main

import (
	"crypto/hmac"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/RyuaNerin/go-krypto/lsh256"
)

func main() {
	keyHex := flag.String("k", "", "Secret key.")
	target := flag.String("f", "", "Target file. ('-' for STDIN)")
	flag.Parse()
	if len(os.Args) < 2 {
		fmt.Println("LSH-HMAC - ALBANESE Lab (c) 2020-2021\n")
		fmt.Println("Usage of", os.Args[0]+":")
		fmt.Printf("%s [-k <secret>] -f <file.ext>\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}
	key, err := hex.DecodeString(*keyHex)
	if err != nil {
		panic(err)
	}
	var file io.Reader
	if *target == "-" {
		file = os.Stdin
	} else {
		file, err = os.Open(*target)
		if err != nil {
			log.Fatalf("failed opening file: %s", err)
		}
	}
	h := hmac.New(lsh256.New, key)
	if _, err = io.Copy(h, file); err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	fmt.Println("MAC:", hex.EncodeToString(h.Sum(nil)))
}
