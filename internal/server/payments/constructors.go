package payments

import "github.com/google/btree"

func newValueSet() *ValueSet {
	return &ValueSet{
		cities:         btree.NewG(3, func(a, b string) bool { return a < b }),
		shops:          btree.NewG(3, func(a, b string) bool { return a < b }),
		paymentMethods: btree.NewG(3, func(a, b string) bool { return a < b }),
		items:          btree.NewG(3, func(a, b string) bool { return a < b }),
	}
}

func newOrder(quantity, unitPrice uint, item string) *Order {
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
		orders:        btree.NewG(3, func(a, b *Order) bool { return a.item < b.item }),
	}
}

func newPaymentForSearches(date string) *Payment {
	return &Payment{date: date}
}

func NewAllPayment() *AllPayments {
	return &AllPayments{
		valueSet: newValueSet(),
		payments: btree.NewG(3, func(a, b *Payment) bool { return a.date < b.date }),
	}
}
