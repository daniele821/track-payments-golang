package cli

import "payment/internal/server/payments"

func Run() {
	addFlags().execute(payments.NewAllPayments())
}
