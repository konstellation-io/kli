package version

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/konstellation-io/kli/cmdutil"
	"github.com/spf13/cobra"
)

// NewVersionCmd creates a new command to handle 'version' subcommands.
func NewVersionCmd(f cmdutil.CmdFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Manage KRE Version",
		Example: heredoc.Doc(`
			$ kli kre version ls --runtime runtime1234
		`),
	}

	cmd.PersistentFlags().StringP("server", "s", f.Config().DefaultServer, "KRE server to use")
	cmd.PersistentFlags().StringP("runtime", "r", "", "Filter for specific runtime")
	_ = cmd.MarkPersistentFlagRequired("runtime")

	cmd.AddCommand(
		NewListCmd(f),
		NewActionCmd(f, "start"),
		NewActionCmd(f, "stop"),
		NewActionCmd(f, "publish"),
		NewActionCmd(f, "unpublish"),
	)

	return cmd
}
