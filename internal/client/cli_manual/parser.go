package cli_manual

import (
	"errors"
	"fmt"
	"os"
	"payment/internal/server/payments"
	"strings"
)

func errMsgForWrongArgs(args []string, wrongAt int) error {
	return errors.New(fmt.Sprintf("invalid args for %s: %s", strings.Join(args[:wrongAt], " "), strings.Join(args[wrongAt:], " ")))
}

func parseParamsAndRun(allPayments payments.AllPayments) error {
	return parseAndRun(allPayments, os.Args[1:])
}

func parseAndRun(allPayments payments.AllPayments, args []string) error {
	if len(args) == 0 {
		return nil
	}
	arg0 := args[0]
	arg1 := args[1]
	switch {
	case matchEveryLenght(arg0, "insert"):
		switch {
		}
	case matchEveryLenght(arg0, "list"):
		switch {
		case matchEveryLenghtFromAnyWords(arg1, []string{"city", "cities"}):
		case matchEveryLenghtFromAnyWords(arg1, []string{"shops"}):
		case matchEveryLenghtFromAnyWords(arg1, []string{"methods"}):
		case matchEveryLenghtFromAnyWords(arg1, []string{"items"}):
		case matchEveryLenghtFromAnyWords(arg1, []string{"payments"}):
		case matchEveryLenghtFromAnyWords(arg1, []string{"orders"}):
		default:
			return errMsgForWrongArgs(args, 1)
		}
	case matchEveryLenght(arg0, "update"):
		switch {
		}
	case matchEveryLenght(arg0, "delete"):
		switch {
		}
	case matchEveryLenght(arg0, "help"):
		helpMsg()
	default:
		return errMsgForWrongArgs(args, 0)
	}
	return nil
}
