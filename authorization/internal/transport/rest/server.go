package rest

import (
	"authorization/internal/transport/rest/forms"
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
	e.Pre(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			jwtToken, err := c.Cookie("_token")
			if err == nil {
				c.Request().Header.Set("Authorization", "Bearer "+jwtToken.Value)
			}
			tenant := c.Request().Header.Get("x-org")
			c.Set("tenant", tenant)
			return next(c)
		}
	})
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} ${method} ${uri} ${status} ${latency_human} ${bytes_in} ${bytes_out}\n",
	}))
	e.Static("/public", "storage/public")
	apiGroup := e.Group("/api")
	InitRoutes(apiGroup)
	e.Logger.Fatal(e.Start(":" + port))
}
