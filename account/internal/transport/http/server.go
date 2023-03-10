package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartHttpServer(port string) {
	e := echo.New()
	if port == "" {
		port = "5000"
	}
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} ${method} ${uri} ${status} ${latency_human} ${bytes_in} ${bytes_out}\n",
	}))
	e.Static("/public", "storage/public")
	apiGroup := e.Group("/api")
	InitRoutes(apiGroup)
	e.Logger.Fatal(e.Start(":" + port))
}
