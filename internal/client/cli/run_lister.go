package cli

import (
	"fmt"
	"payment/internal/server/payments"
	"strconv"
	"strings"
)

func listGeneric(dataType string, data payments.ReadOnlyBTree[string], from, to *string) {
	if data.Len() == 0 {
		fmt.Printf("There are no %s\n", dataType)
		return
	}
	fmt.Printf("Here all the %s:\n", dataType)
	maxLen := len(strconv.Itoa(data.Len()))
	index := 0
	data.AscendRange(from, to, true, true, func(item string) bool {
		fmt.Printf("%-*d | %s\n", maxLen, index, item)
		index += 1
		return true
	})
}

func strToPayment(str *string) *payments.Payment {
	if str == nil {
		return nil
	}
	payment := payments.NewPaymentForSearches(*str)
	return &payment
}

func listPayments(data payments.ReadOnlyBTree[payments.Payment], from, to *string) {
	fromPayment, toPayment, fromInclude, toInclude := strToPayment(from), strToPayment(to), true, true
	if data.Len() == 0 {
		fmt.Printf("There are no payments!\n")
		return
	}
	totPrice := 0
	count := 0
	fmt.Printf("Here's all payments:\n")
	prevDate := ""
	data.AscendRange(fromPayment, toPayment, fromInclude, toInclude, func(item payments.Payment) bool {
		count += 1
		price := item.TotalPrice()
		totPrice += price
		date := item.Date()
		if date[:10] == prevDate {
			date = "          " + date[10:]
		}
		if strings.TrimSpace(date[:10]) != "" {
			prevDate = date[:10]
		}
		fmt.Printf("%s | %s %s %s %d.%02d€\n", date, item.City(), item.Shop(), item.PaymentMethod(), price/100, price%100)
		return true
	})
}

func listDetails(data payments.ReadOnlyBTree[payments.Payment], from, to *string) {
	fromPayment, toPayment, fromInclude, toInclude := strToPayment(from), strToPayment(to), true, true
	if data.Len() == 0 {
		fmt.Printf("There are no payments!\n")
		return
	}
	totPrice := 0
	count := 0
	fmt.Printf("Here's all payments:\n")
	prevDate := ""
	data.AscendRange(fromPayment, toPayment, fromInclude, toInclude, func(item payments.Payment) bool {
		count += 1
		price := item.TotalPrice()
		totPrice += price
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
			fmt.Printf("                 | x%d %s %d.%02d€\n", item.Quantity(), item.Item(), price/100, price%100)
			return true
		})
		return true
	})
}

func listAggregated(data payments.ReadOnlyBTree[payments.Payment], from, to *string) {
	panic("TODO")
}
