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

func (valueSet *ValueSet) String() string {
	cityStr := fmtBtree(valueSet.cities, func(item string) string { return item })
	shopStr := fmtBtree(valueSet.shops, func(item string) string { return item })
	payMStr := fmtBtree(valueSet.paymentMethods, func(item string) string { return item })
	itemStr := fmtBtree(valueSet.items, func(item string) string { return item })
	return fmt.Sprintf("(cities: %s, shops: %s, paymentMethods: %s, items: %s)", cityStr, shopStr, payMStr, itemStr)
}

func (order Order) String() string {
	return fmt.Sprintf("(item: %s, quantity: %d, unitPrice: %d)", order.pointer.item, order.pointer.quantity, order.pointer.unitPrice)
}

func (payment Payment) String() string {
	ordersStr := fmtBtree(payment.pointer.orders, func(item Order) string { return item.String() })
	return fmt.Sprintf("(city: %s, shop: %s, paymentMethod: %s, date: %s, description: %s, orders: %s)",
		payment.pointer.city, payment.pointer.shop, payment.pointer.paymentMethod, payment.pointer.date, payment.pointer.description, ordersStr)
}

func (allPayments *AllPayments) String() string {
	paymentStr := fmtBtree(allPayments.payments, func(item Payment) string { return item.String() })
	valueSetStr := allPayments.valueSet.String()
	return fmt.Sprintf("(payments: %s, valueSet: %s)", paymentStr, valueSetStr)
}
