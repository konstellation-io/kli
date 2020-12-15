package api

import (
	"context"
	"fmt"

	"github.com/machinebox/graphql"
)

func (s *ServerClient) StartVersion(versionID, comment string) error {
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

	ctx, cancel := context.WithTimeout(context.Background(), s.cfg.DefaultRequestTimeout)
	defer cancel()

	err := s.client.Run(ctx, req, &respData)
	if err != nil {
		return fmt.Errorf("error calling GraphQL: %s", err) //nolint:goerr113
	}

	return nil
}
