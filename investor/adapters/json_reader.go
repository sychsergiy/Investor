package adapters

import (
	"encoding/json"
	"investor/entities/payment"
	"io/ioutil"
	"os"
)

type Reader interface {
	Read() ([]byte, error)
}

type FileReader struct {
	Path string
}

type PaymentRecord struct {
	Id             string       `json:"id"`
	AssetAmount    float32      `json:"asset_amount"`
	AbsoluteAmount float32      `json:"absolute_amount"`
	AssetId        string       `json:"asset_id"` // todo: asset id
	Type           payment.Type `json:"type"`
	CreationDate   string       `json:"creation_date"`
}

func (reader FileReader) Read() ([]byte, error) {
	file, err := os.Open(reader.Path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return byteValue, err
}

type PaymentsJsonFileReader struct {
	FileReader Reader
}

func (reader PaymentsJsonFileReader) Read() ([]PaymentRecord, error) {
	content, err := reader.FileReader.Read()

	if err != nil {
		return nil, err
	}

	var payments []PaymentRecord

	err = json.Unmarshal(content, &payments)
	if err != nil {
		return nil, err
	}
	return payments, nil
}
