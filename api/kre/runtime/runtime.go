package runtime

//go:generate mockgen -source=${GOFILE} -destination=$PWD/mocks/${GOFILE} -package=mocks

import (
	"github.com/machinebox/graphql"

	"github.com/konstellation-io/kli/internal/config"
)

//nolint: golint
type RuntimeInterface interface {
	List() (List, error)
}

type Client struct {
	cfg    *config.Config
	client *graphql.Client
}

func New(cfg *config.Config, client *graphql.Client) RuntimeInterface {
	return &Client{
		cfg,
		client,
	}
}
