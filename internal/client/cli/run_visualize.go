package cli

import (
	"fmt"
	"payment/internal/server/payments"
	"strconv"
)

func visualizeGeneric(dataType string, data payments.ReadOnlyBTree[string], from, to *string) {
	boxData := [][][]string{{{"", dataType}}}
	bodyData := [][]string{}
	index := 0
	data.Ascend(func(item string) bool {
		index += 1
		bodyData = append(bodyData, []string{strconv.Itoa(index), item})
		return true
	}, from, to)
	boxData = append(boxData, bodyData)
	fmt.Print(fmtBox(boxData, 1, 1, nil))
}

func visualizePayment(data payments.ReadOnlyBTree[payments.Payment], from, to *string) {
	fromPayment, toPayment := strToPayment(from), strToPayment(to)
	boxData := [][][]string{{{"", "DATE", "CITY", "SHOP", "METHOD", "PRICE"}}}
	bodyData := [][]string{}
	index := 0
	data.Ascend(func(item payments.Payment) bool {
		index += 1
		bodyData = append(bodyData, []string{strconv.Itoa(index), item.Date(), item.City(), item.Shop(), item.PaymentMethod(), fmt.Sprintf("%.2f€", float64(item.TotalPrice())/100.0)})
		return true
	}, fromPayment, toPayment)
	boxData = append(boxData, bodyData)
	fmt.Print(fmtBox(boxData, 1, 1, nil))
}
