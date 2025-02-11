package payments

import (
	"errors"
	"fmt"

	"github.com/google/btree"
)

func compactBtreeToSlice[T any](data *btree.BTreeG[T]) []T {
	res := make([]T, data.Len())
	iter := 0
	data.Ascend(func(item T) bool {
		res[iter] = item
		iter += 1
		return true
	})
	return res
}

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

func (allPayments *AllPayments) GetAllCities() []string {
	return compactBtreeToSlice(allPayments.valueSet.cities)
}

func (allPayments *AllPayments) GetAllShops() []string {
	return compactBtreeToSlice(allPayments.valueSet.shops)
}

func (allPayments *AllPayments) getAllPaymentMethods() []string {
	return compactBtreeToSlice(allPayments.valueSet.paymentMethods)
}

func (allPayments *AllPayments) GetAllItems() []string {
	return compactBtreeToSlice(allPayments.valueSet.items)
}
