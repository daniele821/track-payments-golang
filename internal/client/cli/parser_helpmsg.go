package cli

import (
	"fmt"
	"strings"
)

func parseHelp(args []string) bool {
	if !matchEveryLenght(args[len(args)-1], "help") || len(args) >= 4 {
		return false
	}
	tmpArr := append(args, "", "", "")

	switch {
	case matchEveryLenght(tmpArr[0], "insert"):
		switch {
		case matchEveryLenght(tmpArr[1], "help"):
			fmt.Print(genMsgs([]string{"i[nsert] c[ity|ities]", "i[nsert] s[hops]", "i[nsert] m[ethods]",
				"i[nsert] i[tems]", "i[nsert] p[ayments]"},
				[][2]string{{"insert", "cities"}, {"insert", "shops"}, {"insert", "methods"},
					{"insert", "items"}, {"insert", "payments"}}))
		case matchEveryLenght(tmpArr[1], "cities"):
			fmt.Print(genMsg("i[nsert] c[ity|ities]", [2]string{"insert", "cities"}))
		case matchEveryLenght(tmpArr[1], "shops"):
			fmt.Print(genMsg("i[nsert] s[hops]", [2]string{"insert", "shops"}))
		case matchEveryLenght(tmpArr[1], "methods"):
			fmt.Print(genMsg("i[nsert] m[ethods]", [2]string{"insert", "methods"}))
		case matchEveryLenght(tmpArr[1], "items"):
			fmt.Print(genMsg("i[nsert] i[tems]", [2]string{"insert", "items"}))
		case matchEveryLenght(tmpArr[1], "payments"):
			fmt.Print(genMsg("i[nsert] p[ayments]", [2]string{"insert", "payments"}))
		}

	case matchEveryLenght(tmpArr[0], "list"):
		switch {
		case matchEveryLenght(tmpArr[1], "help"):
			fmt.Print(genMsgs([]string{"l[ist] c[ity|ities]", "l[ist] s[hops]", "l[ist] m[ethods]",
				"l[ist] i[tems]", "l[ist] v[alues]", "l[ist] p[ayments]"},
				[][2]string{{"list", "cities"}, {"list", "shops"}, {"list", "methods"},
					{"list", "items"}, {"list", "values"}, {"list", "payments"}}))
		case matchEveryLenght(tmpArr[1], "cities"):
			fmt.Print(genMsg("l[ist] c[ity|ities]", [2]string{"list", "cities"}))
		case matchEveryLenght(tmpArr[1], "shops"):
			fmt.Print(genMsg("l[ist] s[hops]", [2]string{"list", "shops"}))
		case matchEveryLenght(tmpArr[1], "methods"):
			fmt.Print(genMsg("l[ist] m[ethods]", [2]string{"list", "methods"}))
		case matchEveryLenght(tmpArr[1], "items"):
			fmt.Print(genMsg("l[ist] i[tems]", [2]string{"list", "items"}))
		case matchEveryLenght(tmpArr[1], "values"):
			fmt.Print(genMsg("l[ist] v[alues]", [2]string{"list", "values"}))
		case matchEveryLenght(tmpArr[1], "payments"):
			fmt.Print(genMsg("l[ist] p[ayments]", [2]string{"list", "payments"}))
		}

	case matchEveryLenght(tmpArr[0], "visualize"):
		switch {
		case matchEveryLenght(tmpArr[1], "help"):
			fmt.Print(genMsgs([]string{"v[isualize] c[ity|ities]", "v[isualize] s[hops]", "v[isualize] m[ethods]",
				"v[isualize] i[tems]", "v[isualize] v[alues]", "v[isualize] p[ayments]"},
				[][2]string{{"list", "cities"}, {"list", "shops"}, {"list", "methods"},
					{"list", "items"}, {"list", "values"}, {"list", "payments"}}))
		case matchEveryLenght(tmpArr[1], "cities"):
			fmt.Print(genMsg("v[isualize] c[ity|ities]", [2]string{"list", "cities"}))
		case matchEveryLenght(tmpArr[1], "shops"):
			fmt.Print(genMsg("v[isualize] s[hops]", [2]string{"list", "shops"}))
		case matchEveryLenght(tmpArr[1], "methods"):
			fmt.Print(genMsg("v[isualize] m[ethods]", [2]string{"list", "methods"}))
		case matchEveryLenght(tmpArr[1], "items"):
			fmt.Print(genMsg("v[isualize] i[tems]", [2]string{"list", "items"}))
		case matchEveryLenght(tmpArr[1], "values"):
			fmt.Print(genMsg("v[isualize] v[alues]", [2]string{"list", "values"}))
		case matchEveryLenght(tmpArr[1], "payments"):
			fmt.Print(genMsg("v[isualize] p[ayments]", [2]string{"list", "payments"}))
		}

	case matchEveryLenght(tmpArr[0], "update"):
		switch {
		case matchEveryLenght(tmpArr[1], "help"):
			fmt.Print(genMsgs([]string{"u[pdate] p[ayments]", "u[pdate] o[rders]", "u[pdate] d[etails]"},
				[][2]string{{"update", "payments"}, {"update", "orders"}, {"update", "details"}}))
		case matchEveryLenght(tmpArr[1], "payments"):
			fmt.Print(genMsg("u[pdate] p[ayments]", [2]string{"update", "payments"}))
		case matchEveryLenght(tmpArr[1], "orders"):
			fmt.Print(genMsg("u[pdate] o[rders]", [2]string{"update", "orders"}))
		case matchEveryLenght(tmpArr[1], "details"):
			fmt.Print(genMsg("u[pdate] d[etails]", [2]string{"update", "details"}))
		}

	case matchEveryLenght(tmpArr[0], "delete"):
		switch {
		case matchEveryLenght(tmpArr[1], "help"):
			fmt.Print(genMsgs([]string{"d[elete] p[ayments]", "d[elete] o[rders]"},
				[][2]string{{"delete", "payments"}, {"delete", "orders"}}))
		case matchEveryLenght(tmpArr[1], "payments"):
			fmt.Print(genMsg("d[elete] p[ayments]", [2]string{"delete", "payments"}))
		case matchEveryLenght(tmpArr[1], "orders"):
			fmt.Print(genMsg("d[elete] o[rders]", [2]string{"delete", "orders"}))
		}

	case matchEveryLenght(tmpArr[0], "print"):
		fmt.Print(genMsg("p[rint]", [2]string{"print", ""}))

	case matchEveryLenght(tmpArr[0], "help"):
		helpMsg()
	}

	return true
}

func genMsg(name string, elem [2]string) string {
	acc := strings.Builder{}
	acc.WriteString("Operation: ")
	acc.WriteString(name)
	acc.WriteString(" ")
	acc.WriteString(helpActions[elem])
	acc.WriteString("\n")
	return acc.String()
}

func genMsgs(names []string, elems [][2]string) string {
	acc := strings.Builder{}
	acc.WriteString("Operations:\n")
	for i, elem := range elems {
		acc.WriteString("    ")
		acc.WriteString(names[i])
		acc.WriteString(" ")
		acc.WriteString(helpActions[elem])
		acc.WriteString("\n")
	}
	return acc.String()
}
