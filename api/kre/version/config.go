package version

import (
	"context"
	"fmt"

	"github.com/machinebox/graphql"
)

type Config struct {
	Completed bool
	Vars      []*ConfigVariable
}

type ConfigVariable struct {
	Key   string
	Value string
	Type  ConfigVariableType
}

type ConfigVariableInput struct {
	Key   string
	Value string
}

type ConfigVariableType string

const (
	ConfigVariableTypeVariable ConfigVariableType = "VARIABLE"
	ConfigVariableTypeFile     ConfigVariableType = "FILE"
)

func (e ConfigVariableType) IsValid() bool {
	switch e {
	case ConfigVariableTypeVariable, ConfigVariableTypeFile:
		return true
	}
	return false
}

func (e ConfigVariableType) String() string {
	return string(e)
}

func (c *versionClient) UpdateConfig(versionID string, configVars []ConfigVariableInput) (bool, error) {
	req := graphql.NewRequest(`
	mutation StartVersion($input: UpdateConfigurationInput!) {
		updateVersionConfiguration(input: $input) {
			config {
			  completed
      }
		}
	}
`)
	req.Var("input", map[string]interface{}{"versionId": versionID, "configurationVariables": configVars})

	var respData struct {
		config struct {
			completed bool
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), c.cfg.DefaultRequestTimeout)
	defer cancel()

	err := c.gql.Run(ctx, req, &respData)
	if err != nil {
		return false, fmt.Errorf("error calling GraphQL: %s", err) //nolint:goerr113
	}

	return respData.config.completed, nil
}

func (c *versionClient) GetConfig(versionID string) (*Config, error) {
	req := graphql.NewRequest(`
		query GetVersionConf($versionId: ID!) {
			version(id: $versionId) {
				config {
					completed
					vars{
            key
            value
            type
          }
				}
      }
		}
`)
	req.Var("versionId", versionID)

	var respData struct {
		Version struct {
			Config Config
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), c.cfg.DefaultRequestTimeout)
	defer cancel()

	err := c.gql.Run(ctx, req, &respData)
	if err != nil {
		return nil, fmt.Errorf("error calling GraphQL: %c", err) //nolint:goerr113
	}

	return &respData.Version.Config, nil
}
