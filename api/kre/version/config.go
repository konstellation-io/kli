package version

// Config struct to represent a Version config.
type Config struct {
	Completed bool
	Vars      []*ConfigVariable
}

// ConfigVariable struct to represent a Version config variables.
type ConfigVariable struct {
	Key   string
	Value string
	Type  ConfigVariableType
}

// ConfigVariableInput struct to collect key/values to update Version config.
type ConfigVariableInput struct {
	Key   string
	Value string
}

// ConfigVariableType enum to represent variable types.
type ConfigVariableType string

const (
	// ConfigVariableTypeVariable type VARIABLE.
	ConfigVariableTypeVariable ConfigVariableType = "VARIABLE"
	// ConfigVariableTypeFile type FILE.
	ConfigVariableTypeFile ConfigVariableType = "FILE"
)

// IsValid method to implement enums.
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

// UpdateConfig update a Version config values.
func (c *versionClient) UpdateConfig(versionID string, configVars []ConfigVariableInput) (bool, error) {
	query := `
		mutation StartVersion($input: UpdateConfigurationInput!) {
			updateVersionConfiguration(input: $input) {
				config {
					completed
				}
			}
		}
	`
	vars := map[string]interface{}{
		"input": map[string]interface{}{
			"versionId":              versionID,
			"configurationVariables": configVars,
		},
	}

	var respData struct {
		config struct {
			completed bool
		}
	}

	err := c.client.MakeRequest(query, vars, &respData)

	return respData.config.completed, err
}

func (c *versionClient) GetConfig(versionID string) (*Config, error) {
	query := `
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
  `
	vars := map[string]interface{}{
		"versionId": versionID,
	}

	var respData struct {
		Version struct {
			Config Config
		}
	}

	err := c.client.MakeRequest(query, vars, &respData)

	return &respData.Version.Config, err
}
