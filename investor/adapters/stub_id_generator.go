package adapters

import (
	uuid "github.com/satori/go.uuid"
)

type UUIDGenerator struct {
}

func (g UUIDGenerator) Generate() string {
	id := uuid.NewV4()
	return id.String()
}

func NewUUIDGenerator() UUIDGenerator {
	return UUIDGenerator{}
}
