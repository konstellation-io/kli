package config

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/MakeNowJust/heredoc"
	"github.com/stretchr/testify/require"

	"github.com/konstellation-io/kli/internal/render"
	"github.com/konstellation-io/kli/text"
)

func TestConfig_RenderServerList(t *testing.T) {
	dir := setupConfigDir(t)
	defer cleanConfigDir(t, dir)

	cfg := NewConfigTest()

	err := cfg.AddServer(ServerConfig{
		Name:     "test",
		URL:      "http://test.local",
		APIToken: "12345",
	})
	require.NoError(t, err)

	b := bytes.NewBufferString("")
	r := render.DefaultRenderer(b)

	cfg.RenderServerList(r)

	out, err := ioutil.ReadAll(b)
	require.NoError(t, err)

	expected := text.LinesTrim(heredoc.Doc(`
		SERVER	URL
		test  	http://test.local
	`))
	actual := text.LinesTrim(string(out))
	require.Equal(t, expected, actual)
}
