package payments

import (
	"fmt"
	"testing"
	"time"
)

func TestConstructors(t *testing.T) {
	valueSet := valueSet{}
	allPayments := newAllPayments(valueSet)
	fmt.Println(allPayments.valueSet)

	if err := allPayments.addCity("Milano"); err != nil {
		t.Fatalf("insertion on new city failed (%s)!", err)
	}
	fmt.Println(allPayments.valueSet)

	if err := allPayments.addShop("Coop"); err != nil {
		t.Fatalf("insertion on new shop failed (%s)!", err)
	}
	fmt.Println(allPayments.valueSet)

	if err := allPayments.addPaymentMethod("Contante"); err != nil {
		t.Fatalf("insertion on new payment method failed (%s)!", err)
	}
	fmt.Println(allPayments.valueSet)

	if err := allPayments.addCategory("Cibo"); err != nil {
		t.Fatalf("insertion on new category failed (%s)!", err)
	}
	fmt.Println(allPayments.valueSet)

	if err := allPayments.addItem("Briosche", "Cibo"); err != nil {
		t.Fatalf("insertion on new item failed (%s)!", err)
	}
	fmt.Println(allPayments.valueSet)

}

func TestAllPayments(t *testing.T) {
	valueSet, _ := newValueSet(
		[]string{"Asti", "Cesena", "Milano", "Roma"},
		[]string{"Coop", "Paninaro", "Conad", "Ristorante Indiano"},
		[]string{"Contante", "Postepay", "San Paolo"},
		[]string{"Cibo", "Viaggio"},
		map[string]string{"Pasta": "Cibo", "Biglietto Treno": "Viaggio"},
	)
	allPayments := newAllPayments(valueSet)
	if len(allPayments.payments) != 0 {
		t.Fatal("there shoul not be any payments!")
	}
	fmt.Println(allPayments.payments)

	if err := allPayments.addPayment("Asti", "Coop", "Contante", time.Now()); err != nil || len(allPayments.payments) != 1 {
		t.Fatalf("insertion of new payment failed (%s)!", err)
	}
	fmt.Println(allPayments.payments)

	if err := allPayments.addOrder(0, 12, 1245, "Pasta"); err != nil || len(allPayments.payments[0].orders) != 1 {
		t.Fatalf("insertion of new order failed (%s)!", err)
	}
	fmt.Println(allPayments.payments)

	if err := allPayments.removeOrder(0, "Pasta"); err != nil || len(allPayments.payments[0].orders) != 0 {
		t.Fatalf("deletion of order failed (%s)!", err)
	}
	fmt.Println(allPayments.payments)

	if err := allPayments.removePayment(0); err != nil || len(allPayments.payments) != 0 {
		t.Fatalf("deletion of payment failed (%s)!", err)
	}
	fmt.Println(allPayments.payments)
}
