package cli_manual

import "fmt"

func helpMsg() {
	fmt.Print(`Usage of ./payments

Operations: 
    i[nsert] c[ity|ities] CITY [...]
    i[nsert] s[hops] SHOP [...]
    i[nsert] m[ethods] METHOD [...]
    i[nsert] i[tems] ITEM [...]
    i[nsert] p[ayments] DATE TIME CITY SHOP METHOD DESCRIPTION|* [@ [[DATE2] TIME2] CITY2 SHOP2 METHOD2 DESCRIPTION2|* ...]
    i[nsert] o[rders] DATE TIME ITEM QUANTITY PRICE [@ [[DATE2] TIME2] ITEM2 QUANTITY2 PRICE2 ...]

    l[ist]|v[isualize] c[ity|ities]
    l[ist]|v[isualize] s[hops]
    l[ist]|v[isualize] m[ethods]
    l[ist]|v[isualize] i[tems]
    l[ist]|v[isualize] v[alues]
    l[ist]|v[isualize] p[ayments]
    l[ist]|v[isualize] o[rders]

    u[pdate] p[ayments] DATE TIME [CITY] [SHOP] [METHOD] [DESCRIPTION] [@ ...]
    u[pdate] o[rders] DATE TIME ITEM [QUANTITY] [PRICE] [@ ...]

    d[elete] p[ayments] DATE TIME [@ [DATE2] TIME ...]
    d[elete] o[rders] DATE TIME ITEM [@ [[DATE2] TIME2] ITEM ...]
    
    h[elp]
`)
}
