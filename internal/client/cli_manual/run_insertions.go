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

}

func insertOrders(allPayments payments.AllPayments, data []string) {

}
