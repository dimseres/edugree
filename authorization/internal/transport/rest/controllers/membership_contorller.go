package controllers

import (
	"authorization/internal/helpers"
	"authorization/internal/repositories"
	"authorization/internal/services"
	"authorization/internal/transport/rest/forms"
	"authorization/internal/transport/rest/middlewares"
	"github.com/labstack/echo/v4"
	"strconv"
)

func InitMembershipRoutes(app *echo.Group) {
	protected := app.Group("/")
	protected.Use(middlewares.JwtProtect())
	protected.GET("users", ListUsers, middlewares.TenantGuard, middlewares.CasbinGuard("users", "read"))
	protected.DELETE("users/:id", RemoveMember, middlewares.TenantGuard, middlewares.CasbinGuard("membership", "delete"))
	protected.POST("users/invite", InviteMembers, middlewares.TenantGuard, middlewares.CasbinGuard("membership", "create"))
	protected.PATCH("users/:id/role", ChangeMemberRole, middlewares.TenantGuard, middlewares.CasbinGuard("membership", "update"))
}

func RemoveMember(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(422, echo.Map{
			"error":   true,
			"message": "wrong id",
		})
	}
	repository := repositories.NewMembershipRepository()
	membershipService := services.NewMembershipService(&repository, &helpers.TenantContext{
		Id:     c.Get("tenant_id").(uint),
		Domain: c.Get("tenant").(string),
	})
	err = membershipService.DeleteMember(uint(userId))
	if err != nil {
		errCode := 422
		if err.Error() == "record not found" {
			errCode = 404
		}
		return c.JSON(errCode, echo.Map{
			"error":   true,
			"message": err.Error(),
		})
	}
	return c.JSON(200, echo.Map{
		"error":   false,
		"message": "ok",
	})
}

func ListUsers(c echo.Context) error {
	repository := repositories.NewUserRepository()
	userService := services.NewUserService(&repository, &helpers.TenantContext{
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

func InviteMembers(c echo.Context) error {
	var form forms.InviteMembersForm
	bodyReader := c.Request().Body

	err := helpers.ValidateJsonForm(bodyReader, &form)
	if err != nil {
		return err
	}

	repo := repositories.NewMembershipRepository()
	service := services.NewMembershipService(&repo, helpers.GetDomainContext(c))
	_, err = service.InviteMembers(&form)

	return nil
}

func ChangeMemberRole(c echo.Context) error {

	return nil
}
