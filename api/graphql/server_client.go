package graphql

import (
	"fmt"

	"github.com/machinebox/graphql"

	"github.com/konstellation-io/kli/internal/config"
)

// NewGqlClient initializes a KreClient for a specific server.
func NewGqlClient(cfg *config.Config, server *config.ServerConfig, appVersion string) (*graphql.Client, error) {
	accessToken, err := getAccessToken(cfg, server)
	if err != nil {
		return nil, err
	}

	c := newHTTPClient([]Option{
		addHeader("User-Agent", "Konstellation KLI"),
		addHeader("KLI-Version", appVersion),
		addHeader("Cache-Control", "no-cache"),
		addHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)),
	}...)

	gql := graphql.NewClient(fmt.Sprintf("%s/graphql", server.URL), graphql.WithHTTPClient(c))

	return gql, nil
}
