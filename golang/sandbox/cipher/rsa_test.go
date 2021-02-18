package cipher

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

func ExamplePrivatePEMKey() {
	reader := rand.Reader
	bitSize := 512

	key, _ := rsa.GenerateKey(reader, bitSize)

	fmt.Println(PrivatePEMKey(key))
}

func ExamplePublicPEMKey() {
	reader := rand.Reader
	bitSize := 512

	key, _ := rsa.GenerateKey(reader, bitSize)

	fmt.Println(PublicPEMKey(&key.PublicKey))
}
