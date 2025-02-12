package payments

import (
	"errors"
	"fmt"
)

func (allPayments *AllPayments) Payment(date string) (Payment, error) {
	var paymentEmpty Payment
	if err := allPayments.checks(&date, nil, nil, nil, nil); err != nil {
		return paymentEmpty, err
	}
	payment, foundPayment := allPayments.payments.Get(NewPaymentForSearches(date))
	if !foundPayment {
		return paymentEmpty, errors.New(fmt.Sprintf("payment (%s) not found", date))
	}
	return payment, nil
}

func (allPayments *AllPayments) Order(date, item string) (Order, error) {
	var orderEmpty Order
	if err := allPayments.checks(&date, nil, nil, nil, &item); err != nil {
		return orderEmpty, err
	}
	payment, err := allPayments.Payment(date)
	if err != nil {
		return orderEmpty, err
	}
	order, foundOrder := payment.pointer.orders.Get(NewOrderForSearches(item))
	if !foundOrder {
		return orderEmpty, errors.New(fmt.Sprintf("order (%s, %s) not found", date, item))
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

func (allPayments *AllPayments) Payments() *ReadOnlyBTree[Payment] {
	return &ReadOnlyBTree[Payment]{btree: allPayments.payments}
}

func (payment Payment) City() string {
	return payment.pointer.city
}

func (payment Payment) Shop() string {
	return payment.pointer.shop
}

func (payment Payment) PaymentMethod() string {
	return payment.pointer.paymentMethod
}

func (payment Payment) Date() string {
	return payment.pointer.date
}

func (payment Payment) Description() string {
	return payment.pointer.description
}

func (payment Payment) Orders() *ReadOnlyBTree[Order] {
	return &ReadOnlyBTree[Order]{btree: payment.pointer.orders}
}

func (order Order) Quantity() uint {
	return order.pointer.quantity
}

func (order Order) UnitPrice() uint {
	return order.pointer.unitPrice
}

func (order Order) Item() string {
	return order.pointer.item
}
