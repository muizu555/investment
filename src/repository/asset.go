package repository

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/muizu555/investment/src/domain"
)

func GetAssetSettingsByUserIDANDDate(userID, date string) (domain.AssetSettings, error) {
	database := os.Getenv("DATABASE")
	userName := os.Getenv("USERNAME")
	userPass := os.Getenv("USERPASS")

	dsn := fmt.Sprintf("%s:%s@tcp(mysql:3306)/%s", userName, userPass, database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(`
	SELECT
		SUM(FLOOR(PerFund.QuantitySum * RP.ReferencePrice / 10000)) AS AppraisedAsset,
		SUM(PerFund.PurchasePriceSum) AS PurchasePriceSum,
		SUM(FLOOR(PerFund.QuantitySum * RP.ReferencePrice / 10000) - PerFund.PurchasePriceSum) AS ProfitLoss
	FROM (
		SELECT
			TH.FundID,
			SUM(TH.Quantity) AS QuantitySum,
			SUM(FLOOR(TH.Quantity * RP.ReferencePrice / 10000)) AS PurchasePriceSum
		FROM TradeHistory AS TH
		JOIN ReferencePrices AS RP
		ON
			TH.FundID = RP.FundID AND
			TH.TradeDate = RP.ReferencePriceDate
		WHERE
			TH.UserID = ? AND
			RP.ReferencePriceDate <= ?
		GROUP BY TH.FundID
	) AS PerFund
	JOIN ReferencePrices AS RP
	ON
		PerFund.FundID = RP.FundID
	WHERE
		RP.ReferencePriceDate = ?
`, userID, date, date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var assetSettings domain.AssetSettings
	for rows.Next() {
		var assetSetting domain.AssetSetting
		err := rows.Scan(&assetSetting.AppraisedAsset, &assetSetting.PurchasePriceSum, &assetSetting.ProfitLoss)
		if err != nil {
			return nil, err
		}
		assetSettings = append(assetSettings, assetSetting)
	}
	return assetSettings, nil
}

func GetAssetYearsByUserID(userID, date string) (domain.AssetYearSettings, error) {
	database := os.Getenv("DATABASE")
	userName := os.Getenv("USERNAME")
	userPass := os.Getenv("USERPASS")

	dsn := fmt.Sprintf("%s:%s@tcp(mysql:3306)/%s", userName, userPass, database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(`
	SELECT
		PerYearFund.TradeYear,
		SUM(FLOOR(PerYearFund.QuantitySum * RP.ReferencePrice / 10000)) AS AppraisedAsset,
		SUM(PerYearFund.PurchasePriceSum) AS PurchasePriceSum,
		SUM(FLOOR(PerYearFund.QuantitySum * RP.ReferencePrice / 10000) - PerYearFund.PurchasePriceSum) AS ProfitLoss
	FROM (
		SELECT
			TH.FundID,
			DATE_FORMAT(TH.TradeDate, '%Y') AS TradeYear,
			SUM(TH.Quantity) AS QuantitySum,
			SUM(FLOOR(TH.Quantity * RP.ReferencePrice / 10000)) AS PurchasePriceSum
		FROM TradeHistory AS TH
		JOIN ReferencePrices AS RP
		ON
			TH.FundID = RP.FundID AND
			TH.TradeDate = RP.ReferencePriceDate
		WHERE
			TH.UserID = ? AND
			TH.TradeDate <= ?
		GROUP BY TradeYear, TH.FundID
	) AS PerYearFund
	JOIN ReferencePrices AS RP
	ON
		PerYearFund.FundID = RP.FundID
	WHERE
		RP.ReferencePriceDate = ?
	GROUP BY PerYearFund.TradeYear
	ORDER BY PerYearFund.TradeYear DESC
`, userID, date, date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var assetYearSettings domain.AssetYearSettings
	for rows.Next() {
		var assetYearSetting domain.AssetYearSetting
		err := rows.Scan(&assetYearSetting.TradeYear, &assetYearSetting.AppraisedAsset, &assetYearSetting.PurchasePriceSum, &assetYearSetting.ProfitLoss)
		if err != nil {
			return nil, err
		}
		assetYearSettings = append(assetYearSettings, assetYearSetting)
	}
	return assetYearSettings, nil
}
