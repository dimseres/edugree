package middlewares

import (
	"edugree_auth/internal/helpers"
	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

func JwtProtect() echo.MiddlewareFunc {
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(helpers.JwtAuthClaims)
		},
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"error":   true,
				"message": "Unauthorized",
			})
		},
	}

	return echojwt.WithConfig(config)
}
