package interactors

import (
	assetEntity "investor/entities/asset"
	paymentEntity "investor/entities/payment"
	"investor/interactors/ports"
	"time"
)

type CreatePayment struct {
	Repository  ports.PaymentCreator
	IdGenerator ports.IdGenerator
}

type CreatePaymentModel struct {
	AssetAmount    float32
	AbsoluteAmount float32
	Asset          assetEntity.Asset
	Type           paymentEntity.Type
	CreationDate   time.Time
}

func (pc CreatePayment) Create(paymentModel CreatePaymentModel) (err error) {
	id := pc.IdGenerator.Generate()
	p := paymentEntity.NewPlainPayment(
		id, paymentModel.AssetAmount, paymentModel.AbsoluteAmount,
		paymentModel.Asset, paymentModel.CreationDate, paymentModel.Type,
	)
	// todo: add validation
	err = pc.Repository.Create(p)
	return
}
