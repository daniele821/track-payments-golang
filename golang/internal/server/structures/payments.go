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
	item      string
}

type Payment struct {
	city          string
	shop          string
	paymentMethod string
	date          string
	description   string
	orders        *btree.BTreeG[Order]
}

type AllPayments struct {
	payments *btree.BTreeG[Payment]
	valueSet ValueSet
}
