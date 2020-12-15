package version_test

import (
	"testing"

	"github.com/MakeNowJust/heredoc"
	"github.com/golang/mock/gomock"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"

	"github.com/konstellation-io/kli/api/kre/version"
	"github.com/konstellation-io/kli/internal/config"
	"github.com/konstellation-io/kli/internal/testhelpers"
	"github.com/konstellation-io/kli/mocks"
	cmd "github.com/konstellation-io/kli/pkg/cmd/kre/version"
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

		c := mocks.NewMockKreInterface(ctrl)
		v := mocks.NewMockVersionInterface(ctrl)
		f.EXPECT().KreClient("test").Return(c, nil)
		c.EXPECT().Version().Return(v)
		v.EXPECT().List("runtime1234").Return(version.List{
			{ID: "1234", Name: "greeter-v1", Status: "STARTED"},
			{ID: "6578", Name: "greeter-v2", Status: "STOPPED"},
		}, nil)

		return cmd.NewVersionCmd(f)
	})

	r.Run("version ls --runtime runtime1234 ").
		Contains(heredoc.Doc(`
        ID   NAME       STATUS
			1 1234 greeter-v1 STARTED
			2 6578 greeter-v2 STOPPED
		`))
}

func TestVersionStartCmd(t *testing.T) {
	comment := "test start comment"
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

		c := mocks.NewMockKreInterface(ctrl)
		v := mocks.NewMockVersionInterface(ctrl)
		f.EXPECT().KreClient("test").Return(c, nil)
		c.EXPECT().Version().Return(v)
		v.EXPECT().Start("12345", comment).Return(nil)

		return cmd.NewVersionCmd(f)
	})

	r.Runf("version start 12345 --runtime runtime1234 --message \"%s\"", comment).
		Contains(heredoc.Doc(`
      [✔] Starting version '12345'.
		`))
}

func TestVersionStopCmd(t *testing.T) {
	comment := "test stop comment"
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

		c := mocks.NewMockKreInterface(ctrl)
		v := mocks.NewMockVersionInterface(ctrl)
		f.EXPECT().KreClient("test").Return(c, nil)
		c.EXPECT().Version().Return(v)
		v.EXPECT().Stop("12345", comment).Return(nil)

		return cmd.NewVersionCmd(f)
	})

	r.Runf("version stop 12345 --runtime runtime1234 --message \"%s\"", comment).
		Contains(heredoc.Doc(`
      [✔] Stopping version '12345'.
		`))
}
