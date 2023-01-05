package auth

import "github.com/labstack/echo/v4"

func InitRoutes(app *echo.Group) {
	app.POST("/login", Login)
	app.POST("/logout", Logout)
	app.POST("/register", Register)
	app.POST("/refresh", RefreshToken)
}

func Login(c echo.Context) error {
	return nil
}

func Logout(c echo.Context) error {
	return nil
}

func Register(c echo.Context) error {
	return nil
}

func RefreshToken(c echo.Context) error {
	return nil
}
