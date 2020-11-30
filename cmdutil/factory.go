package cmdutil

//go:generate mockgen -source=${GOFILE} -destination=$PWD/mocks/${GOFILE} -package=mocks

import (
	"github.com/konstellation-io/kli/api"
	"github.com/konstellation-io/kli/internal/config"
	"github.com/konstellation-io/kli/internal/errors"
	"github.com/konstellation-io/kli/internal/logger"
)

type CmdFactory interface {
	IOStreams() *IOStreams
	Config() *config.Config
	Logger() logger.Logger
	ServerClient(string) (api.ServerClienter, error)
}

type Factory struct {
	appVersion string
	ioStreams  *IOStreams
	cfg        *config.Config
	logger     logger.Logger
}

func NewCmdFactory(appVersion string) *Factory {
	io := System()

	return &Factory{
		appVersion: appVersion,
		ioStreams:  io,
		cfg:        config.NewConfig(),
		logger:     logger.NewLogger(io.Out),
	}
}

func (f *Factory) IOStreams() *IOStreams {
	return f.ioStreams
}

func (f *Factory) Config() *config.Config {
	return f.cfg
}

func (f *Factory) Logger() logger.Logger {
	return f.logger
}

func (f *Factory) ServerClient(serverName string) (api.ServerClienter, error) {
	server := f.cfg.GetByServerName(serverName)
	if server == nil {
		return nil, errors.ErrUnknownServerName
	}

	c, err := api.NewServerClient(f.cfg, server, f.appVersion)
	if err != nil {
		return nil, err
	}

	return c, nil
}
