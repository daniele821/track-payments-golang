package cli

import "fmt"

func (f flags) execute() {
	fmt.Println(f.insertAction)
	fmt.Println(f.listAction)
}
