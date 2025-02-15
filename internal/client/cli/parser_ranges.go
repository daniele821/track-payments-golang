package cli

import (
	"errors"
	"fmt"
)

func getNextStr(str string) *string {
	tmp := str + " "
	return &tmp
}

func parseRanges(args []string) (from, to *string, err error) {
	switch {
	case len(args) <= 0:
		return nil, nil, nil
	case matchEveryLenght(args[0], "range"):
		switch {
		case len(args) <= 1:
			return nil, nil, errors.New("missing arg for range filter")
		case matchEveryLenght(args[1], "from"):
			if len(args) != 3 {
				return nil, nil, errors.New("invalid amount of args for range from filter")
			}
			return &args[2], nil, nil
		case matchEveryLenght(args[1], "to"):
			if len(args) != 3 {
				return nil, nil, errors.New("invalid amount of args for range to filter")
			}
			return nil, getNextStr(args[2]), nil
		case matchEveryLenght(args[1], "both"):
			if len(args) != 4 {
				return nil, nil, errors.New("invalid amount of args for range both filter")
			}
			return &args[2], getNextStr(args[3]), nil
		default:
			return nil, nil, errors.New(fmt.Sprintf("invalid arg for range filter: %s", args[1]))
		}
	default:
		return nil, nil, errors.New(fmt.Sprintf("invalid arg to filter: %s", args[0]))
	}
}
