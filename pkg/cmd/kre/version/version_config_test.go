package version

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/MakeNowJust/heredoc"
	"github.com/golang/mock/gomock"
	"github.com/guumaster/logsymbols"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"

	"github.com/konstellation-io/kli/api/kre/version"

	"github.com/konstellation-io/kli/internal/config"
	"github.com/konstellation-io/kli/internal/testhelpers"
	"github.com/konstellation-io/kli/mocks"
)

type testVersionConfigSuite struct {
	ctrl  *gomock.Controller
	mocks versionSuiteMocks
}

type versionSuiteMocks struct {
	kreClient *mocks.MockKreInterface
	version   *mocks.MockVersionInterface
}

func newTestVersionConfigSuite(t *testing.T) *testVersionConfigSuite {
	ctrl := gomock.NewController(t)

	return &testVersionConfigSuite{
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

func TestVersionGetConfigCmd(t *testing.T) {
	s := newTestVersionConfigSuite(t)
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

		return NewVersionCmd(f)
	})

	r.Run("version config 12345").
		Containsf(heredoc.Doc(`
				TYPE     KEY
			1 VARIABLE key1
			2 VARIABLE key2

      [%s] Version config complete
		`), logsymbols.CurrentSymbols().Success)
}

func TestVersionGetConfigWithValuesCmd(t *testing.T) {
	s := newTestVersionConfigSuite(t)
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

		return NewVersionCmd(f)
	})

	r.Run("version config 12345 --show-values").
		Containsf(heredoc.Doc(`
				TYPE     KEY  VALUE
			1 VARIABLE key1 value1
			2 VARIABLE key2 value2

      [%s] Version config complete
		`), logsymbols.CurrentSymbols().Success)
}

func TestVersionSetConfigCmd(t *testing.T) {
	s := newTestVersionConfigSuite(t)
	configVars := []version.ConfigVariableInput{
		{"key": "key1", "value": "value1"},
		{"key": "key2", "value": "value2"},
	}

	r := testhelpers.NewRunner(t, func(f *mocks.MockCmdFactory) *cobra.Command {
		setupVersionConfig(t, f)

		f.EXPECT().KreClient("test").Return(s.mocks.kreClient, nil)
		s.mocks.kreClient.EXPECT().Version().Return(s.mocks.version)
		s.mocks.version.EXPECT().UpdateConfig("12345", configVars).Return(true, nil)

		return NewVersionCmd(f)
	})

	pair1 := fmt.Sprintf("%s=%s", configVars[0]["key"], configVars[0]["value"])
	pair2 := fmt.Sprintf("%s=%s", configVars[1]["key"], configVars[1]["value"])
	r.Runf("version config 12345 --set %s --set %s", pair1, pair2).
		Containsf(heredoc.Doc(`
      [%s] Config completed for version '12345'.
		`), logsymbols.CurrentSymbols().Success)
}

func TestVersionSetConfigErrorEdgeCasesCmd(t *testing.T) {
	s := newTestVersionConfigSuite(t)

	r := testhelpers.NewRunner(t, func(f *mocks.MockCmdFactory) *cobra.Command {
		setupVersionConfig(t, f)

		f.EXPECT().KreClient("test").Return(s.mocks.kreClient, nil)

		return NewVersionCmd(f)
	})

	type testCase struct {
		extraArgs     []string
		expectedError string
	}

	testCases := []testCase{
		{
			[]string{"--set", fmt.Sprintf("%s=%s", "key", "\"test")},
			`invalid argument "key=\"test" for "--set" flag: parse error on line 1, column 4: bare " in non-quoted-field`,
		},
		{
			[]string{"--set", fmt.Sprintf("%s=%s", "key", "\"test test\"")},
			`invalid argument "key=\"test test\"" for "--set" flag: parse error on line 1, column 4: bare " in non-quoted-field`,
		},
		{
			[]string{"--set", fmt.Sprintf("\"%s\"=\"%s\"", "key", `test`)},
			`invalid argument "\"key\"=\"test\"" for "--set" flag: parse error on line 1, column 4: extraneous or missing " in quoted-field`,
		},
	}

	for _, pair := range testCases {
		err := fmt.Errorf(pair.expectedError)
		r.RunArgsE("version config 12345", err, pair.extraArgs...)
	}
}

func TestVersionSetConfigEdgeCasesCmd(t *testing.T) {
	s := newTestVersionConfigSuite(t)

	r := testhelpers.NewRunner(t, func(f *mocks.MockCmdFactory) *cobra.Command {
		setupVersionConfig(t, f)

		f.EXPECT().KreClient("test").Return(s.mocks.kreClient, nil).AnyTimes()
		s.mocks.kreClient.EXPECT().Version().Return(s.mocks.version).AnyTimes()
		s.mocks.version.EXPECT().UpdateConfig("12345", gomock.Any()).Return(true, nil).AnyTimes()

		return NewVersionCmd(f)
	})

	testCases := [][]string{
		{"--set", fmt.Sprintf("%s=%s", "key", "test test")},
		{"--set", fmt.Sprintf("%s=%s", "key", "test\n test")},
		{"--set", fmt.Sprintf("%s=%s", "key", `test`)},
		{"--set", fmt.Sprintf("%s=%s", "key", "test=")},
		// This test should work, but it doesn't.
		// {"--set", fmt.Sprintf("%s=%s", "key", "test \"test")},
	}

	for _, extraArgs := range testCases {
		r.RunArgs("version config 12345", extraArgs...).
			Containsf(heredoc.Doc(`
      [%s] Config completed for version '12345'.
		`), logsymbols.CurrentSymbols().Success)
	}
}

func TestVersionSetConfigFromEnvCmd(t *testing.T) {
	s := newTestVersionConfigSuite(t)
	configVars := []version.ConfigVariableInput{
		{
			"key":   "TEST_VAR",
			"value": "12345\n222",
		},
		{
			"key":   "TEST_VAR2",
			"value": "test\" test",
		},
	}
	r := testhelpers.NewRunner(t, func(f *mocks.MockCmdFactory) *cobra.Command {
		setupVersionConfig(t, f)

		f.EXPECT().KreClient("test").Return(s.mocks.kreClient, nil)
		s.mocks.kreClient.EXPECT().Version().Return(s.mocks.version)
		s.mocks.version.EXPECT().UpdateConfig("12345", configVars).Return(true, nil)

		return NewVersionCmd(f)
	})

	err := os.Setenv("TEST_VAR", heredoc.Doc(`12345
                                                      222`))
	require.NoError(t, err)

	err = os.Setenv("TEST_VAR2", "test\" test")
	require.NoError(t, err)

	r.Run("version config 12345 --set-from-env TEST_VAR --set-from-env TEST_VAR2").
		Containsf(heredoc.Doc(`
      [%s] Config completed for version '12345'.
		`), logsymbols.CurrentSymbols().Success)
}

func TestVersionSetConfigFromEmptyEnvCmd(t *testing.T) {
	s := newTestVersionConfigSuite(t)
	configVars := []version.ConfigVariableInput{
		{
			"key":   "TEST_VAR",
			"value": "",
		},
	}
	r := testhelpers.NewRunner(t, func(f *mocks.MockCmdFactory) *cobra.Command {
		setupVersionConfig(t, f)

		f.EXPECT().KreClient("test").Return(s.mocks.kreClient, nil)
		s.mocks.kreClient.EXPECT().Version().Return(s.mocks.version)
		s.mocks.version.EXPECT().UpdateConfig("12345", configVars).Return(false, nil)

		return NewVersionCmd(f)
	})

	err := os.Unsetenv("TEST_VAR")
	require.NoError(t, err)

	r.Run("version config 12345 --set-from-env TEST_VAR").
		Containsf(heredoc.Doc(`
      [%s] Config updated for version '12345'.
		`), logsymbols.CurrentSymbols().Success)
}

func TestVersionSetConfigFromFileCmd(t *testing.T) {
	s := newTestVersionConfigSuite(t)
	configVars := []version.ConfigVariableInput{
		{
			"key":   "TEST_VAR",
			"value": "123456",
		},
	}
	r := testhelpers.NewRunner(t, func(f *mocks.MockCmdFactory) *cobra.Command {
		setupVersionConfig(t, f)

		f.EXPECT().KreClient("test").Return(s.mocks.kreClient, nil)
		s.mocks.kreClient.EXPECT().Version().Return(s.mocks.version)
		s.mocks.version.EXPECT().UpdateConfig("12345", configVars).Return(true, nil)

		return NewVersionCmd(f)
	})

	d := os.TempDir()
	tempEnvFile, err := ioutil.TempFile(d, "TestVersionSetConfigFromFileCmd")
	require.NoError(t, err)

	defer os.RemoveAll(tempEnvFile.Name())

	_, _ = tempEnvFile.Write([]byte(heredoc.Doc(`
		TEST_VAR=123456
	`)))

	r.Runf("version config 12345 --set-from-file %s", tempEnvFile.Name()).
		Containsf(heredoc.Doc(`
      [%s] Config completed for version '12345'.
		`), logsymbols.CurrentSymbols().Success)
}
