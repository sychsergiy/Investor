package asset

import (
	"encoding/json"
	"investor/adapters/repositories/in_memory"
)

type Unmarshaler struct {
}

func (pu Unmarshaler) Unmarshal(content []byte) (map[string]in_memory.Record, error) {
	var paymentsMap map[string]in_memory.AssetRecord
	err := json.Unmarshal(content, &paymentsMap)
	if err != nil {
		return nil, err
	}
	recordsMap := createRecordsMap(paymentsMap)
	return recordsMap, err
}

func createRecordsMap(payments map[string]in_memory.AssetRecord) map[string]in_memory.Record {
	recordsMap := make(map[string]in_memory.Record)
	for key, value := range payments {
		recordsMap[key] = value
	}
	return recordsMap
}
