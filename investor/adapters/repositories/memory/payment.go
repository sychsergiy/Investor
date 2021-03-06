package memory

import (
	"fmt"
	"investor/entities/asset"
	"investor/entities/payment"
	"sort"
	"time"
)

type AssetFinderByID interface {
	FindByID(id string) (asset.Asset, error)
}

type PaymentRecord struct {
	ID             string       `json:"id"`
	AssetAmount    float32      `json:"asset_amount"`
	AbsoluteAmount float32      `json:"absolute_amount"`
	AssetID        string       `json:"asset_id"`
	Type           payment.Type `json:"type"`
	CreationDate   time.Time    `json:"creation_date"`
}

type PaymentProxy struct {
	record      PaymentRecord
	assetFinder AssetFinderByID
}

func (p PaymentProxy) ID() string {
	return p.record.ID
}

func (p PaymentProxy) AssetAmount() float32 {
	return p.record.AssetAmount
}

func (p PaymentProxy) AbsoluteAmount() float32 {
	return p.record.AbsoluteAmount
}

func (p PaymentProxy) Asset() (asset.Asset, error) {
	return p.assetFinder.FindByID(p.record.AssetID)
}

func (p PaymentProxy) CreationDate() time.Time {
	return p.record.CreationDate
}

func (p PaymentProxy) Type() payment.Type {
	return p.record.Type
}
func NewPaymentProxy(record PaymentRecord, assetFinder AssetFinderByID) PaymentProxy {
	return PaymentProxy{record, assetFinder}
}

func (r PaymentRepository) createPaymentProxy(record PaymentRecord) PaymentProxy {
	return NewPaymentProxy(record, r.assetFinder)
}

func NewPaymentRecord(payment payment.Payment) (pr PaymentRecord, err error) {
	a, err := payment.Asset()
	if err != nil {
		return pr, err
	}
	pr = PaymentRecord{
		ID:             payment.ID(),
		AssetAmount:    payment.AssetAmount(),
		AbsoluteAmount: payment.AbsoluteAmount(),
		AssetID:        a.ID(),
		Type:           payment.Type(),
		CreationDate:   payment.CreationDate(),
	}
	return pr, nil
}

type PaymentAlreadyExistsError struct {
	PaymentID string
}

type PaymentBulkCreateError struct {
	FailedIndex int
	Quantity    int
	Err         error
}

type PaymentDoesntExistsError struct {
	PaymentID string
}

func (e PaymentDoesntExistsError) Error() string {
	return fmt.Sprintf("payment with id %s not found", e.PaymentID)
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
	return fmt.Sprintf("payment with id %s already exists", e.PaymentID)

}

type PaymentRepository struct {
	assetFinder AssetFinderByID
	records     map[string]PaymentRecord
}

func (r *PaymentRepository) Create(payment payment.Payment) error {
	record, err := NewPaymentRecord(payment)
	if err != nil {
		return err
	}
	_, err = r.assetFinder.FindByID(record.AssetID)
	if err != nil {
		return err
	}
	_, idExists := r.records[payment.ID()]
	if idExists {
		return PaymentAlreadyExistsError{PaymentID: payment.ID()}
	}
	r.records[payment.ID()] = record
	return nil
}

func (r *PaymentRepository) CreateBulk(payments []payment.Payment) error {
	for createdCount, p := range payments {
		err := r.Create(p)
		if err != nil {
			return PaymentBulkCreateError{
				FailedIndex: createdCount,
				Quantity:    len(payments),
				Err:         err,
			}
		}
	}
	return nil
}

func (r *PaymentRepository) ListAll() ([]payment.Payment, error) {
	var payments []payment.Payment
	for _, record := range r.records {
		paymentProxy := r.createPaymentProxy(record)
		payments = append(payments, paymentProxy)
	}

	sortByCreationDate(payments)
	return payments, nil
}

func sortByCreationDate(payments []payment.Payment) []payment.Payment {
	sort.Slice(payments, func(i, j int) bool {
		payments[i].CreationDate()
		return payments[i].CreationDate().After(payments[j].CreationDate())
	})
	return payments
}

func (r *PaymentRepository) FindByAssetCategories(
	categories []asset.Category,
	periods []payment.Period,
	paymentTypes []payment.Type,
) ([]payment.Payment, error) {
	payments, err := r.ListAll()
	if err != nil {
		return nil, err
	}
	filter := payment.NewFilter(payments)
	filter, err = filter.ByAssetCategories(categories)
	if err != nil {
		return nil, err
	}
	filtered := filter.ByPeriods(periods).ByTypes(paymentTypes).Payments()
	sortByCreationDate(filtered)
	return filtered, nil
}

func (r *PaymentRepository) FindByAssetNames(
	assetNames []string,
	periods []payment.Period,
	paymentTypes []payment.Type,
) ([]payment.Payment, error) {
	payments, err := r.ListAll()
	if err != nil {
		return nil, err
	}
	filter := payment.NewFilter(payments)
	filter, err = filter.ByAssetNames(assetNames)
	if err != nil {
		return nil, err
	}
	payments = filter.ByPeriods(periods).ByTypes(paymentTypes).Payments()
	sortByCreationDate(payments)
	return payments, nil
}

func (r *PaymentRepository) FindByIDs(ids []string) ([]payment.Payment, error) {
	var payments []payment.Payment

	for _, id := range ids {
		p, ok := r.records[id]
		if !ok {
			return nil, PaymentDoesntExistsError{PaymentID: id}
		}
		payments = append(payments, r.createPaymentProxy(p))
	}
	sortByCreationDate(payments)
	return payments, nil
}

func (r *PaymentRepository) Records() (records []PaymentRecord) {
	for _, record := range r.records {
		records = append(records, record)
	}
	return
}

func (r *PaymentRepository) ConvertRecordToEntity(record PaymentRecord) (p payment.Payment) {
	return r.createPaymentProxy(record)
}

func NewPaymentRepository(assetFinder AssetFinderByID) *PaymentRepository {
	return &PaymentRepository{
		assetFinder,
		make(map[string]PaymentRecord),
	}
}
