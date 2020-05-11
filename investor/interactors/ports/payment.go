package ports

import (
	"investor/entities/asset"
	"investor/entities/payment"
)

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

type PaymentFinderByAssetNames interface {
	FindByAssetNames(
		assetNames []string,
		periods []payment.Period,
		paymentTypes []payment.Type,
	) ([]payment.Payment, error)
}

type PaymentFinderByAssetCategories interface {
	FindByAssetCategories(
		categories []asset.Category,
		periods []payment.Period,
		paymentTypes []payment.Type,
	) ([]payment.Payment, error)
}

type PaymentRepository interface {
	PaymentCreator
	PaymentBulkCreator
	PaymentsLister
	PaymentFinderByIds
	PaymentFinderByAssetNames
	PaymentFinderByAssetCategories
}
