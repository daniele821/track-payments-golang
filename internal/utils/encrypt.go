package utils

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
)

func EncryptFile(plainText, cipherFile, keyFile string) error {
	if _, err := exec.LookPath("openssl"); err != nil {
		return err
	}
	cmd := exec.Command("openssl", "enc", "-aes-256-cbc", "-out", cipherFile, "-k", keyFile, "-pbkdf2")
	cmd.Stdin = bytes.NewBufferString(plainText)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		return errors.New("encryption failed!")
	}
	return nil
}
