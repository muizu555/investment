package repository

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/muizu555/investment/src/domain"
)

func GetAssetSettingsByUserID(userID string) (domain.AssetSettings, error) {
	database := os.Getenv("DATABASE")
	userName := os.Getenv("USERNAME")
	userPass := os.Getenv("USERPASS")

	dsn := fmt.Sprintf("%s:%s@tcp(mysql:3306)/%s", userName, userPass, database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// 現在の日付までの取引を取得
	rows, err := db.Query("SELECT TradeHistory.FundID, TradeHistory.Quantity, TradeHistory.TradeDate, ReferencePrices.ReferencePrice, ReferencePrices.ReferencePriceDate FROM TradeHistory INNER JOIN ReferencePrices ON TradeHistory.FundID = ReferencePrices.FundID AND TradeHistory.TradeDate = ReferencePrices.ReferencePriceDate WHERE TradeHistory.UserID = ? AND TradeHistory.TradeDate <= '2024-06-01'", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var assetSettings domain.AssetSettings
	for rows.Next() {
		var assetSetting domain.AssetSetting
		err := rows.Scan(&assetSetting.FundID, &assetSetting.Quantity, &assetSetting.TradeDate, &assetSetting.ReferencePrice, &assetSetting.ReferencePriceDate)
		if err != nil {
			return nil, err
		}
		assetSettings = append(assetSettings, assetSetting)
	}
	return assetSettings, nil
}

// /ここから関数の戻り値の型を変えることから
// // todo
// // ここで複数形のReferencePricesを返せばよさそう　mapじゃなくて
func GetReferencePricesByFundIDs(fundIDs []string) (domain.ReferencePrices, error) {
	database := os.Getenv("DATABASE")
	userName := os.Getenv("USERNAME")
	userPass := os.Getenv("USERPASS")

	dsn := fmt.Sprintf("%s:%s@tcp(mysql:3306)/%s", userName, userPass, database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// ここでfundIDに含まれているものだけ取得する
	rows, err := db.Query("SELECT FundID, Price, ReferencePriceDate FROM ReferencePrice WHERE FundID IN (?)", fundIDs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	referencePrices := make(domain.ReferencePrices, 0)
	for rows.Next() {
		var referencePrice domain.ReferencePrice
		err := rows.Scan(&referencePrice.FundID, &referencePrice.ReferencePrice)
		if err != nil {
			return nil, err
		}
		referencePrices = append(referencePrices, referencePrice)
	}
	return referencePrices, nil
}

// 　特定の日時のReferencePricesを取得
func GetReferencePricesByDate(date string) (domain.ReferencePrices, error) {
	database := os.Getenv("DATABASE")
	userName := os.Getenv("USERNAME")
	userPass := os.Getenv("USERPASS")

	dsn := fmt.Sprintf("%s:%s@tcp(mysql:3306)/%s", userName, userPass, database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// ここでfundIDに含まれているものだけ取得する
	rows, err := db.Query("SELECT FundID, Price FROM ReferencePrices WHERE Date = ?", date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	referencePrices := make(domain.ReferencePrices, 0)
	for rows.Next() {
		var referencePrice domain.ReferencePrice
		err := rows.Scan(&referencePrice.FundID, &referencePrice.ReferencePrice)
		if err != nil {
			return nil, err
		}
		referencePrices = append(referencePrices, referencePrice)
	}
	return referencePrices, nil
}
