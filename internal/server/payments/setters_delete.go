package payments

import (
	"errors"
	"fmt"
)

func (allPayments *AllPayments) RemovePayment(date string) error {
	if err := allPayments.checks(&date, nil, nil, nil, nil); err != nil {
		return err
	}
	_, found := allPayments.payments.Delete(NewPaymentForSearches(date))
	if !found {
		return errors.New(fmt.Sprintf("payment (%s) was not found", date))
	}
	return nil
}

func (allPayments *AllPayments) RemoveOrder(date, item string) error {
	if err := allPayments.checks(&date, nil, nil, nil, &item); err != nil {
		return err
	}
	payment, err := allPayments.Payment(date)
	if err != nil {
		return err
	}
	_, foundOrder := payment.pointer.orders.Delete(NewOrderForSearches(item))
	if !foundOrder {
		return errors.New(fmt.Sprintf("order (%s, %s) was not found", date, item))
	}
	return nil
}
