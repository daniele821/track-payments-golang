package payments

import (
	"encoding/json"
)

func ConvertToJsonData(input AllPayments) allPaymentsJson {
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
	allPaymentsJson := ConvertToJsonData(allPayments)
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
