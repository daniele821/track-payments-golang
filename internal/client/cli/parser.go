package cli

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
			return insertGeneric("cities", args[2:], allPayments.AddCities)
		case matchEveryLenght(args[1], "shops"):
			return insertGeneric("shops", args[2:], allPayments.AddShops)
		case matchEveryLenght(args[1], "methods"):
			return insertGeneric("methods", args[2:], allPayments.AddPaymentMethods)
		case matchEveryLenght(args[1], "items"):
			return insertGeneric("items", args[2:], allPayments.AddItems)
		case matchEveryLenght(args[1], "payments"):
			return insertPayments(allPayments, args[2:])
		case matchEveryLenght(args[1], "orders"):
			return insertOrders(allPayments, args[2:])
		case matchEveryLenght(args[1], "details"):
			return insertDetails(allPayments, args[2:])
		default:
			return errors.New(fmt.Sprintf("invalid arg for insert: %s", args[1]))
		}
	case matchEveryLenght(args[0], "list"):
		switch {
		case len(args) <= 1:
			return errors.New("missing arg for list")
		case matchEveryLenghtFromAnyWords(args[1], []string{"city", "cities"}):
			listGeneric("cities", allPayments.Cities())
		case matchEveryLenght(args[1], "shops"):
			listGeneric("shops", allPayments.Shops())
		case matchEveryLenght(args[1], "methods"):
			listGeneric("methods", allPayments.PaymentMethods())
		case matchEveryLenght(args[1], "items"):
			listGeneric("items", allPayments.Items())
		case matchEveryLenght(args[1], "values"):
			listGeneric("cities", allPayments.Cities())
			listGeneric("shops", allPayments.Shops())
			listGeneric("methods", allPayments.PaymentMethods())
			listGeneric("items", allPayments.Items())
		case matchEveryLenght(args[1], "payments"):
			listPayments(allPayments.Payments())
		default:
			return errors.New(fmt.Sprintf("invalid arg for list: %s", args[1]))
		}
	case matchEveryLenght(args[0], "visualize"):
		switch {
		case len(args) <= 1:
			return errors.New("missing arg for visualize")
		case matchEveryLenghtFromAnyWords(args[1], []string{"city", "cities"}):
		case matchEveryLenght(args[1], "shops"):
		case matchEveryLenght(args[1], "methods"):
		case matchEveryLenght(args[1], "items"):
		case matchEveryLenght(args[1], "payments"):
		default:
			return errors.New(fmt.Sprintf("invalid arg for visualize: %s", args[1]))
		}
	case matchEveryLenght(args[0], "update"):
		switch {
		case len(args) <= 1:
			return errors.New("missing arg for update")
		case matchEveryLenght(args[1], "payments"):
			return updatePayments(allPayments, args[2:])
		case matchEveryLenght(args[1], "orders"):
			return updateOrders(allPayments, args[2:])
		case matchEveryLenght(args[1], "details"):
			return updateDetails(allPayments, args[2:])
		default:
			return errors.New(fmt.Sprintf("invalid arg for update: %s", args[1]))
		}
	case matchEveryLenght(args[0], "delete"):
		switch {
		case len(args) <= 1:
			return errors.New("missing arg for delete")
		case matchEveryLenght(args[1], "payments"):
			return deletePayments(allPayments, args[2:])
		case matchEveryLenght(args[1], "orders"):
			return deleteOrders(allPayments, args[2:])
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
