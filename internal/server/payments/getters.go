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

func (allPayments *AllPayments) Payments() *ReadOnlyBTree[*Payment] {
	return &ReadOnlyBTree[*Payment]{btree: allPayments.payments}
}

func (payment *Payment) City() string {
	return payment.city
}

func (payment *Payment) Shop() string {
	return payment.shop
}

func (payment *Payment) PaymentMethod() string {
	return payment.paymentMethod
}

func (payment *Payment) Date() string {
	return payment.date
}

func (payment *Payment) Description() string {
	return payment.description
}

func (payment *Payment) Orders() *ReadOnlyBTree[*Order] {
	return &ReadOnlyBTree[*Order]{btree: payment.orders}
}

func (order *Order) Quantity() uint {
	return order.quantity
}

func (order *Order) UnitPrice() uint {
	return order.unitPrice
}

func (order *Order) Item() string {
	return order.item
}
