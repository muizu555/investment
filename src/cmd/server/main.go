package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/muizu555/investment/src/domain"
)

func main() {
	err := godotenv.Load()
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

	getReferencePrices := func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM ReferencePrices")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var referencePrices []domain.ReferencePrices
		for rows.Next() {
			var referencePrice domain.ReferencePrices
			err := rows.Scan(&referencePrice.FundID, &referencePrice.ReferencePrice, &referencePrice.ReferencePriceDate)
			if err != nil {
				log.Fatal(err)
			}
			referencePrices = append(referencePrices, referencePrice)
		}

		for i, referencePrice := range referencePrices {
			if i == 0 {
				fmt.Println(referencePrice.ReferencePriceDate)
			}
			fmt.Fprintf(w, "%s %d %s\n", referencePrice.FundID, referencePrice.ReferencePrice, referencePrice.ReferencePriceDate)
		}
	}

	http.HandleFunc("/reference-prices", getReferencePrices)

	fmt.Println("Server is running at :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
