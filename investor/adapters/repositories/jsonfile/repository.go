package jsonfile

import (
	"investor/adapters/repositories/in_memory"
	"investor/helpers/file"
	"log"
)

type RecordUnmarshaler interface {
	Unmarshal([]byte) (map[string]in_memory.Record, error)
}

type Repository struct {
	unmarshaler RecordUnmarshaler
	jsonFile    file.IJsonFile
	repository  in_memory.Repository
}

func (r Repository) Create(record in_memory.Record) (err error) {
	if err = r.repository.Create(record); err != nil {
		return
	} else {
		err = r.dump()
		return
	}
}

func (r Repository) CreateBulk(records []in_memory.Record) (int, error) {
	count, err := r.repository.CreateBulk(records)
	if err != nil {
		return count, err
	}
	err = r.dump()
	return count, err
}

func (r Repository) dump() error {
	created, err := file.CreateIfNotExists(r.jsonFile)
	if created {
		log.Printf("\nWARNING: file: %s doesn't exists. Create empty.\n", r.jsonFile.Path())
	}
	err = r.jsonFile.WriteJson(r.repository.Records)
	return err
}

func (r Repository) restore() error {
	created, err := file.CreateIfNotExists(r.jsonFile)
	if created {
		log.Printf("\nWARNING: file: %s doesn't exists. Create empty.\n", r.jsonFile.Path())
	}
	content, err := r.jsonFile.Read()
	if err != nil {
		return err
	}

	recordsMap, err := r.unmarshaler.Unmarshal(content)

	var records []in_memory.Record
	for _, value := range recordsMap {
		records = append(records, value)
	}

	if err != nil {
		return err
	} else {
		_, err := r.repository.CreateBulk(records)
		if err != nil {
			return err
		}
	}
	return nil
}

func NewRepository(jsonFile file.IJsonFile, unmarshaler RecordUnmarshaler) Repository {
	repo := Repository{unmarshaler, jsonFile, in_memory.NewRepository()}
	err := repo.restore()
	if err != nil {
		log.Fatal(err)
	}
	return repo
}
