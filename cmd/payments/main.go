package main

import (
	"fmt"
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
	jsonPath := filepath.Join(append([]string{jsonDir}, "payments.json")...)

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
