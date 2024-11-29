package crypto

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

func Sign(data []byte) ([]byte, error) {
	privateKeyBytes, err := os.ReadFile("keys/private.pem")
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(privateKeyBytes)
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	hash := crypto.SHA256.New()
	hash.Write(data)
	hashed := hash.Sum(nil)

	return rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
}
