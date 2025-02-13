package cli_manual

import "fmt"

func helpMsg() {
	fmt.Print(`Usage of ./payments

Operations: 
    i[nsert] c[ity|ities] CITY ...
    i[nsert] s[hops] SHOP ...
    i[nsert] m[ethods] METHOD ...
    i[nsert] i[tems] ITEM ...
    i[nsert] p[ayments] DATE CITY SHOP METHOD [DESCRIPTION] [@ ITEM1 QUANTITY1 PRICE1 @ ...]
    i[nsert] o[rders] DATE ITEM QUANTITY PRICE [@ ITEM2 QUANTITY2 PRICE2 @ ...] 

    l[ist] c[ity|ities]
    l[ist] s[hops]
    l[ist] m[ethods]
    l[ist] i[tems]
    l[ist] p[ayments]
    l[ist] o[rders]

    u[pdate] p[ayments] DATE [CITY] [SHOP] [METHOD] [DESCRIPTION] [@ DATE2 [CITY2] [SHOP2] [METHOD2] [DESCRIPTION2] @ ...]
    u[pdate] o[rders] DATE ITEM [QUANTITY] [PRICE] [@ DATE2 ITEM2 [QUANTITY2] [PRICE2] @ ...]

    d[elete] p[ayments] DATE [DATE2 ...]
    d[elete] o[rders] DATE ITEM [@ [DATE2] ITEM2 @ ...]
    
    h[elp]
`)
}
