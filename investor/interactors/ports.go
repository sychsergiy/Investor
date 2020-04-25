package interactors

import (
	paymentEntity "investor/entities/payment"
)

type PaymentSaver interface {
	Create(payment paymentEntity.Payment) error
}

type IdGenerator interface {
	Generate() string
}
