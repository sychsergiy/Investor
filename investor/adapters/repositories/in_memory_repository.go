package repositories

import (
	"fmt"
)

type Record interface {
	Id() string
}

type InMemoryRepository struct {
	records map[string]Record
}

type RecordAlreadyExistsError struct {
	RecordId string
}

func (e RecordAlreadyExistsError) Error() string {
	return fmt.Sprintf("payment with id %s already exists", e.RecordId)

}

func (r InMemoryRepository) Create(record Record) error {
	_, idExists := r.records[record.Id()]
	if idExists {
		return RecordAlreadyExistsError{RecordId: record.Id()}
	} else {
		r.records[record.Id()] = record
		return nil
	}
}

func (r InMemoryRepository) CreateBulk(records []Record) (int, error) {
	var createdCount int
	for createdCount, record := range records {
		_, idExists := r.records[record.Id()]
		if idExists {
			return createdCount, RecordAlreadyExistsError{RecordId: record.Id()}
		} else {
			r.records[record.Id()] = record
		}
	}
	return createdCount, nil
}

func NewInMemoryRepository() InMemoryRepository {
	return InMemoryRepository{records: make(map[string]Record)}
}
