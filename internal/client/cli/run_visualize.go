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
		month, day, time := dateTime.Format("2006 January"), dateTime.Format("02 Mon"), dateTime.Format("15:04")
		monthFmt, dayFmt, timeFmt := month, day, time
		if dayOld != "" {
			if month == monthOld {
				monthFmt = ""
				if day == dayOld {
					dayFmt = ""
				} else {
					boxData = append(boxData, bodyData)
					bodyData = [][]string{}
				}
			} else {
				// index = 1
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

func visualizeDetail(data payments.ReadOnlyBTree[payments.Payment], from, to *string) {
	fromPayment, toPayment, fromInclude, toInclude := strToPayment(from), strToPayment(to), true, true
	boxData := [][][]string{{{"", "MONTH", "DAY", "TIME", "CITY", "SHOP", "METHOD", "PRICE", "ITEM", "QUANTITY", "PRICE"}}}
	bodyData := [][]string{}
	index := 0
	monthOld, dayOld := "", ""
	data.AscendRange(fromPayment, toPayment, fromInclude, toInclude, func(item payments.Payment) bool {
		index += 1
		dateTime, _ := time.Parse("2006/01/02 15:04", item.Date())
		month, day, time := dateTime.Format("2006 January"), dateTime.Format("02 Mon"), dateTime.Format("15:04")
		monthFmt, dayFmt, timeFmt := month, day, time
		if dayOld != "" {
			if month == monthOld {
				monthFmt = ""
				if day == dayOld {
					dayFmt = ""
				} else {
					boxData = append(boxData, bodyData)
					bodyData = [][]string{}
				}
			} else {
				// index = 1
				boxData = append(boxData, bodyData)
				bodyData = [][]string{}
			}
		}
		dayOld, monthOld = day, month
		orders := [][]string{}
		item.Orders().Ascend(func(item payments.Order) bool {
			orders = append(orders, []string{item.Item(), strconv.Itoa(item.Quantity()), strPrice(item.UnitPrice())})
			return true
		})
		if len(orders) == 0 {
			orders = append(orders, []string{"", "", ""})
		}
		bodyData = append(bodyData, []string{strconv.Itoa(index), monthFmt, dayFmt, timeFmt, item.City(), item.Shop(),
			item.PaymentMethod(), fmt.Sprintf("%.2f€", float64(item.TotalPrice())/100.0), orders[0][0], orders[0][1], orders[0][2]})
		for i := 1; i < len(orders); i++ {
			bodyData = append(bodyData, []string{"", "", "", "", "", "", "", "", orders[i][0], orders[i][1], orders[i][2]})
		}
		return true
	})
	boxData = append(boxData, bodyData)
	fmt.Print(fmtBox(boxData, 1, 1, nil))
}

func visualizeAggregated(data payments.ReadOnlyBTree[payments.Payment], from, to *string) {
	boxData := [][][]string{}
	boxData = append(boxData, [][]string{{"PERIOD", "TOTAL", "PAYMENTS", "AVG-PAYMENT", "DAYS", "AVG-DAILY", "MIN", "MAX"}})
	for _, row := range getAllAggregated(data, from, to) {
		if row != nil {
			boxData = append(boxData, [][]string{row})
		}
	}
	fmt.Print(fmtBox(boxData, 1, 1, nil))
}

func getAggregated(data payments.ReadOnlyBTree[payments.Payment], name, from, to string) []string {
	toMinute := to + " 99:99"
	count := 0
	totPrice := 0
	minFound, maxFound := false, false
	min, max := 0, 0
	data.AscendRange(strToPayment(&from), strToPayment(&toMinute), true, false, func(item payments.Payment) bool {
		count++
		totPrice += item.TotalPrice()
		if !minFound || min > item.TotalPrice() {
			min = item.TotalPrice()
			minFound = true
		}
		if !maxFound || max < item.TotalPrice() {
			max = item.TotalPrice()
			maxFound = true
		}
		return true
	})
	if count == 0 {
		return nil
	}
	totPriceStr := strPrice(totPrice)
	avg := fmt.Sprintf("%.2f€", float64(totPrice)/100.0/float64(count))
	countStr := strconv.Itoa(count)
	dayCount := days(from, to)
	dayCountStr := strconv.Itoa(dayCount)
	avgDaily := fmt.Sprintf("%.2f€", float64(totPrice)/100.0/float64(dayCount))
	minStr := strPrice(min)
	maxStr := strPrice(max)
	return []string{name, totPriceStr, countStr, avg, dayCountStr, avgDaily, minStr, maxStr}
}

func days(from, to string) int {
	fromDate, err := time.Parse("2006/01/02", from)
	if err != nil {
		panic(err)
	}
	toDate, err := time.Parse("2006/01/02", to)
	if err != nil {
		panic(err)
	}
	return int(toDate.Sub(fromDate).Hours()/24) + 1
}

func getAllAggregated(data payments.ReadOnlyBTree[payments.Payment], from, to *string) [][]string {
	lines := [][]string{}

	fromStr, toStr := "", ""
	data.Ascend(func(item payments.Payment) bool {
		fromStr = item.Date()[:10]
		return false
	})
	data.Descend(func(item payments.Payment) bool {
		toStr = item.Date()[:10]
		return false
	})

	fromDate, err := time.Parse("2006/01/02", fromStr)
	if err != nil {
		panic(err)
	}
	toDate, err := time.Parse("2006/01/02", toStr)
	if err != nil {
		panic(err)
	}
	currentDate := fromDate

	// monthly stats
	for !currentDate.After(toDate) {
		year, month, _ := currentDate.Date()
		firstDayMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
		nextMonth := month + 1
		nextYear := year
		if nextMonth > time.December {
			nextMonth = time.January
			nextYear++
		}
		firstDayOfNextMonth := time.Date(nextYear, nextMonth, 1, 0, 0, 0, 0, time.UTC)
		lastDayMonth := firstDayOfNextMonth.AddDate(0, 0, -1)

		lines = append(lines, getAggregated(data, firstDayMonth.Format("2006 January"), firstDayMonth.Format("2006/01/02"), lastDayMonth.Format("2006/01/02")))

		currentDate = firstDayOfNextMonth
	}

	// complessive stats
	lines = append(lines, getAggregated(data, "TOTAL", fromStr, toStr))

	return lines
}
