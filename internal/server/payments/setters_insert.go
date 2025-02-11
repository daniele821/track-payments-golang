package payments

import (
	"errors"
	"fmt"
	"strings"

	"github.com/google/btree"
)

func insertAll(typeData string, valueSet *btree.BTreeG[string], elems ...string) error {
	duplicates := []string{}
	inserted := []string{}
	for _, elem := range elems {
		if _, replaced := valueSet.ReplaceOrInsert(elem); replaced {
			duplicates = append(duplicates, elem)
		} else {
			inserted = append(inserted, elem)
		}
	}
	if len(duplicates) == 0 {
		return nil
	}
	// revert changes on error
	for _, elem := range inserted {
		if _, found := valueSet.Delete(elem); !found {
			panic("UNREACHABLE CODE: an element though to be present, was actually not!")
		}
	}
	return errors.New(fmt.Sprintf("invalid %s: there are duplicates (%s)", typeData, strings.Join(duplicates, ", ")))
}

func (allPayments *AllPayments) AddCities(cities ...string) error {
	return insertAll("cities", allPayments.valueSet.cities, cities...)
}

func (allPayments *AllPayments) AddShops(shops ...string) error {
	return insertAll("shops", allPayments.valueSet.shops, shops...)
}

func (allPayments *AllPayments) AddPaymentMethods(paymentMethods ...string) error {
	return insertAll("paymentMethods", allPayments.valueSet.paymentMethods, paymentMethods...)
}

func (allPayments *AllPayments) AddItems(items ...string) error {
	return insertAll("items", allPayments.valueSet.items, items...)
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
	payment, err := allPayments.Payment(date)
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
