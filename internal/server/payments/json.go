package payments

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
	City          string  `json:"city"`
	Shop          string  `json:"shop"`
	PaymentMethod string  `json:"paymentMethod"`
	Date          string  `json:"date"`
	Description   string  `json:"description"`
	Orders        []Order `json:"orders"`
}

type allPaymentsJson struct {
	Payments []Payment `json:"payments"`
	ValueSet ValueSet  `json:"valueSet"`
}

func convertToJsonData(input AllPayments) (output allPaymentsJson) {
	return output
}

func convertFromJsonData(input allPaymentsJson) (output AllPayments, err error) {
	return output, nil
}

func NewAllPaymentsFromJson(allPaymentsJson string) (AllPayments, error) {
	panic("TODO: convert map to allPayments")
}

func (AllPayments AllPayments) DumpJson(indent bool) (string, error) {
	panic("TODO: convert allPayments to map")
}
