package domain

type FundIDAndSum struct {
	FundID string
	// 銘柄の所持口数
	Sum int
}

type FundIDAndSums []FundIDAndSum

// domain.TradeHistoriesをレシーバとしてそこからFundIDSumを取り出すメソッドを実装
// すでに出ているFundIDは重複せずにすでにあるFundIDAndSum構造体のSUMに加算する
func (ths TradeHistories) GetFundIDAndSums() FundIDAndSums {
	fundIDAndSums := make([]FundIDAndSum, 0)
	for _, th := range ths {
		if fundIDAndSum, ok := findFundIDAndSum(fundIDAndSums, th.FundID); ok {
			fundIDAndSum.Sum += th.Quantity
		} else {
			fundIDAndSums = append(fundIDAndSums, FundIDAndSum{
				FundID: th.FundID,
				Sum:    th.Quantity,
			})
		}
	}
	return fundIDAndSums
}

func findFundIDAndSum(fundIDAndSums []FundIDAndSum, fundID string) (*FundIDAndSum, bool) {
	for i, fundIDAndSum := range fundIDAndSums {
		if fundIDAndSum.FundID == fundID {
			return &fundIDAndSums[i], true
		}
	}
	return nil, false
}

// fundIDAndSumsをレシーバとしてFundIDAndSumのFundIDを取り出すメソッドを実装
func (fids FundIDAndSums) GetFundIDs() []string {
	fundIDs := make([]string, 0)
	for _, fid := range fids {
		fundIDs = append(fundIDs, fid.FundID)
	}
	return fundIDs
}
