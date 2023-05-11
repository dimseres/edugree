package rest

import (
	"authorization/config"
	"authorization/internal/repositories"
	"authorization/internal/transport/rest/forms"
	"authorization/internal/transport/rest/middlewares"
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo-contrib/prometheus"
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

	repo := repositories.NewAuthRepository()
	err := repo.SetInitialCache()
	if err != nil {
		config.GetLogger().Error(err)
		panic(err)
	}

	c := jaegertracing.New(e, nil)
	defer c.Close()
	p := prometheus.NewPrometheus("echo", nil)
	e.Use()
	p.Use(e)

	e.Validator = forms.NewFormValidator()
	e.Pre(middlewares.InitRequest)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} ${method} ${uri} ${status} ${latency_human} ${bytes_in} ${bytes_out}\n",
	}))
	e.Static("/public", "storage/public")
	group := e.Group("")
	InitRoutes(group)
	e.Logger.Fatal(e.Start(":" + port))

}
