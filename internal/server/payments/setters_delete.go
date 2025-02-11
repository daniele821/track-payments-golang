package payments

import (
	"errors"
)

func (allPayments *AllPayments) RemovePayment(date string) error {
	if err := allPayments.checks(&date, nil, nil, nil, nil); err != nil {
		return err
	}
	_, found := allPayments.payments.Delete(newPaymentForSearches(date))
	if !found {
		return errors.New("payment was not found")
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
	_, foundOrder := payment.orders.Delete(newOrderForSearches(item))
	if !foundOrder {
		return errors.New("order not found")
	}
	return nil
}
