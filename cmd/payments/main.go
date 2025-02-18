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
		if args := os.Args[1:]; len(args) == 1 && args[0] == "print" {
			jsonDir, _ := utils.GetExeDir()
			encryptedFile := filepath.Join([]string{jsonDir, cipherJsonFile}...)
			encryptedBytes, _ := os.ReadFile(encryptedFile)
			fmt.Println(string(encryptedBytes))
			return
		}
		fmt.Println(err)
		os.Exit(1)
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
	cipherKeyPath := filepath.Join([]string{jsonDir, cipherKeyFile}...)
	cipherJsonPath := filepath.Join([]string{jsonDir, cipherJsonFile}...)
	jsonLocalPath := filepath.Join([]string{jsonDir, jsonLocalFile}...)

	// load from local file or server encrypted one
	var allPayments payments.AllPayments
	var storedData string
	if _, found := os.LookupEnv("LOCAL"); found {
		allPayments, _ = payments.NewAllPaymentsFromjsonFile(jsonLocalPath)
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
	if _, found := os.LookupEnv("DRYRUN"); !found {
		newStoredData, err := allPayments.DumpJson(false)
		if err != nil {
			return err
		}
		if newStoredData != storedData {
			if err := utils.EncryptFile(newStoredData, cipherJsonPath, cipherKeyPath); err != nil {
				return err
			}
		}
		if err := allPayments.DumpJsonToFile(jsonLocalPath, true); err != nil {
			return err
		}
	}

	return nil
}
