package ports

import "strconv"


type StubIdGenerator struct {
	counter int
}

func (sig StubIdGenerator) incrementCounter() {
	sig.counter += 1
}

func (sig StubIdGenerator) Generate() string {
	sig.incrementCounter()
	return strconv.Itoa(sig.counter)
}

func NewStubIdGenerator() StubIdGenerator {
	return StubIdGenerator{0}
}
