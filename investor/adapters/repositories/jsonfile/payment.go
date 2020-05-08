package jsonfile

import (
	"fmt"
	"investor/adapters/repositories/in_memory"
	paymentEntity "investor/entities/payment"
)

type PaymentRepository struct {
	storage    *Storage
	repository *in_memory.PaymentRepository
	restored   bool
}

func (r *PaymentRepository) CreateBulk(payments []paymentEntity.Payment) error {
	err := r.restore()
	if err != nil {
		return err
	}

	err = r.repository.CreateBulk(payments)
	if err != nil {
		return fmt.Errorf("payments bulk create failed: %w", err)
	}
	err = r.dump()
	return err
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
	records := r.repository.Records()
	err := r.storage.UpdatePayments(records)
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
	records, err := r.storage.RetrievePayments()
	if err != nil {
		return err
	}

	payments, err := r.convertRecordsToEntities(records)
	if err != nil {
		return fmt.Errorf("failed convert reocrds to entities: %w", err)
	}

	err = r.repository.CreateBulk(payments)
	if err != nil {
		err = fmt.Errorf("restore payments failed, storage file malformed: %w", err)
	} else {
		r.restored = true
	}
	return err
}

func (r *PaymentRepository) convertRecordsToEntities(records []in_memory.PaymentRecord) (
	payments []paymentEntity.Payment, err error,
) {
	for _, record := range records {
		payment, err := r.repository.ConvertRecordToEntity(record)
		if err != nil {
			return payments, err
		} else {
			payments = append(payments, payment)
		}
	}
	return payments, err
}

func (r *PaymentRepository) ListAll() ([]paymentEntity.Payment, error) {
	err := r.restore()
	if err != nil {
		return nil, fmt.Errorf("failed to list all payments: %w", err)
	}
	return r.repository.ListAll()
}

func NewPaymentRepository(storage *Storage, assetFinder in_memory.AssetFinderById) *PaymentRepository {
	return &PaymentRepository{storage, in_memory.NewPaymentRepository(assetFinder), false}
}
