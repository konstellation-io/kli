package logger

import (
	"io"

	"github.com/guumaster/cligger"
)

// NOTE: This module imports cligger interface here to enable mocks for testing.

// Logger interface exported from cligger module.
type Logger interface {
	cligger.Logger
}

// NewLogger creates a cligger instance with the given output writer.
func NewLogger(w io.Writer) Logger {
	return cligger.NewLoggerWithWriter(w)
}
