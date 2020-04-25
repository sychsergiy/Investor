package ports

import "investor/entities"

type Identifier int

type PaymentSaver interface {
	Save(payment entities.Payment) Identifier
}
