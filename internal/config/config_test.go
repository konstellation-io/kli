package config

import (
	"fmt"
	"path"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/konstellation-io/kli/text"
)

func TestNewConfig(t *testing.T) {
	dir := setupConfigDir(t)
	defer cleanConfigDir(t, dir)

	cfg := NewConfig()

	// NOTE: This formatting is for testing on windows platform
	f := fmt.Sprintf("%s%s%s", filepath.FromSlash(path.Join(dir, "konstellation-io", "kli")), string(filepath.Separator), "config.yml")
	expected := &Config{
		filename:              f,
		DefaultRequestTimeout: DefaultRequestTimeout,
		DefaultServer:         "",
		ServerList:            []ServerConfig{},
	}

	require.Equal(t, cfg, expected)
}

func TestConfig_AddServer(t *testing.T) {
	dir := setupConfigDir(t)
	defer cleanConfigDir(t, dir)

	cfg := NewConfig()

	newServer := ServerConfig{
		Name:     "local",
		URL:      "http://test.local",
		APIToken: "12345",
	}

	err := cfg.AddServer(newServer)
	require.NoError(t, err)
	require.Equal(t, cfg.ServerList, []ServerConfig{
		newServer,
	})
}

func TestConfig_GetByServerName(t *testing.T) {
	dir := setupConfigDir(t)
	defer cleanConfigDir(t, dir)

	cfg := NewConfig()

	newServer := &ServerConfig{
		Name:     "getby",
		URL:      "http://test.local",
		APIToken: "12345",
	}

	err := cfg.AddServer(*newServer)
	require.NoError(t, err)
	require.Equal(t, cfg.GetByServerName("getby"), newServer)
}

func TestConfig_SetDefaultServer(t *testing.T) {
	dir := setupConfigDir(t)
	defer cleanConfigDir(t, dir)

	cfg := NewConfigTest()

	newServer := ServerConfig{
		Name:     "Default SERVER",
		URL:      "http://test.local",
		APIToken: "12345",
	}

	err := cfg.AddServer(newServer)
	require.NoError(t, err)

	err = cfg.SetDefaultServer(newServer.Name)
	require.NoError(t, err)

	require.Equal(t, cfg.DefaultServer, text.Normalize(newServer.Name))
}
