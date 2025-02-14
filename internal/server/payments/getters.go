package payments

import (
	"errors"
	"fmt"
)

func (allPayments AllPayments) Payment(date string) (Payment, error) {
	var paymentEmpty Payment
	if err := allPayments.checks(&date, nil, nil, nil, nil); err != nil {
		return paymentEmpty, err
	}
	payment, foundPayment := allPayments.p.payments.Get(NewPaymentForSearches(date))
	if !foundPayment {
		return paymentEmpty, errors.New(fmt.Sprintf("payment (%s) not found", date))
	}
	return payment, nil
}

func (allPayments AllPayments) Order(date, item string) (Order, error) {
	var orderEmpty Order
	if err := allPayments.checks(&date, nil, nil, nil, &item); err != nil {
		return orderEmpty, err
	}
	payment, err := allPayments.Payment(date)
	if err != nil {
		return orderEmpty, err
	}
	order, foundOrder := payment.p.orders.Get(NewOrderForSearches(item))
	if !foundOrder {
		return orderEmpty, errors.New(fmt.Sprintf("order (%s, %s) not found", date, item))
	}
	return order, nil
}

func (allPayments AllPayments) Cities() ReadOnlyBTree[string] {
	return ReadOnlyBTree[string]{btree: allPayments.p.valueSet.p.cities}
}

func (allPayments AllPayments) Shops() ReadOnlyBTree[string] {
	return ReadOnlyBTree[string]{btree: allPayments.p.valueSet.p.shops}
}

func (allPayments AllPayments) PaymentMethods() ReadOnlyBTree[string] {
	return ReadOnlyBTree[string]{btree: allPayments.p.valueSet.p.paymentMethods}
}

func (allPayments AllPayments) Items() ReadOnlyBTree[string] {
	return ReadOnlyBTree[string]{btree: allPayments.p.valueSet.p.items}
}

func (allPayments AllPayments) Payments() ReadOnlyBTree[Payment] {
	return ReadOnlyBTree[Payment]{btree: allPayments.p.payments}
}

func (payment Payment) City() string {
	return payment.p.city
}

func (payment Payment) Shop() string {
	return payment.p.shop
}

func (payment Payment) PaymentMethod() string {
	return payment.p.paymentMethod
}

func (payment Payment) Date() string {
	return payment.p.date
}

func (payment Payment) Orders() ReadOnlyBTree[Order] {
	return ReadOnlyBTree[Order]{btree: payment.p.orders}
}

func (payment Payment) TotalPrice() int {
	acc := 0
	payment.Orders().Ascend(func(item Order) bool {
		acc += item.Quantity() * item.UnitPrice()
		return true
	}, nil, nil)
	return acc
}

func (order Order) Quantity() int {
	return order.p.quantity
}

func (order Order) UnitPrice() int {
	return order.p.unitPrice
}

func (order Order) Item() string {
	return order.p.item
}
