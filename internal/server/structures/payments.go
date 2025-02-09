package structures

import (
	"time"

	"github.com/google/btree"
)

type ValueSet struct {
	cities         *btree.BTreeG[string]
	shops          *btree.BTreeG[string]
	paymentMethods *btree.BTreeG[string]
	items          *btree.BTreeG[string]
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
	date          time.Time
	description   string
	orders        *btree.BTreeG[Order]
}

type AllPayments struct {
	payments *btree.BTreeG[Payment]
	valueSet ValueSet
}

func NewValueSet() ValueSet {
	valueSet := ValueSet{
		cities:         btree.NewG(3, func(a, b string) bool { return a < b }),
		shops:          btree.NewG(3, func(a, b string) bool { return a < b }),
		paymentMethods: btree.NewG(3, func(a, b string) bool { return a < b }),
		items:          btree.NewG(3, func(a, b string) bool { return a < b }),
	}
	return valueSet
}
