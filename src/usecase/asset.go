package usecase

import (
	"time"

	"github.com/muizu555/investment/src/domain"
	"github.com/muizu555/investment/src/repository"
)

func GetUserAssets(userID string) (*domain.Asset, error) {
	// あるuserIDのユーザーが持っている現在の日付までの取引を取得
	asset, err := repository.GetAssetSettingsByUserIDANDDate(userID, "2024-06-01")
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

	// TODO: 後で返すデータの型をつくる ポインタ型にするかどうか
	return &domain.Asset{
		Data:         time.Now().Format("2006-01-02"),
		CurrentValue: currentValue,
		CurrentPL:    currentPL,
	}, nil
}
