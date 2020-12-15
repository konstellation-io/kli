package version

//go:generate mockgen -source=${GOFILE} -destination=$PWD/mocks/${GOFILE} -package=mocks

import (
	"github.com/machinebox/graphql"

	"github.com/konstellation-io/kli/internal/config"
)

//nolint: golint
type VersionInterface interface {
	List(string) (List, error)
}

type Client struct {
	cfg *config.Config
	gql *graphql.Client
}

func New(cfg *config.Config, gql *graphql.Client) VersionInterface {
	return &Client{
		cfg,
		gql,
	}
}