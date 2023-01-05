package controllers

import (
	"edugree_auth/internal/controllers/auth"
	"edugree_auth/internal/controllers/permissions"
	"edugree_auth/internal/controllers/users"
	"github.com/labstack/echo/v4"
)

func InitRoutes(app *echo.Group) {
	auth.InitRoutes(app.Group("/auth"))
	permissions.InitRoutes(app.Group("/permission"))
	users.InitRoutes(app.Group("/users"))
}
