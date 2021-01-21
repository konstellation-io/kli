package version

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"

	"github.com/konstellation-io/kli/cmd/factory"
)

// NewVersionCmd creates a new command to handle 'version' subcommands.
func NewVersionCmd(f factory.CmdFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Manage KRE Version",
		Example: heredoc.Doc(`
			$ kli kre version ls
		`),
	}

	cmd.PersistentFlags().StringP("server", "s", f.Config().DefaultServer, "KRE server to use")

	cmd.AddCommand(
		NewListCmd(f),
		NewActionCmd(f, "start"),
		NewActionCmd(f, "stop"),
		NewActionCmd(f, "publish"),
		NewActionCmd(f, "unpublish"),
		NewConfigCmd(f),
		NewCreateCmd(f),
	)

	return cmd
}
