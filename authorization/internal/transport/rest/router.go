package rest

import (
	"authorization/internal/transport/rest/controllers"
	"github.com/labstack/echo/v4"
)

func InitRoutes(app *echo.Group) {
	controllers.InitAuthRoutes(app.Group("/auth"))
	controllers.InitUserRoutes(app.Group("/users"))
	controllers.InitOrganizationRoutes(app.Group("/organization"))
	//permissions.InitRoutes(app.Group("/permissions"))
}
