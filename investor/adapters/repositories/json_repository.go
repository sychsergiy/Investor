package repositories

import (
	"investor/helpers/file"
	"log"
)

type RecordUnmarshaler interface {
	Unmarshal([]byte) (map[string]Record, error)
}

type JsonFileRepository struct {
	unmarshaler RecordUnmarshaler
	jsonFile    file.IJsonFile
	repository  InMemoryRepository
}

func (r JsonFileRepository) Create(record Record) (err error) {
	if err = r.repository.Create(record); err != nil {
		return
	} else {
		err = r.dump()
		return
	}
}

func (r JsonFileRepository) CreateBulk(records []Record) (int, error) {
	count, err := r.repository.CreateBulk(records)
	if err != nil {
		return count, err
	}
	err = r.dump()
	return count, err
}

func (r JsonFileRepository) dump() error {
	created, err := file.CreateIfNotExists(r.jsonFile)
	if created {
		log.Printf("\nWARNING: file: %s doesn't exists. Create empty.\n", r.jsonFile.Path())
	}
	err = r.jsonFile.WriteJson(r.repository.records)
	return err
}

func (r JsonFileRepository) restore() error {
	created, err := file.CreateIfNotExists(r.jsonFile)
	if created {
		log.Printf("\nWARNING: file: %s doesn't exists. Create empty.\n", r.jsonFile.Path())
	}
	content, err := r.jsonFile.Read()
	if err != nil {
		return err
	}

	recordsMap, err := r.unmarshaler.Unmarshal(content)

	var records []Record
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

func NewJsonFileRepository(jsonFile file.IJsonFile, unmarshaler RecordUnmarshaler) JsonFileRepository {
	repo := JsonFileRepository{unmarshaler, jsonFile, NewInMemoryRepository()}
	err := repo.restore()
	if err != nil {
		log.Fatal(err)
	}
	return repo
}
