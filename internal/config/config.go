package config

import (
	"path"
	"sync"
	"time"

	"github.com/OpenPeeDeeP/xdg"
	"github.com/guumaster/cligger"

	"github.com/konstellation-io/kli/internal/errors"
	"github.com/konstellation-io/kli/text"
)

//nolint:gochecknoglobals
var (
	once sync.Once
	cfg  *Config
)

const (
	DefaultRequestTimeout = 2 * time.Minute
)

// Config holds the configuration values for the application.
type Config struct {
	filename              string
	DefaultRequestTimeout time.Duration
	DefaultServer         string         `yaml:"defaultServer"`
	ServerList            []ServerConfig `yaml:"servers"`
}

type ServerConfig struct {
	Name     string `yaml:"name"`
	URL      string `yaml:"url"`
	APIToken string `yaml:"token"`
}

// NewConfig will read the config.yml file from current user home.
func NewConfig() *Config {
	once.Do(createConfig)

	return cfg
}

func NewConfigTest() *Config {
	createConfig()
	return cfg
}

func createConfig() {
	d := xdg.New("konstellation-io", "kli")

	cfg = &Config{
		filename: path.Join(d.ConfigHome(), "config.yml"),
	}

	err := cfg.readFile()
	if err != nil {
		cligger.Fatal("error reading config: %s", err)
	}

	// Add default values
	cfg.DefaultRequestTimeout = DefaultRequestTimeout
}

func (c *Config) GetByServerName(name string) *ServerConfig {
	n := text.Normalize(name)
	for _, s := range c.ServerList {
		if text.Normalize(s.Name) == n {
			return &s
		}
	}

	return nil
}

func (c *Config) AddServer(server ServerConfig) error {
	exists := c.GetByServerName(server.Name)
	if exists != nil {
		return errors.ErrServerAlreadyExists
	}

	c.ServerList = append(c.ServerList, server)

	return c.Save()
}

func (c *Config) SetDefaultServer(name string) error {
	n := text.Normalize(name)

	srv := c.GetByServerName(n)
	if srv == nil {
		return errors.ErrUnknownServerName
	}

	c.DefaultServer = n

	return c.Save()
}
