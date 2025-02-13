package flags

import "fmt"

func (f FlagParsed) String() string {
	return fmt.Sprintf("flagArgs: %v, FlagOrder: %v", f.flagArgs, f.flagOrder)
}
