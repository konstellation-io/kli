package api

import (
	"context"

	"github.com/machinebox/graphql"
)

type Runtime struct {
	ID     string
	Name   string
	Status string
}

type RuntimeList []Runtime

func (s *ServerClient) ListRuntimes() (RuntimeList, error) {
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
		Runtimes RuntimeList
	}

	err := s.client.Run(context.Background(), req, &respData)
	if err != nil {
		return nil, err
	}

	return respData.Runtimes, nil
}
