package rest

import (
	"authorization/internal/transport/rest/forms"
	"authorization/internal/transport/rest/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

func StartHttpServer(port string) {
	e := echo.New()
	if port == "" {
		port = "5000"
	}
	if os.Getenv("APP_ENV") == "development" {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"*"},
			AllowCredentials: true,
		}))
	}

	e.Validator = forms.NewFormValidator()
	e.Pre(middlewares.InitRequest)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} ${method} ${uri} ${status} ${latency_human} ${bytes_in} ${bytes_out}\n",
	}))
	e.Static("/public", "storage/public")
	apiGroup := e.Group("/api")
	InitRoutes(apiGroup)
	e.Logger.Fatal(e.Start(":" + port))
}
