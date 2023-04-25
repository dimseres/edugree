package controllers

import (
	"authorization/internal/repositories"
	"authorization/internal/transport/rest/middlewares"
	"github.com/labstack/echo/v4"
	"strconv"
)

func InitServiceMessagingController(app *echo.Group) {
	courses := app.Group("/courses")
	courses.Use(middlewares.ServiceMessagingMiddleware("courses"))
	courses.GET("/organization/:id", GetOrganizationData)
}

func GetOrganizationData(c echo.Context) error {
	orgId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(422, echo.Map{
			"error":   true,
			"message": "unknown id",
		})
	}

	repository := repositories.NewOrganizationRepository()
	org, err := repository.GetOrganizationWithMembers(uint(orgId))
	if err != nil {
		return c.JSON(404, echo.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.JSON(200, echo.Map{
		"error": false,
		"data":  org,
	})
}
