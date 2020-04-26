package repositories

import (
	"fmt"
)

type Entity interface {
	Id() string
}

type InMemoryRepository struct {
	records map[string]Entity
}

type EntityAlreadyExistsError struct {
	EntityId string
}

func (e EntityAlreadyExistsError) Error() string {
	return fmt.Sprintf("payment with id %s already exists", e.EntityId)

}

func (r InMemoryRepository) Create(entity Entity) error {
	_, idExists := r.records[entity.Id()]
	if idExists {
		return EntityAlreadyExistsError{EntityId: entity.Id()}
	} else {
		r.records[entity.Id()] = entity
		return nil
	}
}

func (r InMemoryRepository) CreateBulk(entities []Entity) (int, error) {
	var createdCount int
	for createdCount, entity := range entities {
		_, idExists := r.records[entity.Id()]
		if idExists {
			return createdCount, EntityAlreadyExistsError{EntityId: entity.Id()}
		} else {
			r.records[entity.Id()] = entity
		}
	}
	return createdCount, nil
}

func NewInMemoryCreateRepository() InMemoryRepository {
	return InMemoryRepository{records: make(map[string]Entity)}
}
