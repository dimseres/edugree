package helpers

import (
	"authorization/internal/transport/rest/forms"
	"encoding/hex"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/scrypt"
	"net/http"
	"os"
	"strconv"
	"time"
)

func CreatePasswordHash(password string) (string, error) {
	firstSalt, err := scrypt.Key([]byte(password), []byte(os.Getenv("SALT")), 2048, 4, 2, 32)
	passwordHash, err := scrypt.Key([]byte(password), firstSalt, 16384, 8, 1, 32)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(passwordHash), nil
}

type CompareFailed struct {
	Message string
}

func (e *CompareFailed) Error() string {
	return fmt.Sprintf("parse %v:", e.Message)
}

func ComparePasswordAndHash(password string, hashedPassword string) error {
	password, err := CreatePasswordHash(password)
	if err != nil {
		return err
	}
	if password != hashedPassword {
		return &CompareFailed{Message: "password not compared"}
	}
	return nil
}

type JwtData struct {
	UserId     uint                `json:"user_id"`
	Name       string              `json:"name"`
	Membership []JwtMembershipData `json:"membership"`
}

type JwtMembershipData struct {
	Organization   *string   `json:"organization"`
	OrganizationId *uint     `json:"organization_id"`
	Role           *string   `json:"role"`
	Services       *[]string `json:"services"`
}

type JwtAuthClaims struct {
	Data JwtData `json:"data"`
	jwt.RegisteredClaims
}

type JwtCreateError struct {
	Message string
}

func (e *JwtCreateError) Error() string {
	return fmt.Sprintf("JWT creation error: %v", e.Message)
}

func CreateAuthToken(payload JwtData) (error, string) {
	lifetime, err := strconv.Atoi(os.Getenv("JWT_LIFETIME"))

	if err != nil {
		return &JwtCreateError{Message: err.Error()}, ""
	}

	claims := &JwtAuthClaims{
		Data: payload,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(lifetime))),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return err, ""
	}

	return err, t
}

func GetValidatedForm(form any, c echo.Context) interface{} {

	return form
}

func EchoControllerValidationHelper(c echo.Context, form interface{}) error {
	err := c.Bind(form)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"error":   true,
			"message": "bad request",
		})
	}
	var validator = forms.NewFormValidator()
	err = validator.Validate(form)
	if err != nil {
		return err
	}
	return nil
}

func NewError(message string) *CustomError {
	err := CustomError{err: message}
	return &err
}

type CustomError struct {
	err string
}

func (e *CustomError) Error() string {
	return e.err
}
