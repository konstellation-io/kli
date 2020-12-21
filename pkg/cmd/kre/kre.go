package kre

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"

	"github.com/konstellation-io/kli/cmd/factory"
	"github.com/konstellation-io/kli/pkg/cmd/kre/version"

	kre "github.com/konstellation-io/kli/pkg/cmd/kre/runtime"
)

// NewKRECmd creates a new command to handle 'kre' keyword.
func NewKRECmd(f factory.CmdFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kre",
		Short: "Manage KRE",
		Example: heredoc.Doc(`
			$ kli kre runtime ls
		`),
	}

	cmd.AddCommand(
		kre.NewRuntimeCmd(f),
		version.NewVersionCmd(f),
	)

	return cmd
}
