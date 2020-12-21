package errors

import (
	"errors"
)

var (
	// ErrUnknownServerName used when server name is unknown.
	ErrUnknownServerName = errors.New("unknown server name")
	// ErrNoServerConf used when no server is passed as a flag or configured as default.
	ErrNoServerConf = errors.New("no server configured as default or passed as flag")
	// ErrServerAlreadyExists used when adding a server with duplicated name.
	ErrServerAlreadyExists = errors.New("server already exists")
	// ErrUnknownVersionAction used when the user request an unknown action.
	ErrUnknownVersionAction = errors.New("version action unknown")
)
