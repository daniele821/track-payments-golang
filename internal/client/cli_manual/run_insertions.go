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
	for index, splittedData := range splitter(data) {
		if len(splittedData) < 5 {
			fmt.Printf("not enough parameters to insert the %dth payment (%s)\n", index, strings.Join(splittedData, ", "))
			return
		}
		date := splittedData[0] + " " + splittedData[1]
		city := splittedData[2]
		shop := splittedData[3]
		method := splittedData[4]
		description := strings.Join(splittedData[4:], " ")
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
