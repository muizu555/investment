package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/muizu555/investment/src/repository"
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

	err = repository.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	db := repository.GetDB()
	defer db.Close()

	// TradeHistoryテーブルにデータを一度に挿入
	var valueStrings1 []string
	var valueArgs1 []interface{}
	for i, row := range rows1 {
		if i == 0 {
			continue
		}
		valueStrings1 = append(valueStrings1, "(?, ?, ?, ?)")
		valueArgs1 = append(valueArgs1, row[0], row[1], row[2], row[3])
	}
	insertStmt1 := fmt.Sprintf("INSERT INTO TradeHistory (UserID, FundID, Quantity, TradeDate) VALUES %s",
		strings.Join(valueStrings1, ","))
	_, err = db.Exec(insertStmt1, valueArgs1...)
	if err != nil {
		log.Fatal(err)
	}

	// ReferencePricesテーブルにデータを一度に挿入
	var valueStrings2 []string
	var valueArgs2 []interface{}
	for i, row := range rows2 {
		if i == 0 {
			continue
		}
		valueStrings2 = append(valueStrings2, "(?, ?, ?)")
		valueArgs2 = append(valueArgs2, row[0], row[1], row[2])
	}
	insertStmt2 := fmt.Sprintf("INSERT INTO ReferencePrices (FundID, ReferencePrice, ReferencePriceDate) VALUES %s",
		strings.Join(valueStrings2, ","))
	_, err = db.Exec(insertStmt2, valueArgs2...)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("データの挿入が完了しました。")
}
