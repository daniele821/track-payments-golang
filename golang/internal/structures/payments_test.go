package structures

import (
	"reflect"
	"testing"
	"time"
)

func TestConstructors(t *testing.T) {
	valueSet := ValueSet{}
	allPayments := NewAllPayments(valueSet)

	if err := allPayments.AddCity("Milano"); err != nil {
		t.Fatalf("insertion on new city failed (%s)!", err)
	}

	if err := allPayments.AddShop("Coop"); err != nil {
		t.Fatalf("insertion on new shop failed (%s)!", err)
	}

	if err := allPayments.AddPaymentMethod("Contante"); err != nil {
		t.Fatalf("insertion on new payment method failed (%s)!", err)
	}

	if err := allPayments.AddCategory("Cibo"); err != nil {
		t.Fatalf("insertion on new category failed (%s)!", err)
	}

	if err := allPayments.AddItem("Briosche", "Cibo"); err != nil {
		t.Fatalf("insertion on new item failed (%s)!", err)
	}

}

func TestAllPayments(t *testing.T) {
	valueSet, _ := NewValueSet(
		[]string{"Asti", "Cesena", "Milano", "Roma"},
		[]string{"Coop", "Paninaro", "Conad", "Ristorante Indiano"},
		[]string{"Contante", "Postepay", "San Paolo"},
		[]string{"Cibo", "Viaggio"},
		map[string]string{"Pasta": "Cibo", "Biglietto Treno": "Viaggio"},
	)
	allPayments := NewAllPayments(valueSet)
	if len(allPayments.Payments) != 0 {
		t.Fatal("there shoul not be any payments!")
	}

	if err := allPayments.AddPayment("Asti", "Coop", "Contante", time.Now()); err != nil || len(allPayments.Payments) != 1 {
		t.Fatalf("insertion of new payment failed (%s)!", err)
	}

	if err := allPayments.AddOrder(0, 12, 1245, "Pasta"); err != nil || len(allPayments.Payments[0].Orders) != 1 {
		t.Fatalf("insertion of new order failed (%s)!", err)
	}

	if err := allPayments.RemoveOrder(0, "Pasta"); err != nil || len(allPayments.Payments[0].Orders) != 0 {
		t.Fatalf("deletion of order failed (%s)!", err)
	}

	if err := allPayments.RemovePayment(0); err != nil || len(allPayments.Payments) != 0 {
		t.Fatalf("deletion of payment failed (%s)!", err)
	}
}

func TestJson(t *testing.T) {
	valueSet, _ := NewValueSet(
		[]string{"Asti", "Cesena", "Milano", "Roma"},
		[]string{"Coop", "Paninaro", "Conad", "Ristorante Indiano"},
		[]string{"Contante", "Postepay", "San Paolo"},
		[]string{"Cibo", "Viaggio"},
		map[string]string{"Pasta": "Cibo", "Biglietto Treno": "Viaggio"},
	)
	allPayments := NewAllPayments(valueSet)
	allPayments.AddPayment("Asti", "Coop", "Contante", time.Now())
	allPayments.AddOrder(0, 12, 1245, "Pasta")

	paymentsJson, err := allPayments.GenerateJson(false)
	if err != nil {
		t.Fatalf("conversion to json failed (%s)!", err)
	}

	allPayments2, err := NewAllPaymentsFromJson(paymentsJson)
	if err != nil {
		t.Fatalf("conversion to json failed (%s)!", err)
	}

	if !reflect.DeepEqual(allPayments, allPayments2) {
		t.Fatal("conversion from struct to json to struct has modified the struct!")
	}
}
