package cli

import (
	"strings"
)

type FlagParsed struct {
	flagArgs  map[string][]string
	flagOrder []string
}

func ParseFlags(args []string) FlagParsed {
	flagArgs := map[string][]string{}
	flagOrder := []string{}

	tmpFlag := ""
	tmpArg := []string{}

	for _, arg := range append(args, "--") {
		if strings.HasPrefix(arg, "-") {
			flagOrder = append(flagOrder, arg)

			flagArgs[tmpFlag] = tmpArg
			tmpFlag = arg
			tmpArg = []string{}
		} else {
			tmpArg = append(tmpArg, arg)
		}
	}

	return FlagParsed{flagArgs: flagArgs, flagOrder: flagOrder[:len(flagOrder)-1]}
}
