package ports

type IdGenerator interface {
	Generate() string
}
