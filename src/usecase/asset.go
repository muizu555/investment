package usecase

import (
	"strconv"

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

func GetUserAssetYears(userID, date string) (*domain.AssetResponse, error) {
	// あるuserIDのユーザーが持っている取引の年を取得
	assetYearSettings, err := repository.GetAssetYearsByUserID(userID, date)
	if err != nil {
		return nil, err
	}
	assets := make(domain.AssetYears, len(assetYearSettings))
	for i, assetYearSetting := range assetYearSettings {
		year, _ := strconv.Atoi(assetYearSetting.TradeYear)
		assets[i] = domain.AssetYear{
			Year:         year,
			CurrentValue: assetYearSetting.AppraisedAsset,
			CurrentPL:    assetYearSetting.ProfitLoss,
		}
	}

	return &domain.AssetResponse{
		Date:   date,
		Assets: assets,
	}, nil
}
