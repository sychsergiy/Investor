package in_memory

import (
	"investor/entities/payment"
	"time"
)

func paymentTypesContains(paymentTypes []payment.Type, paymentType payment.Type) bool {
	for _, pt := range paymentTypes {
		if pt == paymentType {
			return true
		}
	}
	return false
}

func FilterByType(payments []payment.Payment, paymentTypes []payment.Type) (filtered []payment.Payment) {
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

func FilterByPeriod(payments []payment.Payment, periods []payment.Period) (filtered []payment.Payment) {
	if len(periods) == 0 {
		return payments
	}
	for _, payment_ := range payments {
		for _, period := range periods {
			if periodContains(period, payment_.CreationDate()) {
				filtered = append(filtered, payment_)
				break
			}
		}
	}
	return filtered
}

func periodContains(p payment.Period, date time.Time) bool {
	if date.After(p.From()) && date.Before(p.Until()) {
		return true
	}
	return false
}
