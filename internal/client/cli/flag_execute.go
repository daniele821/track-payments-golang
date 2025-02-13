package cli

import "payment/internal/server/payments"

func (f flags) execute(allPayments payments.AllPayments) {
	insertAct := *f.insertAction
	listAct := *f.listAction
	updateAct := *f.updateAction
	deleteAct := *f.deleteAction
	if insertAct != "" && listAct == "" && updateAct == "" && deleteAct == "" {
		switch insertAct {
		case "city":
		case "shop":
		case "method":
		case "item":
		case "payment":
		case "order":
		}
	} else if insertAct == "" && listAct != "" && updateAct == "" && deleteAct == "" {
		switch listAct {
		case "payment":
		case "order":
		}
	} else if insertAct == "" && listAct == "" && updateAct != "" && deleteAct == "" {
		switch updateAct {
		case "payment":
		case "order":
		}
	} else if insertAct == "" && listAct == "" && updateAct == "" && deleteAct != "" {
		switch deleteAct {
		case "payment":
		case "order":
		}
	}
}
