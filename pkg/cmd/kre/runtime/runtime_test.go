package runtime_test

import (
	"testing"

	"github.com/MakeNowJust/heredoc"
	"github.com/golang/mock/gomock"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"

	"github.com/konstellation-io/kli/api"
	"github.com/konstellation-io/kli/internal/config"
	"github.com/konstellation-io/kli/internal/testhelpers"
	"github.com/konstellation-io/kli/mocks"
	"github.com/konstellation-io/kli/pkg/cmd/kre/runtime"
)

func TestRuntimeListCmd(t *testing.T) {
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
		c.EXPECT().ListRuntimes().Return(api.RuntimeList{
			{ID: "greetings", Name: "greetings", Status: ""},
			{ID: "int-tests", Name: "Integration Tests", Status: ""},
		}, nil)

		return runtime.NewRuntimeCmd(f)
	})

	r.Run("runtime ls").
		Contains(heredoc.Doc(`
			  ID        NAME
			1 greetings greetings
			2 int-tests Integration Tests
		`))
}
func TestRuntimeListEmptyCmd(t *testing.T) {
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
		c.EXPECT().ListRuntimes().Return(api.RuntimeList{}, nil)

		return runtime.NewRuntimeCmd(f)
	})

	r.Run("runtime ls").
		Contains(heredoc.Doc(`
			[ℹ] No runtimes found.
		`))
}
