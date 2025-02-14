package cli_manual

import (
	"errors"
	"fmt"
	"payment/internal/server/payments"
	"strconv"
	"strings"
)

func insertGeneric(dataType string, data []string, insertFunc func(data ...string) error) error {
	if len(data) == 0 {
		return errors.New(fmt.Sprintf("no %s passed\n", dataType))
	}
	if err := insertFunc(data...); err != nil {
		fmt.Printf("%s insertion failed: %s\n", dataType, err)
	} else {
		fmt.Printf("successfully inserted %s (%s)\n", dataType, strings.Join(data, ", "))
	}
	return nil
}

func insertPayments(allPayments payments.AllPayments, data []string) error {
	okMsg := []string{}
	dateStr, timeStr := getDateAndTime()
	for index, splittedData := range splitter(data) {
		if len(splittedData) != 6 {
			return errors.New(fmt.Sprintf("invalid amount of parameters to insert the %dth payment (%s)\n", index, strings.Join(splittedData, ", ")))
		}
		dateStr = fillDataIfEmpty(splittedData[0], dateStr)
		timeStr = fillDataIfEmpty(splittedData[1], timeStr)
		city := splittedData[2]
		shop := splittedData[3]
		method := splittedData[4]
		description := fillDataIfEmpty(splittedData[5], "")
		dateFinal := dateStr + " " + timeStr
		if err := allPayments.AddPayment(city, shop, method, dateFinal, description); err != nil {
			return errors.New(fmt.Sprintf("payment (%d) insertion failed: %s\n", index, err))
		} else {
			okMsg = append(okMsg, fmt.Sprintf("successfully inserted payment (%s, %s, %s, %s, %s)\n", dateFinal, city, shop, method, description))
		}
	}
	fmt.Println(strings.Join(okMsg, ""))
	return nil
}

func insertOrders(allPayments payments.AllPayments, data []string) error {
	okMsg := []string{}
	dateStr, timeStr := getDateAndTime()
	for index, splittedData := range splitter(data) {
		if len(splittedData) != 5 {
			return errors.New(fmt.Sprintf("invalid amount of parameters to insert the %dth order (%s)\n", index, strings.Join(splittedData, ", ")))
		}
		dateStr = fillDataIfEmpty(splittedData[0], dateStr)
		timeStr = fillDataIfEmpty(splittedData[1], timeStr)
		item := splittedData[2]
		quantity := splittedData[3]
		quantityInt, err := strconv.Atoi(quantity)
		if err != nil {
			return errors.New(fmt.Sprintf("order (%d) insertion failed: quantity (%s) is not an integer\n", index, quantity))
		}
		price := splittedData[4]
		priceInt, err := parsePrice(price)
		if err != nil {
			return errors.New(fmt.Sprintf("order (%d) insertion failed: invalid price value (%s) \n", index, price))
		}
		dateFinal := dateStr + " " + timeStr
		if err := allPayments.AddOrder(quantityInt, priceInt, item, dateFinal); err != nil {
			return errors.New(fmt.Sprintf("order (%d) insertion failed: %s\n", index, err))
		} else {
			okMsg = append(okMsg, fmt.Sprintf("successfully inserted order (%s, %s, %d, %s)\n", dateFinal, item, quantityInt, "TODO"))
		}
	}
	fmt.Println(strings.Join(okMsg, ""))
	return nil
}

func insertDetails(allPayments payments.AllPayments, data []string) error {
	for _, splittedData := range splitter(data) {
		panic(splittedData)
	}
	return nil
}
