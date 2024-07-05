package usecase

import (
	"github.com/muizu555/investment/src/domain"
	"github.com/muizu555/investment/src/repository"
)

func GetUserAssets(userID, date string) (*domain.Asset, error) {
	// あるuserIDのユーザーが持っている現在の日付までの取引を取得
	assetSettings, err := repository.GetAssetSettingsByUserIDANDDate(userID, date)
	if err != nil {
		return nil, err
	}

	// TODO: 後で返すデータの型をつくる ポインタ型にするかどうか
	return &domain.Asset{
		Data:         date,
		CurrentValue: assetSettings[0].AppraisedAsset,
		CurrentPL:    assetSettings[0].ProfitLoss,
	}, nil
}
