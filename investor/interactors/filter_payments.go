package interactors

import (
	"investor/entities/payment"
	"investor/interactors/ports"
)

type FilterPayments struct {
	filterFactory PaymentIdsFilterFactory
}

func NewFilterPayments(factory PaymentIdsFilterFactory) FilterPayments {
	return FilterPayments{factory}
}

type FilterPaymentsRequest struct {
	PaymentIds []string
}

type FilterPaymentsResponse struct {
	Payments []payment.Payment
}

type PaymentIdsFilterFactory struct {
	paymentFinder ports.PaymentFinderByIds
}

func (f PaymentIdsFilterFactory) Create() *PaymentIdsFilter {
	return NewPaymentIdsFilter(f.paymentFinder)
}

func NewPaymentIdsFilterFactory(finder ports.PaymentFinderByIds) PaymentIdsFilterFactory {
	return PaymentIdsFilterFactory{finder}
}

func (f FilterPayments) Filter(model FilterPaymentsRequest) (FilterPaymentsResponse, error) {
	if len(model.PaymentIds) != 0 {
		payments, err := f.filterFactory.Create().Payments(model.PaymentIds)
		if err != nil {
			return FilterPaymentsResponse{}, err
		} else {
			return FilterPaymentsResponse{payments}, nil
		}
	} else {
		panic("not implemented")
	}
}

type PaymentIdsFilter struct {
	paymentFinder ports.PaymentFinderByIds
}

func (f PaymentIdsFilter) Payments(ids []string) ([]payment.Payment, error) {
	return f.paymentFinder.FindByIds(ids)
}
func NewPaymentIdsFilter(paymentFinder ports.PaymentFinderByIds) *PaymentIdsFilter {
	return &PaymentIdsFilter{paymentFinder: paymentFinder}
}

//type FilterPaymentsRequest struct {
//	Period          payment.Period
//	AssetCategories []string
//	AssetNames      []string
//
//	paymentsFinder // new interface to filter payments
//}

//func (f FilterPayments) FilterPayments(model FilterPaymentsRequest) ([]payment.Payment, error) {
//
//}
