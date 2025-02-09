package structures

import (
	"fmt"
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

// COMPARISON METHODS
func (payment Payment) GreaterThan(otherPayment Payment) bool {
	return payment.date.After(otherPayment.date)
}
func (order Order) GreaterThan(otherOrder Order) bool {
	return order.item > otherOrder.item
}

// STRING METHODS
func (valueSet ValueSet) String() string {
	return "TODO"
}
func (order Order) String() string {
	return fmt.Sprintf("item: %s, quantity: %d, unitPrice: %d", order.item, order.quantity, order.unitPrice)
}
func (payment Payment) String() string {
	return "TODO"
}
func (allPayments AllPayments) String() string {
	return "TODO"
}

// CONSTRUCTORS
func NewAllPayment() AllPayments {
	return AllPayments{
		valueSet: ValueSet{
			cities:         btree.NewG(3, func(a, b string) bool { return a < b }),
			shops:          btree.NewG(3, func(a, b string) bool { return a < b }),
			paymentMethods: btree.NewG(3, func(a, b string) bool { return a < b }),
			items:          btree.NewG(3, func(a, b string) bool { return a < b }),
		},
		payments: btree.NewG(3, func(a, b Payment) bool { return a.GreaterThan(b) }),
	}
}
