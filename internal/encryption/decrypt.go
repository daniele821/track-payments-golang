package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"os"
)

func DecryptFile(cipherFile, keyFile string) (string, error) {
	// Reading ciphertext file
	cipherText, err := os.ReadFile(cipherFile)
	if err != nil {
		return "", err
	}

	// Reading key
	key, err := os.ReadFile(keyFile)
	if err != nil {
		return "", err
	}

	// Creating block of algorithm
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Creating GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Deattached nonce and decrypt
	nonce := cipherText[:gcm.NonceSize()]
	cipherText = cipherText[gcm.NonceSize():]
	plainTextByte, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}

	return string(plainTextByte), nil
}
