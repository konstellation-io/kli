package kre_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/konstellation-io/kli/api/kre/runtime"
)

func TestRuntimeList(t *testing.T) {
	srv, cfg, client := gqlMockServer(t, "", `
			{
					"data": {
							"runtimes": [
									{
											"id": "runtime-1",
											"name": "test",
											"description": "test",
											"status": "STARTED",
											"creationDate": "2020-11-20T09:04:11Z",
											"publishedVersion": null
									}
							]
					}
			}
	`)
	defer srv.Close()

	runtimeCli := runtime.New(cfg, client)

	list, err := runtimeCli.List()
	require.NoError(t, err)
	require.Len(t, list, 1)
	require.Equal(t, list[0], runtime.Runtime{
		ID:     "runtime-1",
		Name:   "test",
		Status: "STARTED",
	})
}
