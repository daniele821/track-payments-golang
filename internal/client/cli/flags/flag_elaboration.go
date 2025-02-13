package flags

import (
	"errors"
	"strings"
)

func (f FlagParsed) GetFlag(flag string, hasParam bool) (found bool, word *string, err error) {
	if strings.HasPrefix(flag, "--") {
		words := f.flagArgs[flag]
		if len(words) >= 2 {
			return false, nil, errors.New("multiple words found for " + flag)
		}
	}
	return false, nil, nil
}

func (f FlagParsed) empty() bool {
	return len(f.flagArgs) == 0
}

func (f FlagParsed) Conclude() error {
	if f.empty() {
		return nil
	}
	errStr := "invalid flag: "
	for key := range f.flagArgs {
		if strings.HasPrefix(key, "--") {
			errStr += key
		} else {
			errStr += key[:2]
		}
	}
	return errors.New(errStr)
}
