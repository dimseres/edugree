package rest

import (
	"edugree_auth/internal/transport/rest/auth"
	"edugree_auth/internal/transport/rest/permissions"
	"edugree_auth/internal/transport/rest/users"
	"github.com/labstack/echo/v4"
)

func InitRoutes(app *echo.Group) {
	auth.InitRoutes(app.Group("/auth"))
	permissions.InitRoutes(app.Group("/permissions"))
	users.InitRoutes(app.Group("/users"))
}
