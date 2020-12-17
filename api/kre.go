package api

import (
	"github.com/konstellation-io/kli/api/graphql"
	"github.com/konstellation-io/kli/api/kre/runtime"
	"github.com/konstellation-io/kli/api/kre/version"
	"github.com/konstellation-io/kli/internal/config"
)

// KRE object to implement access to KRE API.
type KRE struct {
	cfg        *config.Config
	gqlManager *graphql.GqlManager
	runtime    runtime.RuntimeInterface
	version    version.VersionInterface
}

// Runtime access to methods to interact with Runtimes.
func (a *KRE) Runtime() runtime.RuntimeInterface {
	return a.runtime
}

// Version access to methods to interact with Versions.
func (a *KRE) Version() version.VersionInterface {
	return a.version
}

// NewKreClient creates an API client instance.
func NewKreClient(cfg *config.Config, server *config.ServerConfig, appVersion string) (*KRE, error) {
	g, err := graphql.NewGqlManager(cfg, server, appVersion)
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
