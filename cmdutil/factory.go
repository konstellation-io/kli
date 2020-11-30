package cmdutil

//go:generate mockgen -source=${GOFILE} -destination=$PWD/mocks/${GOFILE} -package=mocks

import (
	"github.com/konstellation-io/kli/api"
	"github.com/konstellation-io/kli/internal/config"
	"github.com/konstellation-io/kli/internal/errors"
	"github.com/konstellation-io/kli/internal/logger"
	"github.com/konstellation-io/kli/pkg/iostreams"
)

type CmdFactory interface {
	IOStreams() *iostreams.IOStreams
	Config() *config.Config
	Logger() logger.Logger
	ServerClient(string) (api.ServerClienter, error)
}

type Factory struct {
	appVersion string
	ioStreams  *iostreams.IOStreams
	cfg        *config.Config
	logger     logger.Logger
}

func NewCmdFactory(appVersion string) *Factory {
	io := iostreams.System()

	return &Factory{
		appVersion: appVersion,
		ioStreams:  io,
		cfg:        config.NewConfig(),
		logger:     logger.NewLogger(io.Out),
	}
}

func (f *Factory) IOStreams() *iostreams.IOStreams {
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
