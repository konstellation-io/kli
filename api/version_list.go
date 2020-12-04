package api

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

// VersionList contains a list of  Version.
type VersionList []Version

// ListVersions calls to KRE API and returns a list of Version entities.
func (s *ServerClient) ListVersions(runtimeID string) (VersionList, error) {
	req := graphql.NewRequest(`
	query GetVersionConfStatus($runtimeId: ID!) {
		versions(runtimeId: $runtimeId) {
			id
			name
			status
		}
	}
`)
	req.Var("runtimeId", runtimeID)

	var respData struct {
		Versions VersionList
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.cfg.DefaultRequestTimeout)
	defer cancel()

	err := s.client.Run(ctx, req, &respData)
	if err != nil {
		return nil, fmt.Errorf("error calling GraphQL: %s", err) //nolint:goerr113
	}

	return respData.Versions, nil
}

/*
query: query GetVersionConfStatus($runtimeId: ID!) {
  versions(runtimeId: $runtimeId) {
    id
    name
    status
  }
}

*/
