package cmdutil

import (
	"github.com/konstellation-io/kli/api"
	"github.com/konstellation-io/kli/internal/config"
	"github.com/konstellation-io/kli/pkg/iostreams"
)

type Factory struct {
	IOStreams    *iostreams.IOStreams
	Config       func() *config.Config
	ServerClient func(serverName string) (*api.ServerClient, error)
}
