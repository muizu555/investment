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

	file, err := os.Open("trade_history.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	rows, err := r.ReadAll() // csvを一度に全て読み込む
	if err != nil {
		log.Fatal(err)
	}

	database := os.Getenv("DATABASE")
	userName := os.Getenv("USERNAME")
	userPass := os.Getenv("USERPASS")

	// DSNを構築
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", userName, userPass, database)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// テーブルにデータを挿入
	insertStmt := "INSERT INTO TradeHistory (UserID, FundID, Quantity, TradeDate) VALUES (?, ?, ?, ?)"
	for i, row := range rows {
		if i == 0 {
			continue
		}
		_, err := db.Exec(insertStmt, row[0], row[1], row[2], row[3])
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("データの挿入が完了しました。")
}
