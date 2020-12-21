package server

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"

	"github.com/konstellation-io/kli/cmd/factory"
)

// NewServerCmd creates a new command to handle 'server' subcommands.
func NewServerCmd(f factory.CmdFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server <command>",
		Short: "Manage servers for kli",
		Example: heredoc.Doc(`
			$ kli server ls
			$ kli server add my-server http://api.kre.local TOKEN_12345
			$ kli server default my-server
		`),
	}

	cmd.AddCommand(
		NewListCmd(f),
		NewDefaultCmd(f),
		NewAddCmd(f),
	)

	return cmd
}
