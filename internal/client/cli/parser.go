package cli

import (
	"errors"
	"fmt"
	"payment/internal/server/payments"
)

func ParseAndRun(allPayments payments.AllPayments, args []string) error {
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
		if len(args) <= 1 {
			return errors.New("missing arg for list")
		}
		from, to, err := parseRanges(args[2:])
		if err != nil {
			return err
		}
		switch {
		case matchEveryLenghtFromAnyWords(args[1], []string{"city", "cities"}):
			listGeneric("cities", allPayments.Cities(), from, to)
		case matchEveryLenght(args[1], "shops"):
			listGeneric("shops", allPayments.Shops(), from, to)
		case matchEveryLenght(args[1], "methods"):
			listGeneric("methods", allPayments.PaymentMethods(), from, to)
		case matchEveryLenght(args[1], "items"):
			listGeneric("items", allPayments.Items(), from, to)
		case matchEveryLenght(args[1], "values"):
			listGeneric("cities", allPayments.Cities(), from, to)
			listGeneric("shops", allPayments.Shops(), from, to)
			listGeneric("methods", allPayments.PaymentMethods(), from, to)
			listGeneric("items", allPayments.Items(), from, to)
		case matchEveryLenght(args[1], "payments"):
			listPayments(allPayments.Payments(), from, to)
		default:
			return errors.New(fmt.Sprintf("invalid arg for list: %s", args[1]))
		}
	case matchEveryLenght(args[0], "visualize"):
		if len(args) <= 1 {
			return errors.New("missing arg for list")
		}
		from, to, err := parseRanges(args[2:])
		if err != nil {
			return err
		}
		switch {
		case matchEveryLenghtFromAnyWords(args[1], []string{"city", "cities"}):
			visualizeGeneric("CITY", allPayments.Cities(), from, to)
		case matchEveryLenght(args[1], "shops"):
			visualizeGeneric("SHOP", allPayments.Shops(), from, to)
		case matchEveryLenght(args[1], "methods"):
			visualizeGeneric("METHOD", allPayments.PaymentMethods(), from, to)
		case matchEveryLenght(args[1], "items"):
			visualizeGeneric("ITEM", allPayments.Items(), from, to)
		case matchEveryLenght(args[1], "values"):
			visualizeGeneric("CITY", allPayments.Cities(), from, to)
			visualizeGeneric("SHOP", allPayments.Shops(), from, to)
			visualizeGeneric("METHOD", allPayments.PaymentMethods(), from, to)
			visualizeGeneric("ITEM", allPayments.Items(), from, to)
		case matchEveryLenght(args[1], "payments"):
			visualizePayment(allPayments.Payments(), from, to)
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
	case matchEveryLenght(args[0], "print"):
		jsonData, err := allPayments.DumpJson(true)
		if err != nil {
			return nil
		}
		fmt.Println(jsonData)
	default:
		return errors.New(fmt.Sprintf("invalid arg: %s", args[0]))
	}
	return nil
}
