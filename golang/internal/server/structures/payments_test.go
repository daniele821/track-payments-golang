package structures_test

import (
	"payment/internal/server/structures"
	"reflect"
	"testing"
)

func TestJson(t *testing.T) {
	valueSet, _ := structures.NewValueSet(
		[]string{"Asti", "Cesena", "Milano", "Roma"},
		[]string{"Coop", "Paninaro", "Conad", "Ristorante Indiano"},
		[]string{"Contante", "Postepay", "San Paolo"},
		[]string{"Cibo", "Viaggio"},
		map[string]string{"Pasta": "Cibo", "Biglietto Treno": "Viaggio"},
	)
	allPayments := structures.NewAllPayments(valueSet)
	allPayments.AddPayment("Asti", "Coop", "Contante", "2025-12-03 12:34")
	allPayments.AddOrder(0, 12, "2025-12-03 12:34", "Pasta")

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
