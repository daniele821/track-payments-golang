package flags

import "fmt"

func (f FlagParsed) String() string {
	return fmt.Sprintf("flagArgs: %v", f.flagArgs)
}
