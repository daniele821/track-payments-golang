package climanual

import "fmt"

func helpMsg() {
	fmt.Print(`Usage of ./payments

Operations: 
    i[nsert] c[ity] CITY ...
    i[nsert] s[hop] SHOP ...
    i[nsert] m[ethod] METHOD ...
    i[nsert] i[tem] ITEM ...
    i[nsert] p[ayment] DATE CITY SHOP METHOD [DESCRIPTION] [@ ITEM1 QUANTITY1 PRICE1 @ ...]
    i[nsert] o[rder] DATE ITEM QUANTITY PRICE [@ ITEM2 QUANTITY2 PRICE2 @ ...] 

    l[ist] c[ity|ities]
    l[ist] s[hops]
    l[ist] m[ethods]
    l[ist] i[tems]
    l[ist] p[ayments]
    l[ist] o[rders]

    u[pdate] p[ayment] DATE [CITY] [SHOP] [METHOD] [DESCRIPTION] [@ DATE2 [CITY2] [SHOP2] [METHOD2] [DESCRIPTION2] @ ...]
    u[pdate] o[rder] DATE ITEM [QUANTITY] [PRICE] [@ DATE2 ITEM2 [QUANTITY2] [PRICE2] @ ...]

    d[elete] p[ayment] DATE [DATE2 ...]
    d[elete] o[rder] DATE ITEM [@ [DATE2] ITEM2 @ ...]
    
    h[elp]
`)
}
