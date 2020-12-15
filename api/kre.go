package api

import (
	gql "github.com/machinebox/graphql"

	"github.com/konstellation-io/kli/api/graphql"
	"github.com/konstellation-io/kli/api/kre/runtime"
	"github.com/konstellation-io/kli/api/kre/version"
	"github.com/konstellation-io/kli/internal/config"
)

type KRE struct {
	cfg     *config.Config
	client  *gql.Client
	runtime runtime.RuntimeInterface
	version version.VersionInterface
}

func (a *KRE) Runtime() runtime.RuntimeInterface {
	return a.runtime
}

func (a *KRE) Version() version.VersionInterface {
	return a.version
}

func NewKreClient(cfg *config.Config, server *config.ServerConfig, appVersion string) (*KRE, error) {
	g, err := graphql.NewGqlClient(cfg, server, appVersion)
	if err != nil {
		return nil, err
	}

	return &KRE{
		cfg,
		g,
		runtime.New(cfg, g),
		version.New(cfg, g),
	}, nil
}
