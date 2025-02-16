package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"payment/internal/client/cli"
	"payment/internal/server/payments"
	"strings"
)

const cipherKeyFile = ".cipher_key"
const cipherJsonFile = "payments-cipher.json"

func main() {
	if err := runner(); err != nil {
		fmt.Println(err)
	}
}

func runner() error {
	args := os.Args[1:]

	// print help message
	if cli.ParseHelp(args) {
		return nil
	}

	// get paths
	jsonDir, err := getExeDir()
	if err != nil {
		return err
	}
	cipherKeyPath := filepath.Join(append([]string{jsonDir}, cipherKeyFile)...)
	cipherJsonPath := filepath.Join(append([]string{jsonDir}, cipherJsonFile)...)

	// decrypt file and create data structure
	storedData, err := decryptFile(cipherJsonPath, cipherKeyPath)
	if err != nil {
		fmt.Printf("data decryption failed: %s\n", err)
		fmt.Printf("Do you want to OVERWRITE the file with empty data? ")
		scanner := bufio.NewScanner(os.Stdin)
	outerLoop:
		for scanner.Scan() {
			input := scanner.Text()
			switch strings.ToLower(input) {
			case "y":
				break outerLoop
			case "n":
				return errors.New("cipher file couldn't be decrypted")
			default:
				fmt.Printf("invalid answer (y/n): ")
			}
		}
	}
	allPayments, _ := payments.NewAllPaymentsFromJson(storedData)

	// run cli tool
	if err := cli.ParseAndRun(allPayments, args); err != nil {
		return err
	}

	// save changes to encrypted file
	newStoredData, err := allPayments.DumpJson(false)
	if newStoredData != storedData {
		if err := encryptFile(newStoredData, cipherJsonPath, cipherKeyPath); err != nil {
			return err
		}
	}

	return nil
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

func encryptFile(plainText, cipherFile, keyFile string) error {
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

func decryptFile(cipherFile, keyFile string) (string, error) {
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
