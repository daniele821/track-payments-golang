package payments

import "encoding/json"

func NewAllPaymentsFromJson(allPaymentsJson string) (*AllPayments, error) {
	var jsonParsed map[string]any
	err := json.Unmarshal([]byte(allPaymentsJson), &jsonParsed)
	if err != nil {
		return nil, err
	}
	panic("TODO")
}

func (AllPayments *AllPayments) DumpJson(indent bool) string {
	panic("TODO")
}
