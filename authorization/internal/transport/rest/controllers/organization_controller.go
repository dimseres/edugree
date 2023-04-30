package controllers

import (
	"authorization/internal/helpers"
	"authorization/internal/repositories"
	"authorization/internal/services"
	"authorization/internal/transport/rest/forms"
	"authorization/internal/transport/rest/middlewares"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func InitOrganizationRoutes(app *echo.Group) {
	protected := app.Group("/")
	protected.Use(middlewares.JwtProtect())
	protected.POST("", CreateOrganization)
	protected.GET(":id", GetOrganization)
}

func CreateOrganization(c echo.Context) error {
	var form forms.OrganizationCreate

	err := helpers.EchoControllerValidationHelper(c, &form)
	if err != nil {
		return err
	}

	repository := repositories.NewOrganizationRepository()
	service := services.NewOrganizationService(&repository)

	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*helpers.JwtAuthClaims)

	userId := claims.Data.UserId
	repository.StartTransaction()
	//defer func() {
	//	if r := recover(); r != nil {
	//		repository.RollbackTransaction()
	//	}
	//}()
	organization, err := service.CreateOrganization(&form, userId)

	if err != nil {
		repository.RollbackTransaction()
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	repository.EndTransaction()

	return c.JSON(http.StatusOK, echo.Map{
		"id":          organization.Id,
		"title":       organization.Title,
		"email":       organization.Email,
		"description": organization.Description,
		"avatar":      organization.Avatar,
	})
}

func GetOrganization(c echo.Context) error {
	_id := c.Param("id")
	id, err := strconv.Atoi(_id)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"error":   true,
			"message": "wrong id",
		})
	}

	repo := repositories.NewOrganizationRepository()
	service := services.NewOrganizationService(&repo)

	org, err := service.GetOrganization(uint(id))

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	if org.Id == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{
			"error":   true,
			"message": "not found",
		})
	}

	return c.JSON(200, org)
}
