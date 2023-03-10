package rest

import (
	"authorization/internal/transport/rest/auth"
	"authorization/internal/transport/rest/permissions"
	"authorization/internal/transport/rest/users"
	"github.com/labstack/echo/v4"
)

func InitRoutes(app *echo.Group) {
	auth.InitRoutes(app.Group("/auth"))
	permissions.InitRoutes(app.Group("/permissions"))
	users.InitRoutes(app.Group("/users"))
}
