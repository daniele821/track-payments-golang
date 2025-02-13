package flags

import (
	"errors"
	"maps"
	"slices"
	"strings"
)

type FlagParsed struct {
	flagArgs map[string][]string
}

func (f FlagParsed) FlagArgsCopy() map[string][]string {
	return maps.Clone(f.flagArgs)
}

func convertWordIntoLetters(word string) (letters []string) {
	letters = make([]string, len(word))
	for i, letter := range word {
		letters[i] = string(letter)
	}
	return letters
}

func NewFlagParsed(args []string) (FlagParsed, error) {
	flagArgs := map[string][]string{}

	flagEmpty := FlagParsed{}
	letters := []string{}

	tmpFlag := ""
	tmpArg := []string{}

	for _, arg := range append(args, "--") {
		if strings.HasPrefix(arg, "-") {
			if tmpFlag == "" {
				if len(tmpArg) != 0 {
					return flagEmpty, errors.New("invalid words before flags: " + strings.Join(tmpArg, " "))
				}
			}

			if strings.HasPrefix(arg, "--") {
				if _, exists := flagArgs[tmpFlag]; exists {
					return flagEmpty, errors.New("duplicated word flag: " + tmpFlag)
				}
			} else {
				tmpLetters := convertWordIntoLetters(arg[1:])
				for _, tmpLetter := range tmpLetters {
					if slices.Contains(letters, tmpLetter) {
						return flagEmpty, errors.New("duplicated letter flag: -" + tmpLetter)
					}
				}
				letters = append(letters, tmpLetters...)
			}

			flagArgs[tmpFlag] = tmpArg
			tmpFlag = arg
			tmpArg = []string{}
		} else {
			tmpArg = append(tmpArg, arg)
		}
	}

	delete(flagArgs, "")

	return FlagParsed{flagArgs: flagArgs}, nil
}
