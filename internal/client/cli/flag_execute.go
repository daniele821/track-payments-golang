package cli

import (
	"fmt"
	"payment/internal/server/payments"
	"strconv"
	"time"
)

func insert(typeData, data string, addFunc func(string) error) {
	if data == "" {
		fmt.Printf("no %s was passed\n", typeData)
	} else {
		if err := addFunc(data); err != nil {
			fmt.Printf("%s insertion failed: %s\n", typeData, err)
		} else {
			fmt.Printf("successfully added %s (%s)\n", typeData, data)
		}
	}
	return
}

func insertPayment(allPayments payments.AllPayments, flags flags) {
	date := *flags.dateData
	if len(date) == 5 {
		date = time.Now().Format("2006/01/02") + " " + date
	}
	city := *flags.cityData
	method := *flags.methodData
	shop := *flags.shopData
	description := *flags.descriptionData
	switch {
	case date == "":
		fmt.Printf("no date was passed")
	case city == "":
		fmt.Printf("no city was passed")
	case method == "":
		fmt.Printf("no method was passed")
	case shop == "":
		fmt.Printf("no shop was passed")
	default:
		if err := allPayments.AddPayment(city, shop, method, date, description); err != nil {
			fmt.Printf("payment insertion failed: %s\n", err)
		} else {
			fmt.Printf("successfully added payment (date: %s, city: %s, shop: %s, method: %s, description: %s)\n", date, city, shop, method, description)
		}

	}
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

func listPayments(data payments.ReadOnlyBTree[payments.Payment]) {
	if data.Len() == 0 {
		fmt.Printf("There are no payments!\n")
		return
	}
	fmt.Printf("Here's all payments:\n")
	prevDate := ""
	data.Ascend(func(item payments.Payment) bool {
		price := item.TotalPrice()
		date := item.Date()
		if date[:10] == prevDate {
			date = "          " + date[10:]
		}
		prevDate = date[:10]
		fmt.Printf("%s | %s %s %s %d.%02d€\n", date, item.City(), item.Shop(), item.PaymentMethod(), price/100, price%100)
		item.Orders().Ascend(func(item payments.Order) bool {
			price := item.UnitPrice()
			fmt.Printf("                 | %s x%d %d.%02d\n", item.Item(), item.Quantity(), price/100, price%100)
			return true
		}, nil, nil)
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
			insertPayment(allPayments, f)
		case "order":
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
			listPayments(allPayments.Payments())
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
