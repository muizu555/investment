package router

import (
	"github.com/labstack/echo/v4"
	"github.com/muizu555/investment/src/handler"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/:user_id/trades", handler.GetTradeCount)
}
