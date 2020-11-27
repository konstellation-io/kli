package root

import (
	"github.com/spf13/cobra"

	"github.com/konstellation-io/kli/cmdutil"
	"github.com/konstellation-io/kli/pkg/cmd/kre"
	"github.com/konstellation-io/kli/pkg/cmd/server"
	versionCmd "github.com/konstellation-io/kli/pkg/cmd/version"
)

func NewRootCmd(f cmdutil.CmdFactory, version, buildDate string) *cobra.Command {
	cmd := &cobra.Command{
		Use:           "kli [command] [subcommand] [flags]",
		Short:         "Konstellation CLI",
		Long:          `Use Konstellation API from the command line.`,
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	cmd.SetOut(f.IOStreams().Out)
	cmd.SetErr(f.IOStreams().ErrOut)

	// Hide help command. only --help
	cmd.SetHelpCommand(&cobra.Command{
		Hidden: true,
	})

	// Child commands
	cmd.AddCommand(versionCmd.NewVersionCmd(f, version, buildDate))
	cmd.AddCommand(server.NewServerCmd(f))
	cmd.AddCommand(kre.NewKRECmd(f))

	return cmd
}
