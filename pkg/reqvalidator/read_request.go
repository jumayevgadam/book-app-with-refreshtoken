package reqvalidator

import (
	"log"
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/jumayevgadam/book-app-with-refreshtoken/pkg/errlst"
	"github.com/labstack/echo/v4"
)

var (
	validate   *validator.Validate
	phoneRegex = `^\+993(6[0-5]|7[0-5])\d{7}$`
	emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
)

func init() {
	validate = validator.New()

	// Register custom phone validation
	err := validate.RegisterValidation("phone", func(fl validator.FieldLevel) bool {
		return regexp.MustCompile(phoneRegex).MatchString(fl.Field().String())
	})
	if err != nil {
		log.Printf("[reqvalidator][init]: unable to put validator for phone number: %v", err)
	}

	// Register custom email validation.
	err = validate.RegisterValidation("email", func(fl validator.FieldLevel) bool {
		return regexp.MustCompile(emailRegex).MatchString(fl.Field().String())
	})
	if err != nil {
		log.Printf("[reqvalidator][init]: unable to put validator for email: %v", err)
	}
}

// ReadRequest parses and validates the request body.
func ReadRequest(ctx echo.Context, request any) error {
	// Parse the request body into the provided struct.
	err := ctx.Bind(request)
	if err != nil {
		return errlst.ErrBadQueryParams
	}

	// Validate the parsed struct.
	err = validate.StructCtx(ctx.Request().Context(), request)
	if err != nil {
		return errlst.ErrFieldValidation
	}

	return nil
}
