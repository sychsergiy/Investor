package storage

import "Investor/asset/payment"

type Storage interface {
	AddPayment(payment payment.Payment)
	RetrieveAllPayments() []payment.Payment
}
