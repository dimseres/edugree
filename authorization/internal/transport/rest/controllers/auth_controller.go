package controllers

import (
	"authorization/internal/casbin"
	"authorization/internal/dto"
	"authorization/internal/helpers"
	"authorization/internal/repositories"
	"authorization/internal/services"
	"authorization/internal/transport/rest/forms"
	"authorization/internal/transport/rest/middlewares"
	"authorization/internal/transport/rest/responses"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func InitAuthRoutes(app *echo.Group) {
	app.POST("/login", Login)
	app.POST("/register", Register)
	app.POST("/refresh", RefreshToken)

	protected := app.Group("/")
	protected.Use(middlewares.JwtProtect())
	protected.POST("logout", Logout)
	protected.POST("setTenant", SetTenant)
	protected.GET("permissions", GetPermissionList)
	//protected.POST("profile", Profile)
	protected.GET("test", TestFunc, middlewares.CasbinGuard("courses", "read"))
}

type UserOutputDTO struct {
	Id         uint    `json:"id"`
	Email      string  `gorm:"not null;unique" json:"email"`
	Phone      string  `gorm:"size:256;not null;unique" json:"phone"`
	FullName   string  `gorm:"size:512;not null" json:"full_name"`
	Avatar     *string `json:"avatar"`
	Bio        *string `gorm:"type:text" json:"bio"`
	Active     bool    `gorm:"not null;default:true" json:"active"`
	Membership []struct {
	} `json:"membership"`
}

func TestFunc(c echo.Context) error {
	return c.String(200, "PASSED")
}

func Login(c echo.Context) error {
	domain := c.Get("tenant").(string)

	var form forms.UserLoginForm

	err := helpers.EchoControllerValidationHelper(c, &form)

	if err != nil {
		return err
	}

	repository := repositories.NewAuthRepository()
	service := services.NewAuthService(&repository)

	err, user := service.SignIn(form.Email, form.Password)

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

	SetAuthCookies(c, token, refresh)

	return c.JSON(http.StatusOK, echo.Map{
		"error": false,
		"user":  responses.NewUserResponse(user),
		//"token": token,
	})
}

func Logout(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	//claims := user.Claims.(*helpers.JwtAuthClaims)
	//userId := claims.Data.UserId

	refresh, _ := c.Cookie("_ref")
	if refresh == nil && refresh.Value != "" {
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

	SetAuthCookies(c, "-", "-")
	return c.JSON(http.StatusOK, echo.Map{
		"error":   false,
		"message": "success",
	})
}

func Register(c echo.Context) error {
	var form forms.UserRegistrationForm

	err := helpers.EchoControllerValidationHelper(c, &form)

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
		"user":  responses.NewUserResponse(user),
	})
}

func RefreshToken(c echo.Context) error {
	refresh, _ := c.Cookie("_ref")
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

	SetAuthCookies(c, newToken.Token, newToken.Refresh)

	return c.JSON(http.StatusOK, echo.Map{
		"error": false,
		"user":  responses.NewUserResponse(newToken.User),
		"token": newToken.Token,
	})
}

func SetAuthCookies(c echo.Context, jwtToken string, refreshToken string) {
	c.SetCookie(&http.Cookie{
		Name:  "_ref",
		Value: refreshToken,
		Path:  "/",
		//Domain:   c.Request().Host,
		Expires: time.Now().Add(services.REFRESH_LIFETIME),
		//Secure:   true,
		HttpOnly: true,
	})

	c.SetCookie(&http.Cookie{
		Name:  "_token",
		Value: jwtToken,
		Path:  "/",
		//Domain:   c.Request().Host,
		Expires: time.Now().Add(services.JWT_LIFETIME),
		//Secure:   true,
		HttpOnly: true,
	})
}

func SetTenant(c echo.Context) error {
	var form forms.SetTenantForm
	err := helpers.EchoControllerValidationHelper(c, &form)

	if err != nil {
		return err
	}

	repository := repositories.NewMembershipRepository()
	service := services.NewMembershipService(&repository)

	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*helpers.JwtAuthClaims)

	userId := claims.Data.UserId

	member, err := service.GetMembershipData(userId, form.TenantId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error":   true,
			"message": err,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"error": false,
		"data":  responses.NewSetTenantResponse(member),
	})
}

func GetPermissionList(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*helpers.JwtAuthClaims)

	userId := claims.Data.UserId
	organization := c.Get("tenant").(string)

	membership := claims.Data.Membership
	valid := false
	for _, member := range membership {
		if *member.Organization == organization {
			valid = true
		}
	}

	if !valid {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error":   true,
			"message": "Unauthorized",
		})
	}

	casb, err := casbin.GetEnforcer().GetAllDomains()

	if err != nil {
		panic(err)
	}

	fmt.Println(userId, organization, casb)

	return nil
}
