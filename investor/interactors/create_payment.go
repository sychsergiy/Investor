package interactors

import (
	"fmt"
	"investor/entities"
	"investor/entities/asset"
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
	Asset          asset.Asset
	Type           entities.PaymentType
	CreationDate   time.Time
}

func (pc PaymentCreator) cratePaymentInstance(paymentModel CreatePaymentModel, id string) (payment entities.Payment) {
	if paymentModel.Type == entities.Return {
		payment = entities.NewReturnPayment(
			id, paymentModel.AssetAmount, paymentModel.AbsoluteAmount,
			paymentModel.Asset, paymentModel.CreationDate,
		)
	} else if paymentModel.Type == entities.Invest {
		payment = entities.NewInvestmentPayment(
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
	payment := pc.cratePaymentInstance(paymentModel, id)
	// todo: add validation
	err = pc.Storage.Create(payment)
	return
}
