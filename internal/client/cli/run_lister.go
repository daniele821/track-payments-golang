package cli

import (
	"fmt"
	"payment/internal/server/payments"
	"strconv"
	"strings"
)

func listGeneric(dataType string, data payments.ReadOnlyBTree[string]) {
	if data.Len() == 0 {
		fmt.Printf("There are no %s\n", dataType)
		return
	}
	fmt.Printf("Here all the %s:\n", dataType)
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
		if strings.TrimSpace(date[:10]) != "" {
			prevDate = date[:10]
		}
		fmt.Printf("%s | %s %s %s %d.%02d€\n", date, item.City(), item.Shop(), item.PaymentMethod(), price/100, price%100)
		item.Orders().Ascend(func(item payments.Order) bool {
			price := item.UnitPrice()
			fmt.Printf("                 | %s x%d %d.%02d€\n", item.Item(), item.Quantity(), price/100, price%100)
			return true
		}, nil, nil)
		return true
	}, nil, nil)

}
