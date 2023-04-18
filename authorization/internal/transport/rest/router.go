package rest

import (
	"authorization/internal/transport/rest/controllers"
	"authorization/internal/transport/rest/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/url"
)

func InitRoutes(app *echo.Group) {
	controllers.InitAuthRoutes(app.Group("/auth"))
	controllers.InitUserRoutes(app.Group("/users"))
	controllers.InitOrganizationRoutes(app.Group("/organization"))

	tenant := app.Group("")
	tenant.Use(middlewares.JwtProtect())
	tenant.Use(middlewares.TenantGuard)

	courseUrl, err := url.Parse("http://localhost:8000")
	if err != nil {
		panic(err)
	}
	coursesUrls := []*middleware.ProxyTarget{
		{
			URL: courseUrl,
		},
	}
	courseProxyConf := middleware.ProxyConfig{
		Skipper:  nil,
		Balancer: middleware.NewRoundRobinBalancer(coursesUrls),
		Rewrite: map[string]string{
			"/api/courses*": "/api/v1/$1",
		},
	}
	tenant.Group("/courses", middleware.ProxyWithConfig(courseProxyConf))
}
