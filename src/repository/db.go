package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var db *sql.DB

func InitDB() error {
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
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	return nil
}

func GetDB() *sql.DB {
	return db
}
