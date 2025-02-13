package cli

import (
	"strings"
)

type Flags struct {
	flagArgs  map[string][]string
	flagOrder []string
}

func ParseFlags(args []string) Flags {
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

	return Flags{flagArgs: flagArgs, flagOrder: flagOrder[:len(flagOrder)-1]}
}
