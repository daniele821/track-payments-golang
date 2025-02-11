package payments

import (
	"errors"

	"github.com/google/btree"
)

func insertAll(valueSet *btree.BTreeG[string], elems ...string) (duplicates []string) {
	for _, elem := range elems {
		if old, replaced := valueSet.ReplaceOrInsert(elem); replaced {
			duplicates = append(duplicates, old)
		}
	}
	if len(duplicates) == 0 {
		return nil
	}
	return duplicates
}

func (allPayments *AllPayments) AddCities(cities ...string) (duplicates []string) {
	return insertAll(allPayments.valueSet.cities, cities...)
}

func (allPayments *AllPayments) AddShops(shops ...string) (duplicates []string) {
	return insertAll(allPayments.valueSet.shops, shops...)
}

func (allPayments *AllPayments) AddPaymentMethods(paymentMethods ...string) (duplicates []string) {
	return insertAll(allPayments.valueSet.paymentMethods, paymentMethods...)
}

func (allPayments *AllPayments) AddItems(items ...string) (duplicates []string) {
	return insertAll(allPayments.valueSet.items, items...)
}

func (allPayments *AllPayments) AddPayment(city, shop, paymentMethod, date, description string) error {
	if err := allPayments.checks(&date, &city, &shop, &paymentMethod, nil); err != nil {
		return err
	}
	payment := newPayment(city, shop, paymentMethod, date, description)
	if allPayments.payments.Has(payment) {
		return errors.New("invalid date: already exists")
	}
	if _, replaced := allPayments.payments.ReplaceOrInsert(payment); replaced {
		panic("UNREACHABLE CODE: already check payment wasn't already inserted!")
	}
	return nil
}

func (allPayments *AllPayments) AddOrder(quantity, unitPrice uint, item, date string) error {
	if err := allPayments.checks(&date, nil, nil, nil, &item); err != nil {
		return err
	}
	order := newOrder(quantity, unitPrice, item)
	payment, err := allPayments.GetPayment(date)
	if err != nil {
		return err
	}
	if payment.orders.Has(newOrderForSearches(item)) {
		return errors.New("order item was already inserted")
	}
	if _, replaced := payment.orders.ReplaceOrInsert(order); replaced {
		panic("UNREACHABLE CODE: already checked order wasn't already inserted!")
	}
	return nil
}
