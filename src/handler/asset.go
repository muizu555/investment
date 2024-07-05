package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/muizu555/investment/src/usecase"
)

// TODO: handlerで実際にclientに対してデータを返しているのはいいのか...
func GetUserAssets(c echo.Context) error {
	userID := c.Param("user_id")
	// 日本標準時 (JST) のロケーションを取得
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return nil
	}

	currentTime := time.Now().In(jst)
	date := currentTime.Format("2006-01-02")

	assets, err := usecase.GetUserAssets(userID, date)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, assets)
}
