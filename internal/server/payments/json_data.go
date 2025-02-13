package payments

type ValueSetJson struct {
	Cities         []string `json:"cities"`
	Shops          []string `json:"shops"`
	PaymentMethods []string `json:"paymentMethods"`
	Items          []string `json:"items"`
}

type OrderJson struct {
	Quantity  int    `json:"quantity"`
	UnitPrice int    `json:"unitPrice"`
	Item      string `json:"item"`
}

type PaymentJson struct {
	City          string      `json:"city"`
	Shop          string      `json:"shop"`
	PaymentMethod string      `json:"paymentMethod"`
	Date          string      `json:"date"`
	Description   string      `json:"description"`
	Orders        []OrderJson `json:"orders"`
}

type AllPaymentsJson struct {
	Payments []PaymentJson `json:"payments"`
	ValueSet ValueSetJson  `json:"valueSet"`
}
