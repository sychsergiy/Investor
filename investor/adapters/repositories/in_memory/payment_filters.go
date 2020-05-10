package in_memory

import "investor/entities/payment"

func FilterPaymentsByType(payments []payment.Payment, paymentType payment.Type) (filtered []payment.Payment) {
	for _, p := range payments {
		if p.Type() == paymentType {
			filtered = append(filtered, p)
		}
	}
	return filtered
}
