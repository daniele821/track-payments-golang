package payments

type valueSet struct {
	cities        []string
	shops         []string
	paymentMethod []string
	items         []string
	category      []string
	itemCat       map[string]string
}

type order struct {
	quantity   int
	unit_price int
	item       string
	category   string
	valueSet   valueSet
}

type payment struct {
	city          string
	shop          string
	paymentMethod string
	orders        []order
	valueSet      valueSet
}
