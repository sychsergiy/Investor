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
	Id       string
	Category Category
	Name     string
}

func NewAsset(id string, category Category, name string) Asset {
	return Asset{id, category, name}
}
