package middlewares

import (
	"authorization/internal/helpers"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
)

func TenantGuard(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tenant := c.Get("tenant")
		if tenant == "" {
			return c.JSON(http.StatusNotFound, echo.Map{
				"error":   true,
				"message": "empty tenant",
			})
		}

		token := c.Get("user").(*jwt.Token)
		claims := token.Claims.(*helpers.JwtAuthClaims)

		hasTenant := false
		for _, member := range claims.Data.Membership {
			fmt.Println(tenant == member.Organization, tenant, member.Organization)
			if tenant == *member.Organization {
				hasTenant = true
				break
			}
		}

		if !hasTenant {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"error":   true,
				"message": "access restricted",
			})
		}

		return next(c)
	}
}
