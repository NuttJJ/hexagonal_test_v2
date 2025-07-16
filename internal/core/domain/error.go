package domain

import (
	"errors"
	"strings"
)

var (
	// ErrInternal is an error for when an internal service fails to process the request
	ErrInternal = errors.New("internal server error")

	// ErrDataNotFound is an error for when requested data is not found
	ErrDataNotFound = errors.New("data not found")

	// ErrNoUpdatedData is an error for when no data is provided to update
	ErrNoUpdatedData = errors.New("no data to update")

	// ErrConflictingData is an error for when data conflicts with existing data
	ErrConflictingData = errors.New("data conflicts with existing data in unique column")

	// ErrInvalidCredentials is an error for when the credentials are invalid
	ErrInvalidCredentials = errors.New("invalid email or password")

	// ErrEmptyAuthorizationHeader is an error for when the authorization header is empty
	ErrEmptyAuthorizationHeader = errors.New("authorization header is not provided")

	// ErrInvalidAuthorizationHeader is an error for when the authorization header is invalid
	ErrInvalidAuthorizationHeader = errors.New("authorization header format is invalid")

	// ErrInvalidAuthorizationType is an error for when the authorization type is invalid
	ErrInvalidAuthorizationType = errors.New("authorization type is not supported")

	// ErrUnauthorized is an error for when the user is unauthorized
	ErrUnauthorized = errors.New("user is unauthorized to access the resource")

	// ErrForbidden is an error for when the user is forbidden to access the resource
	ErrForbidden = errors.New("user is forbidden to access the resource")

	// ErrValidation is an error for when the request body is invalid
	ErrValidation = errors.New("validation error")

	// ErrInvalidEmail is an error for when the email is invalid
	ErrInvalidEmail = errors.New("invalid email")

	// ErrInvalidPassword is an error for when the password is invalid
	ErrInvalidPassword = errors.New("invalid password")

	// ErrPasswordMismatch is an error for when the password and password confirmation do not match
	ErrPasswordMismatch = errors.New("password and password confirmation do not match")

	// ErrUsernameAndPasswordRequired is an error for when the username and password are required
	ErrUsernameAndPasswordRequired = errors.New("username and password are required")

	// ErrOneTokenRequired is an error for when the ONE token is required
	ErrOneTokenRequired = errors.New("ONE token is required")

	// ErrAuthorizationCodeRequired is an error for when the authorization code is required
	ErrAuthorizationCodeRequired = errors.New("authorization code is required")

	// ErrTokenOrBizIDRequired is an error for when the token or business ID is required
	ErrTokenOrBizIDRequired = errors.New("token or business ID is required")

	// ErrTokenTypeAndAccessTokenRequired is an error for when the token type and access token are required
	ErrTokenTypeAndAccessTokenRequired = errors.New("token type and access token are required")

	// ErrInvalidTokenOrTokenExpired is an error for when the token is invalid
	ErrInvalidTokenOrTokenExpired = errors.New("invalid token")

	// ErrSignAccessToken is an error for when the access token cannot be signed
	ErrSignAccessToken = errors.New("error signing access token")

	// ErrEncodeAccessToken is an error for when the access token cannot be encoded
	ErrEncodeAccessToken = errors.New("error encoding access token")
)

// IsUniqueConstraintViolationError checks if the error is a unique constraint violation error
func IsUniqueConstraintViolationError(err error) bool {
	return strings.Contains(err.Error(), "23505")
}
