package asset

type Category int

const (
	PreciousMetal Category = iota
	CryptoCurrency
	Stock
)

func (c Category) String() string {
	return [...]string{"PreciousMetal", "CryptoCurrency", "Stock"}[c]
}

type Asset struct {
	Category Category
	Name     string
}
