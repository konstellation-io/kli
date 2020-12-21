package factory

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/konstellation-io/kli/internal/config"
	"github.com/konstellation-io/kli/internal/testhelpers"
	"github.com/konstellation-io/kli/pkg/errors"
)

func TestNewCmdFactory(t *testing.T) {
	appV := "test-version"
	d := testhelpers.SetupConfigDir(t)

	defer testhelpers.CleanConfigDir(t, d)

	f := NewCmdFactory(appV)

	cfg := &config.Config{
		DefaultRequestTimeout: config.DefaultRequestTimeout,
		DefaultServer:         "local",
		ServerList:            []config.ServerConfig{},
	}

	require.Equal(t, appV, f.appVersion)
	require.Equal(t, cfg.DefaultRequestTimeout, f.cfg.DefaultRequestTimeout)
	require.Equal(t, cfg.ServerList, f.cfg.ServerList)
	require.Equal(t, f.IOStreams().Out, os.Stdout)
}

func TestFactory_KreClientError(t *testing.T) {
	d := testhelpers.SetupConfigDir(t)
	defer testhelpers.CleanConfigDir(t, d)

	f := NewCmdFactory("test-version")

	_, err := f.KreClient("BAD_NAME")
	require.EqualError(t, err, errors.ErrUnknownServerName.Error())
}

func TestFactory_KreClient(t *testing.T) {
	d := testhelpers.SetupConfigDir(t)
	defer testhelpers.CleanConfigDir(t, d)

	f := NewCmdFactory("test-version")
	cfg := f.Config()
	err := cfg.AddServer(config.ServerConfig{
		Name:     "test",
		URL:      "http://test",
		APIToken: "12345",
	})
	require.NoError(t, err)

	c, err := f.KreClient("test")
	require.NoError(t, err)

	require.NotEmpty(t, c)
}
