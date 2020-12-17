package runtime

//go:generate mockgen -source=${GOFILE} -destination=$PWD/mocks/${GOFILE} -package=mocks

import (
	"github.com/konstellation-io/kli/api/graphql"
	"github.com/konstellation-io/kli/internal/config"
)

// RuntimeInterface method to interact with Runtimes.
type RuntimeInterface interface { // nolint: golint
	List() (List, error)
}

// runtimeClient struct to implement methods to interact with Runtimes.
type runtimeClient struct {
	cfg    *config.Config
	client *graphql.GqlManager
}

// New creates a new struct to access Runtimes methods.
func New(cfg *config.Config, client *graphql.GqlManager) RuntimeInterface {
	return &runtimeClient{
		cfg,
		client,
	}
}
