package kre_test

import (
	"fmt"
	"io/ioutil"
	"os"
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
		Name:   "test-v1",
		Status: "STOPPED",
	})
}

func TestVersionStart(t *testing.T) {
	versionName := "123456"
	comment := "test start comment"
	expectedVariables := fmt.Sprintf(`
		{
			"input": {
					"versionName": "%s",
					"comment": "%s"
			}
		}
 `, versionName, comment)

	srv, cfg, client := gqlMockServer(t, expectedVariables, "")
	defer srv.Close()

	c := version.New(cfg, client)

	err := c.Start(versionName, comment)
	require.NoError(t, err)
}

func TestVersionStop(t *testing.T) {
	versionName := "123456"
	comment := "test stop comment"
	expectedVariables := fmt.Sprintf(`
		{
			"input": {
					"versionName": "%s",
					"comment": "%s"
			}
		}
 `, versionName, comment)

	srv, cfg, client := gqlMockServer(t, expectedVariables, "")
	defer srv.Close()

	c := version.New(cfg, client)

	err := c.Stop(versionName, comment)
	require.NoError(t, err)
}

func TestVersionPublish(t *testing.T) {
	versionName := "123456"
	comment := "test publish comment"
	expectedVariables := fmt.Sprintf(`
		{
			"input": {
					"versionName": "%s",
					"comment": "%s"
			}
		}
 `, versionName, comment)

	srv, cfg, client := gqlMockServer(t, expectedVariables, "")
	defer srv.Close()

	c := version.New(cfg, client)

	err := c.Publish(versionName, comment)
	require.NoError(t, err)
}

func TestVersionUnpublish(t *testing.T) {
	versionName := "123456"
	comment := "test unpublish comment"
	expectedVariables := fmt.Sprintf(`
		{
			"input": {
					"versionName": "%s",
					"comment": "%s"
			}
		}
 `, versionName, comment)

	srv, cfg, client := gqlMockServer(t, expectedVariables, "")
	defer srv.Close()

	c := version.New(cfg, client)

	err := c.Unpublish(versionName, comment)
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
					"versionName": "test-v1"
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

func TestVersionCreate(t *testing.T) {
	srv, cfg, client := gqlMockServer(t, "", `
		{
				"data": {
						"createVersion": {
								"name": "test-v1"
						}
				}
		}
	`)
	defer srv.Close()

	c := version.New(cfg, client)

	krtFile, err := ioutil.TempFile("", ".test.krt")
	require.NoError(t, err)

	defer os.RemoveAll(krtFile.Name())

	versionName, err := c.Create(krtFile.Name())
	require.NoError(t, err)
	require.Equal(t, versionName, "test-v1")
}
