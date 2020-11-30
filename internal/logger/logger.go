package logger

//go:generate mockgen -source=${GOFILE} -destination=$PWD/mocks/${GOFILE} -package=mocks

import (
	"io"

	"github.com/guumaster/cligger"
)

type Logger interface {
	cligger.Logger
}

func NewLogger(w io.Writer) Logger {
	return cligger.NewLoggerWithWriter(w)
}
