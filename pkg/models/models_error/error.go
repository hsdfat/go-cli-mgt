package models_error

import "errors"

var (
	MissingAuthHeader = errors.New("missing authorize header")
	InvalidAuthHeader = errors.New("authorize header of request invalid")
	InvalidToken      = errors.New("invalid token")
)
