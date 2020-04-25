package ports

import "investor/entities"

type Identifier int

type Creator interface {
	Create(payment entities.Payment) Identifier
}
