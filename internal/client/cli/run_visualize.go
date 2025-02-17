package cli

import (
	"fmt"
	"payment/internal/server/payments"
	"strconv"
	"time"
)

func visualizeGeneric(dataType string, data payments.ReadOnlyBTree[string], from, to *string) {
	boxData := [][][]string{{{"", dataType}}}
	bodyData := [][]string{}
	index := 0
	data.AscendRange(from, to, true, true, func(item string) bool {
		index += 1
		bodyData = append(bodyData, []string{strconv.Itoa(index), item})
		return true
	})
	boxData = append(boxData, bodyData)
	fmt.Print(fmtBox(boxData, 1, 1, nil))
}

func visualizePayment(data payments.ReadOnlyBTree[payments.Payment], from, to *string) {
	fromPayment, toPayment, fromInclude, toInclude := strToPayment(from), strToPayment(to), true, true
	boxData := [][][]string{{{"", "MONTH", "DAY", "TIME", "CITY", "SHOP", "METHOD", "PRICE"}}}
	bodyData := [][]string{}
	index := 0
	monthOld, dayOld := "", ""
	data.AscendRange(fromPayment, toPayment, fromInclude, toInclude, func(item payments.Payment) bool {
		index += 1
		dateTime, _ := time.Parse("2006/01/02 15:04", item.Date())
		month, day, time := dateTime.Format("January 2006"), dateTime.Format("02"), dateTime.Format("15:04")
		monthFmt, dayFmt, timeFmt := month, day, time
		if dayOld != "" {
			if month == monthOld {
				monthFmt = ""
			}
			if day == dayOld {
				dayFmt = ""
			} else {
				boxData = append(boxData, bodyData)
				bodyData = [][]string{}
			}
		}
		dayOld, monthOld = day, month
		bodyData = append(bodyData, []string{strconv.Itoa(index), monthFmt, dayFmt, timeFmt, item.City(), item.Shop(), item.PaymentMethod(), fmt.Sprintf("%.2f€", float64(item.TotalPrice())/100.0)})
		return true
	})
	boxData = append(boxData, bodyData)
	fmt.Print(fmtBox(boxData, 1, 1, nil))
}
