package in_memory

import (
	"fmt"
	"investor/entities/asset"
	paymentEntity "investor/entities/payment"
	"sort"
	"time"
)

type AssetFinderById interface {
	FindById(id string) (*asset.Asset, error)
}

type PaymentRecord struct {
	Id             string             `json:"id"`
	AssetAmount    float32            `json:"asset_amount"`
	AbsoluteAmount float32            `json:"absolute_amount"`
	AssetId        string             `json:"asset_id"`
	Type           paymentEntity.Type `json:"type"`
	CreationDate   time.Time          `json:"creation_date"`
}

type PaymentProxy struct {
	p           PaymentRecord
	assetFinder AssetFinderById
}

func (p PaymentProxy) Id() string {
	return p.p.Id
}

func (p PaymentProxy) AssetAmount() float32 {
	return p.p.AssetAmount
}

func (p PaymentProxy) AbsoluteAmount() float32 {
	return p.p.AbsoluteAmount
}

func (p PaymentProxy) Asset() (*asset.Asset, error) {
	return p.assetFinder.FindById(p.p.AssetId)
}

func (p PaymentProxy) CreationDate() time.Time {
	return p.p.CreationDate
}

func (p PaymentProxy) Type() paymentEntity.Type {
	return p.p.Type
}
func NewPaymentProxy(record PaymentRecord, assetFinder AssetFinderById) PaymentProxy {
	return PaymentProxy{record, assetFinder}
}

type PaymentProxyFactory struct {
	assetFinder AssetFinderById
}

func (f PaymentProxyFactory) Create(record PaymentRecord) PaymentProxy {
	return NewPaymentProxy(record, f.assetFinder)
}

func NewPaymentRecord(payment paymentEntity.Payment) (pr PaymentRecord, err error) {
	a, err := payment.Asset()
	if err != nil {
		return
	}
	pr = PaymentRecord{
		Id:             payment.Id(),
		AssetAmount:    payment.AssetAmount(),
		AbsoluteAmount: payment.AbsoluteAmount(),
		AssetId:        a.Id,
		Type:           payment.Type(),
		CreationDate:   payment.CreationDate(),
	}
	return pr, nil
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
	factory PaymentProxyFactory
	//assetFinder AssetFinderById
	records map[string]PaymentRecord
}

func (r *PaymentRepository) Create(payment paymentEntity.Payment) error {
	record, err := NewPaymentRecord(payment)
	if err != nil {
		return err
	}
	_, idExists := r.records[payment.Id()]
	if idExists {
		return PaymentAlreadyExistsError{PaymentId: payment.Id()}
	} else {
		r.records[payment.Id()] = record
		return nil
	}
}

func (r *PaymentRepository) CreateBulk(payments []paymentEntity.Payment) error {
	for createdCount, payment := range payments {
		record, err := NewPaymentRecord(payment)
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
			r.records[payment.Id()] = record
		}
	}
	return nil
}

func (r *PaymentRepository) ListAll() ([]paymentEntity.Payment, error) {
	var payments []paymentEntity.Payment
	for _, record := range r.records {
		paymentProxy := r.factory.Create(record)
		payments = append(payments, paymentProxy)
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

func (r *PaymentRepository) ConvertRecordToEntity(record PaymentRecord) (p paymentEntity.Payment) {
	return r.factory.Create(record)
}

func NewPaymentRepository(assetFinder AssetFinderById) *PaymentRepository {
	return &PaymentRepository{
		PaymentProxyFactory{assetFinder: assetFinder},
		make(map[string]PaymentRecord),
	}
}
