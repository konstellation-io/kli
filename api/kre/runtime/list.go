package runtime

import (
	"context"
	"fmt"

	"github.com/machinebox/graphql"
)

// Runtime represents a Runtime entity in KRE.
type Runtime struct {
	ID     string
	Name   string
	Status string
}

// List contains a list of  Runtime.
type List []Runtime

// ListRuntimes calls to KRE API and returns a list of Runtime entities.
func (s *Client) List() (List, error) {
	req := graphql.NewRequest(`
	query GetRuntimes {
		runtimes {
			id
			name
			status
		}
	}
`)

	var respData struct {
		Runtimes List
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.cfg.DefaultRequestTimeout)
	defer cancel()

	err := s.client.Run(ctx, req, &respData)
	if err != nil {
		return nil, fmt.Errorf("error calling GraphQL: %s", err) //nolint:goerr113
	}

	return respData.Runtimes, nil
}
