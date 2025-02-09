package structures

type ValueSet struct {
	cities         []*string
	shops          []*string
	paymentMethods []*string
	items          []*string
}

type Order struct {
	quantity  int
	unitPrice int
	item      *string
}

type Payment struct {
	city          *string
	shop          *string
	paymentMethod *string
	date          string
	description   string
	orders        []Order
}

type AllPayments struct {
	payments []Payment
	valueSet ValueSet
}
