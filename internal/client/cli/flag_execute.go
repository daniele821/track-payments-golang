package cli

import (
	"fmt"
	"payment/internal/server/payments"
	"strconv"
)

func insert(typeData, data string, addFunc func(string) error) {
	if data == "" {
		fmt.Printf("no %s was passed\n", typeData)
	} else {
		if err := addFunc(data); err != nil {
			fmt.Printf("%s insertion failed: %s\n", typeData, err)
		} else {
			fmt.Printf("successfully added %s %s\n", typeData, data)
		}
	}
	return
}

func listRaw(typeData string, data payments.ReadOnlyBTree[string]) {
	if data.Len() == 0 {
		fmt.Printf("There are no %s!\n", typeData)
		return
	}
	fmt.Printf("Here's all %s:\n", typeData)
	maxLen := len(strconv.Itoa(data.Len()))
	index := 0
	data.Ascend(func(item string) bool {
		fmt.Printf("%-*d | %s\n", maxLen, index, item)
		index += 1
		return true
	}, nil, nil)
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
		default:
			fmt.Printf("invalid action type (%s)\n", insertAct)
		}
	} else if insertAct == "" && listAct != "" && updateAct == "" && deleteAct == "" {
		switch listAct {
		case "city":
			listRaw("cities", allPayments.Cities())
		case "shop":
			listRaw("shops", allPayments.Shops())
		case "method":
			listRaw("methods", allPayments.PaymentMethods())
		case "item":
			listRaw("items", allPayments.Items())
		case "payment":
		case "order":
		default:
			fmt.Printf("invalid action type (%s)\n", listAct)
		}
	} else if insertAct == "" && listAct == "" && updateAct != "" && deleteAct == "" {
		switch updateAct {
		case "payment":
		case "order":
		default:
			fmt.Printf("invalid action type (%s)\n", updateAct)
		}
	} else if insertAct == "" && listAct == "" && updateAct == "" && deleteAct != "" {
		switch deleteAct {
		case "payment":
		case "order":
		default:
			fmt.Printf("invalid action type (%s)\n", deleteAct)
		}
	}
}
