package cli

import (
	"fmt"
	"payment/internal/server/payments"
)

func insert(typeData, data string, addFunc func(string) error) {
	if data == "" {
		fmt.Printf("no %s was passed\n", typeData)
	} else {
		if err := addFunc(data); err != nil {
			fmt.Printf("%s insertion failed: %s\n", typeData, err)
		}
		fmt.Printf("successfully added %s %s\n", typeData, data)
	}
	return
}

func (f flags) execute(allPayments payments.AllPayments) {
	insertAct := *f.insertAction
	listAct := *f.listAction
	updateAct := *f.updateAction
	deleteAct := *f.deleteAction
	if insertAct != "" && listAct == "" && updateAct == "" && deleteAct == "" {
		switch insertAct {
		case "city":
			insert("city", *f.cityData, func(s string) error { return allPayments.AddCities(s) })
		case "shop":
			insert("shop", *f.shopData, func(s string) error { return allPayments.AddShops(s) })
		case "method":
			insert("method", *f.methodData, func(s string) error { return allPayments.AddPaymentMethods(s) })
		case "item":
			insert("item", *f.itemData, func(s string) error { return allPayments.AddItems(s) })
		case "payment":
		case "order":
		}
	} else if insertAct == "" && listAct != "" && updateAct == "" && deleteAct == "" {
		switch listAct {
		case "payment":
		case "order":
		}
	} else if insertAct == "" && listAct == "" && updateAct != "" && deleteAct == "" {
		switch updateAct {
		case "payment":
		case "order":
		}
	} else if insertAct == "" && listAct == "" && updateAct == "" && deleteAct != "" {
		switch deleteAct {
		case "payment":
		case "order":
		}
	}
}
