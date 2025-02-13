package main

import (
	"fmt"
	"payment/internal/client/cli_goflag"
)

const jsonPath string = "payments.json"

func main() {
	if err := cli_goflag.Run(jsonPath); err != nil {
		fmt.Printf("execution failed: %s\n", err)
	}
}
