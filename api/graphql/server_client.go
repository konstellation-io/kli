package graphql

import (
	"context"
	"fmt"

	"github.com/machinebox/graphql"

	"github.com/konstellation-io/kli/internal/config"
)

// GqlManager struct to implement access to GraphQL endpoints with gql client.
type GqlManager struct {
	cfg    *config.Config
	client *graphql.Client
}

// MakeRequest call to GraphQL endpoint.
func (g *GqlManager) MakeRequest(query string, vars map[string]interface{}, respData interface{}) error {
	req := graphql.NewRequest(query)

	for k, v := range vars {
		req.Var(k, v)
	}

	ctx, cancel := context.WithTimeout(context.Background(), g.cfg.DefaultRequestTimeout)
	defer cancel()

	err := g.client.Run(ctx, req, respData)
	if err != nil {
		return fmt.Errorf("graphql error: %w", err)
	}

	return nil
}

// NewGqlManager creates an instance of GqlManager that takes cares of authentication.
func NewGqlManager(cfg *config.Config, server *config.ServerConfig, appVersion string) (*GqlManager, error) {
	accessToken, err := getAccessToken(cfg, server)
	if err != nil {
		return nil, err
	}

	c := newHTTPClient([]option{
		addHeader("User-Agent", "Konstellation KLI"),
		addHeader("KLI-Version", appVersion),
		addHeader("Cache-Control", "no-cache"),
		addHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)),
	}...)

	gql := graphql.NewClient(fmt.Sprintf("%s/graphql", server.URL), graphql.WithHTTPClient(c))

	return &GqlManager{
		cfg,
		gql,
	}, nil
}
