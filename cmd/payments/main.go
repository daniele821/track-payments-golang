package main

import (
	"fmt"
	"payment/internal/client/cli"
)

const jsonPath string = "payments.json"

func main() {
	if err := cli.Run(jsonPath); err != nil {
		fmt.Printf("%s\n", err)
	}
}
