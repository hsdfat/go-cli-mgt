package models_error

import "errors"

var (
	MissingAuthHeader = errors.New("missing authorize header")
	InvalidAuthHeader = errors.New("authorize header of request invalid")
	InvalidToken      = errors.New("invalid token")
)

var (
	ErrNotFoundUser = errors.New("user not found")
	ErrDisableUser  = errors.New("user have been disable")
	ErrEnableUser   = errors.New("user have been active")
)

var (
	ErrNotFoundNe = errors.New("network element not found")
)

var (
	ErrNotFoundUserNe = errors.New("user do not have permission with ne")
)

var (
	ErrNotFoundRole = errors.New("role not found")
)

var (
	ErrNotFoundUserRole = errors.New("user do not have this role")
)

var (
	ErrNotFoundHistory = errors.New("history command not found")
)
