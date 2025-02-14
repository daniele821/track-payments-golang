package cli

import (
	"payment/internal/server/payments"
)

func Run(jsonPath string) error {
	// load allPayments data
	allPayments, err := payments.NewAllPaymentsFromjsonFile(jsonPath)
	if err != nil {
		return err
	}

	// run program
	if err := parseParamsAndRun(allPayments); err != nil {
		return err
	}

	// save to json file
	if err := allPayments.DumpJsonToFile(jsonPath, true); err != nil {
		return err
	}

	return nil
}
