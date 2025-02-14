package payments

import (
	"reflect"
	"testing"
)

func TestJsonConversion(t *testing.T) {
	allPayments := NewAllPayments()
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
	if err := allPayments.AddPayment("Cesena", "Coop", "Contante", "2025/01/02 23:45"); err != nil {
		t.Fatalf("operation (add payment) failed: %s", err)
	}
	if err := allPayments.AddOrder(1, 8900, "Pizza", "2025/01/02 23:45"); err != nil {
		t.Fatalf("operation (add order1) failed: %s", err)
	}
	if err := allPayments.AddOrder(2, 1200, "Briosche", "2025/01/02 23:45"); err != nil {
		t.Fatalf("operation (add order2) failed: %s", err)
	}

	payment, err := allPayments.Payment("2025/01/02 23:45")
	if err != nil {
		t.Fatalf("error getting payment: %s", err)
	} else if payment.TotalPrice() != 11_300 {
		t.Fatalf("invalid total price calculations: expected 11300, actual %d", payment.TotalPrice())
	}

	expected := allPayments
	tmpStr, err := allPayments.DumpJson(false)
	if err != nil {
		t.Fatalf("json conversion failed: %s", err)
	}
	actual, err := NewAllPaymentsFromJson(tmpStr)
	if err != nil {
		t.Fatalf("json parsing failed: %s", err)
	}

	if !reflect.DeepEqual(convertToJsonData(actual), convertToJsonData(expected)) {
		t.Fatalf("json conversion modified data:\nexpected: %s\nactual: %s", expected, actual)
	}
}
