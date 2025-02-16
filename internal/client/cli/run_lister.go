package cli

import (
	"fmt"
	"math"
	"payment/internal/server/payments"
	"strconv"
	"strings"
	"time"
)

func listGeneric(dataType string, data payments.ReadOnlyBTree[string], from, to *string) {
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
	}, from, to)
}

func strToPayment(str *string) *payments.Payment {
	if str == nil {
		return nil
	}
	payment := payments.NewPaymentForSearches(*str)
	return &payment
}

func listPayments(data payments.ReadOnlyBTree[payments.Payment], from, to *string) {
	fromPayment, toPayment := strToPayment(from), strToPayment(to)
	if data.Len() == 0 {
		fmt.Printf("There are no payments!\n")
		return
	}
	totPrice := 0
	count := 0
	fromDate := ""
	toDate := ""
	data.Ascend(func(item payments.Payment) bool {
		fromDate = item.Date()
		return false
	}, fromPayment, toPayment)
	data.Descend(func(item payments.Payment) bool {
		toDate = item.Date()
		return false
	}, fromPayment, toPayment)
	fmt.Printf("Here's all payments:\n")
	prevDate := ""
	data.Ascend(func(item payments.Payment) bool {
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
			fmt.Printf("                 | %s x%d %d.%02d€\n", item.Item(), item.Quantity(), price/100, price%100)
			return true
		}, nil, nil)
		return true
	}, fromPayment, toPayment)
	fmt.Printf("\n%s - %s\n", fromDate, toDate)
	fmt.Printf("total price: %.2f€\n", float64(totPrice)/100.0)
	fmt.Printf("payments: %d\n", count)
	tmp1, _ := time.Parse("2006/01/02", fromDate[:10])
	tmp2, _ := time.Parse("2006/01/02", toDate[:10])
	days := int(math.RoundToEven((tmp2.Sub(tmp1).Hours() / 24) + 1))
	fmt.Printf("days: %d\n", days)
	fmt.Printf("average daily payment: %.2f€\n", float64(totPrice)/100.0/float64(days))
}
