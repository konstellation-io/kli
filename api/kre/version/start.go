package version

import (
	"context"
	"fmt"

	"github.com/machinebox/graphql"
)

func (c *versionClient) Start(versionID, comment string) error {
	req := graphql.NewRequest(`
	mutation StartVersion($input: StartVersionInput!) {
		startVersion(input: $input) {
			id
			status
		}
	}
`)
	req.Var("input", map[string]string{"versionId": versionID, "comment": comment})

	var respData struct {
		Status string
	}

	ctx, cancel := context.WithTimeout(context.Background(), c.cfg.DefaultRequestTimeout)
	defer cancel()

	err := c.gql.Run(ctx, req, &respData)
	if err != nil {
		return fmt.Errorf("error calling GraphQL: %s", err) //nolint:goerr113
	}

	return nil
}
