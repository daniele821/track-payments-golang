package main

import (
	"fmt"
	"payment/internal/client/cli"
)

func main() {
	if err := cli.Run(); err != nil {
		fmt.Printf("failed: %s", err)
	}
}
