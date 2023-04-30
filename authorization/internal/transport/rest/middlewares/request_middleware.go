package middlewares

import (
	"authorization/internal/repositories"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func InitRequest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_uuid, err := uuid.NewRandom()
		if err != nil {
			return c.JSON(500, echo.Map{
				"error":   true,
				"message": "Internal error",
			})
		}
		c.Request().Header.Set("X-REQUEST-ID", _uuid.String())

		jwtToken, err := c.Cookie("_token")
		if err == nil {
			c.Request().Header.Set("Authorization", "Bearer "+jwtToken.Value)
		}

		tenant := c.Request().Header.Get("X-ORG")
		c.Set("tenant", tenant)
		if tenant != "" {
			rep := repositories.NewCacheRepository()
			org, err := rep.GetTenantInfo(tenant)
			if err != nil {
				switch err.Error() {
				case "redis: nil":
					return c.JSON(404, echo.Map{
						"error":   true,
						"message": "Unknown tenant",
					})
				default:
					return c.JSON(500, echo.Map{
						"error":   true,
						"message": "Internal Error",
					})
				}

			}
			c.Set("tenant_id", org.Id)
		}
		return next(c)
	}
}
