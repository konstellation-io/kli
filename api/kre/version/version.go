package version

//go:generate mockgen -source=${GOFILE} -destination=$PWD/mocks/${GOFILE} -package=mocks

import (
	"github.com/machinebox/graphql"

	"github.com/konstellation-io/kli/internal/config"
)

//nolint: golint
type VersionInterface interface {
	List(string) (List, error)
	Start(versionID, comment string) error
	Stop(versionID, comment string) error
	Publish(versionID, comment string) error
	Unpublish(versionID, comment string) error
	GetConfig(versionID string) (*Config, error)
	UpdateConfig(versionID string, configVars []ConfigVariableInput) (bool, error)
}

type versionClient struct {
	cfg *config.Config
	gql *graphql.Client
}

func New(cfg *config.Config, gql *graphql.Client) VersionInterface {
	return &versionClient{
		cfg,
		gql,
	}
}
