package permissions

import (
	"authorization/internal/transport/rest/middlewares"
	"github.com/labstack/echo/v4"
)

func InitRoutes(group *echo.Group) {
	group.Use(middlewares.JwtProtect())
	group.POST("/", createPermission)
	group.GET("/", getPermissionList)
}

func createPermission(c echo.Context) error {
	return nil
}

func getPermissionList(c echo.Context) error {
	return c.JSON(200, struct {
		Slug        string `json:"slug"`
		Title       string `json:"title"`
		Description string `json:"description"`
	}{
		Slug:        "Hello",
		Description: "World",
		Title:       "!",
	})
}
