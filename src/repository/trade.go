package repository

import (
	_ "github.com/go-sql-driver/mysql"
)

// TODO: repositoryの名前が少し変な気がする...(ex.GetTradeHistoryCounts?)
// 全体でDIをしてわざわざdbを作らないようにする
func GetTradeCountByUserID(userID string) (int, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM TradeHistory WHERE UserID = ?", userID).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
