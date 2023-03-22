package middlewares

import (
	"authorization/internal/helpers"
	"authorization/internal/repositories"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CasbinGuard(obj string, action string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(*helpers.JwtAuthClaims)

			//domain := c.Request().Host
			domain := "example.org"
			//domain := "feelgoodinc"

			domainRole := ""

			for _, membership := range claims.Data.Membership {
				if *membership.Organization == domain {
					domainRole = *membership.Role
				}
			}

			enforcer := repositories.NewEnforcerRepository()
			permission, err := enforcer.EnforcePermission(domainRole, obj, action, domain)

			if err != nil {
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"error":   true,
					"message": err.Error(),
				})
			}

			if !permission {
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"error":   true,
					"message": "access denied",
				})
			}

			fmt.Println(permission)

			return next(c)
		}
	}
}
