package version

import (
	"context"
	"fmt"

	"github.com/machinebox/graphql"
)

// Version represents a Version entity in KRE.
type Version struct {
	ID     string
	Name   string
	Status string
}

// List contains a list of  Version.
type List []Version

// ListVersions calls to KRE API and returns a list of Version entities.
func (c *Client) List(runtimeID string) (List, error) {
	req := graphql.NewRequest(`
	query GetVersions($runtimeId: ID!) {
		versions(runtimeId: $runtimeId) {
			id
			name
			status
		}
	}
`)
	req.Var("runtimeId", runtimeID)

	var respData struct {
		Versions List
	}

	ctx, cancel := context.WithTimeout(context.Background(), c.cfg.DefaultRequestTimeout)
	defer cancel()

	err := c.gql.Run(ctx, req, &respData)
	if err != nil {
		return nil, fmt.Errorf("error calling GraphQL: %c", err) //nolint:goerr113
	}

	return respData.Versions, nil
}
