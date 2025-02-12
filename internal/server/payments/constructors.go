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

func newOrder(quantity, unitPrice uint, item string) *order {
	return &order{
		quantity:  quantity,
		unitPrice: unitPrice,
		item:      item,
	}
}

func newOrderForSearches(item string) *order {
	return &order{item: item}
}

func newPayment(city, shop, paymentMethod, date, description string) *payment {
	return &payment{
		city:          city,
		shop:          shop,
		paymentMethod: paymentMethod,
		date:          date,
		description:   description,
		orders:        btree.NewG(3, func(a, b *order) bool { return a.item < b.item }),
	}
}

func newPaymentForSearches(date string) *payment {
	return &payment{date: date}
}

func NewAllPayments() *AllPayments {
	return &AllPayments{
		valueSet: newValueSet(),
		payments: btree.NewG(3, func(a, b *payment) bool { return a.date < b.date }),
	}
}
