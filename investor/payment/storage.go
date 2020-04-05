package payment

type Identifier int

type Creator interface {
	Create(payment Payment) Identifier
}
