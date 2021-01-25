package krt

import (
	"testing"

	"github.com/MakeNowJust/heredoc"
	"github.com/golang/mock/gomock"
	"github.com/guumaster/logsymbols"
	"github.com/spf13/cobra"

	"github.com/konstellation-io/kli/internal/testhelpers"
	"github.com/konstellation-io/kli/mocks"
)

func TestNewKRTCreateCmd(t *testing.T) {
	r := testhelpers.NewRunner(t, func(f *mocks.MockCmdFactory) *cobra.Command {
		c := gomock.NewController(t)
		krt := mocks.NewMockKrtTooler(c)
		f.EXPECT().Krt().Return(krt)

		krt.EXPECT().Build("/test/krt", "test.krt").Return(nil)

		return NewKRTCmd(f)
	})

	r.Run("krt create /test/krt test.krt").
		Containsf(heredoc.Doc(`
	  [%s] New KRT file created.
  `), logsymbols.CurrentSymbols().Success)
}

func TestNewValidateCmd(t *testing.T) {
	r := testhelpers.NewRunner(t, func(f *mocks.MockCmdFactory) *cobra.Command {
		c := gomock.NewController(t)
		krt := mocks.NewMockKrtTooler(c)

		f.EXPECT().Krt().Return(krt)
		krt.EXPECT().Validate("/test/krt.yaml").Return(nil)

		return NewKRTCmd(f)
	})
	r.Run("krt validate /test/krt.yaml").
		Containsf(heredoc.Doc(`
	  [%s] Krt file is valid.
  `), logsymbols.CurrentSymbols().Success)
}

func TestNewKRTCreateCmdNoTarget(t *testing.T) {
	r := testhelpers.NewRunner(t, func(f *mocks.MockCmdFactory) *cobra.Command {
		c := gomock.NewController(t)
		krt := mocks.NewMockKrtTooler(c)
		f.EXPECT().Krt().Return(krt)

		krt.EXPECT().Build("/test/krt", "").Return(nil)

		return NewKRTCmd(f)
	})

	r.Run("krt create /test/krt test.krt").
		Containsf(heredoc.Doc(`
	  [%s] New KRT file created.
  `), logsymbols.CurrentSymbols().Success)
}
