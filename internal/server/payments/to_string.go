package payments

import (
	"fmt"
	"strings"

	"github.com/google/btree"
)

func fmtBtree[T any](btree *btree.BTreeG[T], strconv func(item T) string) string {
	acc := []string{}
	btree.Ascend(func(item T) bool {
		acc = append(acc, strconv(item))
		return true
	})
	return "[" + strings.Join(acc, " ") + "]"
}

func (valueSet ValueSet) String() string {
	cityStr := fmtBtree(valueSet.p.cities, func(item string) string { return item })
	shopStr := fmtBtree(valueSet.p.shops, func(item string) string { return item })
	payMStr := fmtBtree(valueSet.p.paymentMethods, func(item string) string { return item })
	itemStr := fmtBtree(valueSet.p.items, func(item string) string { return item })
	return fmt.Sprintf("(cities: %s, shops: %s, paymentMethods: %s, items: %s)", cityStr, shopStr, payMStr, itemStr)
}

func (order Order) String() string {
	return fmt.Sprintf("(item: %s, quantity: %d, unitPrice: %d)", order.p.item, order.p.quantity, order.p.unitPrice)
}

func (payment Payment) String() string {
	ordersStr := fmtBtree(payment.p.orders, func(item Order) string { return item.String() })
	return fmt.Sprintf("(city: %s, shop: %s, paymentMethod: %s, date: %s, description: %s, orders: %s)",
		payment.p.city, payment.p.shop, payment.p.paymentMethod, payment.p.date, payment.p.description, ordersStr)
}

func (allPayments AllPayments) String() string {
	paymentStr := fmtBtree(allPayments.p.payments, func(item Payment) string { return item.String() })
	valueSetStr := allPayments.p.valueSet.String()
	return fmt.Sprintf("(payments: %s, valueSet: %s)", paymentStr, valueSetStr)
}
