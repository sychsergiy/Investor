package asset

import (
	"testing"
)

func TestUnmarshaler_Unmarshal(t *testing.T) {
	unmarshaler := Unmarshaler{}
	// todo: fix creation_date parsing
	payload := []byte("{\"1\": {\"category\":1, \"name\": \"test\", \"id\": \"1\"}}")
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
