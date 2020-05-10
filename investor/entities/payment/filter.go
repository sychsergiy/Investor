package payment

type Filter interface {
	Payments() ([]Payment, error)
}
