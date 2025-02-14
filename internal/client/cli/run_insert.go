package cli

import (
	"errors"
	"fmt"
	"payment/internal/server/payments"
	"strconv"
	"strings"
)

func insertGeneric(dataType string, data []string, insertFunc func(data ...string) error) error {
	if len(data) == 0 {
		return errors.New(fmt.Sprintf("no %s passed", dataType))
	}
	if err := insertFunc(data...); err != nil {
		return errors.New(fmt.Sprintf("%s insertion failed: %s\n", dataType, err))
	} else {
		fmt.Printf("successfully inserted %s (%s)\n", dataType, strings.Join(data, ", "))
	}
	return nil
}

func insertPayments(allPayments payments.AllPayments, data []string) error {
	okMsg := []string{}
	dateStr, timeStr := getDateAndTime()
	for index, splittedData := range splitter(data) {
		if len(splittedData) != 5 {
			return errors.New(fmt.Sprintf("invalid amount of parameters to insert the %dth payment (%s)", index, strings.Join(splittedData, ", ")))
		}
		dateStr = fillDataIfEmpty(splittedData[0], dateStr)
		timeStr = fillDataIfEmpty(splittedData[1], timeStr)
		city := splittedData[2]
		shop := splittedData[3]
		method := splittedData[4]
		dateFinal := dateStr + " " + timeStr
		if err := allPayments.AddPayment(city, shop, method, dateFinal); err != nil {
			return errors.New(fmt.Sprintf("payment (%d) insertion failed: %s\n", index, err))
		} else {
			okMsg = append(okMsg, fmt.Sprintf("successfully inserted payment (%s, %s, %s, %s)\n", dateFinal, city, shop, method))
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
			return errors.New(fmt.Sprintf("invalid amount of parameters to insert the %dth order (%s)", index, strings.Join(splittedData, ", ")))
		}
		dateStr = fillDataIfEmpty(splittedData[0], dateStr)
		timeStr = fillDataIfEmpty(splittedData[1], timeStr)
		item := splittedData[2]
		quantity := splittedData[3]
		quantityInt, err := strconv.Atoi(quantity)
		if err != nil {
			return errors.New(fmt.Sprintf("order (%d) insertion failed: quantity (%s) is not an integer", index, quantity))
		}
		price := splittedData[4]
		priceInt, err := parsePrice(price)
		if err != nil {
			return errors.New(fmt.Sprintf("order (%d) insertion failed: invalid price value (%s)", index, price))
		}
		dateFinal := dateStr + " " + timeStr
		if err := allPayments.AddOrder(quantityInt, priceInt, item, dateFinal); err != nil {
			return errors.New(fmt.Sprintf("order (%d) insertion failed: %s\n", index, err))
		} else {
			okMsg = append(okMsg, fmt.Sprintf("successfully inserted order (%s, %s, %d, %.2f)\n", dateFinal, item, quantityInt, float64(priceInt)/100.0))
		}
	}
	fmt.Println(strings.Join(okMsg, ""))
	return nil
}

func insertDetails(allPayments payments.AllPayments, data []string) error {
	okMsg := []string{}
	dateStr, timeStr := getDateAndTime()
	for index, splittedData := range splitter(data) {
		if index == 0 {
			if len(splittedData) != 5 {
				return errors.New(fmt.Sprintf("invalid amount of parameters to insert payment (%s)", strings.Join(splittedData, ", ")))
			}
			dateStr = fillDataIfEmpty(splittedData[0], dateStr)
			timeStr = fillDataIfEmpty(splittedData[1], timeStr)
			city := splittedData[2]
			shop := splittedData[3]
			method := splittedData[4]
			dateFinal := dateStr + " " + timeStr
			if err := allPayments.AddPayment(city, shop, method, dateFinal); err != nil {
				return errors.New(fmt.Sprintf("payment insertion failed: %s\n", err))
			} else {
				okMsg = append(okMsg, fmt.Sprintf("successfully inserted payment (%s, %s, %s, %s)\n", dateFinal, city, shop, method))
			}
		} else {
			if len(splittedData) != 3 {
				return errors.New(fmt.Sprintf("invalid amount of parameters to insert the %dth order (%s)", index-1, strings.Join(splittedData, ", ")))
			}
			item := splittedData[0]
			quantity := splittedData[1]
			quantityInt, err := strconv.Atoi(quantity)
			if err != nil {
				return errors.New(fmt.Sprintf("order (%d) insertion failed: quantity (%s) is not an integer", index-1, quantity))
			}
			price := splittedData[2]
			priceInt, err := parsePrice(price)
			if err != nil {
				return errors.New(fmt.Sprintf("order (%d) insertion failed: invalid price value (%s)", index-1, price))
			}
			dateFinal := dateStr + " " + timeStr
			if err := allPayments.AddOrder(quantityInt, priceInt, item, dateFinal); err != nil {
				return errors.New(fmt.Sprintf("order (%d) insertion failed: %s\n", index, err))
			} else {
				okMsg = append(okMsg, fmt.Sprintf("successfully inserted order (%s, %s, %d, %.2f)\n", dateFinal, item, quantityInt, float64(priceInt)/100.0))
			}
		}
	}
	fmt.Println(strings.Join(okMsg, ""))
	return nil
}
