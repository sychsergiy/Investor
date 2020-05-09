package ports

import "investor/entities/payment"

type PaymentBulkCreator interface {
	CreateBulk(payments []payment.Payment) (int, error)
}

type PaymentCreator interface {
	Create(payment payment.Payment) error
}

type PaymentsLister interface {
	ListAll() ([]payment.Payment, error)
}

type PaymentRepository interface {
	PaymentCreator
	PaymentBulkCreator
	PaymentsLister
}
