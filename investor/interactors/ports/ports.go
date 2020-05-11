package ports

type IDGenerator interface {
	Generate() string
}
