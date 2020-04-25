package ports

import (
	"investor/entities/payment"
)

type Identifier int

type PaymentSaver interface {
	Create(payment payment.Payment) error
}
