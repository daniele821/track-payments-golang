package cli

import "fmt"

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

    l[ist]|v[isualize] c[ity|ities]
    l[ist]|v[isualize] s[hops]
    l[ist]|v[isualize] m[ethods]
    l[ist]|v[isualize] i[tems]
    l[ist]|v[isualize] v[alues]
    l[ist]|v[isualize] p[ayments]

    u[pdate] p[ayments] DATE|* TIME|* CITY|* SHOP|* METHOD|* [@ ...]
    u[pdate] o[rders] DATE|* TIME|* ITEM QUANTITY PRICE [@ ...]
    u[pdate] d[etails] DATE|* TIME|* CITY|* SHOP|* METHOD|* [@ ITEM QUANTITY PRICE] [@ ...]

    d[elete] p[ayments] DATE|* TIME|* [@...]
    d[elete] o[rders] DATE|* TIME|* ITEM [@ ...]
    
    h[elp]
`)
}
