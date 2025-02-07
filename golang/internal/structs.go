package payments

import (
	"encoding/json"
	"errors"
	"payment/internal/utils"
	"slices"
	"time"
)

type valueSet struct {
	Cities         []string          `json:"cities"`
	Shops          []string          `json:"shops"`
	PaymentMethods []string          `json:"paymentMethods"`
	Categories     []string          `json:"categories"`
	Items          map[string]string `json:"items"`
}

type order struct {
	Quantity  int    `json:"quantity"`
	UnitPrice int    `json:"unitPrice"`
	Item      string `json:"item"`
}

type payment struct {
	City          string    `json:"city"`
	Shop          string    `json:"shop"`
	PaymentMethod string    `json:"paymentMethod"`
	Date          time.Time `json:"date"`
	Orders        []order   `json:"orders"`
}

type allPayments struct {
	Payments []payment `json:"payments"`
	ValueSet valueSet  `json:"valueSet"`
}

func newAllPaymentsFromJson(paymentsJson string) (allPayments, error) {
	var tmpPayments allPayments
	err := json.Unmarshal([]byte(paymentsJson), &tmpPayments)
	if err != nil {
		return allPayments{}, err
	}

	// do checks: manually rebuild the allPayments, to have automagical checks!
	tmpVal := tmpPayments.ValueSet
	valueSet, err := newValueSet(tmpVal.Cities, tmpVal.Shops, tmpVal.PaymentMethods, tmpVal.Categories, tmpVal.Items)
	if err != nil {
		return allPayments{}, err
	}
	payments := newAllPayments(valueSet)
	for indexPayment, payment := range tmpPayments.Payments {
		if err := payments.addPayment(payment.City, payment.Shop, payment.PaymentMethod, payment.Date); err != nil {
			return payments, err
		}
		for _, order := range tmpPayments.Payments[indexPayment].Orders {
			if err := payments.addOrder(indexPayment, order.Quantity, order.UnitPrice, order.Item); err != nil {
				return payments, err
			}
		}
	}

	return payments, nil
}

func (allPayments allPayments) generateJson(indent bool) (string, error) {
	if indent {
		paymentJson, err := json.MarshalIndent(allPayments, "", "  ")
		if err != nil {
			return "", err
		}
		return string(paymentJson), nil
	}
	paymentJson, err := json.Marshal(allPayments)
	if err != nil {
		return "", err
	}
	return string(paymentJson), nil
}

func newValueSet(cities, shops, methods, categories []string, itemCat map[string]string) (valueSet, error) {
	valueSet := valueSet{
		Cities:         cities,
		Shops:          shops,
		PaymentMethods: methods,
		Categories:     categories,
		Items:          itemCat,
	}
	if utils.HasDuplicates(cities) {
		return valueSet, errors.New("invalid cities: there are duplicated values!")
	}
	if utils.HasDuplicates(shops) {
		return valueSet, errors.New("invalid shops: there are duplicated values!")
	}
	if utils.HasDuplicates(methods) {
		return valueSet, errors.New("invalid methods: there are duplicated values!")
	}
	if utils.HasDuplicates(categories) {
		return valueSet, errors.New("invalid categories: there are duplicated values!")
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
		ValueSet: valueSet,
	}
}

func (allPayments *allPayments) addCity(city string) error {
	if slices.Contains(allPayments.ValueSet.Cities, city) {
		return errors.New("invalid city: already present in the valueset!")
	}
	allPayments.ValueSet.Cities = append(allPayments.ValueSet.Cities, city)
	return nil
}

func (allPayments *allPayments) addShop(shop string) error {
	if slices.Contains(allPayments.ValueSet.Shops, shop) {
		return errors.New("invalid shop: already present in the valueset!")
	}
	allPayments.ValueSet.Shops = append(allPayments.ValueSet.Shops, shop)
	return nil
}

func (allPayments *allPayments) addPaymentMethod(paymentMethod string) error {
	if slices.Contains(allPayments.ValueSet.PaymentMethods, paymentMethod) {
		return errors.New("invalid payment method: already present in the valueset!")
	}
	allPayments.ValueSet.PaymentMethods = append(allPayments.ValueSet.PaymentMethods, paymentMethod)
	return nil
}

func (allPayments *allPayments) addCategory(category string) error {
	if slices.Contains(allPayments.ValueSet.Categories, category) {
		return errors.New("invalid category: already present in the valueset!")
	}
	allPayments.ValueSet.Categories = append(allPayments.ValueSet.Categories, category)
	return nil
}

func (allPayments *allPayments) addItem(item, category string) error {
	if _, ok := allPayments.ValueSet.Items[item]; ok {
		return errors.New("invalid item: already present in the valueset!")
	}
	if !slices.Contains(allPayments.ValueSet.Categories, category) {
		return errors.New("invalid category: not in the valueset!")
	}
	if allPayments.ValueSet.Items == nil {
		allPayments.ValueSet.Items = map[string]string{}
	}
	allPayments.ValueSet.Items[item] = category
	return nil
}

func (allPayments *allPayments) addPayment(city, shop, paymentMethod string, date time.Time) error {
	if !slices.Contains(allPayments.ValueSet.Cities, city) {
		return errors.New("invalid city: not in the valueset!")
	}
	if !slices.Contains(allPayments.ValueSet.Shops, shop) {
		return errors.New("invalid shop: not in the valueset!")
	}
	if !slices.Contains(allPayments.ValueSet.PaymentMethods, paymentMethod) {
		return errors.New("invalid payment payment method: not in the valueset!")
	}
	if date.After(time.Now()) {
		return errors.New("invalid date: date in the future!")
	}
	allPayments.Payments = append(allPayments.Payments, payment{
		City: city, Shop: shop, PaymentMethod: paymentMethod, Date: date.UTC(), Orders: []order{},
	})
	return nil
}

func (allPayments *allPayments) removePayment(paymentIndex int) error {
	if paymentIndex < 0 || paymentIndex >= len(allPayments.Payments) {
		return errors.New("invalid index: out of bound!")
	}
	allPayments.Payments = slices.Delete(allPayments.Payments, paymentIndex, paymentIndex+1)
	return nil
}

func (allPayments *allPayments) addOrder(paymentIndex, quantity, unitPrice int, item string) error {
	if _, ok := allPayments.ValueSet.Items[item]; !ok {
		return errors.New("invalid item: not in the valueset!")
	}
	if paymentIndex < 0 || paymentIndex >= len(allPayments.Payments) {
		return errors.New("invalid index: out of bound!")
	}
	for _, order := range allPayments.Payments[paymentIndex].Orders {
		if order.Item == item {
			return errors.New("invalid item: duplicate value!")
		}
	}
	allPayments.Payments[paymentIndex].Orders = append(allPayments.Payments[paymentIndex].Orders, order{
		Quantity: quantity, UnitPrice: unitPrice, Item: item,
	})
	return nil
}

func (allPayments *allPayments) removeOrder(paymentIndex int, item string) error {
	if paymentIndex < 0 || paymentIndex >= len(allPayments.Payments) {
		return errors.New("invalid index: out of bounds!")
	}
	allPayments.Payments[paymentIndex].Orders = slices.DeleteFunc(allPayments.Payments[paymentIndex].Orders, func(elem order) bool { return elem.Item == item })
	return nil
}
