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

type InvestPayment struct {
	Payment
}

type ReturnPayment struct {
	Payment
}

func NewInvestPayment(absAmount, currAmount float32, date time.Time) InvestPayment {
	return InvestPayment{Payment{absAmount, currAmount, date, Invest}}
}

func NewReturnPayment(absAmount, currencyAmount float32, date time.Time) ReturnPayment {
	return ReturnPayment{Payment{absAmount, currencyAmount, date, Return}}
}
