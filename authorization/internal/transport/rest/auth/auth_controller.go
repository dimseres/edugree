package auth

import (
	"authorization/internal/dto"
	"authorization/internal/repositories"
	"authorization/internal/services"
	"authorization/internal/transport/rest/forms"
	"authorization/internal/transport/rest/middlewares"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func InitRoutes(app *echo.Group) {
	app.POST("/login", Login)
	app.POST("/register", Register)
	app.POST("/refresh", RefreshToken)

	protected := app.Group("/")
	protected.Use(middlewares.JwtProtect())
	protected.POST("logout", Logout)
}

func Login(c echo.Context) error {
	username := c.FormValue("email")
	password := c.FormValue("password")
	domain := c.FormValue("domain")

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

	err, token := service.CreateJwtToken(user, domain)

	if err != nil {
		return err
	}

	refresh, err := service.CreateRefreshToken(token, user)

	if err != nil {
		return err
	}
	refreshCookie, _ := c.Cookie("refresh")
	fmt.Println(refreshCookie)

	c.SetCookie(&http.Cookie{
		Name:  "refresh",
		Value: refresh,
		Path:  "/",
		//Domain:   c.Request().Host,
		Expires: time.Now().Add(services.REFRESH_LIFETIME),
		//Secure:   true,
		HttpOnly: true,
	})

	return c.JSON(http.StatusOK, echo.Map{
		"error": false,
		"user":  user,
		"token": token,
	})
}

func Logout(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	//claims := user.Claims.(*helpers.JwtAuthClaims)
	//userId := claims.Data.UserId

	refresh, _ := c.Cookie("refresh")
	if refresh == nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"error":   true,
			"message": "unknown refresh token",
		})
	}

	repository := repositories.NewAuthRepository()
	service := services.NewAuthService(&repository)

	err := service.Logout(user.Raw, refresh.Value)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"error":   false,
		"message": "success",
	})
}

func Register(c echo.Context) error {
	var form forms.UserRegistrationForm

	err := c.Bind(&form)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"error":   true,
			"message": "bad request",
		})
	}
	var validator = forms.NewFormValidator()
	err = validator.Validate(&form)
	if err != nil {
		return err
	}

	repository := repositories.NewUserRepository()
	service := services.NewUserService(&repository)

	user, err := service.CreateUser(&dto.CreateUserDTO{
		Email:    form.Email,
		Phone:    form.Phone,
		Password: form.Password,
		FullName: form.FullName,
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"error": false,
		"user":  user,
	})
}

func RefreshToken(c echo.Context) error {
	refresh, _ := c.Cookie("refresh")
	if refresh == nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"error":   true,
			"message": "unknown refresh token",
		})
	}

	repository := repositories.NewAuthRepository()
	service := services.NewAuthService(&repository)

	newToken, err := service.GenerateTokenFromRefresh(refresh.Value)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	c.SetCookie(&http.Cookie{
		Name:  "refresh",
		Value: newToken.Refresh,
		Path:  "/",
		//Domain:   c.Request().Host,
		Expires: time.Now().Add(services.REFRESH_LIFETIME),
		//Secure:   true,
		HttpOnly: true,
	})

	return c.JSON(http.StatusOK, echo.Map{
		"error": false,
		"user":  newToken.User,
		"token": newToken.Token,
	})
}
