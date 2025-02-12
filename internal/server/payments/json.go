package payments

import (
	"encoding/json"

	"github.com/google/btree"
)

func btreeToSlice[T, E any](data *btree.BTreeG[T], mapper func(item T) E) []E {
	acc := make([]E, data.Len())
	index := 0
	data.Ascend(func(item T) bool {
		acc[index] = mapper(item)
		index += 1
		return true
	})
	return acc
}

func convertOrder(item Order) map[string]any {
	res := map[string]any{}
	res["quantity"] = item.Quantity()
	res["item"] = item.Item()
	res["unitPrice"] = item.UnitPrice()
	return res
}

func convertPayment(item Payment) map[string]any {
	res := map[string]any{}
	res["city"] = item.City()
	res["shop"] = item.Shop()
	res["paymentMethod"] = item.PaymentMethod()
	res["date"] = item.Date()
	res["description"] = item.Description()
	res["orders"] = btreeToSlice(item.pointer.orders, convertOrder)
	return res
}

func NewAllPaymentsFromJson(allPaymentsJson string) (*AllPayments, error) {
	var jsonParsed any
	err := json.Unmarshal([]byte(allPaymentsJson), &jsonParsed)
	if err != nil {
		return nil, err
	}
	panic("TODO: convert map to allPayments")
}

func (AllPayments *AllPayments) DumpJson(indent bool) (string, error) {
	simplifiedData := map[string]any{}
	simplifiedData["valueSet"] = map[string]any{
		"cities":         btreeToSlice(AllPayments.valueSet.cities, func(item string) string { return item }),
		"shops":          btreeToSlice(AllPayments.valueSet.shops, func(item string) string { return item }),
		"paymentMethods": btreeToSlice(AllPayments.valueSet.paymentMethods, func(item string) string { return item }),
		"items":          btreeToSlice(AllPayments.valueSet.items, func(item string) string { return item }),
	}
	simplifiedData["payments"] = btreeToSlice(AllPayments.payments, convertPayment)

	// panic("TODO: convert allPayments to map")

	var jsonRes []byte
	var err error
	if indent {
		jsonRes, err = json.MarshalIndent(simplifiedData, "", "  ")
	} else {
		jsonRes, err = json.Marshal(simplifiedData)
	}
	if err != nil {
		return "", err
	}
	return string(jsonRes), nil
}
