package flags

import (
	"errors"
	"maps"
	"strings"
)

type FlagParsed struct {
	flagArgs map[string][]string
}

func (f FlagParsed) FlagArgsCopy() map[string][]string {
	return maps.Clone(f.flagArgs)
}

func NewFlagParsed(args []string) (FlagParsed, error) {
	flagArgs := map[string][]string{}
	flagOrder := []string{}

	flagEmpty := FlagParsed{}
	letters := []string{}

	tmpFlag := ""
	tmpArg := []string{}

	for _, arg := range append(args, "--") {
		if strings.HasPrefix(arg, "-") {
			if arg == "" {
				if len(tmpArg) != 0 {
					return flagEmpty, errors.New("invalid words before flags: " + strings.Join(tmpArg, " "))
				}
			} else {
				flagOrder = append(flagOrder, arg)
			}

			if strings.HasPrefix(arg, "--") {
				if _, exists := flagArgs[tmpFlag]; exists {
					return flagEmpty, errors.New("duplicated word flag: ")
				}
			} else {
				letters = append(letters, arg)
			}

			flagArgs[tmpFlag] = tmpArg
			tmpFlag = arg
			tmpArg = []string{}
		} else {
			tmpArg = append(tmpArg, arg)
		}
	}

	return FlagParsed{flagArgs: flagArgs}, nil
}
