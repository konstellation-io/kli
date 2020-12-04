package api

//go:generate mockgen -source=${GOFILE} -destination=$PWD/mocks/${GOFILE} -package=mocks

import (
	"fmt"

	"github.com/machinebox/graphql"

	"github.com/konstellation-io/kli/api/http"
	"github.com/konstellation-io/kli/internal/config"
)

// ServerClienter interface containing all available methods of the API.
type ServerClienter interface {
	ListRuntimes() (RuntimeList, error)
	ListVersions(runtimeID string) (VersionList, error)
}

// ServerClient facilitates making HTTP requests to a Konstellation API.
type ServerClient struct {
	cfg    *config.Config
	srv    *config.ServerConfig
	client *graphql.Client
}

// NewServerClient initializes a ServerClient for a specific server.
func NewServerClient(cfg *config.Config, server *config.ServerConfig, appVersion string) (ServerClienter, error) {
	accessToken, err := http.GetAccessToken(cfg, server)
	if err != nil {
		return nil, err
	}

	c := http.NewHTTPClient([]http.ClientOption{
		http.AddHeader("User-Agent", "Konstellation KLI"),
		http.AddHeader("KLI-Version", appVersion),
		http.AddHeader("Cache-Control", "no-cache"),
		http.AddHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)),
	}...)

	gql := graphql.NewClient(fmt.Sprintf("%s/graphql", server.URL), graphql.WithHTTPClient(c))

	return &ServerClient{
		cfg:    cfg,
		srv:    server,
		client: gql,
	}, nil
}
