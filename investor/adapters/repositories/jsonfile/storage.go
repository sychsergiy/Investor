package jsonfile

import (
	"encoding/json"
	"fmt"
	"investor/adapters/repositories/in_memory"
	"investor/helpers/file"
	"log"
)

type Data struct {
	Assets   []in_memory.AssetRecord   `json:"assets"`
	Payments []in_memory.PaymentRecord `json:"payments"`
}

type Storage struct {
	jsonFile file.IJSONFile
	data     *Data
}

func (s Storage) RetrievePayments() ([]in_memory.PaymentRecord, error) {
	err := s.restore()
	if err != nil {
		return nil, err
	}
	return s.data.Payments, nil

}

func (s Storage) RetrieveAssets() ([]in_memory.AssetRecord, error) {
	err := s.restore()
	if err != nil {
		return nil, err
	}
	return s.data.Assets, nil

}

func (s *Storage) UpdatePayments(payments []in_memory.PaymentRecord) error {
	err := s.restore()
	if err != nil {
		return err
	}
	s.data.Payments = payments
	return s.dump()
}

func (s *Storage) UpdateAssets(assets []in_memory.AssetRecord) error {
	err := s.restore()
	if err != nil {
		return err
	}
	s.data.Assets = assets
	return s.dump()
}

func (s Storage) dump() error {
	err := s.ensureFileExists()
	if err != nil {
		return fmt.Errorf("storage dump: %w", err)
	}

	err = s.jsonFile.WriteJSON(s.data)
	if err != nil {
		return fmt.Errorf("writing json storage file: %w", err)
	}
	return nil
}

func (s Storage) ensureFileExists() error {
	created, err := file.CreateIfNotExists(s.jsonFile)
	if created {
		log.Printf("\nWARNING: file: %s doesn't exists. Create empty.\n", s.jsonFile.Path())
		return s.jsonFile.WriteJSON(s.data)
	}
	if err != nil {
		return fmt.Errorf("ensure file exists: %w", err)
	}
	return err
}

func (s *Storage) restore() error {
	err := s.ensureFileExists()
	if err != nil {
		return fmt.Errorf("storage restore: %w", err)
	}

	var data Data
	content, err := s.jsonFile.Read()
	if err != nil {
		return fmt.Errorf("reading json storage file: %w", err)
	}

	err = json.Unmarshal(content, &data)
	if err != nil {
		return fmt.Errorf("unmarshaling json storage file content: %w", err)
	}
	s.data.Payments = data.Payments
	s.data.Assets = data.Assets
	return nil
}

func NewStorage(jsonFile file.JSONFile) *Storage {
	return &Storage{
		jsonFile: jsonFile,
		data:     &Data{[]in_memory.AssetRecord{}, []in_memory.PaymentRecord{}},
	}
}
