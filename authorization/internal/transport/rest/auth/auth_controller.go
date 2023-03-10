package auth

import (
	"authorization/internal/models"
	"authorization/internal/repositories"
	"authorization/internal/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

func InitRoutes(app *echo.Group) {
	app.POST("/login", Login)
	app.POST("/logout", Logout)
	app.POST("/register", Register)
	app.POST("/refresh", RefreshToken)
}

func Login(c echo.Context) error {
	username := c.FormValue("email")
	password := c.FormValue("password")

	if username == "" || password == "" {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"error":   true,
			"message": "password and login required",
		})
	}

	repository := repositories.NewAuthRepository()
	service := services.NewAuthService(&repository)

	err, user := service.SignIn(username, password)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	err, token := service.CreateJwtToken(user)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"error": false,
		"user": models.PublicUser{
			BaseModel: user.BaseModel,
			BaseUser:  user.BaseUser,
			Role:      user.Role,
			Token:     nil,
		},
		"token": token,
	})
}

func Logout(c echo.Context) error {
	return nil
}

func Register(c echo.Context) error {
	return nil
}

func RefreshToken(c echo.Context) error {
	return nil
}
