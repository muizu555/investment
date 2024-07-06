package domain

import (
	"fmt"
	"time"
)

func GetNowDate() (string, error) {
	// 日本標準時 (JST) のロケーションを取得
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return "", err
	}

	currentTime := time.Now().In(jst)
	date := currentTime.Format("2006-01-02")

	return date, nil
}
