package main

import (
	"fmt"
	"os"
	"path/filepath"
	"payment/internal/client/cli"
	"payment/internal/server/payments"
	"payment/internal/utils"
)

const cipherKeyFile = ".cipher_key"
const cipherJsonFile = "payments.json.cipher"
const jsonLocalFile = "payments.json"

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
	jsonDir, err := utils.GetExeDir()
	if err != nil {
		return err
	}
	cipherKeyPath := filepath.Join(append([]string{jsonDir}, cipherKeyFile)...)
	cipherJsonPath := filepath.Join(append([]string{jsonDir}, cipherJsonFile)...)
	jsonLocalPath := filepath.Join(append([]string{jsonDir}, jsonLocalFile)...)

	// load from local file or server encrypted one
	var allPayments payments.AllPayments
	var storedData string
	if _, found := os.LookupEnv("LOCAL"); found {
		allPayments, err = payments.NewAllPaymentsFromjsonFile(jsonLocalPath)
	} else {
		storedData, err = utils.DecryptFile(cipherJsonPath, cipherKeyPath)
		if err != nil {
			return err
		}
		allPayments, _ = payments.NewAllPaymentsFromJson(storedData)
	}

	// run cli tool
	if err := cli.ParseAndRun(allPayments, args); err != nil {
		return err
	}

	// save changes to encrypted file
	newStoredData, err := allPayments.DumpJson(false)
	if newStoredData != storedData {
		if err := utils.EncryptFile(newStoredData, cipherJsonPath, cipherKeyPath); err != nil {
			return err
		}
	}
	if err := allPayments.DumpJsonToFile(jsonLocalPath, true); err != nil {
		return err
	}

	return nil
}
