package usecase

import (
	"time"

	"github.com/muizu555/investment/src/domain"
	"github.com/muizu555/investment/src/repository"
)

func GetUserAssets(userID string) (*domain.Asset, error) {
	// あるuserIDのユーザーが持っている現在の日付までの取引を取得
	trades, err := repository.GetTradesByUserID(userID)
	if err != nil {
		return nil, err
	}

	// ここまでは一旦ok
	// ここでfundIDsだけじゃなくて各々の所持口数も取得したい
	fundIDSums := trades.GetFundIDAndSums() // この関数あってる？
	fundIDs := fundIDSums.GetFundIDs()

	referencePrices, err := repository.GetReferencePricesByFundIDs(fundIDs)
	if err != nil {
		return nil, err
	}

	// 現在の日付(2024-06-01)の取引の参照価格を取得
	todaysReferencePrices, err := repository.GetReferencePricesByDate("2024-06-01")
	if err != nil {
		return nil, err
	}

	//domainのエラーは...
	currentValue, currentPL := domain.CalculateAssets(trades, todaysReferencePrices)

	// TODO: 後で返すデータの型をつくる ポインタ型にするかどうか
	return &domain.Asset{
		Data:         time.Now().Format("2006-01-02"),
		CurrentValue: currentValue,
		CurrentPL:    currentPL,
	}, nil
}
