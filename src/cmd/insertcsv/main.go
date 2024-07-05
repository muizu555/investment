package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// trade_history.csvファイルを開く
	file1, err := os.Open("trade_history.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file1.Close()

	// CSVリーダーを作成
	r1 := csv.NewReader(file1)
	rows1, err := r1.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// reference_prices.csvファイルを開く
	file2, err := os.Open("reference_prices.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file2.Close()

	// CSVリーダーを作成
	r2 := csv.NewReader(file2)
	rows2, err := r2.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// 環境変数から接続情報を取得
	database := os.Getenv("DATABASE")
	userName := os.Getenv("USERNAME")
	userPass := os.Getenv("USERPASS")

	// DSNを構築
	dsn := fmt.Sprintf("%s:%s@tcp(mysql:3306)/%s", userName, userPass, database)

	// MySQLデータベースに接続
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// TradeHistoryテーブルにデータを挿入
	insertStmt1 := "INSERT INTO TradeHistory (UserID, FundID, Quantity, TradeDate) VALUES (?, ?, ?, ?)"
	for i, row := range rows1 {
		if i == 0 {
			continue
		}
		_, err := db.Exec(insertStmt1, row[0], row[1], row[2], row[3])
		if err != nil {
			log.Fatal(err)
		}
	}

	// ReferencePricesテーブルにデータを挿入
	insertStmt2 := "INSERT INTO ReferencePrices (FundID, ReferencePrice, ReferencePriceDate) VALUES (?, ?, ?)"
	for i, row := range rows2 {
		if i == 0 {
			continue
		}
		_, err := db.Exec(insertStmt2, row[0], row[1], row[2])
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("データの挿入が完了しました。")
}
