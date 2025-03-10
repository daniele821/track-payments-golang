package utils

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
)

func DecryptFile(cipherFile, keyFile string) (string, error) {
	if _, err := exec.LookPath("openssl"); err != nil {
		return "", err
	}

	// Capture decrypted output
	var stdout bytes.Buffer
	cmd := exec.Command("openssl", "enc", "-d", "-aes-256-cbc", "-in", cipherFile, "-k", keyFile, "-pbkdf2")
	cmd.Stdout = &stdout
	cmd.Stderr = os.Stderr // Capture errors

	// Run command
	if err := cmd.Run(); err != nil {
		return "", errors.New("decryption failed: " + err.Error())
	}

	return stdout.String(), nil
}
