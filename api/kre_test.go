package api_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/konstellation-io/kli/api"
	"github.com/konstellation-io/kli/internal/config"
	"github.com/konstellation-io/kli/internal/testhelpers"
)

func TestNewKreClient(t *testing.T) {
	d := testhelpers.SetupConfigDir(t)
	defer testhelpers.CleanConfigDir(t, d)

	cfg := config.NewConfig()

	srv := config.ServerConfig{
		Name:     "test",
		URL:      "http://test",
		APIToken: "12345",
	}
	err := cfg.AddServer(srv)
	require.NoError(t, err)

	k := api.NewKreClient(cfg, &srv, "test-version")

	require.NotEmpty(t, k.Version())
}
