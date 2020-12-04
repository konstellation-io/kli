package version_test

import (
	"errors"
	"testing"

	"github.com/MakeNowJust/heredoc"
	"github.com/golang/mock/gomock"
	"github.com/konstellation-io/kli/api"
	"github.com/konstellation-io/kli/internal/config"
	"github.com/konstellation-io/kli/internal/testhelpers"
	"github.com/konstellation-io/kli/mocks"
	"github.com/konstellation-io/kli/pkg/cmd/kre/version"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

func TestVersionListCmd(t *testing.T) {
	r := testhelpers.NewRunner(t, func(f *mocks.MockCmdFactory) *cobra.Command {
		ctrl := gomock.NewController(t)
		cfg := f.Config()

		err := cfg.AddServer(config.ServerConfig{
			Name:     "test",
			URL:      "http://test.local",
			APIToken: "12346",
		})
		require.NoError(t, err)
		err = cfg.SetDefaultServer("test")
		require.NoError(t, err)

		c := mocks.NewMockServerClienter(ctrl)
		f.EXPECT().ServerClient("test").Return(c, nil)
		c.EXPECT().ListVersions("runtime1234").Return(api.VersionList{
			{ID: "1234", Name: "greeter-v1", Status: "STARTED"},
			{ID: "6578", Name: "greeter-v2", Status: "STOPPED"},
		}, nil)

		return version.NewVersionCmd(f)
	})

	r.Run("version ls --runtime runtime1234").
		Contains(heredoc.Doc(`
        NAME       STATUS
			1 greeter-v1 STARTED
			2 greeter-v2 STOPPED
		`))
}

var errNoRuntime = errors.New("required flag(s) \"runtime\" not set")

func TestVersionListNoRuntime(t *testing.T) {
	r := testhelpers.NewRunner(t, func(f *mocks.MockCmdFactory) *cobra.Command {
		ctrl := gomock.NewController(t)
		cfg := f.Config()

		err := cfg.AddServer(config.ServerConfig{
			Name:     "test",
			URL:      "http://test.local",
			APIToken: "12346",
		})
		require.NoError(t, err)
		err = cfg.SetDefaultServer("test")
		require.NoError(t, err)

		c := mocks.NewMockServerClienter(ctrl)
		f.EXPECT().ServerClient("test").Return(c, nil)

		return version.NewVersionCmd(f)
	})

	r.RunE("version ls", errNoRuntime)
}
