package biz

import (
	v1 "api-getway/api/getway/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)
