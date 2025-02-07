package payments

import (
	"errors"
	"slices"
	"time"
)

type valueSet struct {
	cities        []string
	shops         []string
	paymentMethod []string
	categories    []string
	itemCat       map[string]string
}

type order struct {
	quantity  int
	unitPrice int
	item      string
}

type payment struct {
	city          string
	shop          string
	paymentMethod string
	date          time.Time
	orders        []order
}

type allPayments struct {
	payments []payment
	valueSet valueSet
}

func newValueSet(cities, shops, methods, categories []string, itemCat map[string]string) (valueSet, error) {
	valueSet := valueSet{
		cities:        cities,
		shops:         shops,
		paymentMethod: methods,
		categories:    categories,
		itemCat:       itemCat,
	}
	for _, category := range itemCat {
		if !slices.Contains(categories, category) {
			return valueSet, errors.New("invalid category!")
		}
	}
	return valueSet, nil
}

func newAllPayments(valueSet valueSet) allPayments {
	return allPayments{
		valueSet: valueSet,
	}
}

func (payments *allPayments) addPayment(city, shop, paymentMethod string, date time.Time) error {
	if !slices.Contains(payments.valueSet.cities, city) {
		return errors.New("invalid city!")
	}
	if !slices.Contains(payments.valueSet.shops, shop) {
		return errors.New("invalid shop!")
	}
	if !slices.Contains(payments.valueSet.paymentMethod, paymentMethod) {
		return errors.New("invalid payment method!")
	}
	payments.payments = append(payments.payments, payment{
		city: city, shop: shop, paymentMethod: paymentMethod, date: date, orders: []order{},
	})
	return nil
}

func (payments *allPayments) addOrder(paymentIndex, quantity, unitPrice int, item string) error {
	if _, ok := payments.valueSet.itemCat[item]; !ok {
		return errors.New("invalid item!")
	}
	if paymentIndex <= 0 || paymentIndex >= len(payments.payments) {
		return errors.New("invalid index!")
	}
	payments.payments[paymentIndex].orders = append(payments.payments[paymentIndex].orders, order{
		quantity: quantity, unitPrice: unitPrice, item: item,
	})
	return nil
}
