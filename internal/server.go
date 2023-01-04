package internal

import "github.com/labstack/echo/v4"

func StartHttpServer(port string) {
	e := echo.New()
	if port == "" {
		port = "5000"
	}
	e.Static("/public", "storage/public")
	e.Logger.Fatal(e.Start(":" + port))
}
