package internal

import (
	"edugree_auth/internal/transport/rest"
	"github.com/labstack/echo/v4"
)

func StartHttpServer(port string) {
	e := echo.New()
	if port == "" {
		port = "5000"
	}
	e.Static("/public", "storage/public")
	apiGroup := e.Group("/api")
	rest.InitRoutes(apiGroup)
	e.Logger.Fatal(e.Start(":" + port))
}
