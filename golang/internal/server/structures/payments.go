package structures

import (
	"encoding/json"
	"errors"
	"fmt"
	"payment/internal/utils"
	"slices"
	"time"
)

type ValueSet struct {
	Cities         []string          `json:"cities"`
	Shops          []string          `json:"shops"`
	PaymentMethods []string          `json:"paymentMethods"`
	Categories     []string          `json:"categories"`
	Items          map[string]string `json:"items"`
}

type Order struct {
	Quantity  int    `json:"quantity"`
	UnitPrice int    `json:"unitPrice"`
	Item      string `json:"item"`
}

type Payment struct {
	City          string  `json:"city"`
	Shop          string  `json:"shop"`
	PaymentMethod string  `json:"paymentMethod"`
	Date          string  `json:"date"`
	Orders        []Order `json:"orders"`
}

type AllPayments struct {
	Payments []Payment `json:"payments"`
	ValueSet ValueSet  `json:"valueSet"`
}

func NewAllPaymentsFromJson(paymentsJson string) (AllPayments, error) {
	var tmpPayments AllPayments
	err := json.Unmarshal([]byte(paymentsJson), &tmpPayments)
	if err != nil {
		return AllPayments{}, err
	}

	// do checks: manually rebuild the allPayments, to have automagical checks!
	tmpVal := tmpPayments.ValueSet
	valueSet, err := NewValueSet(tmpVal.Cities, tmpVal.Shops, tmpVal.PaymentMethods, tmpVal.Categories, tmpVal.Items)
	if err != nil {
		return AllPayments{}, err
	}
	payments := NewAllPayments(valueSet)
	for indexPayment, payment := range tmpPayments.Payments {
		if err := payments.AddPayment(payment.City, payment.Shop, payment.PaymentMethod, payment.Date); err != nil {
			return payments, err
		}
		for _, order := range tmpPayments.Payments[indexPayment].Orders {
			if err := payments.AddOrder(indexPayment, order.Quantity, order.UnitPrice, order.Item); err != nil {
				return payments, err
			}
		}
	}

	return payments, nil
}

func (allPayments AllPayments) GenerateJson(indent bool) (string, error) {
	var paymentJson []byte
	var err error
	if indent {
		paymentJson, err = json.MarshalIndent(allPayments, "", "  ")
	} else {
		paymentJson, err = json.Marshal(allPayments)
	}
	return string(paymentJson), err
}

func NewValueSet(cities, shops, methods, categories []string, itemCat map[string]string) (ValueSet, error) {
	valueSet := ValueSet{
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

func NewAllPayments(valueSet ValueSet) AllPayments {
	return AllPayments{
		ValueSet: valueSet,
	}
}

func (allPayments *AllPayments) AddCity(city string) error {
	if slices.Contains(allPayments.ValueSet.Cities, city) {
		return errors.New("invalid city: already present in the valueset!")
	}
	allPayments.ValueSet.Cities = append(allPayments.ValueSet.Cities, city)
	return nil
}

func (allPayments *AllPayments) AddShop(shop string) error {
	if slices.Contains(allPayments.ValueSet.Shops, shop) {
		return errors.New("invalid shop: already present in the valueset!")
	}
	allPayments.ValueSet.Shops = append(allPayments.ValueSet.Shops, shop)
	return nil
}

func (allPayments *AllPayments) AddPaymentMethod(paymentMethod string) error {
	if slices.Contains(allPayments.ValueSet.PaymentMethods, paymentMethod) {
		return errors.New("invalid payment method: already present in the valueset!")
	}
	allPayments.ValueSet.PaymentMethods = append(allPayments.ValueSet.PaymentMethods, paymentMethod)
	return nil
}

func (allPayments *AllPayments) AddCategory(category string) error {
	if slices.Contains(allPayments.ValueSet.Categories, category) {
		return errors.New("invalid category: already present in the valueset!")
	}
	allPayments.ValueSet.Categories = append(allPayments.ValueSet.Categories, category)
	return nil
}

func (allPayments *AllPayments) AddItem(item, category string) error {
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

func (allPayments *AllPayments) AddPayment(city, shop, paymentMethod string, date string) error {
	if !slices.Contains(allPayments.ValueSet.Cities, city) {
		return errors.New("invalid city: not in the valueset!")
	}
	if !slices.Contains(allPayments.ValueSet.Shops, shop) {
		return errors.New("invalid shop: not in the valueset!")
	}
	if !slices.Contains(allPayments.ValueSet.PaymentMethods, paymentMethod) {
		return errors.New("invalid payment payment method: not in the valueset!")
	}
	if dateTime, err := time.Parse("2026/01/02 15:04", date); err != nil {
		return errors.New(fmt.Sprintf("invalide date format: %s", err))
	} else if dateTime.After(time.Now()) {
		return errors.New("invalid date: date in the future!")
	}
	// add check date is unique
	allPayments.Payments = append(allPayments.Payments, Payment{
		City: city, Shop: shop, PaymentMethod: paymentMethod, Date: date, Orders: []Order{},
	})
	return nil
}

func (allPayments *AllPayments) RemovePayment(date string) error {
	// use the date as the ID
	// if paymentIndex < 0 || paymentIndex >= len(allPayments.Payments) {
	// 	return errors.New("invalid index: out of bound!")
	// }
	// allPayments.Payments = slices.Delete(allPayments.Payments, paymentIndex, paymentIndex+1)
	return nil
}

func (allPayments *AllPayments) AddOrder(quantity, unitPrice int, date, item string) error {
	// search by date
	// if _, ok := allPayments.ValueSet.Items[item]; !ok {
	// 	return errors.New("invalid item: not in the valueset!")
	// }
	// if paymentIndex < 0 || paymentIndex >= len(allPayments.Payments) {
	// 	return errors.New("invalid index: out of bound!")
	// }
	// for _, order := range allPayments.Payments[paymentIndex].Orders {
	// 	if order.Item == item {
	// 		return errors.New("invalid item: duplicate value!")
	// 	}
	// }
	// allPayments.Payments[paymentIndex].Orders = append(allPayments.Payments[paymentIndex].Orders, Order{
	// 	Quantity: quantity, UnitPrice: unitPrice, Item: item,
	// })
	return nil
}

func (allPayments *AllPayments) RemoveOrder(date, item string) error {
	// search by date
	// if paymentIndex < 0 || paymentIndex >= len(allPayments.Payments) {
	// 	return errors.New("invalid index: out of bounds!")
	// }
	// allPayments.Payments[paymentIndex].Orders = slices.DeleteFunc(allPayments.Payments[paymentIndex].Orders, func(elem Order) bool { return elem.Item == item })
	return nil
}
