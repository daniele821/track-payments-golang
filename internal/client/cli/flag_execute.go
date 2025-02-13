package cli

import (
	"errors"
	"fmt"
	"payment/internal/server/payments"
	"strconv"
	"strings"
	"time"
)

func parsePrice(price string) (int, error) {
	price = strings.TrimSpace(price)
	priceInt, err := strconv.Atoi(price)
	if err == nil {
		return priceInt * 100, nil
	}
	splitted := strings.Split(price, ".")
	if len(splitted) > 2 {
		return 0, errors.New("invalid price")
	}
	realInt, err := strconv.Atoi(splitted[0])
	if err != nil {
		return 0, err
	}
	decInt, err := strconv.Atoi((splitted[1] + "00")[2:])
	if err != nil {
		return 0, err
	}
	return realInt*100 + decInt, nil
}

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
		if len(date) == 5 {
			date = time.Now().Format("2006/01/02") + " " + date
		}
		if err := allPayments.AddPayment(city, shop, method, date, description); err != nil {
			fmt.Printf("payment insertion failed: %s\n", err)
		} else {
			fmt.Printf("successfully added payment (date: %s, city: %s, shop: %s, method: %s, description: %s)\n", date, city, shop, method, description)
		}

	}
}

func insertOrder(allPayments payments.AllPayments, flags flags) {
	date := *flags.dateData
	item := *flags.itemData
	quantity := *flags.quantityData
	price := *flags.priceData
	switch {
	case date == "":
		fmt.Printf("no date was passed")
	case item == "":
		fmt.Printf("no item was passed")
	case quantity == "":
		fmt.Printf("no quantity was passed")
	case price == "":
		fmt.Printf("no unitPrice was passed")
	default:
		if len(date) == 5 {
			date = time.Now().Format("2006/01/02") + " " + date
		}
		quantityInt, err := strconv.Atoi(quantity)
		if err != nil {
			fmt.Printf("invalid quantity value: %s\n", err)
		}
		priceInt, err := parsePrice(price)
		if err != nil {
			fmt.Printf("invalid unitPrice value: %s\n", err)
		}
		if err := allPayments.AddOrder(quantityInt, priceInt, item, date); err != nil {
			fmt.Printf("payment insertion failed: %s\n", err)
		} else {
			fmt.Printf("successfully added order (date: %s, item: %s, quantity: %d, unitPrice: %0.2f€)\n", date, item, quantityInt, float64(priceInt)/100.0)
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
			fmt.Printf("                 | %s x%d %d.%02d€\n", item.Item(), item.Quantity(), price/100, price%100)
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
			insertOrder(allPayments, f)
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
		case "all":
			listRaw("cities", allPayments.Cities())
			fmt.Println()
			listRaw("shops", allPayments.Shops())
			fmt.Println()
			listRaw("methods", allPayments.PaymentMethods())
			fmt.Println()
			listRaw("items", allPayments.Items())
			fmt.Println()
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
