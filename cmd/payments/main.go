package main

import (
	"fmt"
	"payment/internal/client/cli_manual"
)

const jsonPath string = "payments.json"

func main() {
	if err := cli_manual.Run(jsonPath); err != nil {
		fmt.Printf("%s\n", err)
	}
}
