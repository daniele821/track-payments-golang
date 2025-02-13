package cli

import (
	"fmt"
	"os"
	"payment/internal/client/cli/flags"
)

func Run() {
	flags := flags.NewFlagParsed(os.Args[1:])
	fmt.Println(flags)
}
