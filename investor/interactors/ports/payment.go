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

type PaymentFinderByIds interface {
	FindByIds(ids []string) ([]payment.Payment, error)
}

type PaymentFinderByAssetName interface {
	FindByAssetName(name string, period payment.Period) ([]payment.Payment, error)
}

type PaymentRepository interface {
	PaymentCreator
	PaymentBulkCreator
	PaymentsLister
	PaymentFinderByIds
	PaymentFinderByAssetName
}
