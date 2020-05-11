package jsonfile

import (
	"fmt"
	"investor/adapters/repositories/memory"
	"investor/entities/asset"
	paymentEntity "investor/entities/payment"
)

type PaymentRepository struct {
	storage    *Storage
	repository *memory.PaymentRepository
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
		// should be called only once to sync memory storage with file
		return nil
	}
	// read payments from storage file and save in memory
	records, err := r.storage.RetrievePayments()
	if err != nil {
		return err
	}

	payments := r.convertRecordsToEntities(records)

	err = r.repository.CreateBulk(payments)
	if err != nil {
		err = fmt.Errorf("restore payments failed, storage file malformed: %w", err)
	} else {
		r.restored = true
	}
	return err
}

func (r *PaymentRepository) convertRecordsToEntities(records []memory.PaymentRecord) (
	payments []paymentEntity.Payment,
) {
	for _, record := range records {
		payment := r.repository.ConvertRecordToEntity(record)
		payments = append(payments, payment)
	}
	return payments
}

func (r *PaymentRepository) ListAll() ([]paymentEntity.Payment, error) {
	err := r.restore()
	if err != nil {
		return nil, fmt.Errorf("failed to list all payments: %w", err)
	}
	return r.repository.ListAll()
}

func (r *PaymentRepository) FindByIDs(ids []string) ([]paymentEntity.Payment, error) {
	err := r.restore()
	if err != nil {
		return nil, fmt.Errorf("failed to find payments by ids: %v", err)
	}
	return r.repository.FindByIDs(ids)
}

func (r *PaymentRepository) FindByAssetNames(
	assetNames []string,
	periods []paymentEntity.Period,
	paymentTypes []paymentEntity.Type,
) ([]paymentEntity.Payment, error) {
	err := r.restore()
	if err != nil {
		return nil, err
	}
	return r.repository.FindByAssetNames(assetNames, periods, paymentTypes)
}

func (r *PaymentRepository) FindByAssetCategories(
	categories []asset.Category,
	periods []paymentEntity.Period,
	paymentTypes []paymentEntity.Type,
) (filtered []paymentEntity.Payment, err error) {
	err = r.restore()
	if err != nil {
		return nil, err
	}
	return r.repository.FindByAssetCategories(
		categories, periods, paymentTypes,
	)
}

func NewPaymentRepository(storage *Storage, assetFinder memory.AssetFinderByID) *PaymentRepository {
	return &PaymentRepository{storage, memory.NewPaymentRepository(assetFinder), false}
}
