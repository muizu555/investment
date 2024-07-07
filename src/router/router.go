package router

import (
	"github.com/labstack/echo/v4"
	"github.com/muizu555/investment/src/handler"
)

func SetupRoutes(e *echo.Echo) {
	// step 3
	e.GET("/:user_id/trades", handler.GetTradeCount)
	// step 4, 5
	e.GET("/:user_id/assets", handler.GetUserAssets)
	// step 6
	e.GET("/:user_id/assets/byYear", handler.GetUserAssetsByYear)
}
