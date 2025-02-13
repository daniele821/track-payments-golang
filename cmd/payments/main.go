package main

import (
	"fmt"
	"payment/internal/client/cli_goflag"
)

func main() {
	if err := cli_goflag.Run(); err != nil {
		fmt.Println(err)
	}
}
