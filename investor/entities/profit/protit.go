package profit

import (
	"fmt"
	"investor/entities/payment"
	"math"
)

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

func NewFromCoefficient(coefficient float32) FloatProfit {
	return FloatProfit{coefficient}
}

func NewFromPercentage(percentage float32) FloatProfit {
	return FloatProfit{percentage/100 + 1}
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
	return "unable to calculate profit due to zero returned sum"
}

type ReturnedAssetSumMoreThanInvested struct{}

func (e ReturnedAssetSumMoreThanInvested) Error() string {
	return "unable to calculate profit due returned Asset sum more than invested"
}

func calcSumsForPayments(payments []payment.Payment) Sums {
	s := Sums{}
	for _, item := range payments {
		switch item.Type() {
		case payment.Return:
			s.Returned += item.AbsoluteAmount()
			s.ReturnedAsset += item.AssetAmount()
		case payment.Invest:
			s.Invested += item.AbsoluteAmount()
			s.InvestedAsset += item.AssetAmount()
		default:
			panic("unexpected payment type")
		}
	}
	return s
}

type LessThanZeroAssetRestError struct {
	value float32
}

func (e LessThanZeroAssetRestError) Error() string {
	return fmt.Sprintf(
		"not able to calc profit due to less than zero asset rest, value: %f", e.value,
	)
}

// only for payments from one asset
func CalcRateFromDesirableProfit(
	profit Profit, payments []payment.Payment,
) (float32, error) {
	sums := calcSumsForPayments(payments)
	desiredSum := profit.Coefficient() * sums.Invested

	//
	sumShouldBeReturned := desiredSum - sums.Returned
	assetRest := sums.InvestedAsset - sums.ReturnedAsset
	if assetRest <= 0 {
		return 0, LessThanZeroAssetRestError{}
	}
	rate := sumShouldBeReturned / assetRest

	rate = float32(math.Round(float64(rate)*100) / 100)
	return rate, nil
}

type AssetProfit struct {
	AssetName     string
	Profit        Profit
	PaymentsCount int
}

func CalcForAsset(payments []payment.Payment, name string) (AssetProfit, error) {
	// todo: maybe filter payments by asset name
	//  or return error if payment with another asset exists

	// all payments should have the same asset
	sums := calcSumsForPayments(payments)
	profit, err := calcProfitForAsset(sums)
	if err != nil {
		return AssetProfit{}, err
	}
	return AssetProfit{name, profit, len(payments)}, nil
}

func calcProfitForAsset(sums Sums) (Profit, error) {
	// calculate asset profit coefficient
	// find invested capital in absolute amount (USD) and in asset
	// a = find percentage of asset rest after all payments
	// b = find percentage of capital was returned
	// profit coefficient: (1 - b / a)
	// profit coefficient 1 means no profit no benefit, 0.5 means lost 50%, 1.5 earned 50%, 5 means earned 400%

	if sums.Invested == 0 {
		return nil, ZeroInvestedSumError{}
	}

	if sums.ReturnedAsset > sums.InvestedAsset {
		return nil, ReturnedAssetSumMoreThanInvested{}
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

	return NewFromCoefficient(returnPart / assetSpendPart), nil
}
