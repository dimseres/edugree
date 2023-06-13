package middlewares

import (
	"authorization/config"
	"authorization/internal/helpers"
	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
)

func JwtProtect() echo.MiddlewareFunc {
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(helpers.JwtAuthClaims)
		},
		SigningKey: []byte(config.GetConfig("JWT_SECRET")),
		ErrorHandler: func(c echo.Context, err error) error {
			//authRepository := repositories.NewAuthRepository()
			//authService := services.NewAuthService(&authRepository)
			//refreshToken, err := c.Cookie("_ref")
			//if err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"error":   true,
				"message": "Unauthorized",
			})
			//}
			//tokens, err := authService.GenerateTokenFromRefresh(refreshToken.String())
			//if err != nil {
			//	return c.JSON(http.StatusUnauthorized, echo.Map{
			//		"error":   true,
			//		"message": "Unauthorized",
			//	})
			//}
			//helpers.SetAuthCookies(c, tokens.Token, tokens.Refresh)
			//
			//return nil
		},
	}

	return echojwt.WithConfig(config)
}
