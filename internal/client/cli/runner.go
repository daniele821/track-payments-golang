package cli

import (
	"os"
	"payment/internal/server/payments"
)

func Run(jsonPath string) error {
	// load allPayments data
	allPayments, err := payments.NewAllPaymentsFromjsonFile(jsonPath)
	if err != nil {
		return err
	}

	// run program
	if helpAction := parseHelp(os.Args[1:]); !helpAction {
		if err := parseAndRun(allPayments, os.Args[1:]); err != nil {
			return err
		}
	}

	// save to json file
	if err := allPayments.DumpJsonToFile(jsonPath, true); err != nil {
		return err
	}

	return nil
}
