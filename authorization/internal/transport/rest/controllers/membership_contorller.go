package controllers

import (
	"authorization/internal/repositories"
	"authorization/internal/services"
	"authorization/internal/transport/rest/middlewares"
	"github.com/labstack/echo/v4"
	"strconv"
)

func InitMembershipRoutes(app *echo.Group) {
	protected := app.Group("/")
	protected.Use(middlewares.JwtProtect())
	protected.GET("users", ListUsers, middlewares.TenantGuard, middlewares.CasbinGuard("users", "read"))
}

func ListUsers(c echo.Context) error {
	repository := repositories.NewUserRepository()
	userService := services.NewUserService(&repository, &services.TenantContext{
		Id:     c.Get("tenant_id").(uint),
		Domain: c.Get("tenant").(string),
	})
	page, _ := strconv.Atoi(c.QueryParams().Get("page"))

	data, err := userService.GetUsersWithPagination(page, 25)
	if err != nil {
		return c.JSON(500, echo.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.JSON(200, echo.Map{
		"error": false,
		"total": data.Total,
		"pages": data.MaxPage,
		"data":  data.Data,
	})
}
