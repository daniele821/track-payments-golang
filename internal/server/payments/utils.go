package payments

import (
	"errors"
	"time"
)

func parseDate(date string) (time.Time, error) {
	return time.ParseInLocation("2006/01/02 15:04", date, time.UTC)
}

func (allPayments *AllPayments) checks(date, city, shop, paymentMethod, item *string) error {
	if date != nil {
		if _, err := parseDate(*date); err != nil {
			return errors.New("invalid date: " + err.Error())
		}
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
