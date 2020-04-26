package json

import (
	"encoding/json"
	"investor/adapters/repositories"
	"investor/adapters/repositories/payment/in_memory"
)

type PaymentUnmarshaler struct {
}

func (pu PaymentUnmarshaler) Unmarshal(content []byte) (map[string]repositories.Record, error) {
	var paymentsMap map[string]in_memory.PaymentRecord
	err := json.Unmarshal(content, &paymentsMap)
	if err != nil {
		return nil, err
	}
	recordsMap := createRecordsMap(paymentsMap)
	return recordsMap, err
}

func createRecordsMap(payments map[string]in_memory.PaymentRecord) map[string]repositories.Record {
	recordsMap := make(map[string]repositories.Record)
	for key, value := range payments {
		recordsMap[key] = value
	}
	return recordsMap
}
