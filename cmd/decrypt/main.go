package main

import (
	"fmt"
	"os"
	"payment/internal/server/payments"
	"payment/internal/utils"
)

func main() {
	var finalStr string
	if len(os.Args[1:]) != 2 {
		return
	}
	keyFile, encryptedFile := os.Args[1], os.Args[2]
	finalStr, err := utils.DecryptFile(encryptedFile, keyFile)
	if err != nil {
		encryptedBytes, _ := os.ReadFile(encryptedFile)
		finalStr = string(encryptedBytes)
	}

	// indent if file is json
	if allPayments, err := payments.NewAllPaymentsFromJson(finalStr); err == nil {
		if tmpStr, err := allPayments.DumpJson(true); err == nil {
			finalStr = tmpStr
		}
	}

	fmt.Println(finalStr)
}
