package payment

type Profit interface {
	Coefficient() float32
	Percentage() float32
}

type FloatProfit struct {
	value float32
}

func (p FloatProfit) Coefficient() float32 {
	return p.value
}
func (p FloatProfit) Percentage() float32 {
	return (p.value - 1) * 100
}

func NewProfitFromCoefficient(coefficient float32) FloatProfit {
	return FloatProfit{coefficient}
}

func NewProfitFromPercentage(percentage float32) FloatProfit {
	return FloatProfit{percentage / 100 + 1}
}

type Sums struct {
	Invested      float32
	Returned      float32
	InvestedAsset float32
	ReturnedAsset float32
}

type ZeroInvestedSumError struct{}

func (e ZeroInvestedSumError) Error() string {
	return "unable to calculate profit due to zero invested sum"
}

type ZeroAssetReturnedError struct{}

func (e ZeroAssetReturnedError) Error() string {
	return "unable to calculate profit de to zero returned sum"
}

func CalcProfit(sums Sums) (Profit, error) {
	// calculate asset profit coefficient
	// find invested capital in absolute amount (USD) and in asset
	// a = find percentage of asset rest after all payments
	// b = find percentage of capital was returned
	// profit coefficient: (1 - b / a)
	// profit coefficient 1 means no profit no benefit, 0.5 means lost 50%, 1.5 earned 50%, 5 means earned 400%

	if sums.Invested == 0 {
		return nil, ZeroInvestedSumError{}
	}

	assetRestPart := (sums.InvestedAsset - sums.ReturnedAsset) / sums.InvestedAsset
	assetSpendPart := 1 - assetRestPart

	// todo: maybe not return error, but calculate profit
	//  with assumption that it is possible to return USD without Asset lost
	//  calc profit for this specific case
	if assetSpendPart == 0 {
		return nil, ZeroAssetReturnedError{}
	}

	returnPart := sums.Returned / sums.Invested

	return NewProfitFromCoefficient(returnPart / assetSpendPart), nil
}
