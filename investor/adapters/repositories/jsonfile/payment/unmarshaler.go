package payment

import (
	"encoding/json"
	"investor/adapters/repositories/in_memory"
	"investor/adapters/repositories/in_memory/payment"
)

type Unmarshaler struct {
}

func (pu Unmarshaler) Unmarshal(content []byte) (map[string]in_memory.Record, error) {
	var paymentsMap map[string]payment.Record
	err := json.Unmarshal(content, &paymentsMap)
	if err != nil {
		return nil, err
	}
	recordsMap := createRecordsMap(paymentsMap)
	return recordsMap, err
}

func createRecordsMap(payments map[string]payment.Record) map[string]in_memory.Record {
	recordsMap := make(map[string]in_memory.Record)
	for key, value := range payments {
		recordsMap[key] = value
	}
	return recordsMap
}
