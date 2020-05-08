package asset

type PlainAsset struct {
	id       string
	category Category
	name     string
}

func (a PlainAsset) Id() string {
	return a.id
}

func (a PlainAsset) Category() Category {
	return a.category
}

func (a PlainAsset) Name() string {
	return a.name
}

func NewPlainAsset(id string, category Category, name string) *PlainAsset {
	return &PlainAsset{id, category, name}
}
