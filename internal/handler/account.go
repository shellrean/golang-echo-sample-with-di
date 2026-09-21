package handler

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/samber/do/v2"
	"shellrean.id/account/internal/service"
)

type Account struct {
	accountService service.Account
}

func NewAccount(i do.Injector) (Account, error) {
	return Account{
		accountService: do.MustInvoke[service.Account](i),
	}, nil
}

func (a Account) Show(ctx *echo.Context) error {
	acc, err := a.accountService.Show(ctx.Request().Context(), ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, acc)
}
