package payments_test

import (
	"payment/internal/server/payments"
	"reflect"
	"testing"
)

func TestJsonConversion(t *testing.T) {
	allPayments := payments.NewAllPayments()
	if err := allPayments.AddCities("Cesena", "Asti"); err != nil {
		t.Fatalf("operation (add cities) failed: %s", err)
	}
	if err := allPayments.AddItems("Briosche", "Pizza"); err != nil {
		t.Fatalf("operation (add items) failed: %s", err)
	}
	if err := allPayments.AddShops("Coop", "Conad"); err != nil {
		t.Fatalf("operation (add shops) failed: %s", err)
	}
	if err := allPayments.AddPaymentMethods("Contante", "Postepay"); err != nil {
		t.Fatalf("operation (add method) failed: %s", err)
	}
	if err := allPayments.AddPayment("Cesena", "Coop", "Contante", "2025/01/02 23:45", "Testing..."); err != nil {
		t.Fatalf("operation (add payment) failed: %s", err)
	}
	if err := allPayments.AddOrder(3, 4, "Pizza", "2025/01/02 23:45"); err != nil {
		t.Fatalf("operation (add order) failed: %s", err)
	}

	expected := allPayments
	actual, err := payments.NewAllPaymentsFromJson(allPayments.DumpJson(false))
	if err != nil {
		t.Fatalf("json conversion failed: %s", err)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("json conversion modified data:\nexpected: %s\nactual: %s", expected, actual)
	}
}
