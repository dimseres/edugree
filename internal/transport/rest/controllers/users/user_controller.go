package users

import (
	"edugree_auth/internal/repositories/users"
	"github.com/labstack/echo/v4"
	"strconv"
)

func InitRoutes(app *echo.Group) {
	app.GET("/:id", getUserByid)
}

func getUserByid(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(500, err)
	}
	repository := users.NewRepository()
	model := repository.GetUserById(id)
	return c.JSON(200, model)
}
