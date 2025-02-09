package structures

import "github.com/google/btree"

type ValueSet struct {
	cities         []string
	shops          []string
	paymentMethods []string
	items          []string
}

type Order struct {
	quantity  int
	unitPrice int
	item      btree.BTreeG[string]
}

type Payment struct {
	city          string
	shop          string
	paymentMethod string
	date          btree.BTreeG[string]
	description   string
	orders        []Order
}

type AllPayments struct {
	payments []Payment
	valueSet ValueSet
}
