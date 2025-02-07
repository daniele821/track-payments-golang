package payments

import (
	"errors"
	"slices"
	"time"
)

type valueSet struct {
	cities         []string
	shops          []string
	paymentMethods []string
	categories     []string
	itemCat        map[string]string
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
		cities:         cities,
		shops:          shops,
		paymentMethods: methods,
		categories:     categories,
		itemCat:        itemCat,
	}
	for _, category := range itemCat {
		if !slices.Contains(categories, category) {
			return valueSet, errors.New("invalid category: not in the valueset!")
		}
	}
	return valueSet, nil
}

func newAllPayments(valueSet valueSet) allPayments {
	return allPayments{
		valueSet: valueSet,
	}
}

func (allPayments *allPayments) addCity(city string) error {
	if !slices.Contains(allPayments.valueSet.cities, city) {
		return errors.New("invalid city: already present in the valueset!")
	}
	allPayments.valueSet.cities = append(allPayments.valueSet.cities, city)
	return nil
}

func (allPayments *allPayments) addShop(shop string) error {
	if !slices.Contains(allPayments.valueSet.cities, shop) {
		return errors.New("invalid shop: already present in the valueset!")
	}
	allPayments.valueSet.shops = append(allPayments.valueSet.shops, shop)
	return nil
}

func (allPayments *allPayments) addPaymentMethod(paymentMethod string) error {
	if !slices.Contains(allPayments.valueSet.cities, paymentMethod) {
		return errors.New("invalid payment method: already present in the valueset!")
	}
	allPayments.valueSet.paymentMethods = append(allPayments.valueSet.paymentMethods, paymentMethod)
	return nil
}

func (allPayments *allPayments) addCategory(category string) error {
	if !slices.Contains(allPayments.valueSet.cities, category) {
		return errors.New("invalid category: already present in the valueset!")
	}
	allPayments.valueSet.categories = append(allPayments.valueSet.categories, category)
	return nil
}

func (allPayments *allPayments) addPayment(city, shop, paymentMethod string, date time.Time) error {
	if !slices.Contains(allPayments.valueSet.cities, city) {
		return errors.New("invalid city: not in the valueset!")
	}
	if !slices.Contains(allPayments.valueSet.shops, shop) {
		return errors.New("invalid shop: not in the valueset!")
	}
	if !slices.Contains(allPayments.valueSet.paymentMethods, paymentMethod) {
		return errors.New("invalid payment payment method: not in the valueset!")
	}
	if date.After(time.Now()) {
		return errors.New("invalid date: date in the future!")
	}
	allPayments.payments = append(allPayments.payments, payment{
		city: city, shop: shop, paymentMethod: paymentMethod, date: date, orders: []order{},
	})
	return nil
}

func (allPayments *allPayments) removePayment(paymentIndex int) error {
	if paymentIndex < 0 || paymentIndex >= len(allPayments.payments) {
		return errors.New("invalid index: out of bound!")
	}
	allPayments.payments = slices.Delete(allPayments.payments, paymentIndex, paymentIndex+1)
	return nil
}

func (allPayments *allPayments) addOrder(paymentIndex, quantity, unitPrice int, item string) error {
	if _, ok := allPayments.valueSet.itemCat[item]; !ok {
		return errors.New("invalid item: not in the valueset!")
	}
	if paymentIndex < 0 || paymentIndex >= len(allPayments.payments) {
		return errors.New("invalid index: out of bound!")
	}
	for _, order := range allPayments.payments[paymentIndex].orders {
		if order.item == item {
			return errors.New("invalid item: duplicate value!")
		}
	}
	allPayments.payments[paymentIndex].orders = append(allPayments.payments[paymentIndex].orders, order{
		quantity: quantity, unitPrice: unitPrice, item: item,
	})
	return nil
}

func (allPayments *allPayments) removeOrder(paymentIndex int, item string) error {
	if paymentIndex < 0 || paymentIndex >= len(allPayments.payments) {
		return errors.New("invalid index: out of bounds!")
	}
	allPayments.payments[paymentIndex].orders = slices.DeleteFunc(allPayments.payments[paymentIndex].orders, func(elem order) bool { return elem.item == item })
	return nil
}
