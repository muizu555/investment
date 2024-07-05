package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/muizu555/investment/src/usecase"
)

func GetTradeCount(c echo.Context) error {
	userID := c.Param("user_id")
	count, err := usecase.GetTradeCount(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	// TODO: countの型を作るかどうか考える
	return c.JSON(http.StatusOK, map[string]int{"count": count})
}
