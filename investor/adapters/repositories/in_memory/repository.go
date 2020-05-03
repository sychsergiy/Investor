package in_memory

import (
	"fmt"
)

type Record interface {
	Id() string
}

type Repository struct {
	Records map[string]Record
}

type RecordAlreadyExistsError struct {
	RecordId string
}

func (e RecordAlreadyExistsError) Error() string {
	return fmt.Sprintf("payment with id %s already exists", e.RecordId)

}

func (r Repository) Create(record Record) error {
	_, idExists := r.Records[record.Id()]
	if idExists {
		return RecordAlreadyExistsError{RecordId: record.Id()}
	} else {
		r.Records[record.Id()] = record
		return nil
	}
}

func (r Repository) CreateBulk(records []Record) (int, error) {
	var createdCount int
	for createdCount, record := range records {
		_, idExists := r.Records[record.Id()]
		if idExists {
			return createdCount, RecordAlreadyExistsError{RecordId: record.Id()}
		} else {
			r.Records[record.Id()] = record
		}
	}
	return createdCount, nil
}

func NewRepository() Repository {
	return Repository{Records: make(map[string]Record)}
}
