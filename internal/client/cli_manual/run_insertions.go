package cli_manual

import (
	"fmt"
	"payment/internal/server/payments"
	"strconv"
	"strings"
)

func insertGeneric(dataType string, data []string, insertFunc func(data ...string) error) {
	if len(data) == 0 {
		fmt.Printf("no %s passed\n", dataType)
		return
	}
	if err := insertFunc(data...); err != nil {
		fmt.Printf("%s insertion failed: %s\n", dataType, err)
	} else {
		fmt.Printf("successfully inserted %s (%s)\n", dataType, strings.Join(data, ", "))
	}
}

func insertPayments(allPayments payments.AllPayments, data []string) {
	okMsg := []string{}
	dateStr, timeStr := getDateAndTime()
	for index, splittedData := range splitter(data) {
		if len(splittedData) != 6 {
			fmt.Printf("invalid amount of parameters to insert the %dth payment (%s)\n", index, strings.Join(splittedData, ", "))
			return
		}
		dateStr = fillDataIfEmpty(splittedData[0], dateStr)
		timeStr = fillDataIfEmpty(splittedData[1], timeStr)
		city := splittedData[2]
		shop := splittedData[3]
		method := splittedData[4]
		description := fillDataIfEmpty(splittedData[5], "")
		dateFinal := dateStr + " " + timeStr
		if err := allPayments.AddPayment(city, shop, method, dateFinal, description); err != nil {
			fmt.Printf("payment (%d) insertion failed: %s\n", index, err)
			return
		} else {
			okMsg = append(okMsg, fmt.Sprintf("successfully inserted payment (%s, %s, %s, %s, %s)\n", dateFinal, city, shop, method, description))
		}
	}
	fmt.Println(strings.Join(okMsg, ""))
}

func insertOrders(allPayments payments.AllPayments, data []string) {
	okMsg := []string{}
	dateStr, timeStr := getDateAndTime()
	for index, splittedData := range splitter(data) {
		if len(splittedData) != 5 {
			fmt.Printf("invalid amount of parameters to insert the %dth order (%s)\n", index, strings.Join(splittedData, ", "))
			return
		}
		dateStr = fillDataIfEmpty(splittedData[0], dateStr)
		timeStr = fillDataIfEmpty(splittedData[1], timeStr)
		item := splittedData[2]
		quantity := splittedData[3]
		quantityInt, err := strconv.Atoi(quantity)
		if err != nil {
			fmt.Printf("order (%d) insertion failed: quantity (%s) is not an integer\n", index, quantity)
			return
		}
		price := splittedData[4]
		priceInt, err := parsePrice(price)
		if err != nil {
			fmt.Printf("order (%d) insertion failed: invalid price value (%s) \n", index, price)
			return
		}
		dateFinal := dateStr + " " + timeStr
		if err := allPayments.AddOrder(quantityInt, priceInt, item, dateFinal); err != nil {
			fmt.Printf("order (%d) insertion failed: %s\n", index, err)
			return
		} else {
			okMsg = append(okMsg, fmt.Sprintf("successfully inserted order (%s, %s, %d, %s)\n", dateFinal, item, quantityInt, "TODO"))
		}
	}
	fmt.Println(strings.Join(okMsg, ""))
}

func insertDetails(allPayments payments.AllPayments, data []string) {
	for _, splittedData := range splitter(data) {
		panic(splittedData)
	}
}
