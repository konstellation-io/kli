package graphql

import (
	"context"
	"fmt"
	"log"

	"github.com/machinebox/graphql"

	"github.com/konstellation-io/kli/internal/config"
)

// GqlManager struct to implement access to GraphQL endpoints with gql client.
type GqlManager struct {
	appVersion string
	cfg        *config.Config
	server     *config.ServerConfig
	client     *graphql.Client
}

func (g *GqlManager) setupClient() error {
	if g.client != nil {
		return nil
	}

	accessToken, err := getAccessToken(g.cfg, g.server)
	if err != nil {
		return err
	}

	c := newHTTPClient([]option{
		addHeader("User-Agent", "Konstellation KLI"),
		addHeader("KLI-Version", g.appVersion),
		addHeader("Cache-Control", "no-cache"),
		addHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)),
	}...)

	g.client = graphql.NewClient(fmt.Sprintf("%s/graphql", g.server.URL), graphql.WithHTTPClient(c))

	if g.cfg.Debug {
		g.client.Log = func(s string) { log.Println(s) }
	}

	return nil
}

// MakeRequest call to GraphQL endpoint.
func (g *GqlManager) MakeRequest(query string, vars map[string]interface{}, respData interface{}) error {
	err := g.setupClient()
	if err != nil {
		return err
	}

	req := graphql.NewRequest(query)

	for k, v := range vars {
		req.Var(k, v)
	}

	ctx, cancel := context.WithTimeout(context.Background(), g.cfg.DefaultRequestTimeout)
	defer cancel()

	err = g.client.Run(ctx, req, respData)
	if err != nil {
		return fmt.Errorf("graphql error: %w", err)
	}

	return nil
}

// NewGqlManager creates an instance of GqlManager that takes cares of authentication.
func NewGqlManager(cfg *config.Config, server *config.ServerConfig, appVersion string) *GqlManager {
	return &GqlManager{
		appVersion,
		cfg,
		server,
		nil,
	}
}
