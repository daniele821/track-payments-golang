package main

import (
	"fmt"
	"io"
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
	fromStdin := false
	if info, err := os.Stdin.Stat(); info.Mode()&os.ModeCharDevice == 0 {
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			return err
		}
		allPayments, err = payments.NewAllPaymentsFromJson(string(data))
		if err != nil {
			return err
		}
		fromStdin = true
	} else if err != nil {
		return err
	} else if _, found := os.LookupEnv("LOCAL"); found {
		allPayments, err = payments.NewAllPaymentsFromjsonFile(jsonLocalPath)
		if err != nil {
			return err
		}
	} else {
		storedData, err = utils.DecryptFile(cipherJsonPath, cipherKeyPath)
		if err != nil {
			return err
		}
		allPayments, err = payments.NewAllPaymentsFromJson(storedData)
		if err != nil {
			return err
		}
	}

	// run cli tool
	if err := cli.ParseAndRun(allPayments, args); err != nil {
		return err
	}

	// save changes to encrypted file
	if _, found := os.LookupEnv("DRYRUN"); !found && !fromStdin {
		newStoredData, err := allPayments.DumpJson(false)
		if err != nil {
			return err
		}
		oldDecryptedData, err := utils.DecryptFile(cipherJsonPath, cipherKeyPath)
		if err != nil {
			return err
		}
		if newStoredData != oldDecryptedData {
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
