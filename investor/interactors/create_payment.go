package interactors

import (
	"fmt"
	assetEntity "investor/entities/asset"
	paymentEntity "investor/entities/payment"
	"time"
)

type CreatePayment struct {
	Repository  PaymentCreator
	IdGenerator IdGenerator
}

type CreatePaymentModel struct {
	AssetAmount    float32
	AbsoluteAmount float32
	Asset          assetEntity.Asset
	Type           paymentEntity.Type
	CreationDate   time.Time
}

func (pc CreatePayment) createPaymentInstance(paymentModel CreatePaymentModel, id string) (p paymentEntity.Payment) {
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
		panic(fmt.Sprintf("unexpected adapters type: %d", paymentModel.Type))
	}
	return
}

func (pc CreatePayment) Create(paymentModel CreatePaymentModel) (err error) {
	id := pc.IdGenerator.Generate()
	p := pc.createPaymentInstance(paymentModel, id)
	// todo: add validation
	err = pc.Repository.Create(p)
	return
}
