package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
)

func EncryptFile(plainText, cipherFile, keyFile string) error {
	// Reading key
	key, err := os.ReadFile(keyFile)
	if err != nil {
		return err
	}

	// Creating block of algorithm
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	// Creating GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	// Generating random nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	// Decrypt file
	cipherText := gcm.Seal(nonce, nonce, []byte(plainText), nil)

	// Writing ciphertext file
	err = os.WriteFile(cipherFile, cipherText, 0600)
	if err != nil {
		return err
	}
	return nil
}
