package cli

import (
	"os"
	"path/filepath"
	"payment/internal/server/payments"
)

func Run(jsonPathFromExeDir ...string) error {
	// load json file
	exePath, err := os.Executable()
	if err != nil {
		return err
	}
	exePath, err = filepath.EvalSymlinks(exePath)
	if err != nil {
		return err
	}
	jsonDir := filepath.Dir(exePath)
	jsonPath := filepath.Join(append([]string{jsonDir}, jsonPathFromExeDir...)...)
	jsonDataByte, err := os.ReadFile(jsonPath)
	if err != nil {
		fileCreated, err := os.Create(jsonPath)
		if err != nil {
			return err
		}
		defer fileCreated.Close()
		if _, err := fileCreated.WriteString("{}"); err != nil {
			return err
		}
		jsonDataByte, err = os.ReadFile(jsonPath)
		if err != nil {
			return err
		}
	}
	JsonData := string(jsonDataByte)

	// load all payments from json file
	allPayments, err := payments.NewAllPaymentsFromJson(JsonData)
	if err != nil {
		return err
	}

	// execute program
	if err := parseParamsAndRun(allPayments); err != nil {
		return err
	}

	// dump all payments to json file
	jsonData, err := allPayments.DumpJson(true)
	if err != nil {
		return err
	}
	if err := os.WriteFile(jsonPath, []byte(jsonData), 0644); err != nil {
		return err
	}

	return nil
}
