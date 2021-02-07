package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"

	"sandbox/cipher"
)

func main() {
	reader := rand.Reader
	bitSize := 512
	key, _ := rsa.GenerateKey(reader, bitSize)
	fmt.Println(cipher.PrivatePEMKey(key))
	fmt.Println(cipher.PublicPEMKey(&key.PublicKey))
}
