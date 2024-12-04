package errlst

import "errors"

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
