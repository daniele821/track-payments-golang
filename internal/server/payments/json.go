package payments

import (
	"encoding/json"
)

func NewAllPaymentsFromJson(allPaymentsJson string) (*AllPayments, error) {
	var jsonParsed map[string]any
	err := json.Unmarshal([]byte(allPaymentsJson), &jsonParsed)
	if err != nil {
		return nil, err
	}
	panic("TODO: convert map to allPayments")
}

func (AllPayments *AllPayments) DumpJson(indent bool) (string, error) {
	simplifiedData := map[string]any{}
	simplifiedData["valueSet"] = map[string]any{}
	simplifiedData["payments"] = map[string]any{}

	panic("TODO: convert allPayments to map")

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
