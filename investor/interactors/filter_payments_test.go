package interactors

import (
	"investor/entities/payment"
	"reflect"
	"testing"
)

func TestFilterPayments_Filter(t *testing.T) {
	payments := []payment.Payment{
		payment.CreatePayment("1", 2020),
		payment.CreatePayment("2", 2020),
	}

	factory := NewPaymentIdsFilterFactory(
		PaymentFinderByIdsMock{FindByIdsFunc: func(ids []string) ([]payment.Payment, error) {
			return payments, nil
		}})
	interactor := NewFilterPayments(factory)

	req := FilterPaymentsRequest{[]string{"1", "2"}}
	resp, err := interactor.Filter(req)
	if err != nil {
		t.Errorf("Unexpected err: %+v", err)
	} else {
		if !reflect.DeepEqual(resp.Payments, payments) {
			t.Errorf("Unexpected payments value")
		}
	}
}
