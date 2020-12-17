package version

//go:generate mockgen -source=${GOFILE} -destination=$PWD/mocks/${GOFILE} -package=mocks

import (
	"github.com/konstellation-io/kli/api/graphql"
	"github.com/konstellation-io/kli/internal/config"
)

// VersionInterface method to interact with Versions.
type VersionInterface interface { // nolint: golint
	List(runtimeID string) (List, error)
	Start(versionID, comment string) error
	Stop(versionID, comment string) error
	Publish(versionID, comment string) error
	Unpublish(versionID, comment string) error
	GetConfig(versionID string) (*Config, error)
	UpdateConfig(versionID string, configVars []ConfigVariableInput) (bool, error)
}

type versionClient struct {
	cfg    *config.Config
	client *graphql.GqlManager
}

// New creates a new struct to access Versions methods.
func New(cfg *config.Config, gql *graphql.GqlManager) VersionInterface {
	return &versionClient{
		cfg,
		gql,
	}
}
