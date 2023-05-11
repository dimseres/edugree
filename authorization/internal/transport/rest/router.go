package rest

import (
	"authorization/internal/transport/rest/controllers"
	"authorization/internal/transport/rest/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/url"
	"os"
)

func InitRoutes(app *echo.Group) {
	apiGroup := app.Group("/api")
	controllers.InitUploadRoutes(apiGroup.Group("/uploads"))
	controllers.InitAuthRoutes(apiGroup.Group("/auth"))
	controllers.InitUserRoutes(apiGroup.Group("/users"))
	controllers.InitOrganizationRoutes(apiGroup.Group("/organization"))
	controllers.InitMembershipRoutes(apiGroup.Group("/membership"))
	controllers.InitServiceMessagingController(apiGroup.Group("/messaging"))

	tenant := apiGroup.Group("")
	tenant.Use(middlewares.JwtProtect())
	tenant.Use(middlewares.TenantGuard)

	courseUrl, err := url.Parse(os.Getenv("COURSE_URL"))
	frontendUrl, err := url.Parse("http://localhost:5173")
	if err != nil {
		panic(err)
	}
	coursesUrls := []*middleware.ProxyTarget{
		{
			URL: courseUrl,
		},
	}
	frontendUrls := []*middleware.ProxyTarget{
		{
			URL: frontendUrl,
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
	app.Group("/*", middleware.ProxyWithConfig(middleware.ProxyConfig{
		Skipper:  nil,
		Balancer: middleware.NewRoundRobinBalancer(frontendUrls),
	}))
}
