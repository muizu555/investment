package usecase

import (
	"fmt"
	"strconv"

	"github.com/muizu555/investment/src/domain"
	"github.com/muizu555/investment/src/repository"
)

func GetUserAssets(userID, date string) (*domain.Asset, error) {
	count, _ := repository.GetTradeCountByUserID(userID)
	if count == 0 {
		// 特定のUserIDの取引データがない場合
		return nil, fmt.Errorf("userID %s: %w", userID, domain.ErrNotFound)
	}

	// 現在の日時or指定された日時のReferencePriceが存在するか確認
	exist, err := repository.ExistReferencePriceByDate(date)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, fmt.Errorf("date %s has no ReferencePrices: %w", date, domain.ErrNotFound)
	}

	assetSettings, err := repository.GetAssetSettingsByUserIDAndDate(userID, date)
	if err != nil {
		return nil, err
	}

	return &domain.Asset{
		Data:         date,
		CurrentValue: assetSettings[0].AppraisedAsset,
		CurrentPL:    assetSettings[0].ProfitLoss,
	}, nil
}

func GetUserAssetYears(userID, date string) (*domain.AssetResponse, error) {
	count, _ := repository.GetTradeCountByUserID(userID)
	if count == 0 {
		// 特定のUserIDの取引データがない場合
		return nil, fmt.Errorf("userID %s: %w", userID, domain.ErrNotFound)
	}

	exist, err := repository.ExistReferencePriceByDate(date)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, fmt.Errorf("date %s has no ReferencePrices: %w", date, domain.ErrNotFound)
	}

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
