package jsonfile

import (
	"encoding/json"
	"fmt"
	"investor/entities/asset"
	"investor/entities/payment"
	"investor/helpers/file"
	"log"
)

type AssetsMap map[string]asset.Asset
type PaymentsMap map[string]payment.Payment

type Data struct {
	Assets   AssetsMap   `json:"assets"`
	Payments PaymentsMap `json:"payments"`
}

type Storage struct {
	jsonFile file.IJsonFile
	data     Data
}

func (s Storage) RetrievePayments() (map[string]payment.Payment, error) {
	err := s.restore()
	if err != nil {
		return nil, err
	}
	return s.data.Payments, nil

}

func (s Storage) RetrieveAssets() (map[string]asset.Asset, error) {
	err := s.restore()
	if err != nil {
		return nil, err
	}
	return s.data.Assets, nil

}

func (s *Storage) UpdatePayments(payments map[string]payment.Payment) error {
	s.data.Payments = payments
	return s.dump()
}

func (s *Storage) UpdateAssets(assets map[string]asset.Asset) error {
	s.data.Assets = assets
	return s.dump()
}

func (s Storage) dump() error {
	err := s.ensureFileExists()
	if err != nil {
		return fmt.Errorf("storage dump: %w", err)
	}

	err = s.jsonFile.WriteJson(s.data)
	if err != nil {
		return fmt.Errorf("writing json storage file: %w", err)
	}
	return nil
}

func (s Storage) ensureFileExists() error {
	created, err := file.CreateIfNotExists(s.jsonFile)
	if created {
		log.Printf("\nWARNING: file: %s doesn't exists. Create empty.\n", s.jsonFile.Path())
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
	s.data = data
	return nil
}

func NewStorage(jsonFile file.JsonFile) *Storage {
	return &Storage{
		jsonFile: jsonFile,
		data:     Data{make(map[string]asset.Asset), make(map[string]payment.Payment)},
	}
}
