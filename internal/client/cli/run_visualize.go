package cli

import (
	"fmt"
	"payment/internal/server/payments"
	"strconv"
	"strings"
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

func visualizePayment(data payments.ReadOnlyBTree[payments.Payment], from, to *string, print bool) [][][]string {
	fromPayment, toPayment, fromInclude, toInclude := strToPayment(from), strToPayment(to), true, true
	boxData := [][][]string{{{"", "MONTH", "DAY", "TOTAL", "TIME", "CITY", "SHOP", "METHOD", "PRICE"}}}
	bodyData := [][]string{}
	index := 0
	monthOld, dayOld := "", ""
	dailyTotal := 0
	data.AscendRange(fromPayment, toPayment, fromInclude, toInclude, func(item payments.Payment) bool {
		index += 1
		dateTime, _ := time.Parse("2006/01/02 15:04", item.Date())
		month, day, time := dateTime.Format("2006 January"), dateTime.Format("02 Monday"), dateTime.Format("15:04")
		monthFmt, dayFmt, timeFmt := month, day, time
		dailyTotal += item.TotalPrice()
		if dayOld != "" {
			if month == monthOld {
				monthFmt = ""
				if day == dayOld {
					dayFmt = ""
				} else {
					dailyTotal = item.TotalPrice()
					boxData = append(boxData, bodyData)
					bodyData = [][]string{}
				}
			} else {
				dailyTotal = item.TotalPrice()
				boxData = append(boxData, bodyData)
				bodyData = [][]string{}
			}
		}

		dayOld, monthOld = day, month
		bodyData = append(bodyData, []string{strconv.Itoa(index), monthFmt, dayFmt, "", timeFmt, item.City(), item.Shop(), item.PaymentMethod(), fmt.Sprintf("%.2f€", float64(item.TotalPrice())/100.0)})
		bodyData[0][3] = strPrice(dailyTotal)
		return true
	})
	boxData = append(boxData, bodyData)
	if print {
		fmt.Print(fmtBox2(boxData, 1, 1, []cell{
			{-1, -1, 3, rightAlign},
			{-1, -1, 8, rightAlign},
			{0, 0, -1, centerLeftAlign},
		}))
	}
	return boxData
}

func visualizeTime(data payments.ReadOnlyBTree[payments.Payment], from, to *string, print bool) [][][]string {
	tmpData := visualizePayment(data, from, to, false)
	row := tmpData[0][0]
	newData := [][][]string{{row[:4]}}
	newBox := [][]string{}
	index := 1
	for _, tmpBox := range tmpData[1:] {
		if len(tmpBox) > 0 {
			newRow := tmpBox[0][:4]
			newRow[0] = strconv.Itoa(index)
			if strings.TrimSpace(newRow[1]) != "" && len(newBox) > 0 {
				newData = append(newData, newBox)
				newBox = [][]string{}
			}
			newBox = append(newBox, newRow)
			index += 1
		}
	}
	newData = append(newData, newBox)
	if print {
		fmt.Print(fmtBox(newData, 1, 1, []cell{
			{-1, -1, 3, rightAlign},
			{0, 0, -1, centerLeftAlign},
		}))
	}
	return newData
}

func visualizeDetail(data payments.ReadOnlyBTree[payments.Payment], from, to *string) {
	fromPayment, toPayment, fromInclude, toInclude := strToPayment(from), strToPayment(to), true, true
	boxData := [][][]string{{{"", "MONTH", "DAY", "TOTAL", "TIME", "CITY", "SHOP", "METHOD", "PRICE", "ITEM", "QUANTITY", "PRICE"}}}
	bodyData := [][]string{}
	index := 0
	monthOld, dayOld := "", ""
	dailyTotal := 0
	data.AscendRange(fromPayment, toPayment, fromInclude, toInclude, func(item payments.Payment) bool {
		index += 1
		dateTime, _ := time.Parse("2006/01/02 15:04", item.Date())
		month, day, time := dateTime.Format("2006 January"), dateTime.Format("02 Monday"), dateTime.Format("15:04")
		monthFmt, dayFmt, timeFmt := month, day, time
		dailyTotal += item.TotalPrice()
		if dayOld != "" {
			if month == monthOld {
				monthFmt = ""
				if day == dayOld {
					dayFmt = ""
				} else {
					dailyTotal = item.TotalPrice()
					boxData = append(boxData, bodyData)
					bodyData = [][]string{}
				}
			} else {
				dailyTotal = item.TotalPrice()
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
		bodyData = append(bodyData, []string{strconv.Itoa(index), monthFmt, dayFmt, "", timeFmt, item.City(), item.Shop(),
			item.PaymentMethod(), fmt.Sprintf("%.2f€", float64(item.TotalPrice())/100.0), orders[0][0], orders[0][1], orders[0][2]})
		for i := 1; i < len(orders); i++ {
			bodyData = append(bodyData, []string{"", "", "", "", "", "", "", "", "", orders[i][0], orders[i][1], orders[i][2]})
		}
		bodyData[0][3] = strPrice(dailyTotal)
		return true
	})
	boxData = append(boxData, bodyData)
	fmt.Print(fmtBox2(reorgData(boxData), 1, 1, []cell{
		{-1, -1, 3, rightAlign},
		{-1, -1, 8, rightAlign},
		{-1, -1, 11, rightAlign},
		{0, 0, -1, centerLeftAlign},
	}))
}

func visualizeAggregated(data payments.ReadOnlyBTree[payments.Payment], from, to *string) {
	boxData := [][][]string{}
	boxData = append(boxData, [][]string{{"PERIOD", "FROM", "TO", "", "DAYS", "AVG", "MIN", "MAX", "", "TOTAL"}})
	for _, row := range getAllAggregated(data, from, to) {
		if row != nil {
			boxData = append(boxData, [][]string{row})
		}
	}
	fmt.Print(fmtBox(boxData, 1, 1, []cell{
		{0, 0, -1, centerLeftAlign},
		{-1, -1, 4, rightAlign},
		{-1, -1, 5, rightAlign},
		{-1, -1, 6, rightAlign},
		{-1, -1, 7, rightAlign},
		{-1, -1, 8, rightAlign},
		{-1, -1, 9, rightAlign},
	}))
}

func reorgData(data [][][]string) [][][]string {
	res := [][][]string{}
	tmpBox := [][]string{}
	for _, box := range data {
		for _, row := range box {
			if len(tmpBox) != 0 {
				if strings.TrimSpace(row[5]) != "" {
					res = append(res, tmpBox)
					tmpBox = [][]string{}
				}
			}
			tmpBox = append(tmpBox, row)
		}
		if len(tmpBox) != 0 {
			res = append(res, tmpBox)
			tmpBox = [][]string{}
		}
	}
	return res
}

func getAggregated(data payments.ReadOnlyBTree[payments.Payment], name, from, to string) []string {
	toMinute := to + " 99:99"
	count := 0
	totPrice := 0
	minFound, maxFound, minDFound, maxDFound := false, false, false, false
	dayOld, dayTot := "", 0
	min, max, minD, maxD := 0, 0, 0, 0
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
		day := item.Date()[:10]
		if dayOld != "" && day != dayOld {
			if !minDFound || minD > dayTot {
				minD = dayTot
				minDFound = true
			}
			if !maxDFound || maxD < dayTot {
				maxD = dayTot
				maxDFound = true
			}
			dayTot = 0
		}
		dayTot += item.TotalPrice()
		dayOld = day
		return true
	})
	if !minDFound || minD > dayTot {
		minD = dayTot
		minDFound = true
	}
	if !maxDFound || maxD < dayTot {
		maxD = dayTot
		maxDFound = true
	}
	if count == 0 {
		return nil
	}
	totPriceStr := strPrice(totPrice)
	// avg := fmt.Sprintf("%.2f€", float64(totPrice)/100.0/float64(count))
	// countStr := strconv.Itoa(count)
	dayCount := days(from, to)
	dayCountStr := strconv.Itoa(dayCount)
	avgDaily := fmt.Sprintf("%.2f€", float64(totPrice)/100.0/float64(dayCount))
	// minStr, maxStr := strPrice(min), strPrice(max)
	minDStr, maxDStr := strPrice(minD), strPrice(maxD)
	return []string{name, from, to, "", dayCountStr, avgDaily, minDStr, maxDStr, "", totPriceStr}
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
	if from != nil {
		if len(*from) == 1 {
			fromStr = time.Now().Format("2006/01/02")
			*from = fromStr
		} else if fromDate, err := time.Parse("2006/01/02", *from); err == nil {
			fromStr = fromDate.Format("2006/01/02")
			*from = fromStr
		}
	}
	if to != nil {
		if len(*to) == 1 {
			toStr = time.Now().Format("2006/01/02")
			*to = toStr
		} else if toDate, err := time.Parse("2006/01/02", *to); err == nil {
			toStr = toDate.Format("2006/01/02")
			*to = toStr
		}
	}
	if fromStr == "" {
		data.AscendRange(strToPayment(from), strToPayment(to), true, true, func(item payments.Payment) bool {
			fromStr = item.Date()[:10]
			return false
		})
	}
	if toStr == "" {
		data.DescendRange(strToPayment(to), strToPayment(from), true, true, func(item payments.Payment) bool {
			toStr = item.Date()[:10]
			return false
		})
	}

	if toStr == "" {
		toStr = fromStr
	}
	if fromStr == "" {
		fromStr = toStr
	}
	if fromStr > toStr {
		tmpStr := fromStr
		fromStr = toStr
		toStr = tmpStr
	}

	if fromStr != "" {

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

			if firstDayMonth.Before(fromDate) {
				firstDayMonth = fromDate
			}
			if lastDayMonth.After(toDate) {
				lastDayMonth = toDate
			}

			lines = append(lines, getAggregated(data, firstDayMonth.Format("2006 January"), firstDayMonth.Format("2006/01/02"), lastDayMonth.Format("2006/01/02")))

			currentDate = firstDayOfNextMonth
		}
	}

	// complessive stats
	total := getAggregated(data, "TOTAL", fromStr, toStr)
	if total != nil {
		lines = append(lines, total)
	} else {
		lines = append(lines, []string{"TOTAL", fromStr, toStr, "", strconv.Itoa(days(fromStr, toStr)), strPrice(0), strPrice(0), strPrice(0), "", strPrice(0)})
	}
	return lines
}
