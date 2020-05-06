package payment

import (
	"encoding/json"
	"investor/entities/payment"
)

type Unmarshaler struct {
}

func (pu Unmarshaler) Unmarshal(content []byte) (map[string]payment.Payment, error) {
	var paymentsMap map[string]payment.Payment
	err := json.Unmarshal(content, &paymentsMap)
	if err != nil {
		return nil, err
	}
	return paymentsMap, err
}
