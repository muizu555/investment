package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/muizu555/investment/src/usecase"
)

// TODO: handlerで実際にclientに対してデータを返しているのはいいのか...
func GetUserAssets(c echo.Context) error {
	userID := c.Param("user_id")
	assets, err := usecase.GetUserAssets(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, assets)
}
