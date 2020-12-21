package root

import (
	"testing"

	"github.com/spf13/cobra"

	"github.com/konstellation-io/kli/cmd/iostreams"
	"github.com/konstellation-io/kli/internal/testhelpers"
	"github.com/konstellation-io/kli/mocks"
)

func TestNewVersionCmd(t *testing.T) {
	r := testhelpers.NewRunner(t, func(f *mocks.MockCmdFactory) *cobra.Command {
		f.EXPECT().IOStreams().Return(&iostreams.IOStreams{}).Times(2)

		return NewRootCmd(f, "test-version", "2020-01-01")
	})

	r.Run("kli version").
		Contains("kli version test-version (2020-01-01)")
}
