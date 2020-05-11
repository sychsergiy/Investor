package payment

import "investor/entities/asset"

type Filter struct {
	payments []Payment
}

func NewFilter(payments []Payment) *Filter {
	return &Filter{payments: payments}
}

func (f *Filter) ByAssetNames(assetNames []string) (*Filter, error) {
	filtered, err := FilterByAssetNames(f.payments, assetNames)
	if err != nil {
		return nil, err
	}
	f.payments = filtered
	return f, nil
}

func (f *Filter) ByPeriods(periods []Period) *Filter {
	f.payments = FilterByPeriods(f.payments, periods)
	return f
}

func (f *Filter) ByTypes(types []Type) *Filter {
	f.payments = FilterByTypes(f.payments, types)
	return f
}

func (f *Filter) ByAssetCategories(categories []asset.Category) (*Filter, error) {
	payments, err := FilterByAssetCategories(f.payments, categories)
	if err != nil {
		return nil, err
	}
	f.payments = payments
	return f, nil
}
func (f Filter) Payments() []Payment {
	return f.payments
}
