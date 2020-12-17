package version_test

import (
	"fmt"
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

type testVersionSuite struct {
	ctrl  *gomock.Controller
	mocks versionSuiteMocks
}

type versionSuiteMocks struct {
	kreClient *mocks.MockKreInterface
	version   *mocks.MockVersionInterface
}

func newTestVersionSuite(t *testing.T) *testVersionSuite {
	ctrl := gomock.NewController(t)

	return &testVersionSuite{
		ctrl,
		versionSuiteMocks{
			kreClient: mocks.NewMockKreInterface(ctrl),
			version:   mocks.NewMockVersionInterface(ctrl),
		},
	}
}

func setupVersionConfig(t *testing.T, f *mocks.MockCmdFactory) {
	cfg := f.Config()

	err := cfg.AddServer(config.ServerConfig{
		Name:     "test",
		URL:      "http://test.local",
		APIToken: "12346",
	})
	require.NoError(t, err)
	err = cfg.SetDefaultServer("test")
	require.NoError(t, err)
}

func TestVersionListCmd(t *testing.T) {
	s := newTestVersionSuite(t)
	r := testhelpers.NewRunner(t, func(f *mocks.MockCmdFactory) *cobra.Command {
		setupVersionConfig(t, f)

		f.EXPECT().KreClient("test").Return(s.mocks.kreClient, nil)
		s.mocks.kreClient.EXPECT().Version().Return(s.mocks.version)
		s.mocks.version.EXPECT().List("runtime1234").Return(version.List{
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
	s := newTestVersionSuite(t)
	r := testhelpers.NewRunner(t, func(f *mocks.MockCmdFactory) *cobra.Command {
		setupVersionConfig(t, f)

		f.EXPECT().KreClient("test").Return(s.mocks.kreClient, nil)
		s.mocks.kreClient.EXPECT().Version().Return(s.mocks.version)
		s.mocks.version.EXPECT().Start("12345", comment).Return(nil)

		return cmd.NewVersionCmd(f)
	})

	r.Runf("version start 12345 --message \"%s\"", comment).
		Contains(heredoc.Doc(`
      [✔] Starting version '12345'.
		`))
}

func TestVersionStopCmd(t *testing.T) {
	comment := "test stop comment"
	s := newTestVersionSuite(t)
	r := testhelpers.NewRunner(t, func(f *mocks.MockCmdFactory) *cobra.Command {
		setupVersionConfig(t, f)

		f.EXPECT().KreClient("test").Return(s.mocks.kreClient, nil)
		s.mocks.kreClient.EXPECT().Version().Return(s.mocks.version)
		s.mocks.version.EXPECT().Stop("12345", comment).Return(nil)

		return cmd.NewVersionCmd(f)
	})

	r.Runf("version stop 12345 --message \"%s\"", comment).
		Contains(heredoc.Doc(`
      [✔] Stopping version '12345'.
		`))
}

func TestVersionPublishCmd(t *testing.T) {
	comment := "test publish comment"
	s := newTestVersionSuite(t)
	r := testhelpers.NewRunner(t, func(f *mocks.MockCmdFactory) *cobra.Command {
		setupVersionConfig(t, f)

		f.EXPECT().KreClient("test").Return(s.mocks.kreClient, nil)
		s.mocks.kreClient.EXPECT().Version().Return(s.mocks.version)
		s.mocks.version.EXPECT().Publish("12345", comment).Return(nil)

		return cmd.NewVersionCmd(f)
	})

	r.Runf("version publish 12345 --message \"%s\"", comment).
		Contains(heredoc.Doc(`
      [✔] Publishing version '12345'.
		`))
}

func TestVersionUnpublishCmd(t *testing.T) {
	comment := "test unpublish comment"
	s := newTestVersionSuite(t)
	r := testhelpers.NewRunner(t, func(f *mocks.MockCmdFactory) *cobra.Command {
		setupVersionConfig(t, f)

		f.EXPECT().KreClient("test").Return(s.mocks.kreClient, nil)
		s.mocks.kreClient.EXPECT().Version().Return(s.mocks.version)
		s.mocks.version.EXPECT().Unpublish("12345", comment).Return(nil)

		return cmd.NewVersionCmd(f)
	})

	r.Runf("version unpublish 12345 --message \"%s\"", comment).
		Contains(heredoc.Doc(`
      [✔] Unpublishing version '12345'.
		`))
}

func TestVersionGetConfigCmd(t *testing.T) {
	s := newTestVersionSuite(t)
	versionConfig := &version.Config{
		Completed: true,
		Vars: []*version.ConfigVariable{
			{
				Key:   "key1",
				Value: "value1",
				Type:  version.ConfigVariableTypeVariable,
			},
			{
				Key:   "key2",
				Value: "value2",
				Type:  version.ConfigVariableTypeVariable,
			},
		},
	}
	r := testhelpers.NewRunner(t, func(f *mocks.MockCmdFactory) *cobra.Command {
		setupVersionConfig(t, f)

		f.EXPECT().KreClient("test").Return(s.mocks.kreClient, nil)
		s.mocks.kreClient.EXPECT().Version().Return(s.mocks.version)
		s.mocks.version.EXPECT().GetConfig("12345").Return(versionConfig, nil)

		return cmd.NewVersionCmd(f)
	})

	r.Run("version config 12345").
		Contains(heredoc.Doc(`
				TYPE     KEY  VALUE
			1 VARIABLE key1 value1
			2 VARIABLE key2 value2

      [✔] Version config complete
		`))
}

func TestVersionSetConfigCmd(t *testing.T) {
	s := newTestVersionSuite(t)
	configVars := []version.ConfigVariableInput{
		{"key": "key1", "value": "value1"},
		{"key": "key2", "value": "value2"},
	}

	r := testhelpers.NewRunner(t, func(f *mocks.MockCmdFactory) *cobra.Command {
		setupVersionConfig(t, f)

		f.EXPECT().KreClient("test").Return(s.mocks.kreClient, nil)
		s.mocks.kreClient.EXPECT().Version().Return(s.mocks.version)
		s.mocks.version.EXPECT().UpdateConfig("12345", configVars).Return(true, nil)

		return cmd.NewVersionCmd(f)
	})

	pair1 := fmt.Sprintf("%s=%s", configVars[0]["key"], configVars[0]["value"])
	pair2 := fmt.Sprintf("%s=%s", configVars[1]["key"], configVars[1]["value"])
	r.Runf("version config 12345 --set %s --set %s", pair1, pair2).
		Contains(heredoc.Doc(`
      [✔] Config completed for version '12345'.
		`))
}
