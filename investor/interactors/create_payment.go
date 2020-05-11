package interactors

import (
	assetEntity "investor/entities/asset"
	paymentEntity "investor/entities/payment"
	"investor/interactors/ports"
	"time"
)

type CreatePayment struct {
	repository  ports.PaymentCreator
	idGenerator ports.IDGenerator
}

type CreatePaymentModel struct {
	AssetAmount    float32
	AbsoluteAmount float32
	Asset          assetEntity.Asset
	Type           paymentEntity.Type
	CreationDate   time.Time
}

func (pc CreatePayment) Create(paymentModel CreatePaymentModel) (err error) {
	id := pc.idGenerator.Generate()
	p := paymentEntity.NewPlainPayment(
		id, paymentModel.AssetAmount, paymentModel.AbsoluteAmount,
		paymentModel.Asset, paymentModel.CreationDate, paymentModel.Type,
	)
	// todo: add validation
	err = pc.repository.Create(p)
	return
}

func NewCreatePayment(repository ports.PaymentCreator, idGenerator ports.IDGenerator) CreatePayment {
	return CreatePayment{repository, idGenerator}
}
