package cli

import (
	"errors"
	"strings"
)

type Flags map[string]string

func Parser(args []string) (Flags, error) {
	flags := Flags{}
	lastFlag := "--"
	lastValue := []string{}
	for _, arg := range append(args, "--") {
		if strings.HasPrefix(arg, "--") {

			if _, ok := flags[lastFlag]; ok {
				return flags, errors.New("duplicated flags are not valid!")
			}
			flags[lastFlag] = strings.Join(lastValue, "")
			lastFlag = arg
			lastValue = []string{}
		} else {
			lastValue = append(lastValue, arg)
		}
	}
	return flags, nil
}
