package middlewares

import (
	"authorization/internal/helpers"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
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

		_uuid, err := uuid.NewRandom()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"error":   true,
				"message": "server error",
			})
		}

		organizationToken, err := helpers.GetServiceJwtToken(claims, tenant.(string), _uuid.String())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"error":   true,
				"message": "server error",
			})
		}
		c.Request().Header.Set("X-REQUEST-ID", _uuid.String())
		c.Request().Header.Set("Authorization", "Bearer "+organizationToken)

		return next(c)
	}
}
