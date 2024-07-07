package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/muizu555/investment/src/domain"
	"github.com/muizu555/investment/src/usecase"
)

func GetUserAssets(c echo.Context) error {
	userID := c.Param("user_id")
	date := c.QueryParam("date")
	if date == "" {
		date, _ = domain.GetNowDate()
	}

	assets, err := usecase.GetUserAssets(userID, date)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, assets)
}

func GetUserAssetsByYear(c echo.Context) error {
	userID := c.Param("user_id")
	// 日本標準時 (JST) のロケーションを取得
	date, _ := domain.GetNowDate()

	assets, err := usecase.GetUserAssetYears(userID, date)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, assets)
}
