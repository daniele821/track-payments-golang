package server

import (
	"errors"
	"fmt"
	"strings"

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
	unitPrice int // is the price in euro cents (2.40 euro => 240)
	item      string
}

type Payment struct {
	city          string
	shop          string
	paymentMethod string
	date          string // note: i will consider every date inserted as is, without any conversions!
	description   string
	orders        *btree.BTreeG[*Order]
}

type AllPayments struct {
	payments *btree.BTreeG[*Payment]
	valueSet *ValueSet
}

// COMPARISON METHODS

func (payment *Payment) LessThan(otherPayment *Payment) bool {
	return payment.date < otherPayment.date
}

func (order *Order) LessThan(otherOrder *Order) bool {
	return order.item < otherOrder.item
}

// CONSTRUCTORS

func newValueSet() *ValueSet {
	return &ValueSet{
		cities:         btree.NewG(3, func(a, b string) bool { return a < b }),
		shops:          btree.NewG(3, func(a, b string) bool { return a < b }),
		paymentMethods: btree.NewG(3, func(a, b string) bool { return a < b }),
		items:          btree.NewG(3, func(a, b string) bool { return a < b }),
	}
}

func newOrder(quantity, unitPrice int, item string) *Order {
	return &Order{
		quantity:  quantity,
		unitPrice: unitPrice,
		item:      item,
	}
}

func newPayment(city, shop, paymentMethod, date, description string) *Payment {
	return &Payment{
		city:          city,
		shop:          shop,
		paymentMethod: paymentMethod,
		date:          date,
		description:   description,
	}
}

func NewAllPayment() *AllPayments {
	return &AllPayments{
		valueSet: newValueSet(),
		payments: btree.NewG(3, func(a, b *Payment) bool { return a.LessThan(b) }),
	}
}

// INSERT/DELETE METHODS

func insertAll(elem string, valueSet *btree.BTreeG[string], elems ...string) error {
	duplicates := []string{}
	for _, elem := range elems {
		if valueSet.Has(elem) {
			duplicates = append(duplicates, elem)
		}
	}
	if len(duplicates) != 0 {
		return errors.New("duplicated " + elem + ": " + strings.Join(duplicates, ", "))
	}
	for _, elem := range elems {
		if _, dup := valueSet.ReplaceOrInsert(elem); dup {
			panic("UNREACHABLE CODE: duplicate shouldn't be possible!")
		}
	}
	return nil
}

func (allPayments *AllPayments) AddCities(cities ...string) error {
	return insertAll("cities", allPayments.valueSet.cities, cities...)
}

func (allPayments *AllPayments) AddShops(shops ...string) error {
	return insertAll("shops", allPayments.valueSet.shops, shops...)
}

func (allPayments *AllPayments) AddPaymentMethods(paymentMethods ...string) error {
	return insertAll("paymentMethods", allPayments.valueSet.paymentMethods, paymentMethods...)
}

func (allPayments *AllPayments) AddItems(items ...string) error {
	return insertAll("items", allPayments.valueSet.items, items...)
}

// STRING METHODS

func fmtBtree[T any](btree *btree.BTreeG[T], strconv func(item T) string) string {
	acc := []string{}
	btree.Ascend(func(item T) bool {
		acc = append(acc, strconv(item))
		return true
	})
	return "[" + strings.Join(acc, " ") + "]"
}

func (valueSet *ValueSet) String() string {
	cityStr := fmtBtree(valueSet.cities, func(item string) string { return item })
	shopStr := fmtBtree(valueSet.shops, func(item string) string { return item })
	payMStr := fmtBtree(valueSet.paymentMethods, func(item string) string { return item })
	itemStr := fmtBtree(valueSet.items, func(item string) string { return item })
	return fmt.Sprintf("(cities: %s, shops: %s, paymentMethods: %s, items: %s)", cityStr, shopStr, payMStr, itemStr)
}

func (order *Order) String() string {
	return fmt.Sprintf("(item: %s, quantity: %d, unitPrice: %d)", order.item, order.quantity, order.unitPrice)
}

func (payment *Payment) String() string {
	ordersStr := fmtBtree(payment.orders, func(item *Order) string { return item.String() })
	return fmt.Sprintf("(city: %s, shop: %s, paymentMethod: %s, date: %s, description: %s, orders: %s)",
		payment.city, payment.shop, payment.paymentMethod, payment.date, payment.description, ordersStr)
}

func (allPayments *AllPayments) String() string {
	paymentStr := fmtBtree(allPayments.payments, func(item *Payment) string { return item.String() })
	valueSetStr := allPayments.valueSet.String()
	return fmt.Sprintf("(payments: %s, valueSet: %s)", paymentStr, valueSetStr)
}
