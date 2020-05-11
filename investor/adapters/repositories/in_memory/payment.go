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

type PaymentProxy struct {
	record      PaymentRecord
	assetFinder AssetFinderById
}

func (p PaymentProxy) Id() string {
	return p.record.Id
}

func (p PaymentProxy) AssetAmount() float32 {
	return p.record.AssetAmount
}

func (p PaymentProxy) AbsoluteAmount() float32 {
	return p.record.AbsoluteAmount
}

func (p PaymentProxy) Asset() (asset.Asset, error) {
	return p.assetFinder.FindById(p.record.AssetId)
}

func (p PaymentProxy) CreationDate() time.Time {
	return p.record.CreationDate
}

func (p PaymentProxy) Type() paymentEntity.Type {
	return p.record.Type
}
func NewPaymentProxy(record PaymentRecord, assetFinder AssetFinderById) PaymentProxy {
	return PaymentProxy{record, assetFinder}
}

func (r PaymentRepository) createPaymentProxy(record PaymentRecord) PaymentProxy {
	return NewPaymentProxy(record, r.assetFinder)
}

func NewPaymentRecord(payment paymentEntity.Payment) (pr PaymentRecord, err error) {
	a, err := payment.Asset()
	if err != nil {
		return pr, err
	}
	pr = PaymentRecord{
		Id:             payment.Id(),
		AssetAmount:    payment.AssetAmount(),
		AbsoluteAmount: payment.AbsoluteAmount(),
		AssetId:        a.Id(),
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

type PaymentDoesntExistsError struct {
	PaymentId string
}

func (e PaymentDoesntExistsError) Error() string {
	return fmt.Sprintf("payment with id %s not found", e.PaymentId)
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
	record, err := NewPaymentRecord(payment)
	if err != nil {
		return err
	}
	_, err = r.assetFinder.FindById(record.AssetId)
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
		err := r.Create(payment)
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

func (r *PaymentRepository) ListAll() ([]paymentEntity.Payment, error) {
	var payments []paymentEntity.Payment
	for _, record := range r.records {
		paymentProxy := r.createPaymentProxy(record)
		payments = append(payments, paymentProxy)
	}

	sortByCreationDate(payments)
	return payments, nil
}

func sortByCreationDate(payments []paymentEntity.Payment) []paymentEntity.Payment {
	sort.Slice(payments, func(i, j int) bool {
		payments[i].CreationDate()
		return payments[i].CreationDate().After(payments[j].CreationDate())
	})
	return payments
}

func (r *PaymentRepository) FindByAssetCategories(
	categories []asset.Category,
	periods []paymentEntity.Period,
	paymentTypes []paymentEntity.Type,
) (filtered []paymentEntity.Payment, err error) {
	payments, err := r.ListAll()
	if err != nil {
		return nil, err
	}
	filtered, err = FilterByAssetCategory(payments, categories)
	if err != nil {
		return nil, err
	}
	filtered = FilterByPeriod(filtered, periods)
	filtered = FilterByType(filtered, paymentTypes)
	sortByCreationDate(filtered)
	return
}

func (r *PaymentRepository) FindByAssetNames(
	assetNames []string,
	periods []paymentEntity.Period,
	paymentTypes []paymentEntity.Type,
) ([]paymentEntity.Payment, error) {
	payments, err := r.ListAll()
	if err != nil {
		return nil, err
	}
	payments, err = FilterByAssetNames(payments, assetNames)
	if err != nil {
		return nil, err
	}
	payments = FilterByPeriod(payments, periods)
	payments = FilterByType(payments, paymentTypes)
	sortByCreationDate(payments)
	return payments, nil
}

func (r *PaymentRepository) FindByIds(ids []string) ([]paymentEntity.Payment, error) {
	var payments []paymentEntity.Payment

	for _, id := range ids {
		p, ok := r.records[id]
		if !ok {
			return nil, PaymentDoesntExistsError{PaymentId: id}
		}
		payments = append(payments, NewPaymentProxy(p, r.assetFinder))
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

func (r *PaymentRepository) ConvertRecordToEntity(record PaymentRecord) (p paymentEntity.Payment) {
	return r.createPaymentProxy(record)
}

func NewPaymentRepository(assetFinder AssetFinderById) *PaymentRepository {
	return &PaymentRepository{
		assetFinder,
		make(map[string]PaymentRecord),
	}
}
