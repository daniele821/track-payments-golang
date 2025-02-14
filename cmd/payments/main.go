package main

import (
	"fmt"
	"os"
	"path/filepath"
	"payment/internal/client/cli"
)

func main() {
	exePath, err := os.Executable()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	exePath, err = filepath.EvalSymlinks(exePath)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	jsonDir := filepath.Dir(exePath)
	jsonPath := filepath.Join(append([]string{jsonDir}, "payments.json")...)
	if err := cli.Run(jsonPath); err != nil {
		fmt.Printf("%s\n", err)
		return
	}
}
