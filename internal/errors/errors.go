package errors

import (
	"errors"
)

var (
	ErrUnknownServerName   = errors.New("unknown server name")
	ErrServerAlreadyExists = errors.New("server already exists")
)
