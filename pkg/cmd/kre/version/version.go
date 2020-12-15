package version

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/konstellation-io/kli/cmdutil"
	"github.com/konstellation-io/kli/pkg/cmd/kre/version/list"
	"github.com/konstellation-io/kli/pkg/cmd/kre/version/start"
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
		list.NewListCmd(f),
		start.NewStartCmd(f),
	)

	return cmd
}
