package in_memory

import "investor/entities/payment"

func paymentTypesContains(paymentTypes []payment.Type, paymentType payment.Type) bool {
	for _, pt := range paymentTypes {
		if pt == paymentType {
			return true
		}
	}
	return false
}

func FilterPaymentsByType(payments []payment.Payment, paymentTypes []payment.Type) (filtered []payment.Payment) {
	if len(paymentTypes) == 0 {
		return payments
	}
	for _, p := range payments {
		if paymentTypesContains(paymentTypes, p.Type()) {
			filtered = append(filtered, p)
		}
	}
	return filtered
}
