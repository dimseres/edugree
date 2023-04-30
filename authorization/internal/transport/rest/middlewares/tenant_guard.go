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
		tenant, tenantId := c.Get("tenant"), c.Get("tenant_id")
		if tenant == "" || tenantId == "" {
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

		requestId := c.Request().Header.Get("X-REQUEST-ID")
		organizationToken, err := helpers.GetServiceJwtToken(claims, tenant.(string), requestId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"error":   true,
				"message": "server error",
			})
		}

		c.Request().Header.Set("Authorization", "Bearer "+organizationToken)

		return next(c)
	}
}
