package cli

import (
	"errors"
	"fmt"
	"payment/internal/server/payments"
	"strconv"
	"strings"
)

func updatePayments(allPayments payments.AllPayments, data []string) error {
	okMsg := []string{}
	dateStr, timeStr := getDateAndTime()
	for index, splittedData := range splitter(data) {
		if len(splittedData) != 5 {
			return errors.New(fmt.Sprintf("invalid amount of parameters to update the %dth payment (%s)", index, strings.Join(splittedData, ", ")))
		}
		dateStr = fillDataIfEmpty(splittedData[0], dateStr)
		timeStr = fillDataIfEmpty(splittedData[1], timeStr)
		city := fillDataIfEmptyOpt(&splittedData[2], nil)
		shop := fillDataIfEmptyOpt(&splittedData[3], nil)
		method := fillDataIfEmptyOpt(&splittedData[4], nil)
		dateFinal := dateStr + " " + timeStr
		if err := allPayments.UpdatePayment(dateFinal, city, shop, method); err != nil {
			return errors.New(fmt.Sprintf("payment (%d) update failed: %s\n", index, err))
		} else {
			okMsg = append(okMsg, fmt.Sprintf("successfully updated payment (%s, %v, %v, %v)\n", dateFinal, city, shop, method))
		}
	}
	fmt.Println(strings.Join(okMsg, ""))
	return nil
}

func updateOrders(allPayments payments.AllPayments, data []string) error {
	okMsg := []string{}
	dateStr, timeStr := getDateAndTime()
	for index, splittedData := range splitter(data) {
		if len(splittedData) != 5 {
			return errors.New(fmt.Sprintf("invalid amount of parameters to update the %dth order (%s)", index, strings.Join(splittedData, ", ")))
		}
		dateStr = fillDataIfEmpty(splittedData[0], dateStr)
		timeStr = fillDataIfEmpty(splittedData[1], timeStr)
		item := splittedData[2]
		quantity := splittedData[3]
		quantityInt, err := strconv.Atoi(quantity)
		if err != nil {
			return errors.New(fmt.Sprintf("order (%d) update failed: quantity (%s) is not an integer", index, quantity))
		}
		price := splittedData[4]
		priceInt, err := parsePrice(price)
		if err != nil {
			return errors.New(fmt.Sprintf("order (%d) update failed: invalid price value (%s)", index, price))
		}
		dateFinal := dateStr + " " + timeStr
		if err := allPayments.UpdateOrder(item, dateFinal, &quantityInt, &priceInt); err != nil {
			return errors.New(fmt.Sprintf("order (%d) update failed: %s\n", index, err))
		} else {
			okMsg = append(okMsg, fmt.Sprintf("successfully updated order (%s, %s, %d, %.2f)\n", dateFinal, item, quantityInt, float64(priceInt)/100.0))
		}
	}
	fmt.Println(strings.Join(okMsg, ""))
	return nil
}

func updateDetails(allPayments payments.AllPayments, data []string) error {
	okMsg := []string{}
	dateStr, timeStr := getDateAndTime()
	for index, splittedData := range splitter(data) {
		if index == 0 {
			if len(splittedData) != 5 {
				return errors.New(fmt.Sprintf("invalid amount of parameters to update the payment (%s)", strings.Join(splittedData, ", ")))
			}
			dateStr = fillDataIfEmpty(splittedData[0], dateStr)
			timeStr = fillDataIfEmpty(splittedData[1], timeStr)
			city := fillDataIfEmptyOpt(&splittedData[2], nil)
			shop := fillDataIfEmptyOpt(&splittedData[3], nil)
			method := fillDataIfEmptyOpt(&splittedData[4], nil)
			dateFinal := dateStr + " " + timeStr
			if err := allPayments.UpdatePayment(dateFinal, city, shop, method); err != nil {
				return errors.New(fmt.Sprintf("payment update failed: %s\n", err))
			} else {
				okMsg = append(okMsg, fmt.Sprintf("successfully updated payment (%s, %v, %v, %v)\n", dateFinal, city, shop, method))
			}
		} else {
			if len(splittedData) != 3 {
				return errors.New(fmt.Sprintf("invalid amount of parameters to update the %dth order (%s)", index-1, strings.Join(splittedData, ", ")))
			}
			item := splittedData[0]
			quantity := splittedData[1]
			quantityInt, err := strconv.Atoi(quantity)
			if err != nil {
				return errors.New(fmt.Sprintf("order (%d) update failed: quantity (%s) is not an integer", index-1, quantity))
			}
			price := splittedData[2]
			priceInt, err := parsePrice(price)
			if err != nil {
				return errors.New(fmt.Sprintf("order (%d) update failed: invalid price value (%s)", index-1, price))
			}
			dateFinal := dateStr + " " + timeStr
			if err := allPayments.UpdateOrder(dateFinal, item, &quantityInt, &priceInt); err != nil {
				return errors.New(fmt.Sprintf("order (%d) update failed: %s\n", index-1, err))
			} else {
				okMsg = append(okMsg, fmt.Sprintf("successfully updated order (%s, %s, %d, %.2f)\n", dateFinal, item, quantityInt, float64(priceInt)/100.0))
			}
		}
	}
	fmt.Println(strings.Join(okMsg, ""))
	return nil
}
