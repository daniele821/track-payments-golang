package cli_manual

import (
	"errors"
	"fmt"
	"os"
	"payment/internal/server/payments"
)

func parseParamsAndRun(allPayments payments.AllPayments) error {
	return parseAndRun(allPayments, os.Args[1:])
}

func parseAndRun(allPayments payments.AllPayments, args []string) error {
	if len(args) == 0 {
		return nil
	}
	switch {
	case matchEveryLenght(args[0], "insert"):
		switch {
		case len(args) <= 1:
			return errors.New("missing arg for insert")
		case matchEveryLenghtFromAnyWords(args[1], []string{"city", "cities"}):
		case matchEveryLenght(args[1], "shops"):
		case matchEveryLenght(args[1], "methods"):
		case matchEveryLenght(args[1], "items"):
		case matchEveryLenght(args[1], "payments"):
		case matchEveryLenght(args[1], "orders"):
		default:
			return errors.New(fmt.Sprintf("invalid arg for insert: %s", args[1]))
		}
	case matchEveryLenght(args[0], "list"):
		switch {
		case len(args) <= 1:
			return errors.New("missing arg for list")
		case matchEveryLenghtFromAnyWords(args[1], []string{"city", "cities"}):
		case matchEveryLenght(args[1], "shops"):
		case matchEveryLenght(args[1], "methods"):
		case matchEveryLenght(args[1], "items"):
		case matchEveryLenght(args[1], "payments"):
		case matchEveryLenght(args[1], "orders"):
		default:
			return errors.New(fmt.Sprintf("invalid arg for list: %s", args[1]))
		}
	case matchEveryLenght(args[0], "update"):
		switch {
		case len(args) <= 1:
			return errors.New("missing arg for update")
		case matchEveryLenght(args[1], "payments"):
		case matchEveryLenght(args[1], "orders"):
		default:
			return errors.New(fmt.Sprintf("invalid arg for update: %s", args[1]))
		}
	case matchEveryLenght(args[0], "delete"):
		switch {
		case len(args) <= 1:
			return errors.New("missing arg for delete")
		case matchEveryLenght(args[1], "payments"):
		case matchEveryLenght(args[1], "orders"):
		default:
			return errors.New(fmt.Sprintf("invalid arg for delete: %s", args[1]))
		}
	case matchEveryLenght(args[0], "help"):
		helpMsg()
	default:
		return errors.New(fmt.Sprintf("invalid arg: %s", args[0]))
	}
	return nil
}
