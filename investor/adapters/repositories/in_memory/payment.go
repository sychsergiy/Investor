package in_memory

import (
	"fmt"
	"investor/entities/asset"
	paymentEntity "investor/entities/payment"
	"sort"
	"time"
)

type AssetFinderById interface {
	FindById(id string) (asset.Asset, error)
}

type PaymentRecord struct {
	Id             string             `json:"id"`
	AssetAmount    float32            `json:"asset_amount"`
	AbsoluteAmount float32            `json:"absolute_amount"`
	AssetId        string             `json:"asset_id"`
	Type           paymentEntity.Type `json:"type"`
	CreationDate   time.Time          `json:"creation_date"`
}

func NewPaymentRecord(payment paymentEntity.Payment) PaymentRecord {
	return PaymentRecord{
		Id:             payment.Id(),
		AssetAmount:    payment.AssetAmount(),
		AbsoluteAmount: payment.AbsoluteAmount(),
		AssetId:        payment.Asset().Id,
		Type:           payment.Type(),
		CreationDate:   payment.CreationDate(),
	}
}

type PaymentAlreadyExistsError struct {
	PaymentId string
}

type PaymentBulkCreateError struct {
	FailedIndex int
	Quantity    int
	Err         error
}

func (e PaymentBulkCreateError) Error() string {
	return fmt.Sprintf(
		"payments bulk create %d items failed on index: %d due to err: %v",
		e.Quantity, e.FailedIndex, e.Err,
	)
}

func (e PaymentBulkCreateError) Unwrap() error {
	return e.Err
}

func (e PaymentAlreadyExistsError) Error() string {
	return fmt.Sprintf("payment with id %s already exists", e.PaymentId)

}

type PaymentRepository struct {
	assetFinder AssetFinderById
	records     map[string]PaymentRecord
}

func (r *PaymentRepository) Create(payment paymentEntity.Payment) error {
	_, err := r.assetFinder.FindById(payment.Asset().Id)
	if err != nil {
		return err
	}
	_, idExists := r.records[payment.Id()]
	if idExists {
		return PaymentAlreadyExistsError{PaymentId: payment.Id()}
	} else {
		r.records[payment.Id()] = NewPaymentRecord(payment)
		return nil
	}
}

func (r *PaymentRepository) CreateBulk(payments []paymentEntity.Payment) error {
	for createdCount, payment := range payments {
		_, err := r.assetFinder.FindById(payment.Asset().Id)
		if err != nil {
			return PaymentBulkCreateError{
				FailedIndex: createdCount,
				Quantity:    len(payments),
				Err:         err,
			}
		}

		_, idExists := r.records[payment.Id()]
		if idExists {
			return PaymentBulkCreateError{
				FailedIndex: createdCount,
				Quantity:    len(payments),
				Err:         PaymentAlreadyExistsError{PaymentId: payment.Id()},
			}
		} else {
			r.records[payment.Id()] = NewPaymentRecord(payment)
		}
	}
	return nil
}

func (r *PaymentRepository) ListAll() ([]paymentEntity.Payment, error) {
	var payments []paymentEntity.Payment
	for _, record := range r.records {
		p, err := r.ConvertRecordToEntity(record)
		if err != nil {
			return nil, fmt.Errorf("failed to list payments: %w", err)
		}
		payments = append(payments, p)
	}
	sort.Slice(payments, func(i, j int) bool {
		return payments[i].CreationDate().After(payments[j].CreationDate())
	})
	return payments, nil
}

func (r *PaymentRepository) Records() (records []PaymentRecord) {
	for _, record := range r.records {
		records = append(records, record)
	}
	return
}

func (r *PaymentRepository) ConvertRecordToEntity(record PaymentRecord) (p paymentEntity.Payment, err error) {
	a, err := r.findAssetById(record.AssetId)
	if err != nil {
		err = fmt.Errorf("join asset to payment failed: %w", err)
		return
	}
	p = paymentEntity.NewPlainPayment(
		record.Id, record.AssetAmount, record.AbsoluteAmount, a, record.CreationDate, record.Type,
	)
	return
}

func (r *PaymentRepository) findAssetById(assetId string) (asset.Asset, error) {
	return r.assetFinder.FindById(assetId)
}

func NewPaymentRepository(assetFinder AssetFinderById) *PaymentRepository {
	return &PaymentRepository{assetFinder, make(map[string]PaymentRecord)}
}
