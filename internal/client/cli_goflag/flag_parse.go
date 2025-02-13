package cli_goflag

import (
	"flag"
	"fmt"
)

type flags struct {
	// action flags
	insertAction *string
	listAction   *string
	updateAction *string
	deleteAction *string

	// data flags
	cityData        *string
	shopData        *string
	methodData      *string
	itemData        *string
	descriptionData *string
	quantityData    *string
	priceData       *string
	dateData        *string
}

func addFlags() flags {
	flags := flags{
		insertAction:    flag.String("insert", "", ""),
		listAction:      flag.String("list", "", ""),
		updateAction:    flag.String("update", "", ""),
		deleteAction:    flag.String("delete", "", ""),
		cityData:        flag.String("city", "", ""),
		shopData:        flag.String("shop", "", ""),
		methodData:      flag.String("method", "", ""),
		itemData:        flag.String("item", "", ""),
		descriptionData: flag.String("description", "", ""),
		quantityData:    flag.String("quantity", "", ""),
		priceData:       flag.String("price", "", ""),
		dateData:        flag.String("date", "", ""),
	}
	flag.Usage = helpMsg
	flag.Parse()
	return flags
}

func helpMsg() {
	fmt.Println(`Usage of payments:

    Insertion operations:
        -insert city    -city CITY
        -insert shop    -shop SHOP
        -insert method  -method METHOD
        -insert item    -item ITEM
        -insert payment -date DATE -city CITY -shop SHOP -method METHOD [-description DESCRIPTION]
        -insert order   -date DATE -item ITEM -quantity QUANTITY -price PRICE

    List operations:
        -list city
        -list shop
        -list method
        -list item
        -list payment   
        -list all     

    Update operations:
        -update payment -date DATE [-city CITY] [-shop SHOP] [-method METHOD] [-description DESCRIPTION]
        -update order   -date DATE -item ITEM [-quantity QUANTITY] [-price PRICE]

    Delete operations:
        -delete payment -date DATE
        -delete order   -date DATE -item ITEM
        `)
}
