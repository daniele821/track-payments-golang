package cli

import "fmt"

var helpActions map[[2]string]string = map[[2]string]string{
	{"insert", "cities"}:   "CITY [...]",
	{"insert", "shops"}:    "SHOP [...]",
	{"insert", "methods"}:  "METHOD [...]",
	{"insert", "items"}:    "ITEM [...]",
	{"insert", "payments"}: "DATE|* TIME|* CITY SHOP METHOD [@ ...]",
	{"insert", "orders"}:   "DATE|* TIME|* ITEM QUANTITY PRICE [@ ...]",
	{"insert", "details"}:  "DATE|* TIME|* CITY SHOP METHOD [@ ITEM QUANTITY PRICE] [@ ...]",

	{"list", "cities"}:     "[r[ange] (f[rom] FROM)|(t[o] TO)|(b[oth] FROM TO)]",
	{"list", "shops"}:      "[r[ange] (f[rom] FROM)|(t[o] TO)|(b[oth] FROM TO)]",
	{"list", "methods"}:    "[r[ange] (f[rom] FROM)|(t[o] TO)|(b[oth] FROM TO)]",
	{"list", "items"}:      "[r[ange] (f[rom] FROM)|(t[o] TO)|(b[oth] FROM TO)]",
	{"list", "payments"}:   "[r[ange] (f[rom] FROM)|(t[o] TO)|(b[oth] FROM TO)]",
	{"list", "details"}:    "[r[ange] (f[rom] FROM)|(t[o] TO)|(b[oth] FROM TO)]",
	{"list", "aggregated"}: "[r[ange] (f[rom] FROM)|(t[o] TO)|(b[oth] FROM TO)]",
	{"list", "time"}:       "[r[ange] (f[rom] FROM)|(t[o] TO)|(b[oth] FROM TO)]",

	{"update", "payments"}: "DATE|* TIME|* CITY|* SHOP|* METHOD|* [@ ...]",
	{"update", "orders"}:   "DATE|* TIME|* ITEM QUANTITY PRICE [@ ...]",
	{"update", "details"}:  "DATE|* TIME|* CITY|* SHOP|* METHOD|* [@ ITEM QUANTITY PRICE] [@ ...]",

	{"delete", "payments"}: "DATE|* TIME|* [@...]",
	{"delete", "orders"}:   "DATE|* TIME|* ITEM [@ ...]",

	{"print", ""}: "",
}

func helpMsg() {
	fmt.Print(`Usage of ./payments

Operations: 
    i[nsert] c[ity|ities] CITY [...]
    i[nsert] s[hops] SHOP [...]
    i[nsert] m[ethods] METHOD [...]
    i[nsert] i[tems] ITEM [...]
    i[nsert] p[ayments] DATE|* TIME|* CITY SHOP METHOD [@ ...]
    i[nsert] o[rders] DATE|* TIME|* ITEM QUANTITY PRICE [@ ...]
    i[nsert] d[etails] DATE|* TIME|* CITY SHOP METHOD [@ ITEM QUANTITY PRICE] [@ ...]

    l[ist]|v[isualize] c[ity|ities] [r[ange] (f[rom] FROM)|(t[o] TO)|(b[oth] FROM TO)] 
    l[ist]|v[isualize] s[hops] [r[ange] (f[rom] FROM)|(t[o] TO)|(b[oth] FROM TO)]
    l[ist]|v[isualize] m[ethods] [r[ange] (f[rom] FROM)|(t[o] TO)|(b[oth] FROM TO)]
    l[ist]|v[isualize] i[tems] [r[ange] (f[rom] FROM)|(t[o] TO)|(b[oth] FROM TO)]
    l[ist]|v[isualize] v[alues] [r[ange] (f[rom] FROM)|(t[o] TO)|(b[oth] FROM TO)]
    l[ist]|v[isualize] p[ayments] [r[ange] (f[rom] FROM)|(t[o] TO)|(b[oth] FROM TO)]
    l[ist]|v[isualize] d[etails] [r[ange] (f[rom] FROM)|(t[o] TO)|(b[oth] FROM TO)]
    l[ist]|v[isualize] a[ggregated] [r[ange] (f[rom] FROM)|(t[o] TO)|(b[oth] FROM TO)]
    l[ist]|v[isualize] t[ime] [r[ange] (f[rom] FROM)|(t[o] TO)|(b[oth] FROM TO)]

    u[pdate] p[ayments] DATE|* TIME|* CITY|* SHOP|* METHOD|* [@ ...]
    u[pdate] o[rders] DATE|* TIME|* ITEM QUANTITY PRICE [@ ...]
    u[pdate] d[etails] DATE|* TIME|* CITY|* SHOP|* METHOD|* [@ ITEM QUANTITY PRICE] [@ ...]

    d[elete] p[ayments] DATE|* TIME|* [@...]
    d[elete] o[rders] DATE|* TIME|* ITEM [@ ...]
    
    h[elp]

    p[rint]
`)
}
