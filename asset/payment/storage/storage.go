package storage

import "Investor/asset/payment"

type PaymentsSaver interface {
	SavePayment(payment payment.Payment)
}

type PaymentsRetriever interface {
	RetrieveAllPayments() []payment.Payment
}

type PaymentDeleter interface {
	DeletePayment(id int) bool
}

type Storage interface {
	PaymentsSaver
	PaymentsRetriever
	PaymentDeleter
}
