package main

import (
	"time"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/samber/do/v2"
	"shellrean.id/account/internal/config"
	"shellrean.id/account/internal/connection"
	"shellrean.id/account/internal/handler"
	"shellrean.id/account/internal/routes"
	"shellrean.id/account/internal/service"
)

func main() {
	in := do.New()
	do.Provide(in, config.Load)
	do.Provide(in, handler.NewAccount)
	do.Provide(in, connection.NewDatabase)
	do.Provide(in, service.NewAccount)

	e := echo.New()
	e.Use(middleware.ContextTimeout(15 * time.Second))

	routes.Register(in, e)

	if err := e.Start(":8899"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
