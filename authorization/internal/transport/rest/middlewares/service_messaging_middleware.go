package middlewares

import (
	"authorization/config"
	"github.com/labstack/echo/v4"
)

func requestValidator(service string, token string) bool {
	switch service {
	case "courses":
		return config.GetConfig("COURSE_SERVICE_KEY") == token
	}
	return false
}

func ServiceMessagingMiddleware(service string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Request().Header.Get("X-SERVICE-KEY")
			if !requestValidator(service, token) {
				return c.JSON(403, echo.Map{
					"error":   true,
					"message": "access denied",
				})
			}
			return next(c)
		}
	}
}
