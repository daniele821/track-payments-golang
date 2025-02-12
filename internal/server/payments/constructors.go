package payments

import "github.com/google/btree"

func newValueSet() ValueSet {
	return ValueSet{&valueSet{
		cities:         btree.NewG(3, func(a, b string) bool { return a < b }),
		shops:          btree.NewG(3, func(a, b string) bool { return a < b }),
		paymentMethods: btree.NewG(3, func(a, b string) bool { return a < b }),
		items:          btree.NewG(3, func(a, b string) bool { return a < b }),
	}}
}

func newOrder(quantity, unitPrice uint, item string) Order {
	return Order{&order{
		quantity:  quantity,
		unitPrice: unitPrice,
		item:      item,
	}}
}

func NewOrderForSearches(item string) Order {
	return Order{&order{item: item}}
}

func newPayment(city, shop, paymentMethod, date, description string) Payment {
	return Payment{&payment{
		city:          city,
		shop:          shop,
		paymentMethod: paymentMethod,
		date:          date,
		description:   description,
		orders:        btree.NewG(3, func(a, b Order) bool { return a.p.item < b.p.item }),
	}}
}

func NewPaymentForSearches(date string) Payment {
	return Payment{&payment{date: date}}
}

func NewAllPayments() AllPayments {
	return AllPayments{&allPayments{
		valueSet: newValueSet(),
		payments: btree.NewG(3, func(a, b Payment) bool { return a.p.date < b.p.date }),
	}}
}
