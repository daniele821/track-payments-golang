package payments

import (
	"encoding/json"
)

func ConvertFromJsonData(input allPaymentsJson) (AllPayments, error) {
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
		if err := output.AddPayment(payment.City, payment.Shop, payment.PaymentMethod, date); err != nil {
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
	output, err := ConvertFromJsonData(data)
	if err != nil {
		return outputEmpty, err
	}
	return output, nil
}
