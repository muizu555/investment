package repository

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/muizu555/investment/src/domain"
)

func GetTradesByUserID(userID string) (domain.TradeHistories, error) {
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
	rows, err := db.Query("SELECT UserID, FundID, Quantity, TradeDate FROM TradeHistory WHERE UserID = ? AND TradeDate <= CURDATE()", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var trades domain.TradeHistories
	for rows.Next() {
		var trade domain.TradeHistory
		err := rows.Scan(&trade.UserID, &trade.FundID, &trade.Quantity, &trade.TradeDate)
		if err != nil {
			return nil, err
		}
		trades = append(trades, trade)
	}

	return trades, nil
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
