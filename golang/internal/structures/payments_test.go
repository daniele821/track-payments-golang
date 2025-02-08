package structures_test

import (
	"payment/internal/structures"
	"reflect"
	"testing"
	"time"
)

func TestConstructors(t *testing.T) {
	valueSet := structures.ValueSet{}
	allPayments := structures.NewAllPayments(valueSet)

	t.Run("InsertCity", func(t *testing.T) {
		if err := allPayments.AddCity("Milano"); err != nil {
			t.Fatalf("insertion on new city failed (%s)!", err)
		}
	})

	t.Run("InsertShop", func(t *testing.T) {
		if err := allPayments.AddShop("Coop"); err != nil {
			t.Fatalf("insertion on new shop failed (%s)!", err)
		}
	})

	t.Run("InsertPaymentMethod", func(t *testing.T) {
		if err := allPayments.AddPaymentMethod("Contante"); err != nil {
			t.Fatalf("insertion on new payment method failed (%s)!", err)
		}
	})

	t.Run("InsertCategory", func(t *testing.T) {
		if err := allPayments.AddCategory("Cibo"); err != nil {
			t.Fatalf("insertion on new category failed (%s)!", err)
		}
	})

	t.Run("InsertItem", func(t *testing.T) {
		if err := allPayments.AddItem("Briosche", "Cibo"); err != nil {
			t.Fatalf("insertion on new item failed (%s)!", err)
		}
	})

}

func TestAllPayments(t *testing.T) {
	valueSet, _ := structures.NewValueSet(
		[]string{"Asti", "Cesena", "Milano", "Roma"},
		[]string{"Coop", "Paninaro", "Conad", "Ristorante Indiano"},
		[]string{"Contante", "Postepay", "San Paolo"},
		[]string{"Cibo", "Viaggio"},
		map[string]string{"Pasta": "Cibo", "Biglietto Treno": "Viaggio"},
	)
	allPayments := structures.NewAllPayments(valueSet)

	t.Run("AssertNoPayments", func(t *testing.T) {
		if len(allPayments.Payments) != 0 {
			t.Fatal("there shoul not be any payments!")
		}
	})

	t.Run("InsertPayment", func(t *testing.T) {
		if err := allPayments.AddPayment("Asti", "Coop", "Contante", time.Now()); err != nil || len(allPayments.Payments) != 1 {
			t.Fatalf("insertion of new payment failed (%s)!", err)
		}
	})

	t.Run("InsertOrder", func(t *testing.T) {
		if err := allPayments.AddOrder(0, 12, 1245, "Pasta"); err != nil || len(allPayments.Payments[0].Orders) != 1 {
			t.Fatalf("insertion of new order failed (%s)!", err)
		}
	})

	t.Run("RemoveOrder", func(t *testing.T) {
		if err := allPayments.RemoveOrder(0, "Pasta"); err != nil || len(allPayments.Payments[0].Orders) != 0 {
			t.Fatalf("deletion of order failed (%s)!", err)
		}
	})

	t.Run("RemovePayment", func(t *testing.T) {
		if err := allPayments.RemovePayment(0); err != nil || len(allPayments.Payments) != 0 {
			t.Fatalf("deletion of payment failed (%s)!", err)
		}
	})

}

func TestJson(t *testing.T) {
	valueSet, _ := structures.NewValueSet(
		[]string{"Asti", "Cesena", "Milano", "Roma"},
		[]string{"Coop", "Paninaro", "Conad", "Ristorante Indiano"},
		[]string{"Contante", "Postepay", "San Paolo"},
		[]string{"Cibo", "Viaggio"},
		map[string]string{"Pasta": "Cibo", "Biglietto Treno": "Viaggio"},
	)
	allPayments := structures.NewAllPayments(valueSet)
	allPayments.AddPayment("Asti", "Coop", "Contante", time.Now())
	allPayments.AddOrder(0, 12, 1245, "Pasta")

	paymentsJson, err := allPayments.GenerateJson(false)
	t.Run("ConvertToJson", func(t *testing.T) {
		if err != nil {
			t.Fatalf("conversion to json failed (%s)!", err)
		}
	})

	allPayments2, err := structures.NewAllPaymentsFromJson(paymentsJson)
	t.Run("ParseFromJson", func(t *testing.T) {
		if err != nil {
			t.Fatalf("conversion to json failed (%s)!", err)
		}
	})

	t.Run("CheckJsonConversion", func(t *testing.T) {
		if !reflect.DeepEqual(allPayments, allPayments2) {
			t.Fatal("conversion from struct to json to struct has modified the struct!")
		}
	})
}
