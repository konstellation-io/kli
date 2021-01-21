package version

import (
	"testing"

	"github.com/MakeNowJust/heredoc"
	"github.com/golang/mock/gomock"
	"github.com/guumaster/logsymbols"
	"github.com/spf13/cobra"

	"github.com/konstellation-io/kli/internal/testhelpers"
	"github.com/konstellation-io/kli/mocks"
)

type testVersionSuite struct {
	ctrl  *gomock.Controller
	mocks versionSuiteMocks
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

func TestNewCreateCmd(t *testing.T) {
	r := testhelpers.NewRunner(t, func(f *mocks.MockCmdFactory) *cobra.Command {
		setupVersionConfig(t, f)

		s := newTestVersionSuite(t)
		f.EXPECT().KreClient("test").Return(s.mocks.kreClient, nil).AnyTimes()
		s.mocks.kreClient.EXPECT().Version().Return(s.mocks.version).AnyTimes()
		s.mocks.version.EXPECT().Create("test.krt").Return("1234", nil).AnyTimes()
		return NewVersionCmd(f)
	})

	r.Run("version create test.krt").
		Containsf(heredoc.Doc(`
      [%s] Upload KRT completed, version 1234 created.
		`), logsymbols.CurrentSymbols().Success)
}
