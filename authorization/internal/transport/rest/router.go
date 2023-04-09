package rest

import (
	"authorization/internal/transport/rest/controllers"
	"authorization/internal/transport/rest/middlewares"
	"github.com/labstack/echo/v4"
)

func InitRoutes(app *echo.Group) {
	controllers.InitAuthRoutes(app.Group("/auth"))
	controllers.InitUserRoutes(app.Group("/users"))
	controllers.InitOrganizationRoutes(app.Group("/organization"))
	tenant := app.Group("/")
	tenant.Use(middlewares.JwtProtect())
	tenant.Use(middlewares.TenantGuard)
	tenant.Any("courses/*", func(c echo.Context) error {
		return c.String(200, "redirect to course service")
	})
}
