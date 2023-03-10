package users

import (
	"authorization/internal/repositories"
	"authorization/internal/services"
	"authorization/internal/transport/rest/middlewares"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func InitRoutes(app *echo.Group) {
	app.Use(middlewares.JwtProtect())
	app.GET("/:id", getUserById)
}

func getUserById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(500, err)
	}

	repository := repositories.NewUserRepository()
	service := services.NewUserService(&repository)
	result := service.GetUser(uint(id))

	return c.JSON(http.StatusOK, result)
}
