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

func (allPayments *allPayments) addPayment(city, shop, paymentMethod string, date time.Time) error {
	if !slices.Contains(allPayments.valueSet.cities, city) {
		return errors.New("invalid city!")
	}
	if !slices.Contains(allPayments.valueSet.shops, shop) {
		return errors.New("invalid shop!")
	}
	if !slices.Contains(allPayments.valueSet.paymentMethod, paymentMethod) {
		return errors.New("invalid payment method!")
	}
	allPayments.payments = append(allPayments.payments, payment{
		city: city, shop: shop, paymentMethod: paymentMethod, date: date, orders: []order{},
	})
	return nil
}

func (allPayments *allPayments) removePayment(paymentIndex int) error {
	if paymentIndex < 0 || paymentIndex >= len(allPayments.payments) {
		return errors.New("invalid index!")
	}
	allPayments.payments = slices.Delete(allPayments.payments, paymentIndex, paymentIndex+1)
	return nil
}

func (allPayments *allPayments) addOrder(paymentIndex, quantity, unitPrice int, item string) error {
	if _, ok := allPayments.valueSet.itemCat[item]; !ok {
		return errors.New("invalid item!")
	}
	if paymentIndex < 0 || paymentIndex >= len(allPayments.payments) {
		return errors.New("invalid index!")
	}
	for _, order := range allPayments.payments[paymentIndex].orders {
		if order.item == item {
			return errors.New("duplicate item!")
		}
	}
	allPayments.payments[paymentIndex].orders = append(allPayments.payments[paymentIndex].orders, order{
		quantity: quantity, unitPrice: unitPrice, item: item,
	})
	return nil
}

func (allPayments *allPayments) removeOrder(paymentIndex int, item string) error {
	if paymentIndex < 0 || paymentIndex >= len(allPayments.payments) {
		return errors.New("invalid index!")
	}
	allPayments.payments[paymentIndex].orders = slices.DeleteFunc(allPayments.payments[paymentIndex].orders, func(elem order) bool { return elem.item == item })
	return nil
}
