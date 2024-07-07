package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/muizu555/investment/src/usecase"
)

func GetTradeCount(c echo.Context) error {
	userID := c.Param("user_id")
	count, err := usecase.GetTradeCount(userID)
	// データベースの内部のエラー
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if count == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "no trade data found for the specified user_id"})
	}

	return c.JSON(http.StatusOK, map[string]int{"count": count})
}
