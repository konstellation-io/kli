package factory

import (
	"github.com/konstellation-io/kli/api"
	"github.com/konstellation-io/kli/cmdutil"
	"github.com/konstellation-io/kli/internal/config"
	"github.com/konstellation-io/kli/internal/errors"
	"github.com/konstellation-io/kli/pkg/iostreams"
)

func New(appVersion string) *cmdutil.Factory {
	io := iostreams.System()
	cfg := config.NewConfig()

	return &cmdutil.Factory{
		IOStreams: io,
		Config:    config.NewConfig,
		ServerClient: func(serverName string) (*api.ServerClient, error) {
			server := cfg.GetByServerName(serverName)
			if server == nil {
				return nil, errors.ErrUnknownServerName
			}

			c, err := api.NewServerClient(*server, appVersion)
			if err != nil {
				return nil, err
			}

			return c, nil
		},
	}
}
