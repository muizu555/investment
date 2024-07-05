package repository

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// TODO: repositoryの名前が少し変な気がする...(ex.GetTradeHistoryCounts?)
// 全体でDIをしてわざわざdbを作らないようにする
func GetTradeCount(userID string) (int, error) {
	database := os.Getenv("DATABASE")
	userName := os.Getenv("USERNAME")
	userPass := os.Getenv("USERPASS")

	dsn := fmt.Sprintf("%s:%s@tcp(mysql:3306)/%s", userName, userPass, database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return 0, err
	}
	defer db.Close()

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM TradeHistory WHERE UserID = ?", userID).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
