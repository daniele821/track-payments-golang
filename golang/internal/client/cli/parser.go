package cli

import "strings"

type Flags map[string]string

func Parser(args []string) (Flags, error) {
	flags := Flags{}
	lastFlag := ""
	lastValue := []string{}
	for _, arg := range args {
		if strings.HasPrefix(arg, "--") {
			flags[lastFlag] = strings.Join(lastValue, "")
			lastFlag = arg
		} else {
			lastValue = append(lastValue, arg)
		}
	}
	flags[lastFlag] = strings.Join(lastValue, "")
	return flags, nil
}
