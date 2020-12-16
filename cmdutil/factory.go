package cmdutil

//go:generate mockgen -source=${GOFILE} -destination=$PWD/mocks/${GOFILE} -package=mocks

import (
	"github.com/konstellation-io/kli/api"
	"github.com/konstellation-io/kli/api/kre"
	"github.com/konstellation-io/kli/internal/config"
	"github.com/konstellation-io/kli/internal/errors"
	"github.com/konstellation-io/kli/internal/logger"
)

// CmdFactory interface to define all methods needed during commands creation.
type CmdFactory interface {
	IOStreams() *IOStreams
	Config() *config.Config
	Logger() logger.Logger
	KreClient(string) (kre.KreInterface, error)
}

// Factory contains all data needed during commands execution.
type Factory struct {
	appVersion string
	ioStreams  *IOStreams
	cfg        *config.Config
	logger     logger.Logger
}

// NewCmdFactory returns a new Factory object used during commands execution.
func NewCmdFactory(appVersion string) *Factory {
	io := System()

	return &Factory{
		appVersion: appVersion,
		ioStreams:  io,
		cfg:        config.NewConfig(),
		logger:     logger.NewLogger(io.Out),
	}
}

// IOStreams access to IOStreams object.
func (f *Factory) IOStreams() *IOStreams {
	return f.ioStreams
}

// Config access to Config object.
func (f *Factory) Config() *config.Config {
	return f.cfg
}

// Logger access to Logger object.
func (f *Factory) Logger() logger.Logger {
	return f.logger
}

// KreClient generates a new ServerClient specific for the given server name.
func (f *Factory) KreClient(serverName string) (kre.KreInterface, error) {
	server := f.cfg.GetByServerName(serverName)
	if server == nil {
		return nil, errors.ErrUnknownServerName
	}

	return api.NewKreClient(f.cfg, server, f.appVersion)
}
