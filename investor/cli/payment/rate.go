package payment

import (
	"fmt"
	"investor/entities/payment"
	"investor/entities/profit"
	"investor/interactors"
	"log"
	"strconv"
)

type CalcRateFromProfitCommand struct {
	rateCalculator interactors.CalcRateFromProfit
	assetsLister   interactors.ListAssets
}

func (l CalcRateFromProfitCommand) Execute() {
	l.CalculateRate()
}

func (l CalcRateFromProfitCommand) CalculateRate() {
	assets, err := l.assetsLister.ListAll()
	if err != nil {
		log.Fatal(err)
	}
	assetName := selectAsset(assets).Name()

	desirableProfit := readDesirableProfit()

	req := interactors.CalcRateFromProfitRequest{
		AssetName:       assetName,
		Periods:         []payment.Period{},
		DesirableProfit: desirableProfit,
	}
	resp, err := l.rateCalculator.Calc(req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(
		"Asset rate to get %.2f%% gain From %s: %.3f",
		req.DesirableProfit.Percentage(), resp.AssetName, resp.AssetRate,
	)
}

func readDesirableProfit() profit.Profit {
	fmt.Printf("Enter desirable profit(1.5 means 150 USD returned from 100 USD invested): ")
	profitStr := readFromConsole()

	profitCoef64, err := strconv.ParseFloat(profitStr, 32)
	profitCoef := float32(profitCoef64)

	if err != nil {
		log.Printf("Failed to parse float number from input, error message: %s\b", err)
		log.Println("Retrying ...")
		return readDesirableProfit()
	}
	return profit.NewFromCoefficient(profitCoef)

}

func NewCalcRateFromProfitCommand(
	rateCalculator interactors.CalcRateFromProfit, assetsLister interactors.ListAssets,
) CalcRateFromProfitCommand {
	return CalcRateFromProfitCommand{rateCalculator, assetsLister}
}
