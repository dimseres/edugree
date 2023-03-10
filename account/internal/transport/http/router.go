package rest

import (
	"account/internal/transport/http/account"
	"github.com/labstack/echo/v4"
)

func InitRoutes(app *echo.Group) {
	account.InitRoutes(app.Group("/account"))
}
