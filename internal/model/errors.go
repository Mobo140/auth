package model

import "errors"

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrEndpointsNotFound = errors.New("endpoints not found")
)
