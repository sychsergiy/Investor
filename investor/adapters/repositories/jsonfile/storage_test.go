package jsonfile

import (
	"encoding/json"
	"investor/adapters/repositories/in_memory"
	"investor/helpers/file"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	file.CreateWorkDir()
	code := m.Run()
	file.CleanupWorkDir()
	os.Exit(code)
}

func writeStorageFile(t *testing.T, filename string, data Data) {
	content, _ := json.Marshal(data)
	file.WriteBytesToFile(t, filename, content)
}

func createStorage(filename string) *Storage {
	jsonFile := file.NewJsonFile(file.NewPlainFile(file.GetFilePath(filename)))
	return NewStorage(jsonFile)
}

func getDataMock() Data {
	asset := in_memory.CreateAssetRecord("1", "test")
	pMock := in_memory.CreatePaymentRecord("1", 2020)
	return Data{[]in_memory.AssetRecord{asset}, []in_memory.PaymentRecord{pMock}}
}

func TestStorage_RetrieveAssets(t *testing.T) {
	filename := "test_retrieve_assets.json"
	data := getDataMock()
	writeStorageFile(t, filename, data)
	assets, err := createStorage(filename).RetrieveAssets()

	if err != nil {
		t.Errorf("Unexpected err: %+v", err)
	} else {
		if len(assets) != 1 {
			t.Errorf("One asset epxected")
			if assets[0] != data.Assets[0] {
				t.Errorf("Asset malformed after unmarhaling")
			}
		}
	}
}

func TestStorage_RetrievePayments(t *testing.T) {
	filename := "test_retrieve_payments.json"
	data := getDataMock()
	writeStorageFile(t, filename, data)
	payments, err := createStorage(filename).RetrievePayments()
	if err != nil {
		t.Errorf("Unexpected err: %+v", err)
	} else {
		if len(payments) != 1 {
			t.Errorf("One payment epxected")
		} else {
			if payments[0] != data.Payments[0] {
				t.Errorf("Payment malformed after unmarhaling")
			}
		}
	}
}

func TestStorage_UpdateAssets(t *testing.T) {
	filename := "test_updates_assets.json"
	data := getDataMock()
	err := createStorage(filename).UpdateAssets(data.Assets)
	expectedJson := "{\"assets\":[{\"id\":\"1\",\"category\":0,\"name\":\"test\"}],\"payments\":[]}"
	if err != nil {
		t.Errorf("Unepxected err: %+v", err)
	} else {
		content := file.ReadFile(filename)
		if string(content) != expectedJson {
			t.Errorf("Asset malformed during writing to storage file")
		}
	}
}

func TestStorage_UpdatePayments(t *testing.T) {
	filename := "test_update_payments.json"
	data := getDataMock()
	err := createStorage(filename).UpdatePayments(data.Payments)
	expectedJson := "{\"assets\":[],\"payments\":[{\"id\":\"1\",\"asset_amount\":50,\"absolute_amount\":100,\"asset_id\":\"testAssetID\",\"type\":0,\"creation_date\":\"2019-11-30T00:00:00Z\"}]}"
	if err != nil {
		t.Errorf("Unepxected err: %+v", err)
	} else {
		content := file.ReadFile(filename)
		if string(content) != expectedJson {
			t.Errorf("Asset malformed during writing to storage file")
		}
	}
}
