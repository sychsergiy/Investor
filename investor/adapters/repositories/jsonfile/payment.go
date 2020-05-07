package jsonfile

import (
	"fmt"
	"investor/adapters/repositories/in_memory"
	paymentEntity "investor/entities/payment"
)

type PaymentRepository struct {
	storage    Storage
	repository in_memory.PaymentRepository
	restored   bool
}

func (r *PaymentRepository) CreateBulk(payments []paymentEntity.Payment) (int, error) {
	err := r.restore()
	if err != nil {
		return 0, err
	}

	n, err := r.repository.CreateBulk(payments)
	if err != nil {
		return n, fmt.Errorf("payments bulk create failed: %w", err)
	}
	err = r.dump()
	return n, err
}

func (r *PaymentRepository) Create(payment paymentEntity.Payment) error {
	err := r.restore()
	if err != nil {
		return err
	}

	err = r.repository.Create(payment)
	if err != nil {
		return fmt.Errorf("in memory create payment failed: %w", err)
	}
	return r.dump()
}

func (r *PaymentRepository) dump() error {
	err := r.storage.UpdatePayments(r.repository.ListAll())
	if err != nil {
		err = fmt.Errorf("update payments on json storage failed: %w", err)
	}
	return err
}

func (r *PaymentRepository) restore() error {
	if r.restored {
		// should be called only once to sync in_memory storage with file
		return nil
	}
	// read payments from storage file and save in memory
	payments, err := r.storage.RetrievePayments()
	if err != nil {
		return err
	}
	_, err = r.repository.CreateBulk(payments)
	if err != nil {
		err = fmt.Errorf("restore payments failed, storage file malformed: %w", err)
	} else {
		r.restored = true
	}
	return err
}

func (r *PaymentRepository) ListAll() []paymentEntity.Payment {
	// todo: add error to  ListAll() interface
	_ = r.restore()
	return r.repository.ListAll()
}

func NewPaymentRepository(storage Storage) *PaymentRepository {
	return &PaymentRepository{storage, *in_memory.NewPaymentRepository(), false}
}
