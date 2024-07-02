package domain

// TODO: IDいるかどうか問題(一旦いらない)
type TradeHistory struct {
	UserID    string
	FundID    string
	Quantity  int
	TradeDate string
}

func NewTradeHistory(userID, fundID string, quantity int, tradeDate string) *TradeHistory {
	return &TradeHistory{
		UserID:    userID,
		FundID:    fundID,
		Quantity:  quantity,
		TradeDate: tradeDate,
	}
}
