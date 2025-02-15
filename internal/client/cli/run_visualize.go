package cli

import (
	"fmt"
	"payment/internal/server/payments"
	"strconv"
)

func visualizeGeneric(dataType string, data payments.ReadOnlyBTree[string]) {
	boxData := [][][]string{{{"", dataType}}}
	bodyData := [][]string{}
	index := 0
	data.Ascend(func(item string) bool {
		index += 1
		bodyData = append(bodyData, []string{strconv.Itoa(index), item})
		return true
	}, nil, nil)
	boxData = append(boxData, bodyData)
	fmt.Print(fmtBox(boxData, 1, 1, nil))
}
