package ports

import "investor/entities"

type Identifier int

type PaymentSaver interface {
	Create(payment entities.Payment) error
}
