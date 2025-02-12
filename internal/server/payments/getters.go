package payments

import (
	"errors"
	"fmt"
)

func (allPayments *AllPayments) Payment(date string) (*payment, error) {
	if err := allPayments.checks(&date, nil, nil, nil, nil); err != nil {
		return nil, err
	}
	payment, foundPayment := allPayments.payments.Get(newPaymentForSearches(date))
	if !foundPayment {
		return nil, errors.New(fmt.Sprintf("payment (%s) not found", date))
	}
	return payment, nil
}

func (allPayments *AllPayments) Order(date, item string) (*order, error) {
	if err := allPayments.checks(&date, nil, nil, nil, &item); err != nil {
		return nil, err
	}
	payment, err := allPayments.Payment(date)
	if err != nil {
		return nil, err
	}
	order, foundOrder := payment.orders.Get(newOrderForSearches(item))
	if !foundOrder {
		return nil, errors.New(fmt.Sprintf("order (%s, %s) not found", date, item))
	}
	return order, nil
}

func (allPayments *AllPayments) Cities() *ReadOnlyBTree[string] {
	return &ReadOnlyBTree[string]{btree: allPayments.valueSet.cities}
}

func (allPayments *AllPayments) Shops() *ReadOnlyBTree[string] {
	return &ReadOnlyBTree[string]{btree: allPayments.valueSet.shops}
}

func (allPayments *AllPayments) PaymentMethods() *ReadOnlyBTree[string] {
	return &ReadOnlyBTree[string]{btree: allPayments.valueSet.paymentMethods}
}

func (allPayments *AllPayments) Items() *ReadOnlyBTree[string] {
	return &ReadOnlyBTree[string]{btree: allPayments.valueSet.items}
}

func (allPayments *AllPayments) Payments() *ReadOnlyBTree[*payment] {
	return &ReadOnlyBTree[*payment]{btree: allPayments.payments}
}

func (payment *payment) City() string {
	return payment.city
}

func (payment *payment) Shop() string {
	return payment.shop
}

func (payment *payment) PaymentMethod() string {
	return payment.paymentMethod
}

func (payment *payment) Date() string {
	return payment.date
}

func (payment *payment) Description() string {
	return payment.description
}

func (payment *payment) Orders() *ReadOnlyBTree[*order] {
	return &ReadOnlyBTree[*order]{btree: payment.orders}
}

func (order *order) Quantity() uint {
	return order.quantity
}

func (order *order) UnitPrice() uint {
	return order.unitPrice
}

func (order *order) Item() string {
	return order.item
}
