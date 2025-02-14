package payments

type valueSetJson struct {
	Cities         []string `json:"cities"`
	Shops          []string `json:"shops"`
	PaymentMethods []string `json:"paymentMethods"`
	Items          []string `json:"items"`
}

type orderJson struct {
	Quantity  int    `json:"quantity"`
	UnitPrice int    `json:"unitPrice"`
	Item      string `json:"item"`
}

type paymentJson struct {
	City          string      `json:"city"`
	Shop          string      `json:"shop"`
	PaymentMethod string      `json:"paymentMethod"`
	Date          string      `json:"date"`
	Orders        []orderJson `json:"orders"`
}

type allPaymentsJson struct {
	Payments []paymentJson `json:"payments"`
	ValueSet valueSetJson  `json:"valueSet"`
}
