package forms

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

//var instance *FormValidator

type FormValidator struct {
	validator *validator.Validate
}

type FormError interface {
	GetValidationMessage() string
}

func NewFormValidator() *FormValidator {
	return &FormValidator{validator: validator.New()}
}

type ApiError struct {
	Param   string `json:"param"`
	Message string `json:"message"`
}

func (self *FormValidator) Validate(i interface{}) error {
	err := self.validator.Struct(i)
	if err != nil {
		validationError := err.(validator.ValidationErrors)
		if validationError != nil {
			if errors.As(err, &validationError) {
				out := make([]ApiError, len(validationError))
				for i, fe := range validationError {
					out[i] = ApiError{fe.Field(), msgForTag(fe)}
				}
				return echo.NewHTTPError(http.StatusUnprocessableEntity, echo.Map{
					"error":   true,
					"message": out[0].Param + " " + out[0].Message,
					"errors":  out,
				})
			}
			return echo.NewHTTPError(http.StatusUnprocessableEntity, echo.Map{
				"error":   true,
				"message": "Validation error",
				"errors":  validationError.Error(),
			})
		}
	}

	return nil
}

func msgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	}
	return fe.Error() // default error
}
