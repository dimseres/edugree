package users

import (
	"edugree_auth/internal/repositories"
	"edugree_auth/internal/services"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func InitRoutes(app *echo.Group) {
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
