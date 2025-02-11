package payments

import (
	"errors"
	"fmt"
)

func (allPayments *AllPayments) GetPayment(date string) (*Payment, error) {
	if err := allPayments.checks(&date, nil, nil, nil, nil); err != nil {
		return nil, err
	}
	payment, foundPayment := allPayments.payments.Get(newPaymentForSearches(date))
	if !foundPayment {
		return nil, errors.New(fmt.Sprintf("payment (%s) not found", date))
	}
	return payment, nil
}

func (allPayments *AllPayments) GetOrder(date, item string) (*Order, error) {
	if err := allPayments.checks(&date, nil, nil, nil, &item); err != nil {
		return nil, err
	}
	payment, err := allPayments.GetPayment(date)
	if err != nil {
		return nil, err
	}
	order, foundOrder := payment.orders.Get(newOrderForSearches(item))
	if !foundOrder {
		return nil, errors.New(fmt.Sprintf("order (%s, %s) not found", date, item))
	}
	return order, nil
}
