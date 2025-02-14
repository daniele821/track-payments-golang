package cli_manual

import (
	"errors"
	"fmt"
	"payment/internal/server/payments"
	"strings"
)

func deletePayments(allPayments payments.AllPayments, data []string) error {
	okMsg := []string{}
	dateStr, timeStr := getDateAndTime()
	for index, splittedData := range splitter(data) {
		if len(splittedData) != 2 {
			return errors.New(fmt.Sprintf("invalid amount of parameters to delete the %dth payment (%s)", index, strings.Join(splittedData, ", ")))
		}
		dateStr = fillDataIfEmpty(splittedData[0], dateStr)
		timeStr = fillDataIfEmpty(splittedData[1], timeStr)
		dateFinal := dateStr + " " + timeStr
		if err := allPayments.RemovePayment(dateFinal); err != nil {
			return errors.New(fmt.Sprintf("payment (%d) deletion failed: %s\n", index, err))
		} else {
			okMsg = append(okMsg, fmt.Sprintf("successfully delete payment (%s)\n", dateFinal))
		}
	}
	fmt.Println(strings.Join(okMsg, ""))
	return nil
}

func deleteOrders(allPayments payments.AllPayments, data []string) error {
	okMsg := []string{}
	dateStr, timeStr := getDateAndTime()
	for index, splittedData := range splitter(data) {
		if len(splittedData) != 3 {
			return errors.New(fmt.Sprintf("invalid amount of parameters to delete the %dth order (%s)", index, strings.Join(splittedData, ", ")))
		}
		dateStr = fillDataIfEmpty(splittedData[0], dateStr)
		timeStr = fillDataIfEmpty(splittedData[1], timeStr)
		item := splittedData[2]
		dateFinal := dateStr + " " + timeStr
		if err := allPayments.RemoveOrder(dateFinal, item); err != nil {
			return errors.New(fmt.Sprintf("order (%d) deletion failed: %s\n", index, err))
		} else {
			okMsg = append(okMsg, fmt.Sprintf("successfully deleted order (%s, %s)\n", dateFinal, item))
		}
	}
	fmt.Println(strings.Join(okMsg, ""))
	return nil
}
