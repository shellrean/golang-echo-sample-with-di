package routes

import (
	"github.com/labstack/echo/v5"
	"github.com/samber/do/v2"
	"shellrean.id/account/internal/handler"
)

func Register(i do.Injector, e *echo.Echo) {
	uHandler := do.MustInvoke[handler.Account](i)

	e.GET("/users/:id", uHandler.Show)
}
