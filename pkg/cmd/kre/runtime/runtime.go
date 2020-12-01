package runtime

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"

	"github.com/konstellation-io/kli/cmdutil"
	"github.com/konstellation-io/kli/pkg/cmd/kre/runtime/list"
)

// NewRuntimeCmd creates a new command to handle 'runtime' subcommands.
func NewRuntimeCmd(f cmdutil.CmdFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "runtime",
		Short: "Manage KRE Runtimes",
		Example: heredoc.Doc(`
			$ kli kre runtime ls
		`),
	}

	cmd.AddCommand(
		list.NewListCmd(f),
	)

	return cmd
}
