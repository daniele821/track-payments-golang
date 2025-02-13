package cli_manual

import (
	"fmt"
	"payment/internal/server/payments"
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
		if len(splittedData) < 4 {
			fmt.Printf("not enough parameters to insert the %dth payment (%s)\n", index, strings.Join(splittedData, ", "))
			return
		} else if len(splittedData) > 6 {
			fmt.Printf("too many parameters to insert the %dth payment (%s)\n", index, strings.Join(splittedData, ", "))
			return
		}
		offset := 2 - (6 - len(splittedData))
		switch len(splittedData) {
		case 5:
			timeStr = splittedData[0]
		case 6:
			dateStr = splittedData[0]
			timeStr = splittedData[1]
		}
		date := dateStr + " " + timeStr
		city := splittedData[0+offset]
		shop := splittedData[1+offset]
		method := splittedData[2+offset]
		description := splittedData[3+offset]
		if len(description) <= 1 {
			description = ""
		}
		if err := allPayments.AddPayment(city, shop, method, date, description); err != nil {
			fmt.Printf("payment (%d) insertion failed: %s\n", index, err)
			return
		} else {
			okMsg = append(okMsg, fmt.Sprintf("successfully inserted payment (%s, %s, %s, %s, %s)\n", date, city, shop, method, description))
		}
	}
	fmt.Println(strings.Join(okMsg, ""))
}

func insertOrders(allPayments payments.AllPayments, data []string) {
	for _, splittedData := range splitter(data) {
		panic(splittedData)
	}
}
