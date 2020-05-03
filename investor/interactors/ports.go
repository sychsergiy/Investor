package interactors

import (
	"investor/entities/asset"
	paymentEntity "investor/entities/payment"
)

type AssetCreator interface {
	Create(asset asset.Asset) error
}

type AssetBulkCreator interface {
	CreateBulk(assets []asset.Asset) (int, error)
}

type AssetRepository interface {
	AssetCreator
	AssetBulkCreator
}

type PaymentBulkCreator interface {
	CreateBulk(payments []paymentEntity.Payment) (int, error)
}

type PaymentCreator interface {
	Create(payment paymentEntity.Payment) error
}

type PaymentsLister interface {
	ListAll() []paymentEntity.Payment
}

type PaymentRepository interface {
	PaymentCreator
	PaymentBulkCreator
	PaymentsLister
}

type IdGenerator interface {
	Generate() string
}
