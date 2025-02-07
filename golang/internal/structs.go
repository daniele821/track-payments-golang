package payments

import (
	"errors"
	"slices"
)

type valueSet struct {
	cities        []string
	shops         []string
	paymentMethod []string
	categories    []string
	itemCat       map[string]string
}

type order struct {
	quantity   int
	unit_price int
	item       string
	category   string
	valueSet   valueSet
}

type payment struct {
	city          string
	shop          string
	paymentMethod string
	orders        []order
	valueSet      valueSet
}

func newValueSet(cities, shops, methods, categories []string, itemCat map[string]string) (*valueSet, error) {
	valueSet := valueSet{
		cities:        cities,
		shops:         shops,
		paymentMethod: methods,
		categories:    categories,
		itemCat:       itemCat,
	}
	for _, category := range itemCat {
		if !slices.Contains(categories, category) {
			return nil, errors.New("invalid category!")
		}
	}
	return &valueSet, nil
}
