package json

import (
	"testing"
)

func TestPaymentUnmarshaler_Unmarshal(t *testing.T) {
	unmarshaler := PaymentUnmarshaler{}
	// todo: fix creation_date parsing
	payload := []byte("{\"1\":{\"id\": \"1\", \"asset_amount\": 5, \"absolute_amount\": 10, \"asset_id\": \"2\", \"type\": 0, \"creation_date\": \"2019-10-9 1:2:3\"}}")
	records, err := unmarshaler.Unmarshal(payload)
	if err != nil {
		t.Errorf("Unexpected error during unmarshaling payment")
	}
	record, ok := records["1"]
	if !ok {
		t.Errorf("Expected payment with id 1, but not found in result map")
	}
	if record.Id() != "1" {
		t.Errorf("Wrong record id")
	}

	payload = []byte("{")
	records, err = unmarshaler.Unmarshal(payload)
	if records != nil {
		t.Errorf("Error and nil result exepcted here")
	}

	if err == nil {
		t.Errorf("Unamrshall error expected here")
	}
}
