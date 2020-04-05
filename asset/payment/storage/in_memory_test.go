package storage

import (
	"Investor/asset/payment"
	"testing"
	"time"
)

func TestInMemory_DeletePayment(t *testing.T) {
	s := NewInMemory()
	paymentID := 1
	s.payments[paymentID] = payment.NewInvestPayment(1, 1, time.Now())

	if result := s.DeletePayment(paymentID); result != true {
		t.Errorf("Trying to remove existstent item, result must be true", )
	}
	if _, ok := s.payments[paymentID]; ok {
		t.Errorf("Payment with id = %d must be removed from paymantes map", paymentID)
	}

	if result := s.DeletePayment(0); result != false {
		t.Errorf("Trying to remove not existstent item, result must be false", )
	}
}

func TestInMemory_SavePayment(t *testing.T) {
	s := NewInMemory()
	p := payment.NewInvestPayment(1, 1, time.Now())

	s.SavePayment(p)
	if savedPayment := s.payments[0]; savedPayment != p {
		t.Errorf("Payment must be saved by key: %d", 0)
	}

	s.SavePayment(p)
	if savedPayment := s.payments[1]; savedPayment != p {
		t.Errorf("Second payment must be saved by key: %d", 1)
	}

	if len(s.payments) != 2 {
		t.Errorf("Two payments must be saved to so map")
	}
}

func TestInMemory_RetrieveAllPayments(t *testing.T) {
	s := NewInMemory()
	p1 := payment.NewInvestPayment(1, 1, time.Now())
	p2 := payment.NewInvestPayment(1, 1, time.Now())

	s.payments[0] = p1
	s.payments[1] = p2

	result := s.RetrieveAllPayments()
	if len(result) != 2 {
		t.Errorf("Two payments must be retrieved from map")
	}
}
