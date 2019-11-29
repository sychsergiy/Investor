package payment

import "time"

type Type int

const (
	Invest Type = iota
	Return
)

func (t Type) String() string {
	return [...]string{"Invest", "Return"}[t]
}

type Payment struct {
	AbsoluteAmount float32 // amount in dollars
	CurrencyAmount float32 // amount in crypto currency
	Date           time.Time
	Type           Type
}

func NewInvestPayment(absAmount, currAmount float32, date time.Time) Payment {
	return Payment{absAmount, currAmount, date, Invest}
}

func NewReturnPayment(absAmount, currAmount float32, date time.Time) Payment {
	return Payment{absAmount, currAmount, date, Return}
}
