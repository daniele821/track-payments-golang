package server

import (
	"errors"
	"fmt"
	"strings"
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

// UTILITY METHODS

func parseDate(date string) (time.Time, error) {
	return time.ParseInLocation("2006/01/02 15:04", date, time.UTC)
}

func (payment *Payment) lessThan(otherPayment *Payment) bool {
	return payment.date < otherPayment.date
}

func (order *Order) lessThan(otherOrder *Order) bool {
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

func newOrderForSearches(item string) *Order {
	return &Order{item: item}
}
func newPayment(city, shop, paymentMethod, date, description string) *Payment {
	return &Payment{
		city:          city,
		shop:          shop,
		paymentMethod: paymentMethod,
		date:          date,
		description:   description,
		orders:        btree.NewG(3, func(a, b *Order) bool { return a.lessThan(b) }),
	}
}

func newPaymentForSearches(date string) *Payment {
	return &Payment{date: date}
}

func NewAllPayment() *AllPayments {
	return &AllPayments{
		valueSet: newValueSet(),
		payments: btree.NewG(3, func(a, b *Payment) bool { return a.lessThan(b) }),
	}
}

// INSERT METHODS

func insertAll(valueSet *btree.BTreeG[string], elems ...string) (duplicates []string) {
	for _, elem := range elems {
		if old, replaced := valueSet.ReplaceOrInsert(elem); replaced {
			duplicates = append(duplicates, old)
		}
	}
	if len(duplicates) == 0 {
		return nil
	}
	return duplicates
}

func (allPayments *AllPayments) AddCities(cities ...string) (duplicates []string) {
	return insertAll(allPayments.valueSet.cities, cities...)
}

func (allPayments *AllPayments) AddShops(shops ...string) (duplicates []string) {
	return insertAll(allPayments.valueSet.shops, shops...)
}

func (allPayments *AllPayments) AddPaymentMethods(paymentMethods ...string) (duplicates []string) {
	return insertAll(allPayments.valueSet.paymentMethods, paymentMethods...)
}

func (allPayments *AllPayments) AddItems(items ...string) (duplicates []string) {
	return insertAll(allPayments.valueSet.items, items...)
}

func (allPayments *AllPayments) AddPayment(city, shop, paymentMethod, date, description string) error {
	payment := newPayment(city, shop, paymentMethod, date, description)
	if !allPayments.valueSet.cities.Has(city) {
		return errors.New("invalid city: " + city)
	}
	if !allPayments.valueSet.shops.Has(shop) {
		return errors.New("invalid shop: " + shop)
	}
	if !allPayments.valueSet.paymentMethods.Has(paymentMethod) {
		return errors.New("invalid paymentMethod: " + paymentMethod)
	}
	if _, err := parseDate(date); err != nil {
		return errors.New(fmt.Sprintf("invalid date: %s", err))
	}
	if allPayments.payments.Has(payment) {
		return errors.New("invalid date: already exists")
	}
	if _, replaced := allPayments.payments.ReplaceOrInsert(payment); replaced {
		panic("UNREACHABLE CODE: already check payment wasn't already inserted!")
	}
	return nil
}

func (allPayments *AllPayments) AddOrder(quantity, unitPrice int, item, date string) error {
	order := newOrder(quantity, unitPrice, item)
	oldPayment, found := allPayments.payments.Get(newPaymentForSearches(date))
	if !found {
		return errors.New("payment to insert order into was not found")
	}
	if oldPayment.orders.Has(newOrderForSearches(item)) {
		return errors.New("order item was already inserted")
	}
	if _, replaced := oldPayment.orders.ReplaceOrInsert(order); replaced {
		panic("UNREACHABLE CODE: already checked order wasn't already inserted!")
	}
	return nil
}

// DELETE METHODS

func (allPayments *AllPayments) RemovePayment(date string) error {
	_, found := allPayments.payments.Delete(newPaymentForSearches(date))
	if !found {
		return errors.New("payment was not found")
	}
	return nil
}

func (allPayments *AllPayments) RemoveOrder(date, item string) error {
	payment, foundPayment := allPayments.payments.Get(newPaymentForSearches(date))
	if !foundPayment {
		return errors.New("payment related to the order was not found")
	}
	_, foundOrder := payment.orders.Delete(newOrderForSearches(item))
	if !foundOrder {
		return errors.New("order was not found")
	}
	return nil
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
