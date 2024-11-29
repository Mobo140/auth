package model

import "errors"

var (
	ErrorUserNotFound = errors.New("user not found")
	ErrorEndpointsNotFound = errors.New("endpoints not found")
)

