package cli_manual

import (
	"fmt"
	"payment/internal/server/payments"
	"strconv"
)

func listGeneric(dataType string, data payments.ReadOnlyBTree[string]) {
	if data.Len() == 0 {
		fmt.Printf("There are no %s\n", dataType)
		return
	}
	fmt.Printf("Here all the %s:\n", dataType)
	maxLen := len(strconv.Itoa(data.Len()))
	index := 0
	data.Ascend(func(item string) bool {
		fmt.Printf("%-*d | %s\n", maxLen, index, item)
		index += 1
		return true
	}, nil, nil)
}

func listPayments() {

}

func listOrders() {

}
