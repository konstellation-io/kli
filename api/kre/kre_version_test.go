package kre_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/konstellation-io/kli/api/kre/version"
)

func TestVersionList(t *testing.T) {
	srv, cfg, client := gqlMockServer(t, "", `
		{
				"data": {
						"versions": [
								{
										"id": "123456",
										"name": "test-v1",
										"status": "STOPPED"
								}
						]
				}
		}
	`)
	defer srv.Close()

	c := version.New(cfg, client)

	list, err := c.List()
	require.NoError(t, err)
	require.Len(t, list, 1)
	require.Equal(t, list[0], version.Version{
		ID:     "123456",
		Name:   "test-v1",
		Status: "STOPPED",
	})
}

func TestVersionStart(t *testing.T) {
	versionID := "123456"
	comment := "test start comment"
	expectedVariables := fmt.Sprintf(`
		{
			"input": {
					"versionId": "%s",
					"comment": "%s"
			}
		}
 `, versionID, comment)

	srv, cfg, client := gqlMockServer(t, expectedVariables, "")
	defer srv.Close()

	c := version.New(cfg, client)

	err := c.Start(versionID, comment)
	require.NoError(t, err)
}

func TestVersionStop(t *testing.T) {
	versionID := "123456"
	comment := "test stop comment"
	expectedVariables := fmt.Sprintf(`
		{
			"input": {
					"versionId": "%s",
					"comment": "%s"
			}
		}
 `, versionID, comment)

	srv, cfg, client := gqlMockServer(t, expectedVariables, "")
	defer srv.Close()

	c := version.New(cfg, client)

	err := c.Stop(versionID, comment)
	require.NoError(t, err)
}

func TestVersionPublish(t *testing.T) {
	versionID := "123456"
	comment := "test publish comment"
	expectedVariables := fmt.Sprintf(`
		{
			"input": {
					"versionId": "%s",
					"comment": "%s"
			}
		}
 `, versionID, comment)

	srv, cfg, client := gqlMockServer(t, expectedVariables, "")
	defer srv.Close()

	c := version.New(cfg, client)

	err := c.Publish(versionID, comment)
	require.NoError(t, err)
}

func TestVersionUnpublish(t *testing.T) {
	versionID := "123456"
	comment := "test unpublish comment"
	expectedVariables := fmt.Sprintf(`
		{
			"input": {
					"versionId": "%s",
					"comment": "%s"
			}
		}
 `, versionID, comment)

	srv, cfg, client := gqlMockServer(t, expectedVariables, "")
	defer srv.Close()

	c := version.New(cfg, client)

	err := c.Unpublish(versionID, comment)
	require.NoError(t, err)
}

func TestVersionGetConfig(t *testing.T) {
	srv, cfg, client := gqlMockServer(t, "", `
		{
			"data": {
				"version": {
					"config": {
						"completed": false,
						"vars": [
							{
								"key": "KEY1",
								"value": "value",
								"type": "VARIABLE"
							},
							{
								"key": "KEY2",
								"value": "",
								"type": "VARIABLE"
							}
						]
					}
				}
			}
		}
	`)
	defer srv.Close()

	c := version.New(cfg, client)

	config, err := c.GetConfig("test-v1")
	require.NoError(t, err)
	require.False(t, config.Completed)
	require.Len(t, config.Vars, 2)
	require.EqualValues(t, config, &version.Config{
		Completed: false,
		Vars: []*version.ConfigVariable{
			{
				Key:   "KEY1",
				Value: "value",
				Type:  version.ConfigVariableTypeVariable,
			},
			{
				Key:   "KEY2",
				Value: "",
				Type:  version.ConfigVariableTypeVariable,
			},
		},
	})
}

func TestVersionUpdateConfig(t *testing.T) {
	configVars := []version.ConfigVariableInput{
		{
			"key":   "KEY2",
			"value": "newValue",
		},
	}
	requestVars := `
		{
			"input": {
					"configurationVariables": [
						{ "key": "KEY2", "value": "newValue" }
					],
					"versionId": "test-v1"
			}
		}
	`
	srv, cfg, client := gqlMockServer(t, requestVars, `
		{
			"data": {
				"updateVersionConfiguration": {
					"config": {
						"completed": true
					}
				}
			}
		}
	`)

	defer srv.Close()

	c := version.New(cfg, client)

	completed, err := c.UpdateConfig("test-v1", configVars)
	require.NoError(t, err)
	require.True(t, completed)
}
