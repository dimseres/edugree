package controllers

import (
	"authorization/internal/helpers"
	"authorization/internal/repositories"
	"authorization/internal/services"
	"authorization/internal/transport/rest/forms"
	"authorization/internal/transport/rest/middlewares"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func InitMembershipRoutes(app *echo.Group) {
	protected := app.Group("/")
	protected.Use(middlewares.JwtProtect())
	protected.GET("users", ListUsers, middlewares.TenantGuard, middlewares.CasbinGuard("membership", "read"))
	protected.DELETE("users/:id", RemoveMember, middlewares.TenantGuard, middlewares.CasbinGuard("membership", "delete"))
	protected.POST("invites", InviteMembers, middlewares.TenantGuard, middlewares.CasbinGuard("membership", "create"))
	protected.GET("invites/create", GetInvitesConstants, middlewares.TenantGuard, middlewares.CasbinGuard("membership", "create"))
	protected.GET("invites", GetInviteList, middlewares.TenantGuard, middlewares.CasbinGuard("membership", "read"))
	protected.GET("invites/:link/join", JoinOrganization, middlewares.TenantGuard)
	protected.DELETE("invites/:link", GetInviteList, middlewares.TenantGuard, middlewares.CasbinGuard("membership", "delete"))
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
		return c.JSON(500, echo.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	repo := repositories.NewMembershipRepository()
	service := services.NewMembershipService(&repo, helpers.GetDomainContext(c))
	_, err = service.InviteMembers(&form)
	if err != nil {
		return c.JSON(500, echo.Map{
			"error":   true,
			"message": err.Error(),
		})
	}
	return c.JSON(200, echo.Map{
		"error":   false,
		"message": "ok",
	})
}

func GetInviteList(c echo.Context) error {
	repo := repositories.NewMembershipRepository()
	service := services.NewMembershipService(&repo, helpers.GetDomainContext(c))

	page, _ := strconv.Atoi(c.QueryParams().Get("page"))

	res, err := service.GetInviteList(page, 25)
	if err != nil {
		return c.JSON(500, echo.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.JSON(200, echo.Map{
		"error": false,
		"data":  res,
	})
}

func JoinOrganization(c echo.Context) error {
	link := c.Param("link")
	action := c.QueryParam("action")
	if link == "" {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"error":   true,
			"message": "wrong link",
		})
	}
	repo := repositories.NewMembershipRepository()
	service := services.NewMembershipService(&repo, helpers.GetDomainContext(c))

	res, err := service.JoinOrganization(link, action)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"error":   true,
			"message": "wrong link",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"error": true,
		"data":  res,
	})
}

func ChangeMemberRole(c echo.Context) error {

	return nil
}

func GetInvitesConstants(c echo.Context) error {
	tenant := helpers.GetDomainContext(c)
	availableRoles, ok := helpers.GetCreateAvailableRoles(tenant.Role)
	if !ok {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"error":   true,
			"message": "Unknown role",
		})
	}

	repo := repositories.NewRoleRepository()
	res, err := repo.GetRolesBySlug(availableRoles)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"error": false,
		"data": map[string]interface{}{
			"roles": res,
		},
	})
}
