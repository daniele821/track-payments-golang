package payments

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestConstructors(t *testing.T) {
	valueSet := valueSet{}
	allPayments := newAllPayments(valueSet)
	fmt.Println(allPayments.ValueSet)

	if err := allPayments.addCity("Milano"); err != nil {
		t.Fatalf("insertion on new city failed (%s)!", err)
	}
	fmt.Println(allPayments.ValueSet)

	if err := allPayments.addShop("Coop"); err != nil {
		t.Fatalf("insertion on new shop failed (%s)!", err)
	}
	fmt.Println(allPayments.ValueSet)

	if err := allPayments.addPaymentMethod("Contante"); err != nil {
		t.Fatalf("insertion on new payment method failed (%s)!", err)
	}
	fmt.Println(allPayments.ValueSet)

	if err := allPayments.addCategory("Cibo"); err != nil {
		t.Fatalf("insertion on new category failed (%s)!", err)
	}
	fmt.Println(allPayments.ValueSet)

	if err := allPayments.addItem("Briosche", "Cibo"); err != nil {
		t.Fatalf("insertion on new item failed (%s)!", err)
	}
	fmt.Println(allPayments.ValueSet)

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
	if len(allPayments.Payments) != 0 {
		t.Fatal("there shoul not be any payments!")
	}
	fmt.Println(allPayments.Payments)

	if err := allPayments.addPayment("Asti", "Coop", "Contante", time.Now()); err != nil || len(allPayments.Payments) != 1 {
		t.Fatalf("insertion of new payment failed (%s)!", err)
	}
	fmt.Println(allPayments.Payments)

	if err := allPayments.addOrder(0, 12, 1245, "Pasta"); err != nil || len(allPayments.Payments[0].Orders) != 1 {
		t.Fatalf("insertion of new order failed (%s)!", err)
	}
	fmt.Println(allPayments.Payments)

	if err := allPayments.removeOrder(0, "Pasta"); err != nil || len(allPayments.Payments[0].Orders) != 0 {
		t.Fatalf("deletion of order failed (%s)!", err)
	}
	fmt.Println(allPayments.Payments)

	if err := allPayments.removePayment(0); err != nil || len(allPayments.Payments) != 0 {
		t.Fatalf("deletion of payment failed (%s)!", err)
	}
	fmt.Println(allPayments.Payments)
}

func TestJson(t *testing.T) {
	valueSet, _ := newValueSet(
		[]string{"Asti", "Cesena", "Milano", "Roma"},
		[]string{"Coop", "Paninaro", "Conad", "Ristorante Indiano"},
		[]string{"Contante", "Postepay", "San Paolo"},
		[]string{"Cibo", "Viaggio"},
		map[string]string{"Pasta": "Cibo", "Biglietto Treno": "Viaggio"},
	)
	allPayments := newAllPayments(valueSet)
	allPayments.addPayment("Asti", "Coop", "Contante", time.Now())
	allPayments.addOrder(0, 12, 1245, "Pasta")

	paymentsJson, err := allPayments.generateJson(false)
	if err != nil {
		t.Fatalf("conversion to json failed (%s)!", err)
	}
	fmt.Println(paymentsJson)

	allPayments2, err := newAllPaymentsFromJson(paymentsJson)
	if err != nil {
		t.Fatalf("conversion to json failed (%s)!", err)
	}
	fmt.Println(allPayments2)

	if !reflect.DeepEqual(allPayments, allPayments2) {
		t.Fatal("conversion from struct to json to struct has modified the struct!")
	}
}
