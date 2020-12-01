package errors

import (
	"errors"
)

var (
	// ErrUnknownServerName used when server name is unknown.
	ErrUnknownServerName = errors.New("unknown server name")
	// ErrServerAlreadyExists used when adding a server with duplicated name.
	ErrServerAlreadyExists = errors.New("server already exists")
)
