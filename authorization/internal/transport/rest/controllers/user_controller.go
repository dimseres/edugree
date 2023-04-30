package controllers

import (
	"authorization/internal/helpers"
	"authorization/internal/repositories"
	"authorization/internal/services"
	"authorization/internal/transport/rest/middlewares"
	"authorization/internal/transport/rest/responses"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func InitUserRoutes(app *echo.Group) {
	protected := app.Group("/")
	protected.Use(middlewares.JwtProtect())
	protected.GET("profile", Profile)
	protected.GET("list", ListUsers, middlewares.TenantGuard, middlewares.CasbinGuard("users", "read"))
}

func ListUsers(c echo.Context) error {
	repository := repositories.NewUserRepository()
	userService := services.NewUserService(&repository, &services.TenantContext{
		Id:     c.Get("tenant_id").(uint),
		Domain: c.Get("tenant").(string),
	})
	page, _ := strconv.Atoi(c.QueryParams().Get("page"))

	users, err := userService.GetUsersWithPagination(page, 25)
	if err != nil {
		return c.JSON(500, echo.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.JSON(200, echo.Map{
		"error": false,
		"data":  users,
	})
}

func Profile(c echo.Context) error {
	repository := repositories.NewUserRepository()
	service := services.NewUserService(&repository, nil)

	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*helpers.JwtAuthClaims)

	userId := claims.Data.UserId
	user := service.GetUserWith(userId, &[]string{
		"Membership",
		"Membership.Role",
		"Membership.Organization.Services",
	})

	if user.Id == 0 {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error":   true,
			"message": "something went wrong",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"error": false,
		"user":  responses.NewUserResponse(user),
	})
}
