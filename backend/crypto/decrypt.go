package crypto

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rsa"
	"crypto/sha256"
	"errors"
)

// Decrypt decrypts the given data using the provided AES key.

func Decrypt(encryptedData []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(encryptedData) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	iv := encryptedData[:aes.BlockSize]
	encryptedData = encryptedData[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	plaintext := make([]byte, len(encryptedData))
	stream.XORKeyStream(plaintext, encryptedData)

	return plaintext, nil
}

// VerifySignature verifies the RSA digital signature for the given data

func VerifySignature(data []byte, signature []byte, publicKey *rsa.PublicKey) error {
	hashed := sha256.Sum256(data)

	// Verify the signature using RSA public key
	err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], signature)
	if err != nil {
		return errors.New("signature verification failed")
	}
	return nil
}
