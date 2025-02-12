package payments

import (
	"encoding/json"

	"github.com/google/btree"
)

type valueSetJson struct {
	Cities         []string `json:"cities"`
	Shops          []string `json:"shops"`
	PaymentMethods []string `json:"paymentMethods"`
	Items          []string `json:"items"`
}

type orderJson struct {
	Quantity  uint   `json:"quantity"`
	UnitPrice uint   `json:"unitPrice"`
	Item      string `json:"item"`
}

type paymentJson struct {
	City          string      `json:"city"`
	Shop          string      `json:"shop"`
	PaymentMethod string      `json:"paymentMethod"`
	Date          string      `json:"date"`
	Description   string      `json:"description"`
	Orders        []orderJson `json:"orders"`
}

type allPaymentsJson struct {
	Payments []paymentJson `json:"payments"`
	ValueSet valueSetJson  `json:"valueSet"`
}

func btreeToSlice[T, S any](data *btree.BTreeG[T], mapper func(item T) S) []S {
	acc := make([]S, data.Len())
	index := 0
	data.Ascend(func(item T) bool {
		acc[index] = mapper(item)
		index += 1
		return true
	})
	return acc
}

func mapperIdentity[T any](item T) T {
	return item
}
func mapperOrderJson(item Order) orderJson {
	return orderJson{
		Quantity:  item.Quantity(),
		UnitPrice: item.Quantity(),
		Item:      item.Item(),
	}
}
func mapperPaymentJson(item Payment) paymentJson {
	return paymentJson{
		City:          item.City(),
		Shop:          item.Shop(),
		PaymentMethod: item.PaymentMethod(),
		Date:          item.Date(),
		Description:   item.Description(),
		Orders:        btreeToSlice(item.p.orders, mapperOrderJson),
	}
}

// TO JSON

func convertToJsonData(input AllPayments) allPaymentsJson {
	return allPaymentsJson{
		ValueSet: valueSetJson{
			Cities:         btreeToSlice(input.p.valueSet.p.cities, mapperIdentity),
			Shops:          btreeToSlice(input.p.valueSet.p.shops, mapperIdentity),
			PaymentMethods: btreeToSlice(input.p.valueSet.p.paymentMethods, mapperIdentity),
			Items:          btreeToSlice(input.p.valueSet.p.items, mapperIdentity),
		},
		Payments: btreeToSlice(input.p.payments, mapperPaymentJson),
	}
}
func (allPayments AllPayments) DumpJson(indent bool) (string, error) {
	allPaymentsJson := convertToJsonData(allPayments)
	var jsonRes []byte
	var err error
	if indent {
		jsonRes, err = json.MarshalIndent(allPaymentsJson, "", "  ")
	} else {
		jsonRes, err = json.Marshal(allPaymentsJson)
	}
	if err != nil {
		return "", err
	}
	return string(jsonRes), nil
}

// FROM JSON

func convertFromJsonData(input allPaymentsJson) (AllPayments, error) {
	output := NewAllPayments()
	outputEmpty := NewAllPayments()
	if err := output.AddCities(input.ValueSet.Cities...); err != nil {
		return outputEmpty, err
	}
	if err := output.AddShops(input.ValueSet.Shops...); err != nil {
		return outputEmpty, err
	}
	if err := output.AddPaymentMethods(input.ValueSet.PaymentMethods...); err != nil {
		return outputEmpty, err
	}
	if err := output.AddItems(input.ValueSet.Items...); err != nil {
		return outputEmpty, err
	}
	for _, payment := range input.Payments {
		date := payment.Date
		if err := output.AddPayment(payment.City, payment.Shop, payment.PaymentMethod, date, payment.Description); err != nil {
			return outputEmpty, err
		}
		for _, order := range payment.Orders {
			if err := output.AddOrder(order.Quantity, order.UnitPrice, order.Item, date); err != nil {
				return outputEmpty, err
			}
		}
	}
	return output, nil
}
func NewAllPaymentsFromJson(input string) (AllPayments, error) {
	data := allPaymentsJson{}
	outputEmpty := NewAllPayments()
	if err := json.Unmarshal([]byte(input), &data); err != nil {
		return outputEmpty, err
	}
	output, err := convertFromJsonData(data)
	if err != nil {
		return outputEmpty, err
	}
	return output, nil
}
