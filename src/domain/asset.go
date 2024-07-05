package domain

type Asset struct {
	Data         string
	CurrentValue int
	CurrentPL    int
}

type Assets []Asset

// CalculateAssets は取引履歴と基準価格をもとに、現在の資産価値と損益を計算する
func CalculateAssets(trades TradeHistories, referencePrices ReferencePrices) (int, int) {
	// 現在の資産価値
	currentValue := 0
	// 現在の損益
	currentPL := 0
	// ここで全ての銘柄のfundIDsとそれぞれ何口かを取得
	fundIDAndSums := trades.GetFundIDAndSums()

	for _, fundIDAndSum := range fundIDAndSums {
		currentValue += fundIDAndSum.Sum * findReferencePrice(referencePrices, fundIDAndSum.FundID) / 1000
	}

	currentPL = currentValue - trades.GetTotalAmount()

	return currentValue, currentPL

}

func findReferencePrice(referencePrices ReferencePrices, fundID string) int {
	for _, referencePrice := range referencePrices {
		if referencePrice.FundID == fundID {
			return referencePrice.ReferencePrice
		}
	}
	return 0
}
