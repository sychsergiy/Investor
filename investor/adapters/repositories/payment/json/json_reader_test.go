package json

import (
	"fmt"
	"testing"
)

type FileReaderMock struct {
	readFunc func() ([]byte, error)
}

func (mock FileReaderMock) Read() ([]byte, error) {
	return mock.readFunc()
}

func TestPaymentsJsonFileReader_Read(t *testing.T) {
	readerMock := FileReaderMock{func() ([]byte, error) {
		return nil, fmt.Errorf("test error")
	}}
	reader := PaymentsJsonFileReader{readerMock}
	content, err := reader.Read()
	// todo: check proper error(err != customError()) instead of nil
	if err == nil {
		t.Errorf("File read failure error expected here")
	}
	if content != nil {
		t.Errorf("No content expecte here because of file read failure")
	}

	readerMock = FileReaderMock{func() ([]byte, error) {
		return []byte("[]"), nil
	}}
	reader = PaymentsJsonFileReader{readerMock}
	content, err = reader.Read()
	if err != nil {
		t.Errorf("Unexpected failure read failure error")
	} else {
		if len(content) != 0 {
			t.Errorf("Empty payments list expected here")
		}
	}

	readerMock = FileReaderMock{func() ([]byte, error) {
		return []byte("[{\"id\": \"1\", \"asset_amount\": 5, \"absolute_amount\": 10, \"asset_id\": \"2\", \"type\": 0, \"creation_date\": \"2019-10-9 1:2:3\"}]"), nil
	}}
	reader = PaymentsJsonFileReader{readerMock}
	content, err = reader.Read()
	if err != nil {
		t.Errorf("Unexpected json read error")
	}
	if len(content) != 1 {
		t.Errorf("One payment record expected to be parse")
	} else {
		expectedRecord := PaymentRecord{
			Id: "1", AssetAmount: 5, AbsoluteAmount: 10,
			AssetId: "2", CreationDate: "2019-10-9 1:2:3",
		}
		if content[0] != expectedRecord {
			t.Errorf("Not epxected unmarshalled payment record")
		}
	}
}
