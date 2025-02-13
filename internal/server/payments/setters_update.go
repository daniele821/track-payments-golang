package payments

func (allPayments AllPayments) UpdatePayment(date string, city, shop, paymentMethod, description *string) error {
	if err := allPayments.checks(&date, city, shop, paymentMethod, nil); err != nil {
		return err
	}
	payment, err := allPayments.Payment(date)
	if err != nil {
		return err
	}
	if city != nil {
		payment.p.city = *city
	}
	if shop != nil {
		payment.p.shop = *shop
	}
	if paymentMethod != nil {
		payment.p.paymentMethod = *paymentMethod
	}
	if description != nil {
		payment.p.description = *description
	}
	return nil
}

func (allPayments AllPayments) UpdateOrder(date, item string, quantity, unitPrice *int) error {
	if err := allPayments.checks(&date, nil, nil, nil, &item); err != nil {
		return err
	}
	order, err := allPayments.Order(date, item)
	if err != nil {
		return err
	}
	if quantity != nil {
		order.p.quantity = *quantity
	}
	if unitPrice != nil {
		order.p.unitPrice = *unitPrice
	}
	return nil
}
