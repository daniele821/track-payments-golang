package payments

import "github.com/google/btree"

type valueSet struct {
	cities         *btree.BTreeG[string]
	shops          *btree.BTreeG[string]
	paymentMethods *btree.BTreeG[string]
	items          *btree.BTreeG[string]
}
type ValueSet struct{ pointer *valueSet }

type order struct {
	quantity  uint
	unitPrice uint // is the price in euro cents (2.40 euro => 240)
	item      string
}
type Order struct{ pointer *order }

type payment struct {
	city          string
	shop          string
	paymentMethod string
	date          string // note: i will consider every date inserted as is, without any conversions!
	description   string
	orders        *btree.BTreeG[Order]
}
type Payment struct{ pointer *payment }

type allPayments struct {
	payments *btree.BTreeG[Payment]
	valueSet ValueSet
}
type AllPayments struct{ pointer *allPayments }
