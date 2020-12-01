package server

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"

	"github.com/konstellation-io/kli/cmdutil"
	"github.com/konstellation-io/kli/pkg/cmd/server/add"
	"github.com/konstellation-io/kli/pkg/cmd/server/list"
	"github.com/konstellation-io/kli/pkg/cmd/server/set"
)

// NewServerCmd creates a new command to handle 'server' subcommands.
func NewServerCmd(f cmdutil.CmdFactory) *cobra.Command {
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
		list.NewListCmd(f),
		set.NewDefaultCmd(f),
		add.NewAddCmd(f),
	)

	return cmd
}
