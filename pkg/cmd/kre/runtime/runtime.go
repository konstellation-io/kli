package runtime

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"

	"github.com/konstellation-io/kli/cmdutil"
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

	cmd.PersistentFlags().StringP("server", "s", f.Config().DefaultServer, "KRE server to use")

	cmd.AddCommand(
		NewListCmd(f),
	)

	return cmd
}
