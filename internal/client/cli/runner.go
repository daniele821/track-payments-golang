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

	allPayments, err := payments.NewAllPaymentsFromjsonFile(jsonPath)
	if err != nil {
		return err
	}

	// execute program
	if err := parseParamsAndRun(allPayments); err != nil {
		return err
	}

	if err := allPayments.DumpJsonToFile(jsonPath, true); err != nil {
		return err
	}

	return nil
}
