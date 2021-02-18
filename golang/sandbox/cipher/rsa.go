package cipher

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

func PrivatePEMKey(key *rsa.PrivateKey) string {
	var buf bytes.Buffer
	var privateKey = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	pem.Encode(&buf, privateKey)
	return buf.String()
}

func PublicPEMKey(pubkey *rsa.PublicKey) string {
	var buf bytes.Buffer
	b, _ := x509.MarshalPKIXPublicKey(pubkey)

	var pemkey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: b,
	}

	pem.Encode(&buf, pemkey)
	return buf.String()
}
