package asset

type Plain struct {
	id       string
	category Category
	name     string
}

func (a Plain) ID() string {
	return a.id
}

func (a Plain) Category() Category {
	return a.category
}

func (a Plain) Name() string {
	return a.name
}

func NewPlain(id string, category Category, name string) *Plain {
	return &Plain{id, category, name}
}
