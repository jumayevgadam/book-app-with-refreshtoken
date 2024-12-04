package errlst

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/config"
	"github.com/jumayevgadam/book-app-with-refreshtoken/pkg/logger"
)

// errors are
var (
	// ErrBadRequest is
	ErrBadRequest = errors.New("bad request")
	// ErrBadQueryParams is
	ErrBadQueryParams = errors.New("bad query params")
	// ErrDBConnection is
	ErrDBConnection = errors.New("cannot connect to psqlDB")
	// ErrDBPing is
	ErrDBPing = errors.New("can not give ping to psqlDB")
	// ErrCommitTx is
	ErrCommitTx = errors.New("can not commit completed actions in DB")
	// ErrTypeAssertInTransaction is
	ErrTypeAssertInTransaction = errors.New("error in type assertion to connection.DBOps")
	// ErrBeginTransaction is
	ErrBeginTransaction = errors.New("cannot start transaction")
	// ErrUnauthorized is
	ErrUnauthorized = errors.New("unauthorized")
	// ErrNotFound is
	ErrNotFound = errors.New("not found")
	// ErrConflict is
	ErrConflict = errors.New("conflict occured")
	// ErrForbidden is
	ErrForbidden = errors.New("forbidden")
	// ErrFieldValidation is
	ErrFieldValidation = errors.New("field validation error")
	// ErrNoSuchUser is
	ErrNoSuchUser = errors.New("no such user")
	// ErrNoSuchRole is
	ErrNoSuchRole = errors.New("no roles found for this permission")
	// ErrInternalServer is
	ErrInternalServer = errors.New("internal server error")
	// ErrTransactionFailed is
	ErrTransactionFailed = errors.New("failed to perform transaction")
	// ErrInvalidJWTToken is
	ErrInvalidJWTToken = errors.New("invalid JWT Token")
	// ErrTokenExpired is
	ErrTokenExpired = errors.New("token is expired")
	// ErrInvalidJWTMethod is
	ErrInvalidJWTMethod = errors.New("invalid jwt token method")
	// ErrInvalidJWTClaims is
	ErrInvalidJWTClaims = errors.New("invalid JWT Claims")

	// ErrRange is
	ErrRange = errors.New("value out of range")
	// ErrSyntax is
	ErrSyntax = errors.New("invalid syntax")
)

var _ RestErr = (*RestError)(nil)

// RestErr interface needs capturing errors
type RestErr interface {
	Status() int
	Error() string
	Causes() interface{}
	// AppearedAt() time.Time
}

// RestError struct to implement the RestErr interface
type RestError struct {
	ErrStatus  int         `json:"err_status,omitempty"`
	ErrMessage string      `json:"err_msg,omitempty"`
	ErrCause   interface{} `json:"err_cause,omitempty"`
}

// ---------- IMPLEMENTING RestErr methods -------------
// Status is
func (e RestError) Status() int {
	return e.ErrStatus
}

// Causes is
func (e RestError) Causes() interface{} {
	return e.ErrCause
}

// ErrorMessage is
func (e RestError) Error() string {
	return fmt.Sprintf("status: %d - err_msg: %s - causes: %v - appearedAt: %v",
		e.ErrStatus, e.ErrMessage, e.ErrCause, time.Now())
}

// ------------------- FACTORY FUNCTIONS FOR CREATING ERRORS ---------------------
// NewBadRequestError creates a new 400 bad request error
func NewBadRequestError(cause interface{}) RestErr {
	return &RestError{
		ErrStatus:  http.StatusBadRequest,
		ErrMessage: ErrBadRequest.Error(),
		ErrCause:   cause,
	}
}

// NewNotFoundError creates a new 404 not found error
func NewNotFoundError(cause interface{}) RestErr {
	return &RestError{
		ErrStatus:  http.StatusNotFound,
		ErrMessage: ErrNotFound.Error(),
		ErrCause:   cause,
	}
}

// NewBadQueryParamsError creates a 403 bad query params error
func NewBadQueryParamsError(cause interface{}) RestErr {
	return &RestError{
		ErrStatus:  http.StatusBadRequest,
		ErrMessage: ErrBadQueryParams.Error(),
		ErrCause:   cause,
	}
}

// NewUnauthorizedError creates a 401 unauthorized error
func NewUnauthorizedError(cause interface{}) RestErr {
	return &RestError{
		ErrStatus:  http.StatusUnauthorized,
		ErrMessage: ErrUnauthorized.Error(),
		ErrCause:   cause,
	}
}

// NewInternalServerError creates a 500 internal server error
func NewInternalServerError(cause interface{}) RestErr {
	return &RestError{
		ErrStatus:  http.StatusInternalServerError,
		ErrMessage: ErrInternalServer.Error(),
		ErrCause:   cause,
	}
}

// NewConflictError creates a new 409 conflict error
func NewConflictError(cause interface{}) RestErr {
	return &RestError{
		ErrStatus:  http.StatusConflict,
		ErrMessage: ErrConflict.Error(),
		ErrCause:   cause,
	}
}

// NewForbiddenError creates a new 403 forbidden error
func NewForbiddenError(cause interface{}) RestErr {
	return &RestError{
		ErrStatus:  http.StatusForbidden,
		ErrMessage: ErrForbidden.Error(),
		ErrCause:   cause,
	}
}

// ParseValidatorError parses validation errors and returns corresponding RestErr
func ParseValidatorError(err error) RestErr {
	validationErrs, ok := err.(validator.ValidationErrors)
	if !ok {
		return NewBadRequestError(err.Error())
	}

	// Collect detailed validation error messages
	var errorMessages []string
	for _, fieldErr := range validationErrs {
		// For each validation error, create a message
		errorMessage := fmt.Sprintf("Field validation for %s, failed on the %s tag",
			fieldErr.Field(), fieldErr.Tag())
		// append each error to error messages
		errorMessages = append(errorMessages, errorMessage)
	}

	// Combine all messages into one string and return a BadRequestError
	return NewBadRequestError(strings.Join(errorMessages, ", "))
}

// ParseSQLErrors returns corresponding RestErr
func ParseSQLErrors(err error) RestErr {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		// CLASS 22
		case "22001": // numeric value out of range
			return NewBadRequestError("Numeric" + ErrRange.Error() + pgErr.Message)

		// CLASS 23
		case "23505": // Unique violation
			return NewConflictError("Unique constraint violation: " + ErrConflict.Error() + pgErr.Message)
		case "23503": // Foreign key violation
			return NewBadRequestError("Foreign key violation: " + pgErr.Message)
		case "23502": // Not null violation
			return NewBadRequestError("Not null violation: " + pgErr.Message)

		// CLASS 40
		case "40001": // serialization failure
			return NewConflictError("Serialization error: " + pgErr.Message)
		// CLASS 42
		case "42601": // syntax error
			return NewBadRequestError("Syntax error in sql statements: " + pgErr.Message)
		}
	}

	if strings.Contains(err.Error(), "scany") {
		return NewBadRequestError(err.Error())
	}

	if strings.Contains(err.Error(), "no corresponding field found") {
		return NewBadRequestError(err.Error())
	}

	return NewBadRequestError(err.Error())
}

// ParseErrors parses common errors (like SQL errors) into the RestErr
func ParseErrors(err error) RestErr {
	logger := logger.NewApiLogger(&config.Config{})
	logger.InitLogger()

	switch {
	// pgx specific errors
	case errors.Is(err, pgx.ErrNoRows):
		logger.Warn("Not found error: ", err)
		return NewNotFoundError(err.Error())
	case errors.Is(err, pgx.ErrTooManyRows):
		logger.Warn("Conflict error: ", err)
		return NewConflictError(err.Error())

	// SQLSTATE error
	case strings.Contains(err.Error(), "SQLSTATE"):
		logger.Error("SQL error: ", err)
		return ParseSQLErrors(err)

	// Handle strconv.Atoi errors
	case strings.Contains(err.Error(), ErrSyntax.Error()),
		strings.Contains(err.Error(), ErrRange.Error()):
		logger.Warn("bad request error: ", err)
		return NewBadRequestError(err.Error())

		// Handle Validation errors from go-validator/v10
	case errors.As(err, &validator.ValidationErrors{}):
		logger.Warn("Validation error: ", err)
		return ParseValidatorError(err)

	// Handle Token or Cookie errors
	case
		strings.Contains(strings.ToLower(err.Error()), ErrInvalidJWTToken.Error()),
		strings.Contains(strings.ToLower(err.Error()), ErrInvalidJWTClaims.Error()):
		logger.Warn("Unauthorized error: ", err)
		return NewUnauthorizedError(ErrUnauthorized.Error() + err.Error())

	default:
		// If already a RestErr, return as-is
		if restErr, ok := err.(RestErr); ok {
			logger.Info("Custom RestErr: ", restErr)
			return restErr
		}

		// For any other errors
		logger.Error("Internal server error: ", err)
		return NewInternalServerError(err.Error())
	}
}
