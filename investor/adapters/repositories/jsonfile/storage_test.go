package jsonfile

import (
	"encoding/json"
	assetEntity "investor/entities/asset"
	"investor/entities/payment"
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
	asset := assetEntity.Asset{Id: "1", Category: assetEntity.PreciousMetal, Name: "test"}
	pMock := payment.CreatePayment("1", 2020)
	return Data{AssetsMap{"1": asset}, PaymentsMap{"1": pMock}}
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
			a, ok := assets["1"]
			if !ok {
				t.Errorf("Expected key 1")
			} else {
				if a != data.Assets["1"] {
					t.Errorf("Asset malformed after unmarhaling")
				}
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
			p, ok := payments["1"]
			if !ok {
				t.Errorf("Expected key 1")
			} else {
				if p != data.Payments["1"] {
					t.Errorf("Payment malformed after unmarhaling")
				}
			}
		}
	}
}

func TestStorage_UpdateAssets(t *testing.T) {
	filename := "test_updates_assets.json"
	data := getDataMock()
	err := createStorage(filename).UpdateAssets(data.Assets)
	expectedJson := "{\"assets\":{\"1\":{\"Id\":\"1\",\"Category\":0,\"Name\":\"test\"}},\"payments\":{}}"
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
	expectedJson := "{\"assets\":{},\"payments\":{\"1\":{\"Id\":\"1\",\"AssetAmount\":0,\"AbsoluteAmount\":0,\"Asset\":{\"Id\":\"\",\"Category\":1,\"Name\":\"test\"},\"Type\":1,\"CreationDate\":\"2019-11-30T00:00:00Z\"}}}"
	if err != nil {
		t.Errorf("Unepxected err: %+v", err)
	} else {
		content := file.ReadFile(filename)
		if string(content) != expectedJson {
			t.Errorf("Asset malformed during writing to storage file")
		}
	}
}
