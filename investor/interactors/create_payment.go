package interactors

import (
	"fmt"
	assetEntity "investor/entities/asset"
	paymentEntity "investor/entities/payment"
	"investor/ports"
	"time"
)

type PaymentCreator struct {
	Storage     ports.PaymentSaver // todo: add repository here
	IdGenerator ports.IdGenerator
}

type CreatePaymentModel struct {
	AssetAmount    float32
	AbsoluteAmount float32
	Asset          assetEntity.Asset
	Type           paymentEntity.Type
	CreationDate   time.Time
}

func (pc PaymentCreator) cratePaymentInstance(paymentModel CreatePaymentModel, id string) (p paymentEntity.Payment) {
	if paymentModel.Type == paymentEntity.Return {
		p = paymentEntity.NewReturn(
			id, paymentModel.AssetAmount, paymentModel.AbsoluteAmount,
			paymentModel.Asset, paymentModel.CreationDate,
		)
	} else if paymentModel.Type == paymentEntity.Invest {
		p = paymentEntity.NewInvestment(
			id, paymentModel.AssetAmount, paymentModel.AbsoluteAmount,
			paymentModel.Asset, paymentModel.CreationDate,
		)
	} else {
		panic(fmt.Sprintf("unexpected ports type: %d", paymentModel.Type))
	}
	return
}

func (pc PaymentCreator) Create(paymentModel CreatePaymentModel) (err error) {
	id := pc.IdGenerator.Generate()
	p := pc.cratePaymentInstance(paymentModel, id)
	// todo: add validation
	err = pc.Storage.Create(p)
	return
}
