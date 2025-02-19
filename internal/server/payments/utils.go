package payments

import (
	"errors"
	"time"

	"github.com/google/btree"
)

func parseDate(date string) (time.Time, error) {
	return time.Parse("2006/01/02 15:04", date)
}

func (allPayments *AllPayments) checks(date, city, shop, paymentMethod, item *string) error {
	if date != nil {
		_, err := parseDate(*date)
		if err != nil {
			return errors.New("invalid date: " + err.Error())
		}
		// FUTURE DATA CHECKS:
		// dateNow := time.Now().Format("2006/01/02 15:04")
		// if *date > dateNow {
		// 	return errors.New(fmt.Sprintf("invalid date: future dates are not accepted (input: %s, now: %s)", *date, dateNow))
		// }
	}
	if city != nil {
		if !allPayments.p.valueSet.p.cities.Has(*city) {
			return errors.New("invalid city: " + *city)
		}
	}
	if shop != nil {
		if !allPayments.p.valueSet.p.shops.Has(*shop) {
			return errors.New("invalid shop: " + *shop)
		}
	}
	if paymentMethod != nil {
		if !allPayments.p.valueSet.p.paymentMethods.Has(*paymentMethod) {
			return errors.New("invalid paymentMethod: " + *paymentMethod)
		}
	}
	if item != nil {
		if !allPayments.p.valueSet.p.items.Has(*item) {
			return errors.New("invalid item: " + *item)
		}
	}
	return nil
}

func btreeToSlice[T, S any](data *btree.BTreeG[T], mapper func(item T) S) []S {
	acc := make([]S, data.Len())
	index := 0
	data.Ascend(func(item T) bool {
		acc[index] = mapper(item)
		index += 1
		return true
	})
	return acc
}

func mapperIdentity[T any](item T) T {
	return item
}
func mapperOrderJson(item Order) orderJson {
	return orderJson{
		Quantity:  item.Quantity(),
		UnitPrice: item.UnitPrice(),
		Item:      item.Item(),
	}
}
func mapperPaymentJson(item Payment) paymentJson {
	return paymentJson{
		City:          item.City(),
		Shop:          item.Shop(),
		PaymentMethod: item.PaymentMethod(),
		Date:          item.Date(),
		Orders:        btreeToSlice(item.p.orders, mapperOrderJson),
	}
}
