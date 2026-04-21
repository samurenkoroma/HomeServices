package domain

import "errors"

var (
	// User errors
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrInvalidEmail      = errors.New("invalid email")
	ErrInvalidPassword   = errors.New("invalid password")
	ErrPasswordTooShort  = errors.New("password too short (min 6 characters)")
	ErrEmailRequired     = errors.New("email is required")
	ErrUsernameRequired  = errors.New("username is required")
	ErrPasswordRequired  = errors.New("password is required")
	ErrUserBlocked       = errors.New("user is blocked")
	ErrUserInactive      = errors.New("user is inactive")

	// Auth errors
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrTokenExpired       = errors.New("token expired")
	ErrInvalidToken       = errors.New("invalid token")
	ErrUnauthorized       = errors.New("unauthorized")
	ErrForbidden          = errors.New("forbidden")
)
