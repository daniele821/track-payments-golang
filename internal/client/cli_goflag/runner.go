package cli_goflag

import (
	"payment/internal/client"
	"payment/internal/server/payments"
)

func Run() error {
	return client.Run(func(allPayments payments.AllPayments) error {
		addFlags().execute(allPayments)
		return nil
	})
}
