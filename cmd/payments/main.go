package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"payment/internal/client/cli"
)

func main() {
	args := os.Args[1:]

	// print help message
	if cli.ParseHelp(args) {
		return
	}

	// get paths
	jsonDir, err := getExeDir()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	jsonPath := filepath.Join(append([]string{jsonDir}, "payments-cipher.json")...)

	// run cli tool
	if err := cli.Run(jsonPath); err != nil {
		fmt.Printf("%s\n", err)
		return
	}
}

func getExeDir() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	exePath, err = filepath.EvalSymlinks(exePath)
	if err != nil {
		return "", err
	}
	return filepath.Dir(exePath), nil
}

func encryptFile(plainText, cipherFile, keyFile string) {
	// Reading key
	key, err := os.ReadFile(keyFile)
	if err != nil {
		log.Fatalf("read file err: %v", err.Error())
	}

	// Creating block of algorithm
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("cipher err: %v", err.Error())
	}

	// Creating GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatalf("cipher GCM err: %v", err.Error())
	}

	// Generating random nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatalf("nonce  err: %v", err.Error())
	}

	// Decrypt file
	cipherText := gcm.Seal(nonce, nonce, []byte(plainText), nil)

	// Writing ciphertext file
	err = os.WriteFile(cipherFile, cipherText, 0600)
	if err != nil {
		log.Fatalf("write file err: %v", err.Error())
	}

}

func decryptFile(cipherFile, keyFile string) (plainText string) {
	// Reading ciphertext file
	cipherText, err := os.ReadFile(cipherFile)
	if err != nil {
		log.Fatal(err)
	}

	// Reading key
	key, err := os.ReadFile(keyFile)
	if err != nil {
		log.Fatalf("read file err: %v", err.Error())
	}

	// Creating block of algorithm
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("cipher err: %v", err.Error())
	}

	// Creating GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatalf("cipher GCM err: %v", err.Error())
	}

	// Deattached nonce and decrypt
	nonce := cipherText[:gcm.NonceSize()]
	cipherText = cipherText[gcm.NonceSize():]
	plainTextByte, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		log.Fatalf("decrypt file err: %v", err.Error())
	}

	return string(plainTextByte)
}
